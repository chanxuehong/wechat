package wxa

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 小程序审核撤回
func UndoCodeAudit(clt *core.Client) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/undocodeaudit?access_token="
	var result struct {
		core.Error
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return nil
}
