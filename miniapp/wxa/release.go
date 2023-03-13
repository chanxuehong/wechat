package wxa

import (
	"github.com/bububa/wechat/mp/core"
)

// 发布已通过审核的小程序（仅供第三方代小程序调用）
func Release(clt *core.Client) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/release?access_token="
	var result struct {
		core.Error
	}
	req := map[string]interface{}{}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return nil
}
