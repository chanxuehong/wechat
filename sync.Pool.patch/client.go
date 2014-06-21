package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/chanxuehong/util/pool"
	"net/http"
	"sync"
	"time"
)

// 相对于微信服务器, 主动请求的功能模块都相当于是 Client;
// 这个结构就是封装所有这些请求功能, 并发安全.
//  NOTE: 必须调用 NewClient() 创建对象!
type Client struct {
	appid, appsecret string

	// 当前的 access token.
	// 由于 token 会过期, 另起一个 goroutine tokenService() 定期更新;
	currentToken struct {
		rwmutex sync.RWMutex
		token   string
		err     error // 获取或更新 access token 的时候可能会出错, 错误保存在这里
	}

	// goroutine tokenService() 里有个定时器, 每次触发都会更新 access token,
	// 同时 goroutine tokenService() 监听这个 resetRefreshTokenTickChan,
	// 如果有新的数据, 则重置定时器, 定时时间为 resetRefreshTokenTickChan 传过来的数据;
	// 主要用于用户手动更新 access token 的情况, see Client.TokenRefresh().
	resetRefreshTokenTickChan chan time.Duration

	//  NOTE: pool.Pool 兼容 go1.3+ 的 sync.Pool, 用于 go1.3 以下的环境.
	bufferPool *pool.Pool // 缓存的是 *bytes.Buffer

	httpClient *http.Client // 可以根据自己的需要定制 http.Client
}

// It will default to http.DefaultClient if httpClient == nil.
//  see wechat.CommonHttpClient and wechat.MediaHttpClient
func NewClient(appid, appsecret string, httpClient *http.Client) *Client {
	c := &Client{
		appid:                     appid,
		appsecret:                 appsecret,
		resetRefreshTokenTickChan: make(chan time.Duration),      // 同步 channel
		bufferPool:                pool.New(clientNewBuffer, 16), // 这个常数 16 可以根据实际修改
	}

	if httpClient == nil {
		c.httpClient = http.DefaultClient
	} else {
		c.httpClient = httpClient
	}

	go c.tokenService() // 定时更新 c.token
	c.TokenRefresh()    // *同步*获取 access token

	return c
}

// Client 通用的 json post 请求
func (c *Client) postJSON(_url string, request interface{}, response interface{}) (err error) {
	buf := c.getBufferFromPool()
	// defer c.putBufferToPool(buf) // buf 要快速迭代, 所以不用 defer, 尽量提前释放

	if err = json.NewEncoder(buf).Encode(request); err != nil {
		c.putBufferToPool(buf) //
		return
	}

	resp, err := c.httpClient.Post(_url, "application/json; charset=utf-8", buf)
	c.putBufferToPool(buf) //
	if err != nil {
		return
	}
	// defer resp.Body.Close() // 逻辑简单, 不用 defer

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close() //
		return fmt.Errorf("http.Status: %s", resp.Status)
	}

	if err = json.NewDecoder(resp.Body).Decode(response); err != nil {
		resp.Body.Close() //
		return
	}

	resp.Body.Close() //
	return
}

// Client 通用的 json get 请求
func (c *Client) getJSON(_url string, response interface{}) (err error) {
	resp, err := c.httpClient.Get(_url)
	if err != nil {
		return
	}
	// defer resp.Body.Close() // 逻辑简单, 不用 defer

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close() //
		return fmt.Errorf("http.Status: %s", resp.Status)
	}

	if err = json.NewDecoder(resp.Body).Decode(response); err != nil {
		resp.Body.Close() //
		return
	}

	resp.Body.Close() //
	return
}
