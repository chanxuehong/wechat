package wechat

import (
	"sync"
	"time"
)

type accessToken struct {
	RWMutex    sync.RWMutex `json:"-"`
	TokenValue string
	Expires    int64 // 过期时间戳, unixtime
}

// 获取 access_token, 如果过期返回 ""
func (at *accessToken) Token() (token string) {
	at.RWMutex.RLock()
	if time.Now().Unix()+10 > at.Expires { // 考虑到网络延时, 提前过期
		token = ""
	} else {
		token = at.TokenValue
	}
	at.RWMutex.RUnlock()
	return
}

// 设置新的 access token
//  NOTE: 不做参数检查, 调用者做检查
func (at *accessToken) Update(token string, expiresin int64) {
	at.RWMutex.Lock()
	at.TokenValue = token
	at.Expires = time.Now().Unix() + expiresin
	at.RWMutex.Unlock()
}
