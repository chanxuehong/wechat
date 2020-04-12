package component

import (
	"github.com/chanxuehong/wechat/component/core"
)

// 该API用于使用授权码换取授权公众号或小程序的授权信息，并换取authorizer_access_token和authorizer_refresh_token。 授权码的获取，需要在用户在第三方平台授权页中完成授权流程后，在回调URI中通过URL参数提供给第三方平台方。请注意，由于现在公众号或小程序可以自定义选择部分权限授权给第三方平台，因此第三方平台开发者需要通过该接口来获取公众号或小程序具体授权了哪些权限，而不是简单地认为自己声明的权限就是公众号或小程序授权的权限。
func QueryAuth(clt *core.Client, appId string, authorizationCode string) (authorizationInfo *AuthorizationInfo, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token="
	var result struct {
		core.Error
		AuthorizationInfo *AuthorizationInfo `json:"authorization_info"`
	}
	req := map[string]string{"component_appid": appId, "authorization_code": authorizationCode}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return result.AuthorizationInfo, nil
}
