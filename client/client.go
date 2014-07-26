// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// 相对于微信服务器, 主动请求的功能都封装在 Client 里面;
// Client 并发安全, 一般情况下一个应用维护一个 Client 实例即可!
type Client struct {
	tokenCache   TokenCache
	tokenService TokenService

	// 用于单进程环境, 和上面两个字段互斥
	defaultTokenCacheService

	//  NOTE: require go1.3+ , 如果你的环境不满足这个条件, 可以自己实现一个简单的 Pool,
	//        see github.com/chanxuehong/util/pool
	bufferPool sync.Pool // 缓存的是 *bytes.Buffer

	httpClient *http.Client // 可以根据自己的需要定制 http.Client
}

// 创建一个新的 Client.
//  NOTE: 用于单进程环境, 如果是多进程(多机器)环境, 请用 NewClientEx.
//  如果 httpClient == nil 则默认用 http.DefaultClient,
//  see ../CommonHttpClient 和 ../MediaHttpClient.
func NewClient(appid, appsecret string, httpClient *http.Client) (clt *Client) {
	clt = &Client{
		defaultTokenCacheService: defaultTokenCacheService{
			appid:                     appid,
			appsecret:                 appsecret,
			resetTokenRefreshTickChan: make(chan time.Duration),
		},
		bufferPool: sync.Pool{New: newBuffer},
	}

	if httpClient == nil {
		clt.defaultTokenCacheService.httpClient = http.DefaultClient
		clt.httpClient = http.DefaultClient
	} else {
		clt.defaultTokenCacheService.httpClient = httpClient
		clt.httpClient = httpClient
	}

	clt.defaultTokenCacheService.start()
	return
}

// 只是实现了接口, 没有实现功能
var defaultTokenService TokenService = incompleteTokenService(0)

type incompleteTokenService int

func (incompleteTokenService) TokenRefresh() (token string, err error) {
	err = errors.New("没有实现 TokenRefresh 功能")
	return
}

// 创建一个新的 Client.
// tokenCache 必须实现, tokenService, httpClient 均可为 nil.
//  如果 httpClient == nil 则默认用 http.DefaultClient，
//  see ../CommonHttpClient 和 ../MediaHttpClient。
func NewClientEx(tokenCache TokenCache, tokenService TokenService, httpClient *http.Client) (clt *Client) {
	if tokenCache == nil {
		panic("tokenCache == nil")
	}
	if tokenService == nil {
		tokenService = defaultTokenService
	}

	clt = &Client{
		tokenCache:   tokenCache,
		tokenService: tokenService,
		bufferPool:   sync.Pool{New: newBuffer},
	}

	if httpClient == nil {
		clt.httpClient = http.DefaultClient
	} else {
		clt.httpClient = httpClient
	}

	return
}

// Client 通用的 json post 请求
func (c *Client) postJSON(_url string, request interface{}, response interface{}) (err error) {
	buf := c.getBufferFromPool()
	defer c.putBufferToPool(buf)

	if err = json.NewEncoder(buf).Encode(request); err != nil {
		return
	}

	resp, err := c.httpClient.Post(_url, "application/json; charset=utf-8", buf)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", resp.Status)
	}

	if err = json.NewDecoder(resp.Body).Decode(response); err != nil {
		return
	}

	return
}

// Client 通用的 json get 请求
func (c *Client) getJSON(_url string, response interface{}) (err error) {
	resp, err := c.httpClient.Get(_url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", resp.Status)
	}

	if err = json.NewDecoder(resp.Body).Decode(response); err != nil {
		return
	}

	return
}
