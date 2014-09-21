// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package request

// 点击菜单拉取消息时的事件推送
type MenuClickEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event    string `xml:"Event"    json:"Event"`    // 事件类型，CLICK
	EventKey string `xml:"EventKey" json:"EventKey"` // 事件KEY值，与自定义菜单接口中KEY值对应
}

func (req *Request) MenuClickEvent() (event *MenuClickEvent) {
	event = &MenuClickEvent{
		CommonHead: req.CommonHead,
		Event:      req.Event,
		EventKey:   req.EventKey,
	}
	return
}

// 点击菜单跳转链接时的事件推送
type MenuViewEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event    string `xml:"Event"    json:"Event"`    // 事件类型，VIEW
	EventKey string `xml:"EventKey" json:"EventKey"` // 事件KEY值，设置的跳转URL
}

func (req *Request) MenuViewEvent() (event *MenuViewEvent) {
	event = &MenuViewEvent{
		CommonHead: req.CommonHead,
		Event:      req.Event,
		EventKey:   req.EventKey,
	}
	return
}

// scancode_push：扫码推事件的事件推送
type MenuScanCodePushEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event    string `xml:"Event"    json:"Event"`    // 事件类型，scancode_push
	EventKey string `xml:"EventKey" json:"EventKey"` // 事件KEY值，由开发者在创建菜单时设定

	ScanCodeInfo struct {
		ScanType   string `xml:"ScanType"   json:"ScanType"`   // 扫描类型，一般是qrcode
		ScanResult string `xml:"ScanResult" json:"ScanResult"` // 扫描结果，即二维码对应的字符串信息
	} `xml:"ScanCodeInfo" json:"ScanCodeInfo"` // 扫描信息
}

func (req *Request) MenuScanCodePushEvent() (event *MenuScanCodePushEvent) {
	event = &MenuScanCodePushEvent{
		CommonHead:   req.CommonHead,
		Event:        req.Event,
		EventKey:     req.EventKey,
		ScanCodeInfo: req.ScanCodeInfo,
	}
	return
}

// scancode_waitmsg：扫码推事件且弹出“消息接收中”提示框的事件推送
type MenuScanCodeWaitMsgEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event    string `xml:"Event"    json:"Event"`    // 事件类型，scancode_waitmsg
	EventKey string `xml:"EventKey" json:"EventKey"` // 事件KEY值，由开发者在创建菜单时设定

	ScanCodeInfo struct {
		ScanType   string `xml:"ScanType"   json:"ScanType"`   // 扫描类型，一般是qrcode
		ScanResult string `xml:"ScanResult" json:"ScanResult"` // 扫描结果，即二维码对应的字符串信息
	} `xml:"ScanCodeInfo" json:"ScanCodeInfo"` // 扫描信息
}

func (req *Request) MenuScanCodeWaitMsgEvent() (event *MenuScanCodeWaitMsgEvent) {
	event = &MenuScanCodeWaitMsgEvent{
		CommonHead:   req.CommonHead,
		Event:        req.Event,
		EventKey:     req.EventKey,
		ScanCodeInfo: req.ScanCodeInfo,
	}
	return
}

// pic_sysphoto：弹出系统拍照发图的事件推送
type MenuPicSysPhotoEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event    string `xml:"Event"    json:"Event"`    // 事件类型，pic_sysphoto
	EventKey string `xml:"EventKey" json:"EventKey"` // 事件KEY值，由开发者在创建菜单时设定

	SendPicsInfo struct {
		Count   int `xml:"Count"   json:"Count"` // 发送的图片数量
		PicList []struct {
			PicMd5Sum string `xml:"PicMd5Sum" json:"PicMd5Sum"` // 图片的MD5值，开发者若需要，可用于验证接收到图片
		} `xml:"PicList>item,omitempty" json:"PicList,omitempty"` // 图片列表
	} `xml:"SendPicsInfo" json:"SendPicsInfo"` // 发送的图片信息
}

