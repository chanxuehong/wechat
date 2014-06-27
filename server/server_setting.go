// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong@gmail.com

package server

import (
	"github.com/chanxuehong/wechat/message/request"
	"net/http"
)

// 非法请求的处理函数.
// @err: 具体的错误信息
type InvalidRequestHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)

// 未知消息类型的消息处理函数.
// @msg: 接收到的消息体
type UnknownRequestHandlerFunc func(w http.ResponseWriter, r *http.Request, msg []byte)

// 正常的消息处理函数
type TextRequestHandlerFunc func(http.ResponseWriter, *http.Request, *request.Text)
type ImageRequestHandlerFunc func(http.ResponseWriter, *http.Request, *request.Image)
type VoiceRequestHandlerFunc func(http.ResponseWriter, *http.Request, *request.Voice)
type VoiceRecognitionRequestHandlerFunc func(http.ResponseWriter, *http.Request, *request.VoiceRecognition)
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
type MassSendJobFinishEventHandlerFunc func(http.ResponseWriter, *http.Request, *request.MassSendJobFinishEvent)
type MerchantOrderEventHandlerFunc func(http.ResponseWriter, *http.Request, *request.MerchantOrderEvent)

type ServerSetting struct {
	Token string

	// Invalid request handler
	InvalidRequestHandler InvalidRequestHandlerFunc

	// unknown request handler
	UnknownRequestHandler UnknownRequestHandlerFunc

	// common request handler
	TextRequestHandler             TextRequestHandlerFunc
	ImageRequestHandler            ImageRequestHandlerFunc
	VoiceRequestHandler            VoiceRequestHandlerFunc
	VoiceRecognitionRequestHandler VoiceRecognitionRequestHandlerFunc
	VideoRequestHandler            VideoRequestHandlerFunc
	LocationRequestHandler         LocationRequestHandlerFunc
	LinkRequestHandler             LinkRequestHandlerFunc
	// event handler
	SubscribeEventHandler         SubscribeEventHandlerFunc
	UnsubscribeEventHandler       UnsubscribeEventHandlerFunc
	SubscribeByScanEventHandler   SubscribeByScanEventHandlerFunc
	ScanEventHandler              ScanEventHandlerFunc
	LocationEventHandler          LocationEventHandlerFunc
	MenuClickEventHandler         MenuClickEventHandlerFunc
	MenuViewEventHandler          MenuViewEventHandlerFunc
	MassSendJobFinishEventHandler MassSendJobFinishEventHandlerFunc
	MerchantOrderEventHandler     MerchantOrderEventHandlerFunc
}

