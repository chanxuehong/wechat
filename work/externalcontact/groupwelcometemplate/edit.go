package groupwelcometemplate

import (
	"github.com/chanxuehong/wechat/work/core"
)

type EditRequest struct {
	Template
	TemplateId string `json:"template_id"`
}

// Edit 编辑群欢迎语素材
func Edit(clt *core.Client, req *EditRequest) (err error) {
	const incompleteURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/edit?access_token="

	var result core.Error

	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}
