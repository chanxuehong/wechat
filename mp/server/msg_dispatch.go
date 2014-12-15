// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import "github.com/chanxuehong/wechat/mp/message/passive/request"

// 明文模式
// 消息（事件）分路器, 可以根据实际业务来调整顺序!
func rawMsgDispatch(agent Agent, mixedRequest *request.Request, para *RequestParameters) {
	switch mixedRequest.MsgType {
	case request.MSG_TYPE_TEXT:
		agent.ServeTextMsg(mixedRequest.Text(), para)

	case request.MSG_TYPE_EVENT:
		switch mixedRequest.Event {
		case request.EVENT_TYPE_LOCATION:
			agent.ServeLocationEvent(mixedRequest.LocationEvent(), para)

		case request.EVENT_TYPE_CLICK:
			agent.ServeMenuClickEvent(mixedRequest.MenuClickEvent(), para)

		case request.EVENT_TYPE_VIEW:
			agent.ServeMenuViewEvent(mixedRequest.MenuViewEvent(), para)

		case request.EVENT_TYPE_SCANCODE_PUSH:
			agent.ServeMenuScanCodePushEvent(mixedRequest.MenuScanCodePushEvent(), para)

		case request.EVENT_TYPE_SCANCODE_WAITMSG:
			agent.ServeMenuScanCodeWaitMsgEvent(mixedRequest.MenuScanCodeWaitMsgEvent(), para)

		case request.EVENT_TYPE_PIC_SYSPHOTO:
			agent.ServeMenuPicSysPhotoEvent(mixedRequest.MenuPicSysPhotoEvent(), para)

		case request.EVENT_TYPE_PIC_PHOTO_OR_ALBUM:
			agent.ServeMenuPicPhotoOrAlbumEvent(mixedRequest.MenuPicPhotoOrAlbumEvent(), para)

		case request.EVENT_TYPE_PIC_WEIXIN:
			agent.ServeMenuPicWeixinEvent(mixedRequest.MenuPicWeixinEvent(), para)

		case request.EVENT_TYPE_LOCATION_SELECT:
			agent.ServeMenuLocationSelectEvent(mixedRequest.MenuLocationSelectEvent(), para)

		case request.EVENT_TYPE_TEMPLATESENDJOBFINISH:
			agent.ServeTemplateSendJobFinishEvent(mixedRequest.TemplateSendJobFinishEvent(), para)

		case request.EVENT_TYPE_MASSSENDJOBFINISH:
			agent.ServeMassSendJobFinishEvent(mixedRequest.MassSendJobFinishEvent(), para)

		case request.EVENT_TYPE_MERCHANTORDER:
			agent.ServeMerchantOrderEvent(mixedRequest.MerchantOrderEvent(), para)

		case request.EVENT_TYPE_SUBSCRIBE:
			if mixedRequest.Ticket == "" { // 普通订阅
				agent.ServeSubscribeEvent(mixedRequest.SubscribeEvent(), para)
			} else { // 扫描二维码订阅
				agent.ServeSubscribeByScanEvent(mixedRequest.SubscribeByScanEvent(), para)
			}

		case request.EVENT_TYPE_UNSUBSCRIBE:
			agent.ServeUnsubscribeEvent(mixedRequest.UnsubscribeEvent(), para)

		case request.EVENT_TYPE_SCAN:
			agent.ServeScanEvent(mixedRequest.ScanEvent(), para)

		default: // unknown event type
			agent.ServeUnknownMsg(para)
		}

	case request.MSG_TYPE_LINK:
		agent.ServeLinkMsg(mixedRequest.Link(), para)

	case request.MSG_TYPE_VOICE:
		agent.ServeVoiceMsg(mixedRequest.Voice(), para)

	case request.MSG_TYPE_LOCATION:
		agent.ServeLocationMsg(mixedRequest.Location(), para)

	case request.MSG_TYPE_IMAGE:
		agent.ServeImageMsg(mixedRequest.Image(), para)

	case request.MSG_TYPE_VIDEO:
		agent.ServeVideoMsg(mixedRequest.Video(), para)

	default: // unknown message type
		agent.ServeUnknownMsg(para)
	}
}

