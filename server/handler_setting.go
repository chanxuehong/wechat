// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"github.com/chanxuehong/wechat/message/request"
	"net/http"
)

// 非法请求的处理函数.
//  @err: 具体的错误信息
type InvalidRequestHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)

// 未知消息类型的消息处理函数.
//  @msg: 接收到的消息体
type UnknownRequestHandlerFunc func(w http.ResponseWriter, r *http.Request, msg []byte)

// 正常的消息处理函数
type TextRequestHandlerFunc func(http.ResponseWriter, *http.Request, *request.Text)
type ImageRequestHandlerFunc func(http.ResponseWriter, *http.Request, *request.Image)
type VoiceRequestHandlerFunc func(http.ResponseWriter, *http.Request, *request.Voice)
type VideoRequestHandlerFunc func(http.ResponseWriter, *http.Request, *request.Video)
type LocationRequestHandlerFunc func(http.ResponseWriter, *http.Request, *request.Location)
type LinkRequestHandlerFunc func(http.ResponseWriter, *http.Request, *request.Link)
type SubscribeEventHandlerFunc func(http.ResponseWriter, *http.Request, *request.SubscribeEvent)
type UnsubscribeEventHandlerFunc func(http.ResponseWriter, *http.Request, *request.UnsubscribeEvent)
type SubscribeByScanEventHandlerFunc func(http.ResponseWriter, *http.Request, *request.SubscribeByScanEvent)
type ScanEventHandlerFunc func(http.ResponseWriter, *http.Request, *request.ScanEvent)
type LocationEventHandlerFunc func(http.ResponseWriter, *http.Request, *request.LocationEvent)
type MenuClickEventHandlerFunc func(http.ResponseWriter, *http.Request, *request.MenuClickEvent)
type MenuViewEventHandlerFunc func(http.ResponseWriter, *http.Request, *request.MenuViewEvent)
type MenuScanCodePushEventHandlerFunc func(http.ResponseWriter, *http.Request, *request.MenuScanCodePushEvent)
type MenuScanCodeWaitMsgEventHandlerFunc func(http.ResponseWriter, *http.Request, *request.MenuScanCodeWaitMsgEvent)
type MenuPicSysPhotoEventHandlerFunc func(http.ResponseWriter, *http.Request, *request.MenuPicSysPhotoEvent)
type MenuPicPhotoOrAlbumEventHandlerFunc func(http.ResponseWriter, *http.Request, *request.MenuPicPhotoOrAlbumEvent)
type MenuPicWeixinEventHandlerFunc func(http.ResponseWriter, *http.Request, *request.MenuPicWeixinEvent)
type MenuLocationSelectEventHandlerFunc func(http.ResponseWriter, *http.Request, *request.MenuLocationSelectEvent)
type MassSendJobFinishEventHandlerFunc func(http.ResponseWriter, *http.Request, *request.MassSendJobFinishEvent)
type TemplateSendJobFinishEventHandlerFunc func(http.ResponseWriter, *http.Request, *request.TemplateSendJobFinishEvent)
type MerchantOrderEventHandlerFunc func(http.ResponseWriter, *http.Request, *request.MerchantOrderEvent)

type HandlerSetting struct {
	Token string

	// Invalid request handler
	InvalidRequestHandler InvalidRequestHandlerFunc

	// unknown request handler
	UnknownRequestHandler UnknownRequestHandlerFunc

	// common request handler
	TextRequestHandler     TextRequestHandlerFunc
	ImageRequestHandler    ImageRequestHandlerFunc
	VoiceRequestHandler    VoiceRequestHandlerFunc
	VideoRequestHandler    VideoRequestHandlerFunc
	LocationRequestHandler LocationRequestHandlerFunc
	LinkRequestHandler     LinkRequestHandlerFunc

	// event handler
	SubscribeEventHandler             SubscribeEventHandlerFunc
	UnsubscribeEventHandler           UnsubscribeEventHandlerFunc
	SubscribeByScanEventHandler       SubscribeByScanEventHandlerFunc
	ScanEventHandler                  ScanEventHandlerFunc
	LocationEventHandler              LocationEventHandlerFunc
	MenuClickEventHandler             MenuClickEventHandlerFunc
	MenuViewEventHandler              MenuViewEventHandlerFunc
	MenuScanCodePushEventHandler      MenuScanCodePushEventHandlerFunc
	MenuScanCodeWaitMsgEventHandler   MenuScanCodeWaitMsgEventHandlerFunc
	MenuPicSysPhotoEventHandler       MenuPicSysPhotoEventHandlerFunc
	MenuPicPhotoOrAlbumEventHandler   MenuPicPhotoOrAlbumEventHandlerFunc
	MenuPicWeixinEventHandler         MenuPicWeixinEventHandlerFunc
	MenuLocationSelectEventHandler    MenuLocationSelectEventHandlerFunc
	MassSendJobFinishEventHandler     MassSendJobFinishEventHandlerFunc
	TemplateSendJobFinishEventHandler TemplateSendJobFinishEventHandlerFunc
	MerchantOrderEventHandler         MerchantOrderEventHandlerFunc
}

