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
	}
}

// scancode_push: 扫码推事件的事件推送
type ScanCodePushEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
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
	}
}

// scancode_waitmsg: 扫码推事件且弹出"消息接收中"提示框的事件推送
type ScanCodeWaitMsgEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
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
	}
}

// pic_sysphoto: 弹出系统拍照发图的事件推送
type PicSysPhotoEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
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
	}
}

// pic_photo_or_album: 弹出拍照或者相册发图的事件推送
type PicPhotoOrAlbumEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
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
	}
}

// pic_weixin: 弹出微信相册发图器的事件推送
type PicWeixinEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
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
	}
}

// location_select: 弹出地理位置选择器的事件推送
type LocationSelectEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	EventType core.EventType `xml:"Event"    json:"Event"`    // 事件类型, location_select
	EventKey  string         `xml:"EventKey" json:"EventKey"` // 事件KEY值, 由开发者在创建菜单时设定

	SendLocationInfo *struct {
		LocationX float64 `xml:"Location_X" json:"Location_X"` // 地理位置纬度
		LocationY float64 `xml:"Location_Y" json:"Location_Y"` // 地理位置经度
		Scale     int     `xml:"Scale"      json:"Scale"`      // 精度, 可理解为精度或者比例尺, 越精细的话 scale越高
		Label     string  `xml:"Label"      json:"Label"`      // 地理位置的字符串信息
		PoiName   string  `xml:"Poiname"    json:"Poiname"`    // 朋友圈POI的名字, 可能为空
	} `xml:"SendLocationInfo,omitempty" json:"SendLocationInfo,omitempty"`
}

func GetLocationSelectEvent(msg *core.MixedMsg) *LocationSelectEvent {
	return &LocationSelectEvent{
		MsgHeader:        msg.MsgHeader,
		EventType:        msg.EventType,
		EventKey:         msg.EventKey,
		SendLocationInfo: msg.SendLocationInfo,
	}
}
