package mass2users

import (
	"github.com/bububa/wechat/mp/core"
)

const (
	MsgTypeText   core.MsgType = "text"
	MsgTypeImage  core.MsgType = "image"
	MsgTypeVoice  core.MsgType = "voice"
	MsgTypeVideo  core.MsgType = "mpvideo"
	MsgTypeNews   core.MsgType = "mpnews"
	MsgTypeWxCard core.MsgType = "wxcard"
)

type MsgHeader struct {
	ToUser            []string     `json:"touser,omitempty"`
	MsgType           core.MsgType `json:"msgtype"`
	SendIgnoreReprint int          `json:"send_ignore_reprint"` //  参数设置为1时，文章被判定为转载时，且原创文允许转载时，将继续进行群发操作, 参数设置为0时，文章被判定为转载时，将停止群发操作
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
		MediaId      string `json:"media_id"`
		Title        string `json:"title,omitempty"`
		Description  string `json:"description,omitempty"`
		ThumbMediaId string `json:"thumb_media_id,omitempty"`
	} `json:"mpvideo"`
}

// 新建视频消息.
//
//	NOTE: 对于临时素材, mediaId 应该通过 media.UploadVideo2 得到
func NewVideo(toUser []string, mediaId string, title string, description string, thumbMediaId string) *Video {
	var msg Video
	msg.MsgType = MsgTypeVideo
	msg.ToUser = toUser
	msg.Video.MediaId = mediaId
	msg.Video.Title = title
	msg.Video.Description = description
	msg.Video.ThumbMediaId = thumbMediaId
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
//
//	NOTE: 对于临时素材, mediaId 应该通过 media.UploadNews 得到
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
		CardId  string `json:"card_id"`
		CardExt string `json:"card_ext,omitempty"`
	} `json:"wxcard"`
}

// 新建卡券, 特别注意: 目前该接口仅支持填入非自定义code的卡券和预存模式的自定义code卡券.
//
//	cardExt 可以为空
func NewWxCard(toUser []string, cardId, cardExt string) *WxCard {
	var msg WxCard
	msg.MsgType = MsgTypeWxCard
	msg.ToUser = toUser
	msg.WxCard.CardId = cardId
	msg.WxCard.CardExt = cardExt
	return &msg
}