func (req *Request) MenuPicSysPhotoEvent() (event *MenuPicSysPhotoEvent) {
	event = &MenuPicSysPhotoEvent{
		CommonHead:   req.CommonHead,
		Event:        req.Event,
		EventKey:     req.EventKey,
		SendPicsInfo: req.SendPicsInfo,
	}
	return
}

// pic_photo_or_album：弹出拍照或者相册发图的事件推送
type MenuPicPhotoOrAlbumEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event    string `xml:"Event"    json:"Event"`    // 事件类型，pic_photo_or_album
	EventKey string `xml:"EventKey" json:"EventKey"` // 事件KEY值，由开发者在创建菜单时设定

	SendPicsInfo struct {
		Count   int `xml:"Count"   json:"Count"` // 发送的图片数量
		PicList []struct {
			PicMd5Sum string `xml:"PicMd5Sum" json:"PicMd5Sum"` // 图片的MD5值，开发者若需要，可用于验证接收到图片
		} `xml:"PicList>item,omitempty" json:"PicList,omitempty"` // 图片列表
	} `xml:"SendPicsInfo" json:"SendPicsInfo"` // 发送的图片信息
}

func (req *Request) MenuPicPhotoOrAlbumEvent() (event *MenuPicPhotoOrAlbumEvent) {
	event = &MenuPicPhotoOrAlbumEvent{
		CommonHead:   req.CommonHead,
		Event:        req.Event,
		EventKey:     req.EventKey,
		SendPicsInfo: req.SendPicsInfo,
	}
	return
}

// pic_weixin：弹出微信相册发图器的事件推送
type MenuPicWeixinEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event    string `xml:"Event"    json:"Event"`    // 事件类型，pic_weixin
	EventKey string `xml:"EventKey" json:"EventKey"` // 事件KEY值，由开发者在创建菜单时设定

	SendPicsInfo struct {
		Count   int `xml:"Count"   json:"Count"` // 发送的图片数量
		PicList []struct {
			PicMd5Sum string `xml:"PicMd5Sum" json:"PicMd5Sum"` // 图片的MD5值，开发者若需要，可用于验证接收到图片
		} `xml:"PicList>item,omitempty" json:"PicList,omitempty"` // 图片列表
	} `xml:"SendPicsInfo" json:"SendPicsInfo"` // 发送的图片信息
}

func (req *Request) MenuPicWeixinEvent() (event *MenuPicWeixinEvent) {
	event = &MenuPicWeixinEvent{
		CommonHead:   req.CommonHead,
		Event:        req.Event,
		EventKey:     req.EventKey,
		SendPicsInfo: req.SendPicsInfo,
	}
	return
}

// location_select：弹出地理位置选择器的事件推送
type MenuLocationSelectEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event    string `xml:"Event"    json:"Event"`    // 事件类型，location_select
	EventKey string `xml:"EventKey" json:"EventKey"` // 事件KEY值，由开发者在创建菜单时设定

	SendLocationInfo struct {
		LocationX float64 `xml:"Location_X" json:"Location_X"` // 地理位置纬度
		LocationY float64 `xml:"Location_Y" json:"Location_Y"` // 地理位置经度
		Scale     int     `xml:"Scale"      json:"Scale"`      // 精度，可理解为精度或者比例尺、越精细的话 scale越高
		Label     string  `xml:"Label"      json:"Label"`      // 地理位置的字符串信息
		Poiname   string  `xml:"Poiname"    json:"Poiname"`    // 朋友圈POI的名字，可能为空
	} `xml:"SendLocationInfo" json:"SendLocationInfo"` // 发送的位置信息
}

func (req *Request) MenuLocationSelectEvent() (event *MenuLocationSelectEvent) {
	event = &MenuLocationSelectEvent{
		CommonHead:       req.CommonHead,
		Event:            req.Event,
		EventKey:         req.EventKey,
		SendLocationInfo: req.SendLocationInfo,
	}
	return
}
