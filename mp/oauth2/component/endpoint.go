package component

import (
	"fmt"
	"net/url"

	"github.com/chanxuehong/wechat/oauth2"
)

var _ oauth2.Endpoint = (*Endpoint)(nil)

// Endpoint 实现了 wechat.v2/oauth2.Endpoint 接口.
type Endpoint struct {
	AppId                string
	ComponentAppId       string
	ComponentAccessToken string
}

func NewEndpoint(appId, componentAppId, componentAccessToken string) *Endpoint {
	return &Endpoint{
		AppId:                appId,
		ComponentAppId:       componentAppId,
		ComponentAccessToken: componentAccessToken,
	}
}

func (p *Endpoint) ExchangeTokenURL(code string) string {
	return fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/component/access_token?"+
		"appid=%s&component_appid=%s&component_access_token=%s&code=%s&grant_type=authorization_code",
		url.QueryEscape(p.AppId),
		url.QueryEscape(p.ComponentAppId),
		url.QueryEscape(p.ComponentAccessToken),
		url.QueryEscape(code),
	)
}

func (p *Endpoint) RefreshTokenURL(refreshToken string) string {
	return fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/component/refresh_token?"+
		"appid=%s&component_appid=%s&component_access_token=%s&refresh_token=%s&grant_type=refresh_token",
		url.QueryEscape(p.AppId),
		url.QueryEscape(p.ComponentAppId),
		url.QueryEscape(p.ComponentAccessToken),
		url.QueryEscape(refreshToken),
	)
}

func (p *Endpoint) SessionCodeUrl(code string) string {
	return fmt.Sprintf("https://api.weixin.qq.com/sns/component/jscode2session?"+
		"appid=%s&component_appid=%s&component_access_token=%s&js_code=%s&grant_type=authorization_code",
		url.QueryEscape(p.AppId),
		url.QueryEscape(p.ComponentAppId),
		url.QueryEscape(p.ComponentAccessToken),
		url.QueryEscape(code),
	)
}

// 要授权的帐号类型， 1则商户扫码后，手机端仅展示公众号、2表示仅展示小程序，3表示公众号和小程序都展示。如果为未制定，则默认小程序和公众号都展示。第三方平台开发者可以使用本字段来控制授权的帐号类型。
func (p *Endpoint) LoginUrl(preAuthCode string, redirectUri string, authType uint, bizAppId string) string {
	return fmt.Sprintf("https://mp.weixin.qq.com/cgi-bin/componentloginpage?"+
		"component_appid=%s&pre_auth_code=%s&redirect_uri=%s&auth_type=%d&biz_appid=%s",
		url.QueryEscape(p.ComponentAppId),
		url.QueryEscape(preAuthCode),
		url.QueryEscape(redirectUri),
		url.QueryEscape(authType),
		url.QueryEscape(bizAppId),
	)
}
