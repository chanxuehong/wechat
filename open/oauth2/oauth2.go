package oauth2

import (
	"net/http"
	"net/url"

	mpoauth2 "github.com/chanxuehong/wechat/mp/oauth2"
)

// AuthCodeURL 生成网页授权地址.
//  appId:       公众号的唯一标识
//  redirectURI: 授权后重定向的回调链接地址
//  scope:       应用授权作用域
//  state:       重定向后会带上 state 参数, 开发者可以填写 a-zA-Z0-9 的参数值, 最多128字节
func AuthCodeURL(appId, redirectURI, scope, state string) string {
	return "https://open.weixin.qq.com/connect/qrconnect?appid=" + url.QueryEscape(appId) +
		"&redirect_uri=" + url.QueryEscape(redirectURI) +
		"&response_type=code&scope=" + url.QueryEscape(scope) +
		"&state=" + url.QueryEscape(state) +
		"#wechat_redirect"
}

// Auth 检验授权凭证 access_token 是否有效.
//  accessToken: 网页授权接口调用凭证
//  openId:      用户的唯一标识
//  httpClient:  如果不指定则默认为 util.DefaultHttpClient
func Auth(accessToken, openId string, httpClient *http.Client) (valid bool, err error) {
	return mpoauth2.Auth(accessToken, openId, httpClient)
}
