package component

import (
	"github.com/bububa/wechat/util"
)

// AuthCodeURL 生成网页授权地址.
//
//	appId:          公众号的唯一标识
//	componentAppId: 服务方的appid，在申请创建公众号服务成功后，可在公众号服务详情页找到
//	redirectURI:    授权后重定向的回调链接地址
//	scope:          应用授权作用域
//	state:          重定向后会带上 state 参数, 开发者可以填写 a-zA-Z0-9 的参数值, 最多128字节
func AuthCodeURL(appId, componentAppId, redirectURI, scope, state string) string {
	values := util.GetUrlValues()
	values.Set("appid", appId)
	values.Set("component_appid", componentAppId)
	values.Set("redirect_uri", redirectURI)
	values.Set("state", state)
	values.Set("response_type", "code")
	query := values.Encode()
	util.PutUrlValues(values)
	return util.StringsJoin("https://open.weixin.qq.com/connect/oauth2/authorize?", query, "#wechat_redirect")
}