// 根据另外一个 ServerSetting 来初始化.
// 没有设置的函数将会被初始化为默认处理函数.
//  NOTE: 确保 setting != nil
func (ss *ServerSetting) initialize(setting *ServerSetting) {
	ss.Token = setting.Token

	if setting.InvalidRequestHandler != nil {
		ss.InvalidRequestHandler = setting.InvalidRequestHandler
	} else {
		ss.InvalidRequestHandler = defaultInvalidRequestHandler
	}

	if setting.UnknownRequestHandler != nil {
		ss.UnknownRequestHandler = setting.UnknownRequestHandler
	} else {
		ss.UnknownRequestHandler = defaultUnknownRequestHandler
	}

	// common request handler
	if setting.TextRequestHandler != nil {
		ss.TextRequestHandler = setting.TextRequestHandler
	} else {
		ss.TextRequestHandler = defaultTextRequestHandler
	}
	if setting.ImageRequestHandler != nil {
		ss.ImageRequestHandler = setting.ImageRequestHandler
	} else {
		ss.ImageRequestHandler = defaultImageRequestHandler
	}
	if setting.VoiceRequestHandler != nil {
		ss.VoiceRequestHandler = setting.VoiceRequestHandler
	} else {
		ss.VoiceRequestHandler = defaultVoiceRequestHandler
	}
	if setting.VoiceRecognitionRequestHandler != nil {
		ss.VoiceRecognitionRequestHandler = setting.VoiceRecognitionRequestHandler
	} else {
		ss.VoiceRecognitionRequestHandler = defaultVoiceRecognitionRequestHandler
	}
	if setting.VideoRequestHandler != nil {
		ss.VideoRequestHandler = setting.VideoRequestHandler
	} else {
		ss.VideoRequestHandler = defaultVideoRequestHandler
	}
	if setting.LocationRequestHandler != nil {
		ss.LocationRequestHandler = setting.LocationRequestHandler
	} else {
		ss.LocationRequestHandler = defaultLocationRequestHandler
	}
	if setting.LinkRequestHandler != nil {
		ss.LinkRequestHandler = setting.LinkRequestHandler
	} else {
		ss.LinkRequestHandler = defaultLinkRequestHandler
	}

	// event handler
	if setting.SubscribeEventHandler != nil {
		ss.SubscribeEventHandler = setting.SubscribeEventHandler
	} else {
		ss.SubscribeEventHandler = defaultSubscribeEventHandler
	}
	if setting.UnsubscribeEventHandler != nil {
		ss.UnsubscribeEventHandler = setting.UnsubscribeEventHandler
	} else {
		ss.UnsubscribeEventHandler = defaultUnsubscribeEventHandler
	}
	if setting.SubscribeByScanEventHandler != nil {
		ss.SubscribeByScanEventHandler = setting.SubscribeByScanEventHandler
	} else {
		ss.SubscribeByScanEventHandler = defaultSubscribeByScanEventHandler
	}
	if setting.ScanEventHandler != nil {
		ss.ScanEventHandler = setting.ScanEventHandler
	} else {
		ss.ScanEventHandler = defaultScanEventHandler
	}
	if setting.LocationEventHandler != nil {
		ss.LocationEventHandler = setting.LocationEventHandler
	} else {
		ss.LocationEventHandler = defaultLocationEventHandler
	}
	if setting.MenuClickEventHandler != nil {
		ss.MenuClickEventHandler = setting.MenuClickEventHandler
	} else {
		ss.MenuClickEventHandler = defaultMenuClickEventHandler
	}
	if setting.MenuViewEventHandler != nil {
		ss.MenuViewEventHandler = setting.MenuViewEventHandler
	} else {
		ss.MenuViewEventHandler = defaultMenuViewEventHandler
	}
	if setting.MassSendJobFinishEventHandler != nil {
		ss.MassSendJobFinishEventHandler = setting.MassSendJobFinishEventHandler
	} else {
		ss.MassSendJobFinishEventHandler = defaultMassSendJobFinishEventHandler
	}
	if setting.MerchantOrderEventHandler != nil {
		ss.MerchantOrderEventHandler = setting.MerchantOrderEventHandler
	} else {
		ss.MerchantOrderEventHandler = defaultMerchantOrderEventHandler
	}
}

// 默认的消息处理函数是什么都不做
func defaultInvalidRequestHandler(http.ResponseWriter, *http.Request, error)  {}
func defaultUnknownRequestHandler(http.ResponseWriter, *http.Request, []byte) {}

func defaultTextRequestHandler(http.ResponseWriter, *http.Request, *request.Text)   {}
func defaultImageRequestHandler(http.ResponseWriter, *http.Request, *request.Image) {}
func defaultVoiceRequestHandler(http.ResponseWriter, *http.Request, *request.Voice) {}
func defaultVoiceRecognitionRequestHandler(http.ResponseWriter, *http.Request, *request.VoiceRecognition) {
}
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
func defaultMassSendJobFinishEventHandler(http.ResponseWriter, *http.Request, *request.MassSendJobFinishEvent) {
}
func defaultMerchantOrderEventHandler(http.ResponseWriter, *http.Request, *request.MerchantOrderEvent) {
}
