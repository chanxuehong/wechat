package mass2group

import (
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

type MsgHeader struct {
	Filter struct {
		GroupId int64 `json:"group_id"`
	} `json:"filter"`
	MsgType core.MsgType `json:"msgtype"`
}

type Text struct {
	MsgHeader
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

func NewText(groupId int64, content string) *Text {
	var msg Text
	msg.MsgType = MsgTypeText
	msg.Filter.GroupId = groupId
	msg.Text.Content = content
	return &msg
}

type Image struct {
	MsgHeader
	Image struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
}

func NewImage(groupId int64, mediaId string) *Image {
	var msg Image
	msg.MsgType = MsgTypeImage
	msg.Filter.GroupId = groupId
	msg.Image.MediaId = mediaId
	return &msg
}

type Voice struct {
	MsgHeader
	Voice struct {
		MediaId string `json:"media_id"`
	} `json:"voice"`
}

func NewVoice(groupId int64, mediaId string) *Voice {
	var msg Voice
	msg.MsgType = MsgTypeVoice
	msg.Filter.GroupId = groupId
	msg.Voice.MediaId = mediaId
	return &msg
}

type Video struct {
	MsgHeader
	Video struct {
		MediaId string `json:"media_id"`
	} `json:"mpvideo"`
}

// 新建视频消息
//  NOTE: 对于临时素材, mediaId 应该通过 media.UploadVideo2 得到
func NewVideo(groupId int64, mediaId string) *Video {
	var msg Video
	msg.MsgType = MsgTypeVideo
	msg.Filter.GroupId = groupId
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

// 新建图文消息
//  NOTE: 对于临时素材, mediaId 应该通过 media.UploadNews 得到
func NewNews(groupId int64, mediaId string) *News {
	var msg News
	msg.MsgType = MsgTypeNews
	msg.Filter.GroupId = groupId
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
//  cardExt 可以为空
func NewWxCard(groupId int64, cardId, cardExt string) *WxCard {
	var msg WxCard
	msg.MsgType = MsgTypeWxCard
	msg.Filter.GroupId = groupId
	msg.WxCard.CardId = cardId
	msg.WxCard.CardExt = cardExt
	return &msg
}
