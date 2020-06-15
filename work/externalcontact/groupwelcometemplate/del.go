package groupwelcometemplate

import (
	"github.com/chanxuehong/wechat/work/core"
)

// Del 删除群欢迎语素材
func Del(clt *core.Client, templateId string) (err error) {
	const incompleteURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/del?access_token="

	var result core.Error

	if err = clt.PostJSON(incompleteURL, map[string]string{"template_id": templateId}, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}
