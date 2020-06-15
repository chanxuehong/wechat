package externalcontact

import (
	"github.com/chanxuehong/wechat/work/core"
)

type GroupMsgResult struct {
	ChatId         string `json:"chat_id,omitempty"`
	ExternalUserId string `json:"external_userid,omitempty"`
	UserId         string `json:"userid,omitempty"`
	Status         int    `json:"status,omitempty"`
	SendTime       int64  `json:"send_time,omitempty"`
}

// GetGroupMsgResult 获取企业群发消息发送结果
func GetGroupMsgResult(clt *core.Client, msgId string) (ret []GroupMsgResult, err error) {
	const incompleteURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_group_msg_result??access_token="

	var result struct {
		core.Error
		DetailList []GroupMsgResult `json:"detail_list"`
	}
	if err = clt.PostJSON(incompleteURL, map[string]string{"msgid": msgId}, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	ret = result.DetailList
	return
}
