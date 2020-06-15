package externalcontact

import (
	"github.com/chanxuehong/wechat/work/core"
)

type AddMsgTemplateRequest struct {
	ChatType        string              `json:"chat_type,omitempty"`
	ExternalUserIds []string            `json:"external_userid,omitempty"`
	Sender          string              `json:"sender,omitempty"`
	Text            *TextMessage        `json:"text,omitempty"`
	Image           *ImageMessage       `json:"image,omitempty"`
	Link            *LinkMessage        `json:"link,omitempty"`
	MiniProgram     *MiniProgramMessage `json:"miniprogram,omitempty"`
}

// AddMsgTemplate 添加企业群发消息任务
func AddMsgTemplate(clt *core.Client, req *AddMsgTemplateRequest) (msgId string, failList []string, err error) {
	const incompleteURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_msg_template?access_token="

	var result struct {
		core.Error
		FailList []string `json:"fail_list"`
		MsgId    string   `json:"msgid"`
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	msgId = result.MsgId
	failList = result.FailList
	return
}
