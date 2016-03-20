<<<<<<< HEAD
// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

// 群发给分组的消息数据结构.
package mass2group

const (
	MsgTypeText   = "text"
	MsgTypeImage  = "image"
	MsgTypeVoice  = "voice"
	MsgTypeVideo  = "mpvideo"
	MsgTypeNews   = "mpnews"
	MsgTypeWxCard = "wxcard"
)

type MessageHeader struct {
	Filter struct {
		GroupId int64 `json:"group_id"`
	} `json:"filter"`
	MsgType string `json:"msgtype"`
}

type Text struct {
	MessageHeader
=======
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
>>>>>>> github/v2
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
<<<<<<< HEAD
	MessageHeader
=======
	MsgHeader
>>>>>>> github/v2
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
<<<<<<< HEAD
	MessageHeader
=======
	MsgHeader
>>>>>>> github/v2
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
<<<<<<< HEAD
	MessageHeader
=======
	MsgHeader
>>>>>>> github/v2
	Video struct {
		MediaId string `json:"media_id"`
	} `json:"mpvideo"`
}

// 新建视频消息
<<<<<<< HEAD
//  NOTE: 对于临时素材, mediaId 应该通过 media.Client.CreateVideo 得到
=======
//  NOTE: 对于临时素材, mediaId 应该通过 media.UploadVideo2 得到
>>>>>>> github/v2
func NewVideo(groupId int64, mediaId string) *Video {
	var msg Video
	msg.MsgType = MsgTypeVideo
	msg.Filter.GroupId = groupId
	msg.Video.MediaId = mediaId
	return &msg
}

// 图文消息
type News struct {
<<<<<<< HEAD
	MessageHeader
=======
	MsgHeader
>>>>>>> github/v2
	News struct {
		MediaId string `json:"media_id"`
	} `json:"mpnews"`
}

// 新建图文消息
<<<<<<< HEAD
//  NOTE: 对于临时素材, mediaId 应该通过 media.Client.CreateNews 得到
=======
//  NOTE: 对于临时素材, mediaId 应该通过 media.UploadNews 得到
>>>>>>> github/v2
func NewNews(groupId int64, mediaId string) *News {
	var msg News
	msg.MsgType = MsgTypeNews
	msg.Filter.GroupId = groupId
	msg.News.MediaId = mediaId
	return &msg
}

// 卡券消息
type WxCard struct {
<<<<<<< HEAD
	MessageHeader
	WxCard struct {
		CardId string `json:"card_id"`
=======
	MsgHeader
	WxCard struct {
		CardId  string `json:"card_id"`
		CardExt string `json:"card_ext,omitempty"`
>>>>>>> github/v2
	} `json:"wxcard"`
}

// 新建卡券, 特别注意: 目前该接口仅支持填入非自定义code的卡券和预存模式的自定义code卡券.
<<<<<<< HEAD
func NewWxCard(groupId int64, cardId string) *WxCard {
=======
//  cardExt 可以为空
func NewWxCard(groupId int64, cardId, cardExt string) *WxCard {
>>>>>>> github/v2
	var msg WxCard
	msg.MsgType = MsgTypeWxCard
	msg.Filter.GroupId = groupId
	msg.WxCard.CardId = cardId
<<<<<<< HEAD
=======
	msg.WxCard.CardExt = cardExt
>>>>>>> github/v2
	return &msg
}
