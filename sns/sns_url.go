// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong@gmail.com

package sns

import (
	"net/url"
)

// https://open.weixin.qq.com/connect/oauth2/authorize?appid=APPID
// &redirect_uri=REDIRECT_URI&response_type=code&scope=SCOPE&state=STATE#wechat_redirect
func oauth2AuthURL(appid, redirectURL, scope, state string) string {
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

// https://api.weixin.qq.com/sns/oauth2/access_token?appid=APPID&secret=SECRET
// &code=CODE&grant_type=authorization_code
func oauth2TokenURL(appid, appsecret, code string) string {
	return "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" +
		appid +
		"&secret=" +
		appsecret +
		"&code=" +
		url.QueryEscape(code) +
		"&grant_type=authorization_code"
}

// https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=APPID
// &grant_type=refresh_token&refresh_token=REFRESH_TOKEN
func oauth2RefreshTokenURL(appid, refreshToken string) string {
	return "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=" +
		appid +
		"&grant_type=refresh_token&refresh_token=" +
		refreshToken
}

// https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
func userInfoURL(accessToken, openid, lang string) string {
	return "https://api.weixin.qq.com/sns/userinfo?access_token=" +
		accessToken +
		"&openid=" +
		openid +
		"&lang=" +
		lang
}
