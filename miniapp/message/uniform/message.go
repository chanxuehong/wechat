package uniform

import (
	"github.com/chanxuehong/wechat/mp/message/template"
)

type Message struct {
	ToUser     string      `json:"touser"`                       // 用户openid，可以是小程序的openid，也可以是mp_template_msg.appid对应的公众号的openid
	AppMessage *AppMessage `json:"weapp_template_msg,omitempty"` // 小程序模板消息相关的信息，可以参考小程序模板消息接口; 有此节点则优先发送小程序模板消息
	MpMessage  *MpMessage  `json:"mp_template_msg"`              // 公众号模板消息相关的信息，可以参考公众号模板消息接口；有此节点并且没有weapp_template_msg节点时，发送公众号模板消息
}
type AppMessage struct {
	TemplateId      string                       `json:"template_id"`                // 小程序模板ID
	Page            string                       `json:"page,omitempty"`             // 小程序页面路径
	FormId          string                       `json:"form_id"`                    // 小程序模板消息formid
	Data            map[string]template.DataItem `json:"data"`                       // 小程序模板数据
	EmphasisKeyword string                       `json:"emphasis_keyword,omitempty"` // 小程序模板放大关键词
}

type MpMessage struct {
	AppId       string                       `json:"app_id"`      // 公众号appid，要求与小程序有绑定且同主体
	TemplateId  string                       `json:"template_id"` // 公众号模板id
	Url         string                       `json:"url"`         // 公众号模板消息所要跳转的url。
	MiniProgram string                       `json:"miniprogram"` // 公众号模板消息所要跳转的小程序，小程序的必须与公众号具有绑定关系
	Data        map[string]template.DataItem `json:"data"`        // 公众号模板消息的数据
}
