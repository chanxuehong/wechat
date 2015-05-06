// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

// +build wechatdebug

package corp

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"
)

// access_token 中控服务器接口, see token_server.png
type TokenServer interface {
	// 从中控服务器获取被缓存的 access_token.
	Token() (token string, err error)

	// 请求中控服务器到微信服务器刷新 access_token.
	//
	//  高并发场景下某个时间点可能有很多请求(比如缓存的access_token刚好过期时), 但是我们
	//  不期望也没有必要让这些请求都去微信服务器获取 access_token(有可能导致api超过调用限制),
	//  实际上这些请求只需要一个新的 access_token 即可, 所以建议 TokenServer 从微信服务器
	//  获取一次 access_token 之后的至多5秒内(收敛时间, 视情况而定, 理论上至多5个http或tcp周期)
	//  再次调用该函数不再去微信服务器获取, 而是直接返回之前的结果.
	TokenRefresh() (token string, err error)
}

var _ TokenServer = (*DefaultTokenServer)(nil)

// TokenServer 的简单实现.
//  NOTE:
//  1. 用于单进程环境.
//  2. 因为 DefaultTokenServer 同时也是一个简单的中控服务器, 而不是仅仅实现 TokenServer 接口,
//     所以整个系统只能存在一个 DefaultTokenServer 实例!
type DefaultTokenServer struct {
	corpId     string
	corpSecret string
	httpClient *http.Client

	resetTickerChan chan time.Duration // 用于重置 tokenDaemon 里的 ticker

	tokenGet struct {
		sync.Mutex
		LastTokenInfo tokenInfo // 最后一次成功从微信服务器获取的 access_token 信息
		LastTimestamp int64     // 最后一次成功从微信服务器获取 access_token 的时间戳
	}

	tokenCache struct {
		sync.RWMutex
		Token string
	}
}

// 创建一个新的 DefaultTokenServer.
//  如果 httpClient == nil 则默认使用 http.DefaultClient.
func NewDefaultTokenServer(corpId, corpSecret string,
	httpClient *http.Client) (srv *DefaultTokenServer) {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	srv = &DefaultTokenServer{
		corpId:          corpId,
		corpSecret:      corpSecret,
		httpClient:      httpClient,
		resetTickerChan: make(chan time.Duration),
	}

	const tenYears time.Duration = time.Minute * 5259492 // time.Minute * 60 * 24 * 365.2425 * 10
	go srv.tokenDaemon(tenYears)                         // 启动 tokenDaemon
	return
}

func (srv *DefaultTokenServer) Token() (token string, err error) {
	srv.tokenCache.RLock()
	token = srv.tokenCache.Token
	srv.tokenCache.RUnlock()

	if token != "" {
		return
	}
	return srv.TokenRefresh()
}

func (srv *DefaultTokenServer) TokenRefresh() (token string, err error) {
	tokenInfo, cached, err := srv.getToken()
	if err != nil {
		return
	}
	if !cached {
		srv.resetTickerChan <- time.Duration(tokenInfo.ExpiresIn) * time.Second
	}
	token = tokenInfo.Token
	return
}

func (srv *DefaultTokenServer) tokenDaemon(tickDuration time.Duration) {
NEW_TICK_DURATION:
	ticker := time.NewTicker(tickDuration)

	for {
		select {
		case tickDuration = <-srv.resetTickerChan:
			ticker.Stop()
			goto NEW_TICK_DURATION

		case <-ticker.C:
			tokenInfo, cached, err := srv.getToken()
			if err != nil {
				continue
			}
			if !cached {
				newTickDuration := time.Duration(tokenInfo.ExpiresIn) * time.Second
				if tickDuration != newTickDuration {
					tickDuration = newTickDuration
					ticker.Stop()
					goto NEW_TICK_DURATION
				}
			}
		}
	}
}

type tokenInfo struct {
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"` // 有效时间, seconds
}

// 从微信服务器获取 access_token.
//  同一时刻只能一个 goroutine 进入, 防止没必要的重复获取.
func (srv *DefaultTokenServer) getToken() (token tokenInfo, cached bool, err error) {
	srv.tokenGet.Lock()
	defer srv.tokenGet.Unlock()

	timeNowUnix := time.Now().Unix()

	// 在收敛周期内直接返回最近一次获取的 access_token,
	// 这里的收敛时间设定为2秒, 因为在同一个进程内, 收敛周期为2个http周期
	if n := srv.tokenGet.LastTimestamp; n <= timeNowUnix && timeNowUnix < n+2 {
		// 因为只有成功获取后才会更新 srv.tokenGet.LastTimestamp, 所以这些都是有效数据
		token = tokenInfo{
			Token:     srv.tokenGet.LastTokenInfo.Token,
			ExpiresIn: srv.tokenGet.LastTokenInfo.ExpiresIn - timeNowUnix + n,
		}
		cached = true
		return
	}

	_url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + url.QueryEscape(srv.corpId) +
		"&corpsecret=" + url.QueryEscape(srv.corpSecret)
	httpResp, err := srv.httpClient.Get(_url)
	if err != nil {
		srv.tokenCache.Lock()
		srv.tokenCache.Token = ""
		srv.tokenCache.Unlock()
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		srv.tokenCache.Lock()
		srv.tokenCache.Token = ""
		srv.tokenCache.Unlock()

		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	var result struct {
		Error
		tokenInfo
	}

	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		srv.tokenCache.Lock()
		srv.tokenCache.Token = ""
		srv.tokenCache.Unlock()
		return
	}

	LogInfoln("[WECHAT_DEBUG] request url:", _url)
	LogInfoln("[WECHAT_DEBUG] response json:", string(respBody))

	if err = json.Unmarshal(respBody, &result); err != nil {
		srv.tokenCache.Lock()
		srv.tokenCache.Token = ""
		srv.tokenCache.Unlock()
		return
	}

	if result.ErrCode != ErrCodeOK {
		srv.tokenCache.Lock()
		srv.tokenCache.Token = ""
		srv.tokenCache.Unlock()

		err = &result.Error
		return
	}

	if result.ExpiresIn <= 60 {
		srv.tokenCache.Lock()
		srv.tokenCache.Token = ""
		srv.tokenCache.Unlock()

		err = errors.New("expires_in too small: " + strconv.FormatInt(result.ExpiresIn, 10))
		return
	}

	// 由于企业号的 access_token 会自动续期, 并且不做改变, 这样对于安全不利,
	// 所以这里故意增加1秒, 让其过期, 获取一个不同的 access_token.
	result.ExpiresIn++

	// 更新 tokenGet 信息
	srv.tokenGet.LastTokenInfo = result.tokenInfo
	srv.tokenGet.LastTimestamp = timeNowUnix

	// 更新缓存
	srv.tokenCache.Lock()
	srv.tokenCache.Token = result.tokenInfo.Token
	srv.tokenCache.Unlock()

	token = result.tokenInfo
	return
}
