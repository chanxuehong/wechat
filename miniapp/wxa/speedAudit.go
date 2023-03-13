package wxa

import (
	"github.com/bububa/wechat/mp/core"
)

// 加急审核申请
func SpeedAudit(clt *core.Client, auditId uint64) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/speedupaudit?access_token="
	var result struct {
		core.Error
	}
	req := map[string]uint64{"auditid": auditId}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return nil
}
