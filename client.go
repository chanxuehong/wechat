package wechat

import (
	"sync"
	"time"
)

// 缓存 access token, 每次获取当前的 access token 都是从这个结构里读取.
// see Client.Token()
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

// see Client.RefreshToken() and Client.accessTokenService()
func (at *accessToken) Update(token string, err error) {
	at.rwmutex.Lock()
	at.token = token
	at.err = err
	at.rwmutex.Unlock()
}

// 并发安全
type Client struct {
	appid, appsecret string
	// 缓存当前的 access token, 另起一个 goroutine 定期更新
	accessToken accessToken
	// 负责更新 accessToken 的 goroutine 里有个定时器, 监听这个 resetTickChan,
	// 如果有新的数据, 则重置定时器, 定时时间为 resetTickChan 里的数据.
	// 主要用于用户手动更新 access token 的情况, Client.RefreshToken()
	resetTickChan chan time.Duration
}

func NewClient(appid, appsecret string) *Client {
	c := &Client{
		appid:         appid,
		appsecret:     appsecret,
		resetTickChan: make(chan time.Duration),
	}
	go c.accessTokenService()
	c.RefreshToken() // 同步获取 access token
	return c
}
