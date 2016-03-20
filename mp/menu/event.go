<<<<<<< HEAD
// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package menu

import (
	"github.com/chanxuehong/wechat/mp"
)

const (
	EventTypeClick = "CLICK" // 点击菜单拉取消息时的事件推送
	EventTypeView  = "VIEW"  // 点击菜单跳转链接时的事件推送

	// 请注意, 下面的事件仅支持微信iPhone5.4.1以上版本, 和Android5.4以上版本的微信用户,
	// 旧版本微信用户点击后将没有回应, 开发者也不能正常接收到事件推送.
	EventTypeScanCodePush    = "scancode_push"      // scancode_push: 扫码推事件的事件推送
	EventTypeScanCodeWaitMsg = "scancode_waitmsg"   // scancode_waitmsg: 扫码推事件且弹出"消息接收中"提示框的事件推送
	EventTypePicSysPhoto     = "pic_sysphoto"       // pic_sysphoto: 弹出系统拍照发图的事件推送
	EventTypePicPhotoOrAlbum = "pic_photo_or_album" // pic_photo_or_album: 弹出拍照或者相册发图的事件推送
	EventTypePicWeixin       = "pic_weixin"         // pic_weixin: 弹出微信相册发图器的事件推送
	EventTypeLocationSelect  = "location_select"    // location_select: 弹出地理位置选择器的事件推送
)

// 点击菜单拉取消息时的事件推送
type ClickEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	mp.MessageHeader

	Event    string `xml:"Event"    json:"Event"`    // 事件类型, CLICK
	EventKey string `xml:"EventKey" json:"EventKey"` // 事件KEY值, 与自定义菜单接口中KEY值对应
}

func GetClickEvent(msg *mp.MixedMessage) *ClickEvent {
	return &ClickEvent{
		MessageHeader: msg.MessageHeader,
		Event:         msg.Event,
		EventKey:      msg.EventKey,
	}
}

// 点击菜单跳转链接时的事件推送
type ViewEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	mp.MessageHeader

	Event    string `xml:"Event"    json:"Event"`    // 事件类型, VIEW
	EventKey string `xml:"EventKey" json:"EventKey"` // 事件KEY值, 设置的跳转URL
}

func GetViewEvent(msg *mp.MixedMessage) *ViewEvent {
	return &ViewEvent{
		MessageHeader: msg.MessageHeader,
		Event:         msg.Event,
		EventKey:      msg.EventKey,
=======
package menu

import (
	"github.com/chanxuehong/wechat/mp/core"
)

const (
	EventTypeClick core.EventType = "CLICK" // 点击菜单拉取消息时的事件推送
	EventTypeView  core.EventType = "VIEW"  // 点击菜单跳转链接时的事件推送

	// 请注意, 下面的事件仅支持微信iPhone5.4.1以上版本, 和Android5.4以上版本的微信用户,
	// 旧版本微信用户点击后将没有回应, 开发者也不能正常接收到事件推送.
	EventTypeScanCodePush    core.EventType = "scancode_push"      // 扫码推事件的事件推送
	EventTypeScanCodeWaitMsg core.EventType = "scancode_waitmsg"   // 扫码推事件且弹出"消息接收中"提示框的事件推送
	EventTypePicSysPhoto     core.EventType = "pic_sysphoto"       // 弹出系统拍照发图的事件推送
	EventTypePicPhotoOrAlbum core.EventType = "pic_photo_or_album" // 弹出拍照或者相册发图的事件推送
	EventTypePicWeixin       core.EventType = "pic_weixin"         // 弹出微信相册发图器的事件推送
	EventTypeLocationSelect  core.EventType = "location_select"    // 弹出地理位置选择器的事件推送
)

// CLICK: 点击菜单拉取消息时的事件推送
type ClickEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	EventType core.EventType `xml:"Event"    json:"Event"`    // 事件类型, CLICK
	EventKey  string         `xml:"EventKey" json:"EventKey"` // 事件KEY值, 与自定义菜单接口中KEY值对应
}

func GetClickEvent(msg *core.MixedMsg) *ClickEvent {
	return &ClickEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		EventKey:  msg.EventKey,
	}
}

// VIEW: 点击菜单跳转链接时的事件推送
type ViewEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	EventType core.EventType `xml:"Event"            json:"Event"`            // 事件类型, VIEW
	EventKey  string         `xml:"EventKey"         json:"EventKey"`         // 事件KEY值, 设置的跳转URL
	MenuId    int64          `xml:"MenuId,omitempty" json:"MenuId,omitempty"` // 菜单ID，如果是个性化菜单，则可以通过这个字段，知道是哪个规则的菜单被点击了。
}

