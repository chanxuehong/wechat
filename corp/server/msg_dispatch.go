// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import "github.com/chanxuehong/wechat/corp/message/passive/request"

// 消息（事件）分路器, 可以根据实际业务来调整顺序!
func msgDispatch(agent Agent, msg *request.Request, para *RequestParameters) {
	switch msg.MsgType {
	case request.MSG_TYPE_TEXT:
		agent.ServeTextMsg(msg.Text(), para)

	case request.MSG_TYPE_EVENT:
		switch msg.Event {
		case request.EVENT_TYPE_LOCATION:
			agent.ServeLocationEvent(msg.LocationEvent(), para)

		case request.EVENT_TYPE_CLICK:
			agent.ServeMenuClickEvent(msg.MenuClickEvent(), para)

		case request.EVENT_TYPE_SCANCODE_PUSH:
			agent.ServeMenuScanCodePushEvent(msg.MenuScanCodePushEvent(), para)

		case request.EVENT_TYPE_SCANCODE_WAITMSG:
			agent.ServeMenuScanCodeWaitMsgEvent(msg.MenuScanCodeWaitMsgEvent(), para)

		case request.EVENT_TYPE_PIC_SYSPHOTO:
			agent.ServeMenuPicSysPhotoEvent(msg.MenuPicSysPhotoEvent(), para)

		case request.EVENT_TYPE_PIC_PHOTO_OR_ALBUM:
			agent.ServeMenuPicPhotoOrAlbumEvent(msg.MenuPicPhotoOrAlbumEvent(), para)

		case request.EVENT_TYPE_PIC_WEIXIN:
			agent.ServeMenuPicWeixinEvent(msg.MenuPicWeixinEvent(), para)

		case request.EVENT_TYPE_LOCATION_SELECT:
			agent.ServeMenuLocationSelectEvent(msg.MenuLocationSelectEvent(), para)

		case request.EVENT_TYPE_VIEW:
			agent.ServeMenuViewEvent(msg.MenuViewEvent(), para)

		case request.EVENT_TYPE_SUBSCRIBE:
			agent.ServeSubscribeEvent(msg.SubscribeEvent(), para)

		case request.EVENT_TYPE_UNSUBSCRIBE:
			agent.ServeUnsubscribeEvent(msg.UnsubscribeEvent(), para)

		default: // unknown event type
			agent.ServeUnknownMsg(para)
		}

	case request.MSG_TYPE_VOICE:
		agent.ServeVoiceMsg(msg.Voice(), para)

	case request.MSG_TYPE_LOCATION:
		agent.ServeLocationMsg(msg.Location(), para)

	case request.MSG_TYPE_IMAGE:
		agent.ServeImageMsg(msg.Image(), para)

	case request.MSG_TYPE_VIDEO:
		agent.ServeVideoMsg(msg.Video(), para)

	default: // unknown message type
		agent.ServeUnknownMsg(para)
	}
}
