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

// 相对于微信服务器, 主动请求的功能都封装在 Client 里面;
// Client 并发安全, 一个应用维护一个 Client 实例即可!
type Client struct {
	appid, appsecret string

	tokenCache TokenCache

	// 如果 tokenCache == nil，则 Client 自己负责更新 access token，
	// 下面两个字段 currentToken, resetRefreshTokenTickChan 就是做这个用处的

	// currentToken 缓存当前的 access token，goroutine tokenAutoUpdate() 定期更新。
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
//  这个参数可以参考 ../CommonHttpClient 和 ../MediaHttpClient。
func NewClient(appid, appsecret string, httpClient *http.Client) (clt *Client) {
	clt = &Client{
		appid:                     appid,
		appsecret:                 appsecret,
		resetRefreshTokenTickChan: make(chan time.Duration), // 同步 channel
		bufferPool: sync.Pool{
			New: newBuffer,
		},
	}

	if httpClient == nil {
		clt.httpClient = http.DefaultClient
	} else {
		clt.httpClient = httpClient
	}

	tk, err := clt.getNewToken()
	if err != nil {
		clt.updateCurrentToken("", err)
		go clt.tokenAutoUpdate(time.Minute) // 一分钟后尝试
	} else {
		clt.updateCurrentToken(tk.Token, nil)
		go clt.tokenAutoUpdate(time.Duration(tk.ExpiresIn) * time.Second)
	}

	return
}

// 创建一个新的 Client, 一般用于分布式环境.
//  如果 httpClient == nil 则默认用 http.DefaultClient，
//  这个参数可以参考 ../CommonHttpClient 和 ../MediaHttpClient。
func NewClientEx(appid, appsecret string,
	httpClient *http.Client, tokenCache TokenCache) (clt *Client) {

	if tokenCache == nil {
		panic("tokenCache == nil")
	}

	clt = &Client{
		appid:      appid,
		appsecret:  appsecret,
		tokenCache: tokenCache,
		bufferPool: sync.Pool{
			New: newBuffer,
		},
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
