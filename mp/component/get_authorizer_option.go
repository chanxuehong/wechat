// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://gopkg.in/chanxuehong/wechat.v1 for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/v1/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package component

import (
	"gopkg.in/chanxuehong/wechat.v1/mp"
)

// 获取授权方的选项设置信息.
func (clt *Client) GetAuthorizerOption(authorizerAppId, optionName string) (optionValue string, err error) {
	request := struct {
		ComponentAppId  string `json:"component_appid"`
		AuthorizerAppId string `json:"authorizer_appid"`
		OptionName      string `json:"option_name"`
	}{
		ComponentAppId:  clt.AppId,
		AuthorizerAppId: authorizerAppId,
		OptionName:      optionName,
	}

	var result struct {
		mp.Error
		OptionValue int `json:"option_value"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_option?component_access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	optionValue = string(result.OptionValue)
	return
}
