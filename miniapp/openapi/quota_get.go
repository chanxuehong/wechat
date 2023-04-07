package openapi

import (
	"github.com/bububa/wechat/mp/core"
)

type Quota struct {
	// DailyLimit 当天该账号可调用该接口的次数
	DailyLimit int64 `json:"daily_limit,omitempty"`
	// Used 当天已经调用的次数
	Used int64 `json:"used,omitempty"`
	// Remain 当天剩余调用次数
	Remain int64 `json:"remain,omitempty"`
}

// QuotaGet 查询API调用额度
func QuotaGet(clt *core.Client, cgiPath string) (quota *Quota, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/openapi/quota/get?access_token="
	request := map[string]string{"cgi_path": cgiPath}
	var result struct {
		core.Error
		Quota *Quota `json:"quota,omitempty"`
	}
	if err = clt.PostJSON(incompleteURL, request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return nil, err
	}
	quota = result.Quota
	return
}
