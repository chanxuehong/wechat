package template

import (
	"encoding/json"
)

type TemplateMessage struct {
	ToUser     string `json:"touser"`             // 必须, 接受者OpenID
	TemplateId string `json:"template_id"`        // 必须, 模版ID
	URL        string `json:"url,omitempty"`      // 可选, 用户点击后跳转的URL, 该URL必须处于开发者在公众平台网站中设置的域中
	TopColor   string `json:"topcolor,omitempty"` // 可选, 整个消息的颜色, 可以不设置

	RawJSONData json.RawMessage `json:"data"` // 必须, JSON 格式的 []byte, 满足特定的模板需求
}
