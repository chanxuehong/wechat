package uniform

import (
	"github.com/bububa/wechat/mp/message/template"
)

type Message struct {
	// ToUser 用户openid，可以是小程序的openid，也可以是mp_template_msg.appid对应的公众号的openid
	ToUser string `json:"touser"`
	// WeappTemplateMsg 小程序模板消息相关的信息，可以参考小程序模板消息接口; 有此节点则优先发送小程序模板消息
	WeappTemplateMsg *WeappTemplateMsg `json:"weapp_template_msg,omitempty"`
	// MpTemplateMsg 公众号模板消息相关的信息，可以参考公众号模板消息接口；有此节点并且没有weapp_template_msg节点时，发送公众号模板消息
	MpTemplateMsg *MpTemplateMsg `json:"mp_template_msg"`
}

type WeappTemplateMsg struct {
	// TemplateId 小程序模板ID
	TemplateId string `json:"template_id"`
	// Page 小程序页面路径
	Page string `json:"page,omitempty"`
	// FormId 小程序模板消息formid
	FormId string `json:"form_id"`
	// Data 小程序模板数据
	Data map[string]template.DataItem `json:"data"`
	// EmphasisKeyword 小程序模板放大关键词
	EmphasisKeyword string `json:"emphasis_keyword,omitempty"`
}

type MpTemplateMsg struct {
	// AppId 公众号appid，要求与小程序有绑定且同主体
	AppId string `json:"appid"`
	// TemplateId 公众号模板id
	TemplateId string `json:"template_id"`
	// Url 公众号模板消息所要跳转的url。
	Url string `json:"url,omitempty"`
	// Miniprogram 公众号模板消息所要跳转的小程序，小程序的必须与公众号具有绑定关系
	Miniprogram *Miniprogram `json:"miniprogram,omitempty"`
	// Data 公众号模板消息的数据
	Data map[string]template.DataItem `json:"data"`
}

type Miniprogram struct {
	AppId    string `json:"appid"`
	PagePath string `json:"pagepath"`
}
