package uniform

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 下发小程序和公众号统一的服务消息
func Send(clt *core.Client, msg *Message) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send?access_token="
	var result struct {
		core.Error
	}
	if err = clt.PostJSON(incompleteURL, msg, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return nil
}
