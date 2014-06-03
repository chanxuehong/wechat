package wechat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chanxuehong/wechat/message"
	"net/http"
)

// 发送客服消息功能都一样, 之所以不暴露这个接口是因为怕接收到不合法的参数.
func (c *Client) msgCustomSend(msg interface{}) error {
	token, err := c.Token()
	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	_url := clientMessageCustomSendURL(token)
	resp, err := c.httpClient.Post(_url, postJSONContentType, bytes.NewReader(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("msgCustomSend: %s", resp.Status)
	}

	var result Error
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}
	if result.ErrCode != 0 {
		return &result
	}
	return nil
}

// 发送客服消息, 文本.
func (c *Client) MsgCustomSendText(msg *message.TextResponse) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return c.msgCustomSend(msg)
}

// 发送客服消息, 图片.
func (c *Client) MsgCustomSendImage(msg *message.ImageResponse) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return c.msgCustomSend(msg)
}

// 发送客服消息, 语音.
func (c *Client) MsgCustomSendVoice(msg *message.VoiceResponse) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return c.msgCustomSend(msg)
}

// 发送客服消息, 视频.
func (c *Client) MsgCustomSendVideo(msg *message.VideoResponse) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return c.msgCustomSend(msg)
}

// 发送客服消息, 音乐.
func (c *Client) MsgCustomSendMusic(msg *message.MusicResponse) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return c.msgCustomSend(msg)
}

// 发送客服消息, 图文.
func (c *Client) MsgCustomSendNews(msg *message.NewsResponse) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return c.msgCustomSend(msg)
}