func GetViewEvent(msg *core.MixedMsg) *ViewEvent {
	return &ViewEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		EventKey:  msg.EventKey,
		MenuId:    msg.MenuId,
>>>>>>> github/v2
	}
}

// scancode_push: 扫码推事件的事件推送
type ScanCodePushEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
<<<<<<< HEAD
	mp.MessageHeader

	Event    string `xml:"Event"    json:"Event"`    // 事件类型, scancode_push
	EventKey string `xml:"EventKey" json:"EventKey"` // 事件KEY值, 由开发者在创建菜单时设定

	ScanCodeInfo struct {
		ScanType   string `xml:"ScanType"   json:"ScanType"`   // 扫描类型, 一般是qrcode
		ScanResult string `xml:"ScanResult" json:"ScanResult"` // 扫描结果, 即二维码对应的字符串信息
	} `xml:"ScanCodeInfo" json:"ScanCodeInfo"`
}

func GetScanCodePushEvent(msg *mp.MixedMessage) *ScanCodePushEvent {
	return &ScanCodePushEvent{
		MessageHeader: msg.MessageHeader,
		Event:         msg.Event,
		EventKey:      msg.EventKey,
		ScanCodeInfo:  msg.ScanCodeInfo,
=======
	core.MsgHeader
	EventType core.EventType `xml:"Event"    json:"Event"`    // 事件类型, scancode_push
	EventKey  string         `xml:"EventKey" json:"EventKey"` // 事件KEY值, 由开发者在创建菜单时设定

	ScanCodeInfo *struct {
		ScanType   string `xml:"ScanType"   json:"ScanType"`   // 扫描类型, 一般是qrcode
		ScanResult string `xml:"ScanResult" json:"ScanResult"` // 扫描结果, 即二维码对应的字符串信息
	} `xml:"ScanCodeInfo,omitempty" json:"ScanCodeInfo,omitempty"`
}

func GetScanCodePushEvent(msg *core.MixedMsg) *ScanCodePushEvent {
	return &ScanCodePushEvent{
		MsgHeader:    msg.MsgHeader,
		EventType:    msg.EventType,
		EventKey:     msg.EventKey,
		ScanCodeInfo: msg.ScanCodeInfo,
>>>>>>> github/v2
	}
}

// scancode_waitmsg: 扫码推事件且弹出"消息接收中"提示框的事件推送
type ScanCodeWaitMsgEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
<<<<<<< HEAD
	mp.MessageHeader

	Event    string `xml:"Event"    json:"Event"`    // 事件类型, scancode_waitmsg
	EventKey string `xml:"EventKey" json:"EventKey"` // 事件KEY值, 由开发者在创建菜单时设定

	ScanCodeInfo struct {
		ScanType   string `xml:"ScanType"   json:"ScanType"`   // 扫描类型, 一般是qrcode
		ScanResult string `xml:"ScanResult" json:"ScanResult"` // 扫描结果, 即二维码对应的字符串信息
	} `xml:"ScanCodeInfo" json:"ScanCodeInfo"`
}

func GetScanCodeWaitMsgEvent(msg *mp.MixedMessage) *ScanCodeWaitMsgEvent {
	return &ScanCodeWaitMsgEvent{
		MessageHeader: msg.MessageHeader,
		Event:         msg.Event,
		EventKey:      msg.EventKey,
		ScanCodeInfo:  msg.ScanCodeInfo,
=======
	core.MsgHeader
	EventType core.EventType `xml:"Event"    json:"Event"`    // 事件类型, scancode_waitmsg
	EventKey  string         `xml:"EventKey" json:"EventKey"` // 事件KEY值, 由开发者在创建菜单时设定

	ScanCodeInfo *struct {
		ScanType   string `xml:"ScanType"   json:"ScanType"`   // 扫描类型, 一般是qrcode
		ScanResult string `xml:"ScanResult" json:"ScanResult"` // 扫描结果, 即二维码对应的字符串信息
	} `xml:"ScanCodeInfo,omitempty" json:"ScanCodeInfo,omitempty"`
}

func GetScanCodeWaitMsgEvent(msg *core.MixedMsg) *ScanCodeWaitMsgEvent {
	return &ScanCodeWaitMsgEvent{
		MsgHeader:    msg.MsgHeader,
		EventType:    msg.EventType,
		EventKey:     msg.EventKey,
		ScanCodeInfo: msg.ScanCodeInfo,
>>>>>>> github/v2
	}
}

// pic_sysphoto: 弹出系统拍照发图的事件推送
type PicSysPhotoEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
<<<<<<< HEAD
	mp.MessageHeader

	Event    string `xml:"Event"    json:"Event"`    // 事件类型, pic_sysphoto
	EventKey string `xml:"EventKey" json:"EventKey"` // 事件KEY值, 由开发者在创建菜单时设定

	SendPicsInfo struct {
		Count   int `xml:"Count" json:"Count"` // 发送的图片数量
		PicList []struct {
			PicMD5Sum string `xml:"PicMd5Sum" json:"PicMd5Sum"` // 图片的MD5值, 开发者若需要, 可用于验证接收到图片
		} `xml:"PicList>item,omitempty" json:"PicList,omitempty"` // 图片列表
	} `xml:"SendPicsInfo" json:"SendPicsInfo"` // 发送的图片信息
}

