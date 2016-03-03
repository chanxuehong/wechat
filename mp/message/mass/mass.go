package mass

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 群发结果
type Result struct {
	MsgId int64 `json:"msg_id"` // 消息发送任务的ID

	// 消息的数据ID，该字段只有在群发图文消息时，才会出现。可以用于在图文分析数据接口中，获取到对应的图文消息的数据，
	// 是图文分析数据接口中的msgid字段中的前半部分，详见图文分析数据接口中的msgid字段的介绍。
	MsgDataId int64 `json:"msg_data_id"`
}

// Delete 删除群发.
func Delete(clt *core.Client, msgid int64) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/message/mass/delete?access_token="

	var request = struct {
		MsgId int64 `json:"msg_id"`
	}{
		MsgId: msgid,
	}
	var result core.Error
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}

type Status struct {
	MsgId  int64  `json:"msg_id"`
	Status string `json:"msg_status"` // 消息发送后的状态, SEND_SUCCESS表示发送成功
}

// GetStatus 查询群发消息发送状态.
func GetStatus(clt *core.Client, msgid int64) (status *Status, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/message/mass/get?access_token="

	var request = struct {
		MsgId int64 `json:"msg_id"`
	}{
		MsgId: msgid,
	}
	var result struct {
		core.Error
		Status
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	status = &result.Status
	return
}
