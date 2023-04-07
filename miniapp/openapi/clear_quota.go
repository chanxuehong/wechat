package openapi

import (
	"github.com/bububa/wechat/mp/core"
)

// ClearQuota 重置API调用次数
func ClearQuota(clt *core.Client, appID string) error {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/clear_quota?access_token="
	request := map[string]string{"appid": appID}
	var result core.Error
	if err := clt.PostJSON(incompleteURL, request, &result); err != nil {
		return err
	}
	if result.ErrCode != core.ErrCodeOK {
		return &result
	}
	return nil
}
