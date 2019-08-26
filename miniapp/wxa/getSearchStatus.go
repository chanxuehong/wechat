package wxa

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 查询小程序当前隐私设置（是否可被搜索）
func GetSearchStatus(clt *core.Client, auditId uint64) (status SearchStatus, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/getwxasearchstatus?access_token="
	var result struct {
		core.Error
		Status SearchStatus `json:"status"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return result.Status, nil
}
