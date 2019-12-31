package bot

type MessageType = string

const (
	TEXT     MessageType = "text"
	MARKDOWN MessageType = "markdown"
	IMAGE    MessageType = "image"
	NEWS     MessageType = "news"
)

type Message struct {
	Type     MessageType `json:"msgtype"`
	Text     *Text       `json:"text,omitempty"`
	Markdown *Markdown   `json:"markdown,omitempty"`
	Image    *Image      `json:"image,omitempty"`
	News     *News       `json:"news,omitempty"`
}

// 消息类型，此时固定为image
type Image struct {
	Base64 string `json:"base64"` // 图片内容的base64编码
	Md5    string `json:"md5"`    // 图片内容（base64编码前）的md5值
}
