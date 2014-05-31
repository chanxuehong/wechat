package wechat

import (
	"github.com/chanxuehong/util/pool"
	"net"
	"net/http"
	"sync"
	"time"
)

// 缓存 access token, 以 Client 的一个字段存在;
// 每次获取当前的 access token 都是从这个结构里读取, see Client.Token()
type accessToken struct {
	rwmutex sync.RWMutex
	token   string
	err     error
}

func (at *accessToken) Token() (token string, err error) {
	at.rwmutex.RLock()
	token = at.token
	err = at.err
	at.rwmutex.RUnlock()
	return
}

// see Client.TokenRefresh() and Client.accessTokenService()
func (at *accessToken) Update(token string, err error) {
	at.rwmutex.Lock()
	at.token = token
	at.err = err
	at.rwmutex.Unlock()
}

// 相对于微信服务器, 主动请求的功能模块都相当于是 Client;
// 这个结构就是封装所有这些请求功能, 并发安全.
//  NOTE: 必须调用 NewClient() 创建对象!
type Client struct {
	appid, appsecret string
	httpClient       *http.Client
	// 缓存当前的 access token, 另起一个 goroutine accessTokenService() 定期更新;
	accessToken accessToken
	// goroutine accessTokenService() 里有个定时器, 每次触发都会更新 access token,
	// 同时 goroutine accessTokenService() 监听这个 resetTickChan,
	// 如果有新的数据, 则重置定时器, 定时时间为 resetTickChan 传过来的数据;
	// 主要用于用户手动更新 access token 的情况, Client.TokenRefresh().
	resetTickChan chan time.Duration
	// 对于上传媒体文件, 一般要申请比较大的内存, 所以增加一个内存池;
	// pool.Pool 的接口兼容 sync.Pool.
	bufferPool *pool.Pool
}

func NewClient(appid, appsecret string) *Client {
	const bufferPoolSize = 64 // 不暴露这个选项是为了变更到 sync.Pool 不做大的变动

	c := &Client{
		appid:     appid,
		appsecret: appsecret,

		httpClient: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				Dial: (&net.Dialer{
					Timeout:   5 * time.Second, // 连接超时设置为 5 秒
					KeepAlive: 30 * time.Second,
				}).Dial,
				TLSHandshakeTimeout: 5 * time.Second, // TLS 握手超时设置为 5 秒
			},
			Timeout: 15 * time.Second, // 请求超时时间设置为 15 秒
		},

		resetTickChan: make(chan time.Duration), // 同步 channel
		bufferPool:    pool.New(newBuffer, bufferPoolSize),
	}
	go c.accessTokenService() // 定时更新 c.accessToken
	c.TokenRefresh()          // *同步*获取 access token
	return c
}
