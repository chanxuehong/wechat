package groupwelcometemplate

import (
	"github.com/chanxuehong/wechat/work/externalcontact"
)

type Template struct {
	Text        *externalcontact.TextMessage        `json:"text,omitempty"`
	Image       *externalcontact.ImageMessage       `json:"image,omitempty"`
	Link        *externalcontact.LinkMessage        `json:"link,omitempty"`
	MiniProgram *externalcontact.MiniProgramMessage `json:"miniprogram,omitempty"`
}
