// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"fmt"
	"github.com/chanxuehong/wechat/message/custom"
)

// 发送客服消息, 文本.
func (c *Client) MsgCustomSendText(toUser, content string) error {
	text := custom.Text{
		CommonHead: custom.CommonHead{
			ToUser:  toUser,
			MsgType: custom.MSG_TYPE_TEXT,
		},
	}
	text.Text.Content = content

	return c.msgCustomSend(&text)
}

// 发送客服消息, 图片.
func (c *Client) MsgCustomSendImage(toUser, mediaId string) error {
	image := custom.Image{
		CommonHead: custom.CommonHead{
			ToUser:  toUser,
			MsgType: custom.MSG_TYPE_IMAGE,
		},
	}
	image.Image.MediaId = mediaId

	return c.msgCustomSend(&image)
}

// 发送客服消息, 语音.
func (c *Client) MsgCustomSendVoice(toUser, mediaId string) error {
	voice := custom.Voice{
		CommonHead: custom.CommonHead{
			ToUser:  toUser,
			MsgType: custom.MSG_TYPE_VOICE,
		},
	}
	voice.Voice.MediaId = mediaId

	return c.msgCustomSend(&voice)
}

// 发送客服消息, 视频.
//  title, description 可以为 ""
func (c *Client) MsgCustomSendVideo(toUser, mediaId, title, description string) error {
	video := custom.Video{
		CommonHead: custom.CommonHead{
			ToUser:  toUser,
			MsgType: custom.MSG_TYPE_VIDEO,
		},
	}
	video.Video.Title = title
	video.Video.Description = description
	video.Video.MediaId = mediaId

	return c.msgCustomSend(&video)
}

// 发送客服消息, 音乐.
//  title, description 可以为 ""
func (c *Client) MsgCustomSendMusic(toUser, thumbMediaId, musicURL, HQMusicURL,
	title, description string) error {

	music := custom.Music{
		CommonHead: custom.CommonHead{
			ToUser:  toUser,
			MsgType: custom.MSG_TYPE_MUSIC,
		},
	}
	music.Music.Title = title
	music.Music.Description = description
	music.Music.ThumbMediaId = thumbMediaId
	music.Music.MusicURL = musicURL
	music.Music.HQMusicURL = HQMusicURL

	return c.msgCustomSend(&music)
}

// 发送客服消息, 图文.
//  len(articles) 不能大于 custom.NewsArticleCountLimit
func (c *Client) MsgCustomSendNews(toUser string, articles []custom.NewsArticle) (err error) {
	if len(articles) > custom.NewsArticleCountLimit {
		err = fmt.Errorf("图文消息的文章个数不能超过 %d, 现在为 %d", custom.NewsArticleCountLimit, len(articles))
		return
	}

	news := custom.News{
		CommonHead: custom.CommonHead{
			ToUser:  toUser,
			MsgType: custom.MSG_TYPE_NEWS,
		},
	}
	news.News.Articles = articles

	return c.msgCustomSend(&news)
}

// 发送客服消息功能都一样, 之所以不暴露这个接口是因为怕接收到不合法的参数.
func (c *Client) msgCustomSend(msg interface{}) (err error) {
	var result Error

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := messageCustomSendURL(token)
	if err = c.postJSON(_url, msg, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = result
		return
	}
}
