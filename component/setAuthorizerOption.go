package component

import (
	"github.com/bububa/wechat/component/core"
)

// 该API用于设置授权方的公众号或小程序的选项信息，如：地理位置上报，语音识别开关，多客服开关。注意，设置各项选项设置信息，需要有授权方的授权，详见权限集说明。
func SetAuthorizerOption(clt *core.Client, appId string, authorizerAppId string, optionName string, optionValue string) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/component/ api_set_authorizer_option?component_access_token="
	var result struct {
		core.Error
	}
	req := map[string]string{"component_appid": appId, "authorizer_appid": authorizerAppId, "option_name": optionName, "option_value": optionValue}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return nil
}
