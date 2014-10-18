// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"github.com/chanxuehong/wechat/corp/message/passive/request"
	"net/http"
)

// 消息（事件）分路器, 可以根据实际业务来调整顺序!
func msgDispatch(w http.ResponseWriter, r *http.Request, msg *request.Request, rawXMLMsg []byte, timestamp int64, nonce string, random [16]byte, agent Agent) {
	switch msg.MsgType {
	case request.MSG_TYPE_TEXT:
		agent.ServeTextMsg(w, r, msg.Text(), rawXMLMsg, timestamp, nonce, random)

	case request.MSG_TYPE_EVENT:
		switch msg.Event {
		case request.EVENT_TYPE_LOCATION:
			agent.ServeLocationEvent(w, r, msg.LocationEvent(), rawXMLMsg, timestamp, nonce, random)

		case request.EVENT_TYPE_CLICK:
			agent.ServeMenuClickEvent(w, r, msg.MenuClickEvent(), rawXMLMsg, timestamp, nonce, random)

		case request.EVENT_TYPE_VIEW:
			agent.ServeMenuViewEvent(w, r, msg.MenuViewEvent(), rawXMLMsg, timestamp, nonce, random)

		case request.EVENT_TYPE_SUBSCRIBE:
			agent.ServeSubscribeEvent(w, r, msg.SubscribeEvent(), rawXMLMsg, timestamp, nonce, random)

		case request.EVENT_TYPE_UNSUBSCRIBE:
			agent.ServeUnsubscribeEvent(w, r, msg.UnsubscribeEvent(), rawXMLMsg, timestamp, nonce, random)

		default: // unknown event type
			agent.ServeUnknownMsg(w, r, rawXMLMsg, timestamp, nonce, random)
		}

	case request.MSG_TYPE_VOICE:
		agent.ServeVoiceMsg(w, r, msg.Voice(), rawXMLMsg, timestamp, nonce, random)

	case request.MSG_TYPE_LOCATION:
		agent.ServeLocationMsg(w, r, msg.Location(), rawXMLMsg, timestamp, nonce, random)

	case request.MSG_TYPE_IMAGE:
		agent.ServeImageMsg(w, r, msg.Image(), rawXMLMsg, timestamp, nonce, random)

	case request.MSG_TYPE_VIDEO:
		agent.ServeVideoMsg(w, r, msg.Video(), rawXMLMsg, timestamp, nonce, random)

	default: // unknown message type
		agent.ServeUnknownMsg(w, r, rawXMLMsg, timestamp, nonce, random)
	}
}
