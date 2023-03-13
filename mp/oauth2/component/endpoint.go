package component

import (
	"strconv"

	"github.com/bububa/wechat/oauth2"
	"github.com/bububa/wechat/util"
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
	values := util.GetUrlValues()
	values.Set("appid", p.AppId)
	values.Set("component_appid", p.ComponentAppId)
	values.Set("component_access_token", p.ComponentAccessToken)
	values.Set("code", code)
	values.Set("grand_type", "authorization_code")
	query := values.Encode()
	util.PutUrlValues(values)
	return util.StringsJoin("https://api.weixin.qq.com/sns/oauth2/component/access_token?", query)
}

func (p *Endpoint) RefreshTokenURL(refreshToken string) string {
	values := util.GetUrlValues()
	values.Set("appid", p.AppId)
	values.Set("component_appid", p.ComponentAppId)
	values.Set("component_access_token", p.ComponentAccessToken)
	values.Set("refresh_token", refreshToken)
	values.Set("grand_type", "refresh_token")
	query := values.Encode()
	util.PutUrlValues(values)
	return util.StringsJoin("https://api.weixin.qq.com/sns/oauth2/component/refresh_token?", query)
}

func (p *Endpoint) SessionCodeUrl(code string) string {
	values := util.GetUrlValues()
	values.Set("appid", p.AppId)
	values.Set("component_appid", p.ComponentAppId)
	values.Set("component_access_token", p.ComponentAccessToken)
	values.Set("js_code", code)
	values.Set("grand_type", "authorization_code")
	query := values.Encode()
	util.PutUrlValues(values)
	return util.StringsJoin("https://api.weixin.qq.com/sns/component/jscode2session?", query)
}

// 要授权的帐号类型， 1则商户扫码后，手机端仅展示公众号、2表示仅展示小程序，3表示公众号和小程序都展示。如果为未制定，则默认小程序和公众号都展示。第三方平台开发者可以使用本字段来控制授权的帐号类型。
func (p *Endpoint) LoginUrl(preAuthCode string, redirectUri string, authType uint, bizAppId string) string {
	values := util.GetUrlValues()
	values.Set("component_appid", p.ComponentAppId)
	values.Set("pre_auth_code", preAuthCode)
	values.Set("redirect_uri", redirectUri)
	values.Set("auth_type", strconv.Itoa(int(authType)))
	values.Set("biz_appid", bizAppId)
	query := values.Encode()
	util.PutUrlValues(values)
	return util.StringsJoin("https://mp.weixin.qq.com/cgi-bin/componentloginpage?", query)
}
