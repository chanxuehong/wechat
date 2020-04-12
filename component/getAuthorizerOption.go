package component

import (
	"github.com/chanxuehong/wechat/component/core"
)

// 该API用于获取授权方的公众号或小程序的选项设置信息，如：地理位置上报，语音识别开关，多客服开关。注意，获取各项选项设置信息，需要有授权方的授权，详见权限集说明。
func GetAuthorizerOption(clt *core.Client, appId string, authorizerAppId string, optionName string) (optionValue string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_option?component_access_token="
	var result struct {
		core.Error
		AuthorizerAppId string `json:"authorizer_appid"`
		OptionName      string `json:"option_name"`
		OptionValue     string `json:"option_value"`
	}
	req := map[string]string{"component_appid": appId, "authorizer_appid": authorizerAppId, "option_name": optionName}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return result.OptionValue, nil
}
