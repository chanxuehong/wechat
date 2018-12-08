// 预览消息.
package preview

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// Send 发送消息, msg 是经过 encoding/json.Marshal 得到的结果符合微信消息格式的任何数据结构.
func Send(clt *core.Client, msg interface{}) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/message/mass/preview?access_token="

	var result core.Error
	if err = clt.PostJSON(incompleteURL, msg, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}
