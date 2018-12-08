package device

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 删除设备
func Delete(clt *core.Client, bssid string) (err error) {
	request := struct {
		BSSID string `json:"bssid"`
	}{
		BSSID: bssid,
	}

	var result core.Error

	incompleteURL := "https://api.weixin.qq.com/bizwifi/device/delete?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}
