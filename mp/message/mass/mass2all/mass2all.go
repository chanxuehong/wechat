// 群发消息给所有用户.
package mass2all

import (
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/message/mass"
)

// Send 发送消息, msg 是经过 encoding/json.Marshal 得到的结果符合微信消息格式的任何数据结构.
func Send(clt *core.Client, msg interface{}) (rslt *mass.Result, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token="

	var result struct {
		core.Error
		mass.Result
	}
	if err = clt.PostJSON(incompleteURL, msg, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	rslt = &result.Result
	return
}
