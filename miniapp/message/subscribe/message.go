package subscribe

import (
	"github.com/chanxuehong/wechat/mp/message/template"
)

type Message struct {
	ToUser     string                       `json:"touser"`         // 必须, 接受者OpenID
	TemplateId string                       `json:"template_id"`    // 必须, 模版ID
	Page       string                       `json:"page,omitempty"` // 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。formId；支付场景下，为本次支付的 prepay_id
	Data       map[string]template.DataItem `json:"data"`           // 模板内容，不填则下发空模板。具体格式请参考示例。
}
