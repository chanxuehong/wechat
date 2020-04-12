package wxa

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 解除绑定小程序的体验者
func UnbindTester(clt *core.Client, wechatId string, userstr string) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/unbind_tester?access_token="
	var result struct {
		core.Error
	}
	req := make(map[string]string, 1)
	if wechatId != "" {
		req["wechatid"] = wechatId
	} else if userstr != "" {
		req["userstr"] = userstr
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return nil
}
