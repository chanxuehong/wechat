package wxa

import (
	"github.com/chanxuehong/wechat/component/core"
)

// 绑定微信用户为小程序体验者
func BindTester(clt *core.Client, wechatId string) (userStr string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/bind_tester?access_token="
	var result struct {
		core.Error
		UserStr string `json:"userstr"`
	}
	req := map[string]string{"wechatid": wechatId}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return result.UserStr, nil
}