// 兼容模式, 安全模式
// 消息（事件）分路器, 可以根据实际业务来调整顺序!
func aesMsgDispatch(agent Agent, mixedRequest *request.Request, para *AESRequestParameters) {
	switch mixedRequest.MsgType {
	case request.MSG_TYPE_TEXT:
		agent.ServeAESTextMsg(mixedRequest.Text(), para)

	case request.MSG_TYPE_EVENT:
		switch mixedRequest.Event {
		case request.EVENT_TYPE_LOCATION:
			agent.ServeAESLocationEvent(mixedRequest.LocationEvent(), para)

		case request.EVENT_TYPE_CLICK:
			agent.ServeAESMenuClickEvent(mixedRequest.MenuClickEvent(), para)

		case request.EVENT_TYPE_VIEW:
			agent.ServeAESMenuViewEvent(mixedRequest.MenuViewEvent(), para)

		case request.EVENT_TYPE_SCANCODE_PUSH:
			agent.ServeAESMenuScanCodePushEvent(mixedRequest.MenuScanCodePushEvent(), para)

		case request.EVENT_TYPE_SCANCODE_WAITMSG:
			agent.ServeAESMenuScanCodeWaitMsgEvent(mixedRequest.MenuScanCodeWaitMsgEvent(), para)

		case request.EVENT_TYPE_PIC_SYSPHOTO:
			agent.ServeAESMenuPicSysPhotoEvent(mixedRequest.MenuPicSysPhotoEvent(), para)

		case request.EVENT_TYPE_PIC_PHOTO_OR_ALBUM:
			agent.ServeAESMenuPicPhotoOrAlbumEvent(mixedRequest.MenuPicPhotoOrAlbumEvent(), para)

		case request.EVENT_TYPE_PIC_WEIXIN:
			agent.ServeAESMenuPicWeixinEvent(mixedRequest.MenuPicWeixinEvent(), para)

		case request.EVENT_TYPE_LOCATION_SELECT:
			agent.ServeAESMenuLocationSelectEvent(mixedRequest.MenuLocationSelectEvent(), para)

		case request.EVENT_TYPE_TEMPLATESENDJOBFINISH:
			agent.ServeAESTemplateSendJobFinishEvent(mixedRequest.TemplateSendJobFinishEvent(), para)

		case request.EVENT_TYPE_MASSSENDJOBFINISH:
			agent.ServeAESMassSendJobFinishEvent(mixedRequest.MassSendJobFinishEvent(), para)

		case request.EVENT_TYPE_MERCHANTORDER:
			agent.ServeAESMerchantOrderEvent(mixedRequest.MerchantOrderEvent(), para)

		case request.EVENT_TYPE_SUBSCRIBE:
			if mixedRequest.Ticket == "" { // 普通订阅
				agent.ServeAESSubscribeEvent(mixedRequest.SubscribeEvent(), para)
			} else { // 扫描二维码订阅
				agent.ServeAESSubscribeByScanEvent(mixedRequest.SubscribeByScanEvent(), para)
			}

		case request.EVENT_TYPE_UNSUBSCRIBE:
			agent.ServeAESUnsubscribeEvent(mixedRequest.UnsubscribeEvent(), para)

		case request.EVENT_TYPE_SCAN:
			agent.ServeAESScanEvent(mixedRequest.ScanEvent(), para)

		default: // unknown event type
			agent.ServeAESUnknownMsg(para)
		}

	case request.MSG_TYPE_LINK:
		agent.ServeAESLinkMsg(mixedRequest.Link(), para)

	case request.MSG_TYPE_VOICE:
		agent.ServeAESVoiceMsg(mixedRequest.Voice(), para)

	case request.MSG_TYPE_LOCATION:
		agent.ServeAESLocationMsg(mixedRequest.Location(), para)

	case request.MSG_TYPE_IMAGE:
		agent.ServeAESImageMsg(mixedRequest.Image(), para)

	case request.MSG_TYPE_VIDEO:
		agent.ServeAESVideoMsg(mixedRequest.Video(), para)

	default: // unknown message type
		agent.ServeAESUnknownMsg(para)
	}
}
