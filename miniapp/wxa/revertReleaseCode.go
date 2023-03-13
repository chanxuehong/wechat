package wxa

import (
	"github.com/bububa/wechat/mp/core"
)

// 小程序版本回退（仅供第三方代小程序调用）
func RevertReleaseCode(clt *core.Client) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/revertcoderelease?access_token="
	var result struct {
		core.Error
	}
	if err = clt.PostJSON(incompleteURL, nil, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return nil
}
