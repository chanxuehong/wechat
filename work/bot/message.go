package bot

import (
	"net/http"
)

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

func (this *Message) Send(webhook string) error {
	var buf bytes.Buffer
	err := json.NewEncoder(buf).Encode(this)
	if err != nil {
		return
	}
	resp, err := http.DefaultClient.Post(webhook, "application/json", buf)
	if err != nil {
		return err
	}
	return nil
}
