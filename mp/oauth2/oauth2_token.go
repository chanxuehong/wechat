// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package oauth2

import (
	"errors"
	"time"
)

// 用户相关的 oauth2 token 信息
//  NOTE: 每个用户对应一个这样的结构, 应该缓存起来, 一般缓存在 session 中.
type OAuth2Token struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    int64 // 过期时间, unixtime, 分布式系统要求时间同步

	OpenId string
	Scopes []string // 用户授权的作用域
}

// 判断授权的 access token 是否过期, 过期返回 true, 没有过期返回 false
func (this *OAuth2Token) accessTokenExpired() bool {
	return time.Now().Unix() > this.ExpiresAt
}

var _ TokenCache = new(OAuth2Token)

func (this *OAuth2Token) Token() (tk *OAuth2Token, err error) {
	// 防止用户不小心修改了 tk 而影响了原始的, 返回拷贝
	tk = new(OAuth2Token)
	*tk = *this
	return
}

func (this *OAuth2Token) PutToken(tk *OAuth2Token) (err error) {
	if tk == nil {
		return errors.New("input OAuth2Token is nil")
	}
	*this = *tk
	return
}
