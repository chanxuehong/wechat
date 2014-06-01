package wechat

import (
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
	// 缓存当前的 access token, 另起一个 goroutine tokenService() 定期更新;
	_token struct {
		rwmutex sync.RWMutex
		token   string
		err     error
	}
	// goroutine tokenService() 里有个定时器, 每次触发都会更新 access token,
	// 同时 goroutine tokenService() 监听这个 resetTickChan,
	// 如果有新的数据, 则重置定时器, 定时时间为 resetTickChan 传过来的数据;
	// 主要用于用户手动更新 access token 的情况, see Client.TokenRefresh().
	resetTickChan chan time.Duration
	// 对于上传媒体文件, 一般要申请比较大的内存, 所以增加一个内存池;
	// pool.Pool 的接口兼容 sync.Pool.
	bufferPool *pool.Pool
	httpClient *http.Client // It will default to http.DefaultClient if httpClient == nil.
}

// It will default to http.DefaultClient if httpClient == nil.
func NewClient(appid, appsecret string, httpClient *http.Client) *Client {
	const bufferPoolSize = 64 // 不暴露这个选项是为了变更到 sync.Pool 不做大的变动

	c := &Client{
		appid:         appid,
		appsecret:     appsecret,
		resetTickChan: make(chan time.Duration), // 同步 channel
		bufferPool:    pool.New(newBuffer, bufferPoolSize),
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
