package wechat

import (
	"sync"
	"time"
)

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

func (at *accessToken) Update(token string, err error) {
	at.rwmutex.Lock()
	at.token = token
	at.err = err
	at.rwmutex.Unlock()
}

// 并发安全
type Client struct {
	appid, appsecret string
	accessToken      accessToken
	resetTickChan    chan time.Duration
}

func NewClient(appid, appsecret string) *Client {
	c := &Client{
		appid:         appid,
		appsecret:     appsecret,
		resetTickChan: make(chan time.Duration),
	}
	go c.accessTokenService()
	c.RefreshToken()
	return c
}
