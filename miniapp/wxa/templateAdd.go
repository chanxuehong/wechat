package wxa

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 组合模板并添加至帐号下的个人模板库
func TemplateAdd(clt *core.Client, id string, keywordIds []uint) (templateId string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxopen/template/add?access_token="
	var result struct {
		core.Error
		TemplateId string `json:"template_id"`
	}
	req := map[string]interface{}{"id": id, "keyword_id_list": keywordIds}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return result.TemplateId, nil
}
