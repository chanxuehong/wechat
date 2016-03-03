// 群发给用户列表的消息数据结构.
package mass2users

import (
	"errors"
	"fmt"

	"github.com/chanxuehong/wechat/mp/core"
)

const (
	MsgTypeText   core.MsgType = "text"
	MsgTypeImage  core.MsgType = "image"
	MsgTypeVoice  core.MsgType = "voice"
	MsgTypeVideo  core.MsgType = "mpvideo"
	MsgTypeNews   core.MsgType = "mpnews"
	MsgTypeWxCard core.MsgType = "wxcard"
)

const ToUserCountLimit = 10000

type MsgHeader struct {
	ToUser  []string     `json:"touser,omitempty"` // 长度不能超过 ToUserCountLimit
	MsgType core.MsgType `json:"msgtype"`
}

func (header *MsgHeader) CheckValid() (err error) {
	n := len(header.ToUser)
	if n <= 0 {
		return errors.New("用户列表是空的")
	}
	if n > ToUserCountLimit {
		return fmt.Errorf("用户列表的长度不能超过 %d, 现在为 %d", ToUserCountLimit, n)
	}
	return
}

type Text struct {
	MsgHeader
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

func NewText(toUser []string, content string) *Text {
	var msg Text
	msg.MsgType = MsgTypeText
	msg.ToUser = toUser
	msg.Text.Content = content
	return &msg
}

type Image struct {
	MsgHeader
	Image struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
}

func NewImage(toUser []string, mediaId string) *Image {
	var msg Image
	msg.MsgType = MsgTypeImage
	msg.ToUser = toUser
	msg.Image.MediaId = mediaId
	return &msg
}

type Voice struct {
	MsgHeader
	Voice struct {
		MediaId string `json:"media_id"` // mediaId 通过上传多媒体文件得到
	} `json:"voice"`
}

func NewVoice(toUser []string, mediaId string) *Voice {
	var msg Voice
	msg.MsgType = MsgTypeVoice
	msg.ToUser = toUser
	msg.Voice.MediaId = mediaId
	return &msg
}

type Video struct {
	MsgHeader
	Video struct {
		MediaId string `json:"media_id"`
	} `json:"mpvideo"`
}

// 新建视频消息.
//  NOTE: 对于临时素材, mediaId 应该通过 media.Client.CreateVideo 得到
func NewVideo(toUser []string, mediaId string) *Video {
	var msg Video
	msg.MsgType = MsgTypeVideo
	msg.ToUser = toUser
	msg.Video.MediaId = mediaId
	return &msg
}

// 图文消息
type News struct {
	MsgHeader
	News struct {
		MediaId string `json:"media_id"`
	} `json:"mpnews"`
}

// 新建图文消息.
//  NOTE: 对于临时素材, mediaId 应该通过 media.Client.CreateNews 得到
func NewNews(toUser []string, mediaId string) *News {
	var msg News
	msg.MsgType = MsgTypeNews
	msg.ToUser = toUser
	msg.News.MediaId = mediaId
	return &msg
}

// 卡券消息
type WxCard struct {
	MsgHeader
	WxCard struct {
		CardId string `json:"card_id"`
	} `json:"wxcard"`
}

// 新建卡券, 特别注意: 目前该接口仅支持填入非自定义code的卡券和预存模式的自定义code卡券.
func NewWxCard(toUser []string, cardId string) *WxCard {
	var msg WxCard
	msg.MsgType = MsgTypeWxCard
	msg.ToUser = toUser
	msg.WxCard.CardId = cardId
	return &msg
}
