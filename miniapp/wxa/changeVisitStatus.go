package wxa

import (
	"github.com/bububa/wechat/mp/core"
)

type VisitStatus = string

const (
	OPEN_VISITSTATUS  = "open"
	CLOSE_VISITSTATUS = "close"
)

// 修改小程序线上代码的可见状态（仅供第三方代小程序调用）
func ChangeVisitStatus(clt *core.Client, status VisitStatus) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/change_visitstatus?access_token="
	var result struct {
		core.Error
	}
	req := map[string]string{"action": status}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return nil
}
