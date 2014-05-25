package wechat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chanxuehong/wechat/message"
	"io/ioutil"
	"net/http"
)

// 通用的功能
func (c *Client) customSendResponse(jsonData []byte) error {
	token, err := c.Token()
	if err != nil {
		return err
	}

	url := fmt.Sprintf(customSendMessageUrlFormat, token)
	resp, err := http.Post(url, "application/json; charset=utf-8", bytes.NewReader(jsonData))
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

func (c *Client) CustomSendTextResponse(msg *message.TextResponse) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	msgBody, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return c.customSendResponse(msgBody)
}

func (c *Client) CustomSendImageResponse(msg *message.ImageResponse) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	msgBody, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return c.customSendResponse(msgBody)
}

func (c *Client) CustomSendVoiceResponse(msg *message.VoiceResponse) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	msgBody, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return c.customSendResponse(msgBody)
}

func (c *Client) CustomSendVideoResponse(msg *message.VideoResponse) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	msgBody, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return c.customSendResponse(msgBody)
}

func (c *Client) CustomSendMusicResponse(msg *message.MusicResponse) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	msgBody, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return c.customSendResponse(msgBody)
}

func (c *Client) CustomSendNewsResponse(msg *message.NewsResponse) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	msgBody, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return c.customSendResponse(msgBody)
}
