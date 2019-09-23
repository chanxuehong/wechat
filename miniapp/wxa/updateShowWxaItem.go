package wxa

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 设置小程序隐私设置（是否可被搜索）
func UpdateShowWxaItem(clt *core.Client, appId string, show uint) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/updateshowwxaitem?access_token="
	var result struct {
		core.Error
	}
	req := map[string]interface{}{"wxa_subscribe_biz_flag": show, "appid": appId}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return nil
}
