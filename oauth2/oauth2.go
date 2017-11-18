package oauth2

import (
	"time"
)

type Endpoint interface {
	ExchangeTokenURL(code string) string        // 通过code换取access_token的地址
	RefreshTokenURL(refreshToken string) string // 刷新access_token的地址
}

type TokenStorage interface {
	Token() (*Token, error)
	PutToken(*Token) error
}

type Token struct {
	AccessToken  string `json:"access_token"`            // 网页授权接口调用凭证
	CreatedAt    int64  `json:"created_at"`              // access_token 创建时间, unixtime, 分布式系统要求时间同步, 建议使用 NTP
	ExpiresIn    int64  `json:"expires_in"`              // access_token 接口调用凭证超时时间, 单位: 秒
	RefreshToken string `json:"refresh_token,omitempty"` // 刷新 access_token 的凭证

	OpenId  string `json:"openid,omitempty"`
	UnionId string `json:"unionid,omitempty"`
	Scope   string `json:"scope,omitempty"` // 用户授权的作用域, 使用逗号(,)分隔
}

// Expired 判断 token.AccessToken 是否过期, 过期返回 true, 否则返回 false.
func (token *Token) Expired() bool {
	return time.Now().Unix() >= token.CreatedAt+token.ExpiresIn
}
