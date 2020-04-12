package bot

// 消息类型，此时固定为markdown
type Markdown struct {
	Content string `json:"content"` // markdown内容，最长不超过4096个字节，必须是utf8编码
}

func NewMarkdown(content string) *Message {
	return &Message{
		Type: MARKDOWN,
		Markdown: &Markdown{
			Content: content,
		},
	}
}