// 根据另外一个 HandlerSetting 来初始化, 没有设置的函数将会被初始化为默认处理函数.
//  NOTE: 确保 setting != nil
func (this *HandlerSetting) initialize(setting *HandlerSetting) {
	this.Token = setting.Token

	if setting.InvalidRequestHandler != nil {
		this.InvalidRequestHandler = setting.InvalidRequestHandler
	} else {
		this.InvalidRequestHandler = defaultInvalidRequestHandler
	}

	if setting.UnknownRequestHandler != nil {
		this.UnknownRequestHandler = setting.UnknownRequestHandler
	} else {
		this.UnknownRequestHandler = defaultUnknownRequestHandler
	}

	// common request handler
	if setting.TextRequestHandler != nil {
		this.TextRequestHandler = setting.TextRequestHandler
	} else {
		this.TextRequestHandler = defaultTextRequestHandler
	}
	if setting.ImageRequestHandler != nil {
		this.ImageRequestHandler = setting.ImageRequestHandler
	} else {
		this.ImageRequestHandler = defaultImageRequestHandler
	}
	if setting.VoiceRequestHandler != nil {
		this.VoiceRequestHandler = setting.VoiceRequestHandler
	} else {
		this.VoiceRequestHandler = defaultVoiceRequestHandler
	}
	if setting.VideoRequestHandler != nil {
		this.VideoRequestHandler = setting.VideoRequestHandler
	} else {
		this.VideoRequestHandler = defaultVideoRequestHandler
	}
	if setting.LocationRequestHandler != nil {
		this.LocationRequestHandler = setting.LocationRequestHandler
	} else {
		this.LocationRequestHandler = defaultLocationRequestHandler
	}
	if setting.LinkRequestHandler != nil {
		this.LinkRequestHandler = setting.LinkRequestHandler
	} else {
		this.LinkRequestHandler = defaultLinkRequestHandler
	}

	// event handler
	if setting.SubscribeEventHandler != nil {
		this.SubscribeEventHandler = setting.SubscribeEventHandler
	} else {
		this.SubscribeEventHandler = defaultSubscribeEventHandler
	}
	if setting.UnsubscribeEventHandler != nil {
		this.UnsubscribeEventHandler = setting.UnsubscribeEventHandler
	} else {
		this.UnsubscribeEventHandler = defaultUnsubscribeEventHandler
	}
	if setting.SubscribeByScanEventHandler != nil {
		this.SubscribeByScanEventHandler = setting.SubscribeByScanEventHandler
	} else {
		this.SubscribeByScanEventHandler = defaultSubscribeByScanEventHandler
	}
	if setting.ScanEventHandler != nil {
		this.ScanEventHandler = setting.ScanEventHandler
	} else {
		this.ScanEventHandler = defaultScanEventHandler
	}
	if setting.LocationEventHandler != nil {
		this.LocationEventHandler = setting.LocationEventHandler
	} else {
		this.LocationEventHandler = defaultLocationEventHandler
	}
	if setting.MenuClickEventHandler != nil {
		this.MenuClickEventHandler = setting.MenuClickEventHandler
	} else {
		this.MenuClickEventHandler = defaultMenuClickEventHandler
	}
	if setting.MenuViewEventHandler != nil {
		this.MenuViewEventHandler = setting.MenuViewEventHandler
	} else {
		this.MenuViewEventHandler = defaultMenuViewEventHandler
	}
	if setting.MenuScanCodePushEventHandler != nil {
		this.MenuScanCodePushEventHandler = setting.MenuScanCodePushEventHandler
	} else {
		this.MenuScanCodePushEventHandler = defaultMenuScanCodePushEventHandler
	}
	if setting.MenuScanCodeWaitMsgEventHandler != nil {
		this.MenuScanCodeWaitMsgEventHandler = setting.MenuScanCodeWaitMsgEventHandler
	} else {
		this.MenuScanCodeWaitMsgEventHandler = defaultMenuScanCodeWaitMsgEventHandler
	}
	if setting.MenuPicSysPhotoEventHandler != nil {
		this.MenuPicSysPhotoEventHandler = setting.MenuPicSysPhotoEventHandler
	} else {
		this.MenuPicSysPhotoEventHandler = defaultMenuPicSysPhotoEventHandler
	}
	if setting.MenuPicPhotoOrAlbumEventHandler != nil {
		this.MenuPicPhotoOrAlbumEventHandler = setting.MenuPicPhotoOrAlbumEventHandler
	} else {
		this.MenuPicPhotoOrAlbumEventHandler = defaultMenuPicPhotoOrAlbumEventHandler
	}
	if setting.MenuPicWeixinEventHandler != nil {
		this.MenuPicWeixinEventHandler = setting.MenuPicWeixinEventHandler
	} else {
		this.MenuPicWeixinEventHandler = defaultMenuPicWeixinEventHandler
	}
	if setting.MenuLocationSelectEventHandler != nil {
		this.MenuLocationSelectEventHandler = setting.MenuLocationSelectEventHandler
	} else {
		this.MenuLocationSelectEventHandler = defaultMenuLocationSelectEventHandler
	}
	if setting.MassSendJobFinishEventHandler != nil {
		this.MassSendJobFinishEventHandler = setting.MassSendJobFinishEventHandler
	} else {
		this.MassSendJobFinishEventHandler = defaultMassSendJobFinishEventHandler
	}
	if setting.TemplateSendJobFinishEventHandler != nil {
		this.TemplateSendJobFinishEventHandler = setting.TemplateSendJobFinishEventHandler
	} else {
		this.TemplateSendJobFinishEventHandler = defaultTemplateSendJobFinishEventHandler
	}
	if setting.MerchantOrderEventHandler != nil {
		this.MerchantOrderEventHandler = setting.MerchantOrderEventHandler
	} else {
		this.MerchantOrderEventHandler = defaultMerchantOrderEventHandler
	}
}

