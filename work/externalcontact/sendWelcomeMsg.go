package externalcontact

import (
	"github.com/chanxuehong/wechat/work/core"
)

type SendWelcomeMsgRequest struct {
	WelcomeCode string              `json:"welcome_code,omitempty"`
	Text        *TextMessage        `json:"text,omitempty"`
	Image       *ImageMessage       `json:"image,omitempty"`
	Link        *LinkMessage        `json:"link,omitempty"`
	MiniProgram *MiniProgramMessage `json:"miniprogram,omitempty"`
}

// SendWelcomeMsg 发送新客户欢迎语
func SendWelcomeMsg(clt *core.Client, req *SendWelcomeMsgRequest) (err error) {
	const incompleteURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/send_welcome_msg?access_token="

	var result core.Error
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}
