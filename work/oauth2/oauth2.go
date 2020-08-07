package oauth2

import (
	"net/url"
)

func AuthCodeURL(appId, redirectURI, scope, state string) string {
	return "https://open.weixin.qq.com/connect/oauth2/authorize?appid=" + url.QueryEscape(appId) +
		"&redirect_uri=" + url.QueryEscape(redirectURI) +
		"&response_type=code&scope=" + url.QueryEscape(scope) +
		"&state=" + url.QueryEscape(state) +
		"#wechat_redirect"
}
