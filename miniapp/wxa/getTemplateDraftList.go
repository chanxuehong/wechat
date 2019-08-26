package wxa

import (
	"github.com/chanxuehong/wechat/component/core"
)

// 获取草稿箱内的所有临时代码草稿
func GetTemplateDraftList(clt *core.Client) (list []TemplateDraft, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/gettemplatedraftlist?access_token="
	var result struct {
		core.Error
		List []TemplateDraft `json:"drafttemplate_list"`
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
