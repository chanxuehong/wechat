package wxa

import (
	"github.com/chanxuehong/wechat/component/core"
)

// 获取代码模版库中的所有小程序代码模版
func GetTemplateList(clt *core.Client) (list []Template, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/gettemplatelist?access_token="
	var result struct {
		core.Error
		List []Template `json:"template_list"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return result.List, nil
}
