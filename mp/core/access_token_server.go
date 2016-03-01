package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"
)

// access_token 中控服务器接口.
type AccessTokenServer interface {
	Token() (token string, err error)        // 请求中控服务器返回缓存的 access_token
	TokenRefresh() (token string, err error) // 请求中控服务器刷新 access_token
	IID01332E16DF5011E5A9D5A4DB30FED8E1()    // 接口标识, 没有实际意义
}

var _ AccessTokenServer = (*DefaultAccessTokenServer)(nil)

// DefaultAccessTokenServer 实现了 AccessTokenServer 接口.
//  NOTE:
//  1. 用于单进程环境.
//  2. 因为 DefaultAccessTokenServer 同时也是一个简单的中控服务器, 而不是仅仅实现 AccessTokenServer 接口,
//     所以整个系统只能存在一个 DefaultAccessTokenServer 实例!
type DefaultAccessTokenServer struct {
	appId      string
	appSecret  string
	httpClient *http.Client

	ticker          *time.Ticker       // 用于定时更新 access_token
	resetTickerChan chan time.Duration // 用于重置 ticker

	tokenGet struct {
		sync.Mutex
		lastAccessToken accessToken // 最近一次成功更新 access_token 的数据信息
		lastTimestamp   time.Time   // 最近一次成功更新 access_token 的时间戳
	}

	tokenCache struct {
		sync.RWMutex
		token string // 保存有效的 access_token, 当更新 access_token 失败时此字段置空
	}
}

// 微信服务器返回的 access_token 的数据结构.
type accessToken struct {
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

// NewDefaultAccessTokenServer 创建一个新的 DefaultAccessTokenServer, 如果 httpClient == nil 则默认使用 http.DefaultClient.
func NewDefaultAccessTokenServer(appId, appSecret string, httpClient *http.Client) (srv *DefaultAccessTokenServer) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	srv = &DefaultAccessTokenServer{
		appId:           url.QueryEscape(appId),
		appSecret:       url.QueryEscape(appSecret),
		httpClient:      httpClient,
		ticker:          nil,
		resetTickerChan: make(chan time.Duration),
	}

	go srv.tokenUpdateDaemon(time.Hour * 24 * time.Duration(100+rand.Int63n(200)))
	return
}

func (srv *DefaultAccessTokenServer) IID01332E16DF5011E5A9D5A4DB30FED8E1() {}

func (srv *DefaultAccessTokenServer) Token() (token string, err error) {
	srv.tokenCache.RLock()
	token = srv.tokenCache.token
	srv.tokenCache.RUnlock()

	if token != "" {
		return
	}
	return srv.TokenRefresh()
}

func (srv *DefaultAccessTokenServer) TokenRefresh() (token string, err error) {
	accessToken, cached, err := srv.tokenRefresh()
	if err != nil {
		return
	}
	if !cached {
		srv.resetTickerChan <- time.Duration(accessToken.ExpiresIn) * time.Second
	}
	token = accessToken.Token
	return
}

func (srv *DefaultAccessTokenServer) tokenUpdateDaemon(initTickDuration time.Duration) {
	tickDuration := initTickDuration

NEW_TICK_DURATION:
	srv.ticker = time.NewTicker(tickDuration)
	for {
		select {
		case tickDuration = <-srv.resetTickerChan:
			srv.ticker.Stop()
			goto NEW_TICK_DURATION

		case <-srv.ticker.C:
			accessToken, cached, err := srv.tokenRefresh()
			if err != nil {
				break
			}
			if !cached {
				newTickDuration := time.Duration(accessToken.ExpiresIn) * time.Second
				if tickDuration != newTickDuration {
					tickDuration = newTickDuration
					srv.ticker.Stop()
					goto NEW_TICK_DURATION
				}
			}
		}
	}
}

// tokenRefresh 从微信服务器获取新的 access_token 并存入缓存, 同时返回该 access_token.
func (srv *DefaultAccessTokenServer) tokenRefresh() (token accessToken, cached bool, err error) {
	srv.tokenGet.Lock()
	defer srv.tokenGet.Unlock()

	timeNow := time.Now()

	// 如果在收敛周期内则直接返回最近一次获取的 access_token!
	//
	// 当 access_token 失效时, SDK 内部会尝试刷新 access_token, 这时在短时间(一个收敛周期)内可能会有多个 goroutine 同时进行刷新,
	// 实际上我们没有必要也不能让这些操作都去微信服务器获取新的 access_token, 而只用返回同一个 access_token 即可.
	// 因为 access_token 缓存在内存里, 所以收敛周期为2个http周期, 我们这里设定为2秒.
	if d := timeNow.Sub(srv.tokenGet.lastTimestamp); 0 <= d && d < time.Second*2 {
		token = srv.tokenGet.lastAccessToken
		cached = true
		return
	}

	_url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + srv.appId +
		"&secret=" + srv.appSecret
	httpResp, err := srv.httpClient.Get(_url)
	if err != nil {
		srv.tokenCache.Lock()
		srv.tokenCache.token = ""
		srv.tokenCache.Unlock()
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		srv.tokenCache.Lock()
		srv.tokenCache.token = ""
		srv.tokenCache.Unlock()
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	var result struct {
		Error
		accessToken
	}
	if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
		srv.tokenCache.Lock()
		srv.tokenCache.token = ""
		srv.tokenCache.Unlock()
		return
	}

	if result.ErrCode != ErrCodeOK {
		srv.tokenCache.Lock()
		srv.tokenCache.token = ""
		srv.tokenCache.Unlock()
		err = &result.Error
		return
	}

	// 由于网络的延时, access_token 过期时间留有一个缓冲区
	switch {
	case result.ExpiresIn > 31556952: // 60*60*24*365.2425
		srv.tokenCache.Lock()
		srv.tokenCache.token = ""
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
		srv.tokenCache.token = ""
		srv.tokenCache.Unlock()
		err = errors.New("expires_in too small: " + strconv.FormatInt(result.ExpiresIn, 10))
		return
	}

	// 更新 tokenGet
	srv.tokenGet.lastAccessToken = result.accessToken
	srv.tokenGet.lastTimestamp = timeNow

	// 更新 tokenCache
	srv.tokenCache.Lock()
	srv.tokenCache.token = result.Token
	srv.tokenCache.Unlock()

	token = result.accessToken
	return
}
