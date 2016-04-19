package template

import (
	"encoding/json"

	"github.com/chanxuehong/wechat.v2/mp/core"
)

type TemplateMessage struct {
	ToUser     string          `json:"touser"`             // 必须, 接受者OpenID
	TemplateId string          `json:"template_id"`        // 必须, 模版ID
	URL        string          `json:"url,omitempty"`      // 可选, 用户点击后跳转的URL, 该URL必须处于开发者在公众平台网站中设置的域中
	TopColor   string          `json:"topcolor,omitempty"` // 可选, 整个消息的颜色, 可以不设置
	Data       json.RawMessage `json:"data"`               // 必须, JSON 格式的 []byte, 满足特定的模板需求
}

type TemplateMessage2 struct {
	ToUser     string      `json:"touser"`             // 必须, 接受者OpenID
	TemplateId string      `json:"template_id"`        // 必须, 模版ID
	URL        string      `json:"url,omitempty"`      // 可选, 用户点击后跳转的URL, 该URL必须处于开发者在公众平台网站中设置的域中
	TopColor   string      `json:"topcolor,omitempty"` // 可选, 整个消息的颜色, 可以不设置
	Data       interface{} `json:"data"`               // 必须, struct 或者 *struct, encoding/json.Marshal 后满足格式要求.
}

// 模版内某个 .DATA 的值
type DataItem struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

// 发送模板消息, msg 是经过 encoding/json.Marshal 得到的结果符合微信消息格式的任何数据结构, 一般为 *TemplateMessage 类型.
func Send(clt *core.Client, msg interface{}) (msgid int64, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token="

	var result struct {
		core.Error
		MsgId int64 `json:"msgid"`
	}
	if err = clt.PostJSON(incompleteURL, msg, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	msgid = result.MsgId
	return
}
