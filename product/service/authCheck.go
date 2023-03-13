package service

import (
	"github.com/bububa/wechat/component/core"
)

// AuthCheck 登录验证
// code: 跳转码(小商店服务市场跳转到第三方url里面会带上code)
func AuthCheck(clt *core.Client, code string) (appId string, serviceId string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/product/service/check_auth?component_access_token="

	req := map[string]string{
		"code": code,
	}

	var result struct {
		core.Error
		Data struct {
			AppId     string `json:"appid"`      // 小程序ID
			ServiceId string `json:"service_id"` // 服务ID
		} `json:"data"`
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	appId = result.Data.AppId
	serviceId = result.Data.ServiceId
	return
}
