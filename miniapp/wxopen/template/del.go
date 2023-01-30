package template

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 删除帐号下的某个模板
func Del(clt *core.Client, templateId string) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxopen/template/del?access_token="
	var result struct {
		core.Error
	}
	req := map[string]interface{}{"template_id": templateId}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return nil
}
