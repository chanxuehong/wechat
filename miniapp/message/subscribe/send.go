package subscribe

import (
	"github.com/bububa/wechat/mp/core"
)

// 发送模板消息, msg 是经过 encoding/json.Marshal 得到的结果符合微信消息格式的任何数据结构, 一般为 *TemplateMessage 类型.
func Send(clt *core.Client, msg *Message) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token="
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
