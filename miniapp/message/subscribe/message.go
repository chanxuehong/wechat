package subscribe

import (
	"github.com/bububa/wechat/mp/message/template"
)

type Message struct {
	// ToUser 必须, 接受者OpenID
	ToUser string `json:"touser"`
	// TemplateId 必须, 模版ID
	TemplateId string `json:"template_id"`
	// Page 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转
	Page string `json:"page,omitempty"`
	// MiniProgramState 跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
	MiniprogramState string `json:"miniprogram_state,omitempty"`
	// Date 模板内容，不填则下发空模板。具体格式请参考示例。
	Data map[string]template.DataItem `json:"data"`
	// Lang 进入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
	Lang string `json:"lang,omitempty"`
}
