package request

import (
	"github.com/chanxuehong/wechat/work/message"
)

const (
	MsgTypeEvent message.MsgType = "event" // 事件消息
	// 普通消息类型
	MsgTypeText     message.MsgType = "text"     // 文本消息
	MsgTypeImage    message.MsgType = "image"    // 图片消息
	MsgTypeVoice    message.MsgType = "voice"    // 语音消息
	MsgTypeVideo    message.MsgType = "video"    // 视频消息
	MsgTypeLocation message.MsgType = "location" // 地理位置消息
	MsgTypeLink     message.MsgType = "link"     // 链接消息
)

// 文本消息
type Text struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	MsgId   int64         `xml:"MsgId"   json:"MsgId"`   // 消息id, 64位整型
	Content message.CDATA `xml:"Content" json:"Content"` // 文本消息内容
}

func GetText(msg *message.MixedMsg) *Text {
	return &Text{
		MsgHeader: msg.MsgHeader,
		MsgId:     msg.MsgId,
		Content:   msg.Content,
	}
}

// 图片消息
type Image struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	MsgId   int64         `xml:"MsgId"   json:"MsgId"`   // 消息id, 64位整型
	MediaId message.CDATA `xml:"MediaId" json:"MediaId"` // 图片消息媒体id, 可以调用多媒体文件下载接口拉取数据.
	PicURL  message.CDATA `xml:"PicUrl"  json:"PicUrl"`  // 图片链接
}

func GetImage(msg *message.MixedMsg) *Image {
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
	message.MsgHeader
	MsgId   int64         `xml:"MsgId"   json:"MsgId"`   // 消息id, 64位整型
	MediaId message.CDATA `xml:"MediaId" json:"MediaId"` // 语音消息媒体id, 可以调用多媒体文件下载接口拉取该媒体
	Format  message.CDATA `xml:"Format"  json:"Format"`  // 语音格式, 如amr, speex等
}

func GetVoice(msg *message.MixedMsg) *Voice {
	return &Voice{
		MsgHeader: msg.MsgHeader,
		MsgId:     msg.MsgId,
		MediaId:   msg.MediaId,
		Format:    msg.Format,
	}
}

// 视频消息
type Video struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	MsgId        int64         `xml:"MsgId"        json:"MsgId"`        // 消息id, 64位整型
	MediaId      message.CDATA `xml:"MediaId"      json:"MediaId"`      // 视频消息媒体id, 可以调用多媒体文件下载接口拉取数据.
	ThumbMediaId message.CDATA `xml:"ThumbMediaId" json:"ThumbMediaId"` // 视频消息缩略图的媒体id, 可以调用多媒体文件下载接口拉取数据.
}

func GetVideo(msg *message.MixedMsg) *Video {
	return &Video{
		MsgHeader:    msg.MsgHeader,
		MsgId:        msg.MsgId,
		MediaId:      msg.MediaId,
		ThumbMediaId: msg.ThumbMediaId,
	}
}

// 地理位置消息
type Location struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	MsgId     int64         `xml:"MsgId"      json:"MsgId"`      // 消息id, 64位整型
	LocationX float64       `xml:"Location_X" json:"Location_X"` // 地理位置纬度
	LocationY float64       `xml:"Location_Y" json:"Location_Y"` // 地理位置经度
	Scale     int           `xml:"Scale"      json:"Scale"`      // 地图缩放大小
	Label     message.CDATA `xml:"Label"      json:"Label"`      // 地理位置信息
}

func GetLocation(msg *message.MixedMsg) *Location {
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
	message.MsgHeader
	MsgId       int64         `xml:"MsgId"       json:"MsgId"`       // 消息id, 64位整型
	Title       message.CDATA `xml:"Title"       json:"Title"`       // 消息标题
	Description message.CDATA `xml:"Description" json:"Description"` // 消息描述
	URL         message.CDATA `xml:"Url"         json:"Url"`         // 消息链接
	PicURL      message.CDATA `xml:"PicUrl"  json:"PicUrl"`          // 图片链接
}

func GetLink(msg *message.MixedMsg) *Link {
	return &Link{
		MsgHeader:   msg.MsgHeader,
		MsgId:       msg.MsgId,
		Title:       msg.Title,
		Description: msg.Description,
		URL:         msg.URL,
		PicURL:      msg.PicURL,
	}
}
