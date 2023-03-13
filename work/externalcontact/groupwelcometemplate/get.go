package groupwelcometemplate

import (
	"github.com/bububa/wechat/work/core"
)

// Get 获取群欢迎语素材
func Get(clt *core.Client, templateId string) (ret *Template, err error) {
	const incompleteURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/get?access_token="

	var result struct {
		core.Error
		Template
	}

	if err = clt.PostJSON(incompleteURL, map[string]string{"template_id": templateId}, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	ret = &Template{
		Text:        ret.Text,
		Image:       ret.Image,
		Link:        ret.Link,
		MiniProgram: ret.MiniProgram,
	}
	return
}