func GetPicSysPhotoEvent(msg *mp.MixedMessage) *PicSysPhotoEvent {
	return &PicSysPhotoEvent{
		MessageHeader: msg.MessageHeader,
		Event:         msg.Event,
		EventKey:      msg.EventKey,
		SendPicsInfo:  msg.SendPicsInfo,
=======
	core.MsgHeader
	EventType core.EventType `xml:"Event"    json:"Event"`    // 事件类型, pic_sysphoto
	EventKey  string         `xml:"EventKey" json:"EventKey"` // 事件KEY值, 由开发者在创建菜单时设定

	SendPicsInfo *struct {
		Count   int `xml:"Count" json:"Count"`
		PicList []struct {
			PicMd5Sum string `xml:"PicMd5Sum" json:"PicMd5Sum"`
		} `xml:"PicList>item,omitempty" json:"PicList,omitempty"`
	} `xml:"SendPicsInfo,omitempty" json:"SendPicsInfo,omitempty"`
}

func GetPicSysPhotoEvent(msg *core.MixedMsg) *PicSysPhotoEvent {
	return &PicSysPhotoEvent{
		MsgHeader:    msg.MsgHeader,
		EventType:    msg.EventType,
		EventKey:     msg.EventKey,
		SendPicsInfo: msg.SendPicsInfo,
>>>>>>> github/v2
	}
}

// pic_photo_or_album: 弹出拍照或者相册发图的事件推送
type PicPhotoOrAlbumEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
<<<<<<< HEAD
	mp.MessageHeader

	Event    string `xml:"Event"    json:"Event"`    // 事件类型, pic_photo_or_album
	EventKey string `xml:"EventKey" json:"EventKey"` // 事件KEY值, 由开发者在创建菜单时设定

	SendPicsInfo struct {
		Count   int `xml:"Count" json:"Count"` // 发送的图片数量
		PicList []struct {
			PicMD5Sum string `xml:"PicMd5Sum" json:"PicMd5Sum"` // 图片的MD5值, 开发者若需要, 可用于验证接收到图片
		} `xml:"PicList>item,omitempty" json:"PicList,omitempty"` // 图片列表
	} `xml:"SendPicsInfo" json:"SendPicsInfo"` // 发送的图片信息
}

func GetPicPhotoOrAlbumEvent(msg *mp.MixedMessage) *PicPhotoOrAlbumEvent {
	return &PicPhotoOrAlbumEvent{
		MessageHeader: msg.MessageHeader,
		Event:         msg.Event,
		EventKey:      msg.EventKey,
		SendPicsInfo:  msg.SendPicsInfo,
=======
	core.MsgHeader
	EventType core.EventType `xml:"Event"    json:"Event"`    // 事件类型, pic_photo_or_album
	EventKey  string         `xml:"EventKey" json:"EventKey"` // 事件KEY值, 由开发者在创建菜单时设定

	SendPicsInfo *struct {
		Count   int `xml:"Count" json:"Count"`
		PicList []struct {
			PicMd5Sum string `xml:"PicMd5Sum" json:"PicMd5Sum"`
		} `xml:"PicList>item,omitempty" json:"PicList,omitempty"`
	} `xml:"SendPicsInfo,omitempty" json:"SendPicsInfo,omitempty"`
}

func GetPicPhotoOrAlbumEvent(msg *core.MixedMsg) *PicPhotoOrAlbumEvent {
	return &PicPhotoOrAlbumEvent{
		MsgHeader:    msg.MsgHeader,
		EventType:    msg.EventType,
		EventKey:     msg.EventKey,
		SendPicsInfo: msg.SendPicsInfo,
>>>>>>> github/v2
	}
}

// pic_weixin: 弹出微信相册发图器的事件推送
type PicWeixinEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
<<<<<<< HEAD
	mp.MessageHeader

	Event    string `xml:"Event"    json:"Event"`    // 事件类型, pic_weixin
	EventKey string `xml:"EventKey" json:"EventKey"` // 事件KEY值, 由开发者在创建菜单时设定

	SendPicsInfo struct {
		Count   int `xml:"Count" json:"Count"` // 发送的图片数量
		PicList []struct {
			PicMD5Sum string `xml:"PicMd5Sum" json:"PicMd5Sum"` // 图片的MD5值, 开发者若需要, 可用于验证接收到图片
		} `xml:"PicList>item,omitempty" json:"PicList,omitempty"` // 图片列表
	} `xml:"SendPicsInfo" json:"SendPicsInfo"` // 发送的图片信息
}

