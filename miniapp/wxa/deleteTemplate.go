package wxa

import (
	"github.com/chanxuehong/wechat/component/core"
)

// 删除指定小程序代码模版
func DeleteTemplate(clt *core.Client, templateId uint64) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/deletetemplate?access_token="
	var result struct {
		core.Error
	}
	req := map[string]uint64{"template_id": templateId}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return nil
}
