package openapi

import (
	"github.com/bububa/wechat/mp/core"
)

// ClearQuotaV2 使用AppSecret重置API调用次数
func ClearQuotaV2(clt *core.Client, appID string, secret string) error {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/clear_quota/v2?access_token="
	request := map[string]string{"appid": appID, "appsecret": secret}
	var result core.Error
	if err := clt.PostJSON(incompleteURL, request, &result); err != nil {
		return err
	}
	if result.ErrCode != core.ErrCodeOK {
		return &result
	}
	return nil
}
