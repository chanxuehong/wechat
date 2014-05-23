package wechat

import (
	"github.com/chanxuehong/wechat/message"
)

type Client struct {
	AccessToken string
}

func (c *Client) customSendUrl() string {
	return ""
}

func (c *Client) SendTextResponse(msg *message.TextResponse) {

}

func (c *Client) SendImageResponse(msg *message.ImageResponse) {

}

func (c *Client) SendVoiceResponse(msg *message.VoiceResponse) {

}

func (c *Client) SendVideoResponse(msg *message.VideoResponse) {

}

func (c *Client) SendMusicResponse(msg *message.MusicResponse) {

}

func (c *Client) SendNewsResponse(msg *message.NewsResponse) {

}
