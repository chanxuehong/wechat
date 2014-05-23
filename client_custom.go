package wechat

import (
	"github.com/chanxuehong/wechat/message"
)

func (c *Client) CustomSendTextResponse(msg *message.TextResponse) error {
	return nil
}

func (c *Client) CustomSendImageResponse(msg *message.ImageResponse) error {
	return nil
}

func (c *Client) CustomSendVoiceResponse(msg *message.VoiceResponse) error {
	return nil
}

func (c *Client) CustomSendVideoResponse(msg *message.VideoResponse) error {
	return nil
}

func (c *Client) CustomSendMusicResponse(msg *message.MusicResponse) error {
	return nil
}

func (c *Client) CustomSendNewsResponse(msg *message.NewsResponse) error {
	return nil
}
