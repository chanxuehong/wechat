// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

// +build !wechatdebug

package thirdparty

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/chanxuehong/wechat/corp"
)

type SuiteAccessToken string

// suite_access_token 中控服务器接口.
type SuiteAccessTokenServer interface {
	// 从中控服务器获取被缓存的 suite_access_token.
	Token() (token SuiteAccessToken, err error)

	// 请求中控服务器到微信服务器刷新 suite_access_token.
	//
	//  高并发场景下某个时间点可能有很多请求(比如缓存的suite_access_token刚好过期时), 但是我们
	//  不期望也没有必要让这些请求都去微信服务器获取 suite_access_token(有可能导致api超过调用限制),
	//  实际上这些请求只需要一个新的 suite_access_token 即可, 所以建议 SuiteAccessTokenServer 从微信服务器
	//  获取一次 suite_access_token 之后的至多5秒内(收敛时间, 视情况而定, 理论上至多5个http或tcp周期)
	//  再次调用该函数不再去微信服务器获取, 而是直接返回之前的结果.
	TokenRefresh() (token SuiteAccessToken, err error)
}

var _ SuiteAccessTokenServer = (*DefaultSuiteAccessTokenServer)(nil)

// SuiteAccessTokenServer 的简单实现.
//  NOTE:
//  1. 用于单进程环境.
//  2. 因为 DefaultSuiteAccessTokenServer 同时也是一个简单的中控服务器, 而不是仅仅实现 SuiteAccessTokenServer 接口,
//     所以整个系统只能存在一个 DefaultSuiteAccessTokenServer 实例!
type DefaultSuiteAccessTokenServer struct {
	suiteId           string
	suiteSecret       string
	suiteTicketGetter SuiteTicketGetter
	httpClient        *http.Client

	resetTickerChan chan time.Duration // 用于重置 tokenDaemon 里的 ticker

	tokenGet struct {
		sync.Mutex
		LastTokenInfo suiteAccessTokenInfo // 最后一次成功从微信服务器获取的 suite_access_token 信息
		LastTimestamp int64                // 最后一次成功从微信服务器获取 suite_access_token 的时间戳
	}

	tokenCache struct {
		sync.RWMutex
		Token string
	}
}

// 创建一个新的 DefaultSuiteAccessTokenServer.
//  如果 httpClient == nil 则默认使用 http.DefaultClient.
func NewDefaultSuiteAccessTokenServer(suiteId, suiteSecret string, suiteTicketGetter SuiteTicketGetter,
	httpClient *http.Client) (srv *DefaultSuiteAccessTokenServer) {

	if suiteId == "" {
		panic("empty suiteId")
	}
	if suiteSecret == "" {
		panic("empty suiteSecret")
	}
	if suiteTicketGetter == nil {
		panic("nil suiteTicketGetter")
	}
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	srv = &DefaultSuiteAccessTokenServer{
		suiteId:           suiteId,
		suiteSecret:       suiteSecret,
		suiteTicketGetter: suiteTicketGetter,
		httpClient:        httpClient,
		resetTickerChan:   make(chan time.Duration),
	}

	go srv.tokenDaemon(time.Hour * 24) // 启动 tokenDaemon
	return
}

func (srv *DefaultSuiteAccessTokenServer) Token() (token SuiteAccessToken, err error) {
	srv.tokenCache.RLock()
	token = SuiteAccessToken(srv.tokenCache.Token)
	srv.tokenCache.RUnlock()

	if token != "" {
		return
	}
	return srv.TokenRefresh()
}

func (srv *DefaultSuiteAccessTokenServer) TokenRefresh() (token SuiteAccessToken, err error) {
	suiteAccessTokenInfo, cached, err := srv.getToken()
	if err != nil {
		return
	}
	if !cached {
		srv.resetTickerChan <- time.Duration(suiteAccessTokenInfo.ExpiresIn) * time.Second
	}
	token = SuiteAccessToken(suiteAccessTokenInfo.Token)
	return
}

func (srv *DefaultSuiteAccessTokenServer) tokenDaemon(tickDuration time.Duration) {
NEW_TICK_DURATION:
	ticker := time.NewTicker(tickDuration)

	for {
		select {
		case tickDuration = <-srv.resetTickerChan:
			ticker.Stop()
			goto NEW_TICK_DURATION

		case <-ticker.C:
			suiteAccessTokenInfo, cached, err := srv.getToken()
			if err != nil {
				break
			}
			if !cached {
				newTickDuration := time.Duration(suiteAccessTokenInfo.ExpiresIn) * time.Second
				if tickDuration != newTickDuration {
					tickDuration = newTickDuration
					ticker.Stop()
					goto NEW_TICK_DURATION
				}
			}
		}
	}
}

