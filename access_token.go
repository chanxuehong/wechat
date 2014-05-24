package wechat

import (
	"time"
)

type accessToken struct {
	TokenValue string
	Expires    int64 // unixtime
}

// 获取 access_token, 如果过期返回 ""
func (at *accessToken) Token() string {
	if time.Now().Unix() > at.Expires {
		return ""
	}
	return at.TokenValue
}
