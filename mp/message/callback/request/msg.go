package request

import (
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
)

const (
	// 普通消息类型
	MsgTypeText       core.MsgType = "text"       // 文本消息
	MsgTypeImage      core.MsgType = "image"      // 图片消息
	MsgTypeVoice      core.MsgType = "voice"      // 语音消息
	MsgTypeVideo      core.MsgType = "video"      // 视频消息
	MsgTypeShortVideo core.MsgType = "shortvideo" // 小视频消息
	MsgTypeLocation   core.MsgType = "location"   // 地理位置消息
	MsgTypeLink       core.MsgType = "link"       // 链接消息
)

// 文本消息
type Text struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	MsgId   int64  `xml:"MsgId"   json:"MsgId"`   // 消息id, 64位整型
	Content string `xml:"Content" json:"Content"` // 文本消息内容
}

func GetText(msg *core.MixedMsg) *Text {
	return &Text{
		MsgHeader: msg.MsgHeader,
		MsgId:     msg.MsgId,
		Content:   msg.Content,
	}
}

// 图片消息
type Image struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	MsgId   int64  `xml:"MsgId"   json:"MsgId"`   // 消息id, 64位整型
	MediaId string `xml:"MediaId" json:"MediaId"` // 图片消息媒体id, 可以调用多媒体文件下载接口拉取数据.
	PicURL  string `xml:"PicUrl"  json:"PicUrl"`  // 图片链接
}

func GetImage(msg *core.MixedMsg) *Image {
	return &Image{
		MsgHeader: msg.MsgHeader,
		MsgId:     msg.MsgId,
		MediaId:   msg.MediaId,
		PicURL:    msg.PicURL,
	}
}

// 语音消息
type Voice struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	MsgId   int64  `xml:"MsgId"   json:"MsgId"`   // 消息id, 64位整型
	MediaId string `xml:"MediaId" json:"MediaId"` // 语音消息媒体id, 可以调用多媒体文件下载接口拉取该媒体
	Format  string `xml:"Format"  json:"Format"`  // 语音格式, 如amr, speex等

	// 语音识别结果, UTF8编码,
	// NOTE: 需要开通语音识别功能, 否则该字段为空, 即使开通了语音识别该字段还是有可能为空
	Recognition string `xml:"Recognition,omitempty" json:"Recognition,omitempty"`
}

func GetVoice(msg *core.MixedMsg) *Voice {
	return &Voice{
		MsgHeader:   msg.MsgHeader,
		MsgId:       msg.MsgId,
		MediaId:     msg.MediaId,
		Format:      msg.Format,
		Recognition: msg.Recognition,
	}
}

// 视频消息
type Video struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	MsgId        int64  `xml:"MsgId"        json:"MsgId"`        // 消息id, 64位整型
	MediaId      string `xml:"MediaId"      json:"MediaId"`      // 视频消息媒体id, 可以调用多媒体文件下载接口拉取数据.
	ThumbMediaId string `xml:"ThumbMediaId" json:"ThumbMediaId"` // 视频消息缩略图的媒体id, 可以调用多媒体文件下载接口拉取数据.
}

func GetVideo(msg *core.MixedMsg) *Video {
	return &Video{
		MsgHeader:    msg.MsgHeader,
		MsgId:        msg.MsgId,
		MediaId:      msg.MediaId,
		ThumbMediaId: msg.ThumbMediaId,
	}
}

// 小视频消息
type ShortVideo struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	MsgId        int64  `xml:"MsgId"        json:"MsgId"`        // 消息id, 64位整型
	MediaId      string `xml:"MediaId"      json:"MediaId"`      // 视频消息媒体id, 可以调用多媒体文件下载接口拉取数据.
	ThumbMediaId string `xml:"ThumbMediaId" json:"ThumbMediaId"` // 视频消息缩略图的媒体id, 可以调用多媒体文件下载接口拉取数据.
}

func GetShortVideo(msg *core.MixedMsg) *ShortVideo {
	return &ShortVideo{
		MsgHeader:    msg.MsgHeader,
		MsgId:        msg.MsgId,
		MediaId:      msg.MediaId,
		ThumbMediaId: msg.ThumbMediaId,
	}
}

// 地理位置消息
type Location struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	MsgId     int64   `xml:"MsgId"      json:"MsgId"`      // 消息id, 64位整型
	LocationX float64 `xml:"Location_X" json:"Location_X"` // 地理位置纬度
	LocationY float64 `xml:"Location_Y" json:"Location_Y"` // 地理位置经度
	Scale     int     `xml:"Scale"      json:"Scale"`      // 地图缩放大小
	Label     string  `xml:"Label"      json:"Label"`      // 地理位置信息
}

func GetLocation(msg *core.MixedMsg) *Location {
	return &Location{
		MsgHeader: msg.MsgHeader,
		MsgId:     msg.MsgId,
		LocationX: msg.LocationX,
		LocationY: msg.LocationY,
		Scale:     msg.Scale,
		Label:     msg.Label,
	}
}

// 链接消息
type Link struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	MsgId       int64  `xml:"MsgId"       json:"MsgId"`       // 消息id, 64位整型
	Title       string `xml:"Title"       json:"Title"`       // 消息标题
	Description string `xml:"Description" json:"Description"` // 消息描述
	URL         string `xml:"Url"         json:"Url"`         // 消息链接
}

func GetLink(msg *core.MixedMsg) *Link {
	return &Link{
		MsgHeader:   msg.MsgHeader,
		MsgId:       msg.MsgId,
		Title:       msg.Title,
		Description: msg.Description,
		URL:         msg.URL,
	}
}
