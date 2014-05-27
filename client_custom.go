package wechat

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/chanxuehong/wechat/message"
	"io/ioutil"
	"net/http"
)

// 发送客服消息功能都一样, 之所以不暴露这个接口是因为怕接收到不合法的参数.
func (c *Client) customSendResponse(msg interface{}) error {
	token, err := c.Token()
	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	url := customSendMessageUrlPrefix + token
	resp, err := http.Post(url, postJSONContentType, bytes.NewReader(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result Error
	if err = json.Unmarshal(body, &result); err != nil {
		return err
	}
	if result.ErrCode != 0 {
		return &result
	}
	return nil
}

// 发送客服消息, 文本.
func (c *Client) CustomSendTextResponse(msg *message.TextResponse) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return c.customSendResponse(msg)
}

// 发送客服消息, 图片.
func (c *Client) CustomSendImageResponse(msg *message.ImageResponse) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return c.customSendResponse(msg)
}

// 发送客服消息, 语音.
func (c *Client) CustomSendVoiceResponse(msg *message.VoiceResponse) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return c.customSendResponse(msg)
}

// 发送客服消息, 视频.
func (c *Client) CustomSendVideoResponse(msg *message.VideoResponse) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return c.customSendResponse(msg)
}

// 发送客服消息, 音乐.
func (c *Client) CustomSendMusicResponse(msg *message.MusicResponse) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return c.customSendResponse(msg)
}

// 发送客服消息, 图文.
func (c *Client) CustomSendNewsResponse(msg *message.NewsResponse) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return c.customSendResponse(msg)
}
