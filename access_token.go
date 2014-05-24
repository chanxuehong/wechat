package wechat

import (
	"time"
)

// 从服务器获取 acces_token 成功时返回的消息格式
type AccessTokenResponse struct {
	AccessToken string `json:"access_token"` // 获取到的凭证
	ExpiresIn   int64  `json:"expires_in"`   // 凭证有效时间，单位：秒
}

type accessToken struct {
	TokenValue string
	Expires    int64 // 过期时间戳, unixtime
}

// 获取 access_token, 如果过期返回 ""
func (at *accessToken) Token() string {
	if time.Now().Unix() > at.Expires {
		return ""
	}
	return at.TokenValue
}
