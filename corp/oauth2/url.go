// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package oauth2

import (
	"net/url"
)

// 构造认证页面的URL
//  appid:       企业的CorpID
//  redirectURL: 授权后重定向的回调链接地址，请使用urlencode对链接进行处理
//               用户授权后跳转到 redirectURL?code=CODE&state=STATE
//               用户禁止授权跳转到 redirectURL?state=STATE
//  scope:       应用授权作用域，此时固定为：snsapi_base
//  state:       重定向后会带上state参数，企业可以填写a-zA-Z0-9的参数值
func OAuth2AuthURL(appid, redirectURL, scope, state string) string {
	// https://open.weixin.qq.com/connect/oauth2/authorize?appid=CORPID&redirect_uri=REDIRECT_URI
	// &response_type=code&scope=SCOPE&state=STATE#wechat_redirect

	return "https://open.weixin.qq.com/connect/oauth2/authorize?appid=" +
		appid +
		"&redirect_uri=" +
		url.QueryEscape(redirectURL) +
		"&response_type=code&scope=" +
		url.QueryEscape(scope) +
		"&state=" +
		url.QueryEscape(state) +
		"#wechat_redirect"
}
