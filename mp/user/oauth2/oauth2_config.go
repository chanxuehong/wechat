// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package oauth2

import (
	"strings"
)

type OAuth2Config struct {
	AppId, AppSecret string

	// 应用授权作用域, 多个作用域用逗号(,)分隔;
	// 目前有 snsapi_base, snsapi_userinfo.
	Scope string

	// 用户授权后跳转的目的地址
	// 用户授权后跳转到 RedirectURL?code=CODE&state=STATE
	// 用户禁止授权跳转到 RedirectURL?state=STATE
	RedirectURL string
}

func NewOAuth2Config(AppId, AppSecret, RedirectURL string, Scope ...string) *OAuth2Config {
	return &OAuth2Config{
		AppId:       AppId,
		AppSecret:   AppSecret,
		Scope:       strings.Join(Scope, ","),
		RedirectURL: RedirectURL,
	}
}

// 请求用户授权获取code的地址.
func (cfg *OAuth2Config) AuthCodeURL(state string) string {
	return AuthCodeURL(cfg.AppId, cfg.RedirectURL, cfg.Scope, state)
}
