package template

import (
	"github.com/chanxuehong/wechat/mp/message/template"
)

type Message struct {
	ToUser          string                       `json:"touser"`                     // 必须, 接受者OpenID
	TemplateId      string                       `json:"template_id"`                // 必须, 模版ID
	Page            string                       `json:"page,omitempty"`             // 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
	FormId          string                       `json:"form_id"`                    // 表单提交场景下，为 submit 事件带上的 formId；支付场景下，为本次支付的 prepay_id
	Data            map[string]template.DataItem `json:"data"`                       // 模板内容，不填则下发空模板。具体格式请参考示例。
	EmphasisKeyword string                       `json:"emphasis_keyword,omitempty"` // 模板需要放大的关键词，不填则默认无放大
}
