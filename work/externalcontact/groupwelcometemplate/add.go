package groupwelcometemplate

import (
	"github.com/bububa/wechat/work/core"
)

// Add 添加群欢迎语素材
func Add(clt *core.Client, req *Template) (templateId string, err error) {
	const incompleteURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/add?access_token="

	var result struct {
		core.Error
		TemplateId string `json:"template_id"`
	}

	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	templateId = result.TemplateId
	return
}