// 默认的消息处理函数是什么都不做
func defaultInvalidRequestHandler(http.ResponseWriter, *http.Request, error)  {}
func defaultUnknownRequestHandler(http.ResponseWriter, *http.Request, []byte) {}

func defaultTextRequestHandler(http.ResponseWriter, *http.Request, *request.Text)                  {}
func defaultImageRequestHandler(http.ResponseWriter, *http.Request, *request.Image)                {}
func defaultVoiceRequestHandler(http.ResponseWriter, *http.Request, *request.Voice)                {}
func defaultVideoRequestHandler(http.ResponseWriter, *http.Request, *request.Video)                {}
func defaultLocationRequestHandler(http.ResponseWriter, *http.Request, *request.Location)          {}
func defaultLinkRequestHandler(http.ResponseWriter, *http.Request, *request.Link)                  {}
func defaultSubscribeEventHandler(http.ResponseWriter, *http.Request, *request.SubscribeEvent)     {}
func defaultUnsubscribeEventHandler(http.ResponseWriter, *http.Request, *request.UnsubscribeEvent) {}
func defaultSubscribeByScanEventHandler(http.ResponseWriter, *http.Request, *request.SubscribeByScanEvent) {
}
func defaultScanEventHandler(http.ResponseWriter, *http.Request, *request.ScanEvent)           {}
func defaultLocationEventHandler(http.ResponseWriter, *http.Request, *request.LocationEvent)   {}
func defaultMenuClickEventHandler(http.ResponseWriter, *http.Request, *request.MenuClickEvent) {}
func defaultMenuViewEventHandler(http.ResponseWriter, *http.Request, *request.MenuViewEvent)   {}
func defaultMenuScanCodePushEventHandler(http.ResponseWriter, *http.Request, *request.MenuScanCodePushEvent) {
}
func defaultMenuScanCodeWaitMsgEventHandler(http.ResponseWriter, *http.Request, *request.MenuScanCodeWaitMsgEvent) {
}
func defaultMenuPicSysPhotoEventHandler(http.ResponseWriter, *http.Request, *request.MenuPicSysPhotoEvent) {
}
func defaultMenuPicPhotoOrAlbumEventHandler(http.ResponseWriter, *http.Request, *request.MenuPicPhotoOrAlbumEvent) {
}
func defaultMenuPicWeixinEventHandler(http.ResponseWriter, *http.Request, *request.MenuPicWeixinEvent) {
}
func defaultMenuLocationSelectEventHandler(http.ResponseWriter, *http.Request, *request.MenuLocationSelectEvent) {
}
func defaultMassSendJobFinishEventHandler(http.ResponseWriter, *http.Request, *request.MassSendJobFinishEvent) {
}
func defaultTemplateSendJobFinishEventHandler(http.ResponseWriter, *http.Request, *request.TemplateSendJobFinishEvent) {
}
func defaultMerchantOrderEventHandler(http.ResponseWriter, *http.Request, *request.MerchantOrderEvent) {
}
