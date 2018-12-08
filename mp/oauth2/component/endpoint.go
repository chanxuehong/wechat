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
