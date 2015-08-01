// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package request

import (
	"github.com/zihuxinyu/wechat/corp/chat"
)

const (
// 微信服务器推送过来的消息类型
	MsgTypeText = "text"     // 文本消息
	MsgTypeImage = "image"    // 图片消息
)

type Text struct {
	XMLName struct {} `xml:"xml" json:"-"`

	chat.ItemHeader
	chat.Receiver
	MsgId   int64   `xml:"MsgId" json:"MsgId"`     //消息id, 64位整型
	Content string  `xml:"Content" json:"Content"` //消息内容，支持表情
}

func GetText(msg *chat.MixedMessage) *Text {
	return &Text{
		ItemHeader:msg.CurrentItem.ItemHeader,
		Receiver:msg.CurrentItem.Receiver,
		MsgId:msg.CurrentItem.MsgId,
		Content:msg.CurrentItem.Content,
	}


}


type Image struct {
	XMLName struct {} `xml:"xml" json:"-"`
	chat.ItemHeader
	chat.Receiver

	MsgId   int64  `xml:"MsgId"   json:"MsgId"`   // 消息id, 64位整型
	MediaId string `xml:"MediaId" json:"MediaId"` // 图片媒体文件id, 可以调用获取媒体文件接口拉取数据
	PicURL  string `xml:"PicUrl"  json:"PicUrl"`  // 图片链接
}

func GetImage(msg *chat.MixedMessage) *Image {
	return &Image{
		ItemHeader:msg.CurrentItem.ItemHeader,
		Receiver:msg.CurrentItem.Receiver,
		MsgId:msg.CurrentItem.MsgId,
		MediaId:msg.CurrentItem.MediaId,
		PicURL:msg.CurrentItem.PicUrl,
	}

}

