// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// 对于分布式应用，统一由（逻辑上）一个服务器来获取和更新 access token，其他的服务都从这个
// 服务器上获取 access token 或者请求服务器更新 access token。
// TokenService 就是这个服务器的接口。
type TokenService interface {
	Token() (token string, err error)        // 从服务器获取 access token
	TokenRefresh() (token string, err error) // 请求服务器刷新 access token
}

// 相对于微信服务器, 主动请求的功能模块都相当于是 Client;
// 这个结构就是封装所有这些请求功能, 并发安全.
type Client struct {
	tokenService TokenService

	// 如果 tokenService == nil，则 Client 自己负责更新 access token，
	// 下面 4 个字段就是做这个用的。

	appid, appsecret string

	// 缓存当前的 access token，goroutine tokenAutoUpdate() 定期更新。
	currentToken struct {
		rwmutex sync.RWMutex
		token   string
		err     error // 获取或更新 access token 的时候可能会出错, 错误保存在这里
	}

	// goroutine tokenAutoUpdate() 里有个定时器, 每次触发都会更新 access token,
	// 同时 goroutine tokenAutoUpdate() 监听这个 resetRefreshTokenTickChan,
	// 如果有新的数据, 则重置定时器, 定时时间为 resetRefreshTokenTickChan 传过来的数据;
	// 主要用于用户手动更新 access token 的情况, see Client.TokenRefresh().
	resetRefreshTokenTickChan chan time.Duration

	//  NOTE: require go1.3+ , 如果你的环境不满足这个条件, 可以自己实现一个简单的 Pool,
	//        see github.com/chanxuehong/util/pool
	bufferPool sync.Pool // 缓存的是 *bytes.Buffer

	httpClient *http.Client // 可以根据自己的需要定制 http.Client
}

// 创建一个新的 Client, 一般用于单进程环境.
//  如果 httpClient == nil 则默认用 http.DefaultClient，
//  这个参数可以参考 CommonHttpClient 和 MediaHttpClient。
func NewClient(appid, appsecret string, httpClient *http.Client) *Client {
	c := Client{
		appid:                     appid,
		appsecret:                 appsecret,
		resetRefreshTokenTickChan: make(chan time.Duration), // 同步 channel
		bufferPool: sync.Pool{
			New: newBuffer,
		},
	}

	if httpClient == nil {
		c.httpClient = http.DefaultClient
	} else {
		c.httpClient = httpClient
	}

	tk, err := c.getNewToken()
	if err != nil {
		c.updateCurrentToken("", err)
		go c.tokenAutoUpdate(time.Minute) // 一分钟后尝试
	} else {
		c.updateCurrentToken(tk.Token, nil)
		go c.tokenAutoUpdate(time.Duration(tk.ExpiresIn) * time.Second)
	}

	return &c
}

// 创建一个新的 Client, 一般用于分布式环境.
//  如果 httpClient == nil 则默认用 http.DefaultClient，
//  这个参数可以参考 CommonHttpClient 和 MediaHttpClient。
func NewClientEx(tokenService TokenService, httpClient *http.Client) *Client {
	if tokenService == nil {
		panic("tokenService == nil")
	}

	c := Client{
		tokenService: tokenService,
		bufferPool: sync.Pool{
			New: newBuffer,
		},
	}

	if httpClient == nil {
		c.httpClient = http.DefaultClient
	} else {
		c.httpClient = httpClient
	}

	return &c
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
