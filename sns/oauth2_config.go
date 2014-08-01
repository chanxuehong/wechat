// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package sns

import (
	"strings"
)

// oauth2 相关配置 ( 一般全局只用保存一个变量 )
type OAuth2Config struct {
	AppId     string
	AppSecret string

	// 应用授权作用域，拥有多个作用域用逗号（,）分隔;
	// 目前有 snsapi_base, snsapi_userinfo.
	Scope string

	// 用户授权后跳转的目的地址
	// 用户授权后跳转到 RedirectURL?code=CODE&state=STATE
	// 用户禁止授权跳转到 RedirectURL?state=STATE
	RedirectURL string
}

func NewOAuth2Config(appid, appsecret, redirectURL string, scope ...string) *OAuth2Config {
	return &OAuth2Config{
		AppId:       appid,
		AppSecret:   appsecret,
		Scope:       strings.Join(scope, ","),
		RedirectURL: redirectURL,
	}
}

// 请求用户授权时跳转的地址.
func (cfg *OAuth2Config) AuthCodeURL(state string) string {
	return oauth2AuthURL(cfg.AppId, cfg.RedirectURL, cfg.Scope, state)
}
