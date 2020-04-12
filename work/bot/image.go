package bot

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"io"
	"io/ioutil"
	"net/http"
)

// 消息类型，此时固定为image
type Image struct {
	Base64 string `json:"base64"` // 图片内容的base64编码
	Md5    string `json:"md5"`    // 图片内容（base64编码前）的md5值
}

func NewImage(data []byte) *Message {
	md5Data := md5.Sum(data)
	return &Message{
		Type: IMAGE,
		Image: &Image{
			Base64: base64.URLEncoding.EncodeToString(data),
			Md5:    hex.EncodeToString(md5Data[:]),
		},
	}
}

func NewImageFromReader(r io.Reader) (*Message, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return NewImage(data), nil
}

func NewImageFromURL(link string) (*Message, error) {
	resp, err := http.Get(link)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return NewImageFromReader(resp.Body)
}