func GetPicWeixinEvent(msg *mp.MixedMessage) *PicWeixinEvent {
	return &PicWeixinEvent{
		MessageHeader: msg.MessageHeader,
		Event:         msg.Event,
		EventKey:      msg.EventKey,
		SendPicsInfo:  msg.SendPicsInfo,
=======
	core.MsgHeader
	EventType core.EventType `xml:"Event"    json:"Event"`    // 事件类型, pic_weixin
	EventKey  string         `xml:"EventKey" json:"EventKey"` // 事件KEY值, 由开发者在创建菜单时设定

	SendPicsInfo *struct {
		Count   int `xml:"Count" json:"Count"`
		PicList []struct {
			PicMd5Sum string `xml:"PicMd5Sum" json:"PicMd5Sum"`
		} `xml:"PicList>item,omitempty" json:"PicList,omitempty"`
	} `xml:"SendPicsInfo,omitempty" json:"SendPicsInfo,omitempty"`
}

func GetPicWeixinEvent(msg *core.MixedMsg) *PicWeixinEvent {
	return &PicWeixinEvent{
		MsgHeader:    msg.MsgHeader,
		EventType:    msg.EventType,
		EventKey:     msg.EventKey,
		SendPicsInfo: msg.SendPicsInfo,
>>>>>>> github/v2
	}
}

// location_select: 弹出地理位置选择器的事件推送
type LocationSelectEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
<<<<<<< HEAD
	mp.MessageHeader

	Event    string `xml:"Event"    json:"Event"`    // 事件类型, location_select
	EventKey string `xml:"EventKey" json:"EventKey"` // 事件KEY值, 由开发者在创建菜单时设定

	SendLocationInfo struct {
=======
	core.MsgHeader
	EventType core.EventType `xml:"Event"    json:"Event"`    // 事件类型, location_select
	EventKey  string         `xml:"EventKey" json:"EventKey"` // 事件KEY值, 由开发者在创建菜单时设定

	SendLocationInfo *struct {
>>>>>>> github/v2
		LocationX float64 `xml:"Location_X" json:"Location_X"` // 地理位置纬度
		LocationY float64 `xml:"Location_Y" json:"Location_Y"` // 地理位置经度
		Scale     int     `xml:"Scale"      json:"Scale"`      // 精度, 可理解为精度或者比例尺, 越精细的话 scale越高
		Label     string  `xml:"Label"      json:"Label"`      // 地理位置的字符串信息
		PoiName   string  `xml:"Poiname"    json:"Poiname"`    // 朋友圈POI的名字, 可能为空
<<<<<<< HEAD
	} `xml:"SendLocationInfo" json:"SendLocationInfo"` // 发送的位置信息
}

func GetLocationSelectEvent(msg *mp.MixedMessage) *LocationSelectEvent {
	return &LocationSelectEvent{
		MessageHeader:    msg.MessageHeader,
		Event:            msg.Event,
=======
	} `xml:"SendLocationInfo,omitempty" json:"SendLocationInfo,omitempty"`
}

func GetLocationSelectEvent(msg *core.MixedMsg) *LocationSelectEvent {
	return &LocationSelectEvent{
		MsgHeader:        msg.MsgHeader,
		EventType:        msg.EventType,
>>>>>>> github/v2
		EventKey:         msg.EventKey,
		SendLocationInfo: msg.SendLocationInfo,
	}
}
