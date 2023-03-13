package component

import (
	"github.com/bububa/wechat/component/core"
)

// 该API用于获取授权方的基本信息，包括头像、昵称、帐号类型、认证类型、微信号、原始ID和二维码图片URL。需要特别记录授权方的帐号类型，在消息及事件推送时，对于不具备客服接口的公众号，需要在5秒内立即响应；而若有客服接口，则可以选择暂时不响应，而选择后续通过客服接口来发送消息触达粉丝。
func GetAuthorizerInfo(clt *core.Client, appId string, authorizerAppId string) (authorizerInfo *AuthorizerInfo, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info?component_access_token="
	var result struct {
		core.Error
		AuthorizerInfo *AuthorizerInfo `json:"authorizer_info"`
	}
	req := map[string]string{"component_appid": appId, "authorizer_appid": authorizerAppId}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return result.AuthorizerInfo, nil
}