type suiteAccessTokenInfo struct {
	Token     string `json:"suite_access_token"`
	ExpiresIn int64  `json:"expires_in"` // 有效时间, seconds
}

// 从微信服务器获取 suite_access_token.
//  同一时刻只能一个 goroutine 进入, 防止没必要的重复获取.
func (srv *DefaultSuiteAccessTokenServer) getToken() (token suiteAccessTokenInfo, cached bool, err error) {
	srv.tokenGet.Lock()
	defer srv.tokenGet.Unlock()

	timeNowUnix := time.Now().Unix()

	// 在收敛周期内直接返回最近一次获取的 suite_access_token, 这里的收敛时间设定为4秒.
	if n := srv.tokenGet.LastTimestamp; n <= timeNowUnix && timeNowUnix < n+4 {
		// 因为只有成功获取后才会更新 srv.tokenGet.LastTimestamp, 所以这些都是有效数据
		token = suiteAccessTokenInfo{
			Token:     srv.tokenGet.LastTokenInfo.Token,
			ExpiresIn: srv.tokenGet.LastTokenInfo.ExpiresIn - timeNowUnix + n,
		}
		cached = true
		return
	}

	suiteTicket, err := srv.suiteTicketGetter.GetSuiteTicket(srv.suiteId)
	if err != nil {
		srv.tokenCache.Lock()
		srv.tokenCache.Token = ""
		srv.tokenCache.Unlock()
		return
	}

	request := struct {
		SuiteId     string `json:"suite_id"`
		SuiteSecret string `json:"suite_secret"`
		SuiteTicket string `json:"suite_ticket"`
	}{
		SuiteId:     srv.suiteId,
		SuiteSecret: srv.suiteSecret,
		SuiteTicket: string(suiteTicket),
	}

	requestBuf := textBufferPool.Get().(*bytes.Buffer)
	requestBuf.Reset()
	defer textBufferPool.Put(requestBuf)

	if err = json.NewEncoder(requestBuf).Encode(&request); err != nil {
		srv.tokenCache.Lock()
		srv.tokenCache.Token = ""
		srv.tokenCache.Unlock()
		return
	}
	requestBytes := requestBuf.Bytes()

	url := "https://qyapi.weixin.qq.com/cgi-bin/service/get_suite_token"
	httpResp, err := srv.httpClient.Post(url, "application/json; charset=utf-8", bytes.NewReader(requestBytes))
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
		corp.Error
		suiteAccessTokenInfo
	}

	if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
		srv.tokenCache.Lock()
		srv.tokenCache.Token = ""
		srv.tokenCache.Unlock()
		return
	}

	if result.ErrCode != corp.ErrCodeOK {
		srv.tokenCache.Lock()
		srv.tokenCache.Token = ""
		srv.tokenCache.Unlock()

		err = &result.Error
		return
	}

	// 由于网络的延时, suite_access_token 过期时间留了一个缓冲区
	switch {
	case result.ExpiresIn > 31556952: // 60*60*24*365.2425
		srv.tokenCache.Lock()
		srv.tokenCache.Token = ""
		srv.tokenCache.Unlock()

		err = errors.New("expires_in too large: " + strconv.FormatInt(result.ExpiresIn, 10))
		return
	case result.ExpiresIn > 60*60:
		result.ExpiresIn -= 60 * 10
	case result.ExpiresIn > 60*30:
		result.ExpiresIn -= 60 * 5
	case result.ExpiresIn > 60*5:
		result.ExpiresIn -= 60
	case result.ExpiresIn > 60:
		result.ExpiresIn -= 10
	default:
		srv.tokenCache.Lock()
		srv.tokenCache.Token = ""
		srv.tokenCache.Unlock()

		err = errors.New("expires_in too small: " + strconv.FormatInt(result.ExpiresIn, 10))
		return
	}

	// 更新 tokenGet 信息
	srv.tokenGet.LastTokenInfo = result.suiteAccessTokenInfo
	srv.tokenGet.LastTimestamp = timeNowUnix

	// 更新缓存
	srv.tokenCache.Lock()
	srv.tokenCache.Token = result.suiteAccessTokenInfo.Token
	srv.tokenCache.Unlock()

	token = result.suiteAccessTokenInfo
	return
}
