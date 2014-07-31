// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"encoding/xml"
	"errors"
	"github.com/chanxuehong/wechat/message/request"
	"io"
	"net/http"
	"net/url"
	"sync"
)

// 对于公众号开发模式, 都会要求提供一个 URL 来处理微信服务器推送过来的消息和事件,
// Handler 就是处理推送到这个 URL 上的消息(事件).
//  Handler 实现了 http.Handler 接口, 使用时把 Handler 绑定到 URL 的 path 上即可;
//  Handler 并发安全.
type Handler struct {
	setting HandlerSetting

	// 对于微信服务器推送过来的请求, 处理过程中有些中间状态比较大的变量, 所以可以缓存起来.
	//  NOTE: require go1.3+ , 如果你的环境不满足这个条件, 可以自己实现一个简单的 Pool,
	//        see github.com/chanxuehong/util/pool
	bufferUnitPool sync.Pool
}

func NewHandler(setting *HandlerSetting) (handler *Handler) {
	if setting == nil {
		panic("setting == nil")
	}

	handler = &Handler{
		bufferUnitPool: sync.Pool{New: newBufferUnit},
	}
	handler.setting.initialize(setting)

	return
}

// Handler 实现 http.Handler 接口
func (handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST": // 处理从微信服务器推送过来的消息(事件) ==============================
		var urlValues url.Values
		var signature, timestamp, nonce string
		var err error

		if r.URL == nil {
			handler.setting.InvalidRequestHandler(w, r, errors.New("r.URL == nil"))
			return
		}
		if urlValues, err = url.ParseQuery(r.URL.RawQuery); err != nil {
			handler.setting.InvalidRequestHandler(w, r, err)
			return
		}

		if signature = urlValues.Get("signature"); signature == "" {
			handler.setting.InvalidRequestHandler(w, r, errors.New("signature is empty"))
			return
		}
		if timestamp = urlValues.Get("timestamp"); timestamp == "" {
			handler.setting.InvalidRequestHandler(w, r, errors.New("timestamp is empty"))
			return
		}
		if nonce = urlValues.Get("nonce"); nonce == "" {
			handler.setting.InvalidRequestHandler(w, r, errors.New("nonce is empty"))
			return
		}

		bufferUnit := handler.getBufferUnitFromPool() // *bufferUnit
		defer handler.putBufferUnitToPool(bufferUnit) // important!

		if !checkSignature(signature, timestamp, nonce, handler.setting.Token, bufferUnit.signatureBuf[:]) {
			handler.setting.InvalidRequestHandler(w, r, errors.New("check signature failed"))
			return
		}

		if _, err = io.Copy(bufferUnit.msgBuf, r.Body); err != nil {
			handler.setting.InvalidRequestHandler(w, r, err)
			return
		}

		msgRqstBody := bufferUnit.msgBuf.Bytes()
		msgRqst := &bufferUnit.msgRequest // & 不能丢
		if err = xml.Unmarshal(msgRqstBody, msgRqst); err != nil {
			handler.setting.InvalidRequestHandler(w, r, err)
			return
		}

		// request router, 可一个根据自己的实际业务调整顺序!
		switch msgRqst.MsgType {
		case request.MSG_TYPE_TEXT:
			handler.setting.TextRequestHandler(w, r, msgRqst.Text())

		case request.MSG_TYPE_EVENT:
			// event router
			switch msgRqst.Event {
			case request.EVENT_TYPE_CLICK:
				handler.setting.MenuClickEventHandler(w, r, msgRqst.MenuClickEvent())

			case request.EVENT_TYPE_VIEW:
				handler.setting.MenuViewEventHandler(w, r, msgRqst.MenuViewEvent())

			case request.EVENT_TYPE_LOCATION:
				handler.setting.LocationEventHandler(w, r, msgRqst.LocationEvent())

			case request.EVENT_TYPE_MERCHANTORDER:
				handler.setting.MerchantOrderEventHandler(w, r, msgRqst.MerchantOrderEvent())

			case request.EVENT_TYPE_SUBSCRIBE:
				if msgRqst.Ticket == "" { // 普通订阅
					handler.setting.SubscribeEventHandler(w, r, msgRqst.SubscribeEvent())
				} else { // 扫描二维码订阅
					handler.setting.SubscribeByScanEventHandler(w, r, msgRqst.SubscribeByScanEvent())
				}

			case request.EVENT_TYPE_UNSUBSCRIBE:
				handler.setting.UnsubscribeEventHandler(w, r, msgRqst.UnsubscribeEvent())

			case request.EVENT_TYPE_SCAN:
				handler.setting.ScanEventHandler(w, r, msgRqst.ScanEvent())

			case request.EVENT_TYPE_MASSSENDJOBFINISH:
				handler.setting.MassSendJobFinishEventHandler(w, r, msgRqst.MassSendJobFinishEvent())

			default: // unknown event
				// 因为 msgRqstBody 底层需要缓存, 所以这里需要一个副本
				msgRqstBodyCopy := make([]byte, len(msgRqstBody))
				copy(msgRqstBodyCopy, msgRqstBody)
				handler.setting.UnknownRequestHandler(w, r, msgRqstBodyCopy)
			}

		case request.MSG_TYPE_LINK:
			handler.setting.LinkRequestHandler(w, r, msgRqst.Link())

		case request.MSG_TYPE_VOICE:
			handler.setting.VoiceRequestHandler(w, r, msgRqst.Voice())

		case request.MSG_TYPE_LOCATION:
			handler.setting.LocationRequestHandler(w, r, msgRqst.Location())

		case request.MSG_TYPE_IMAGE:
			handler.setting.ImageRequestHandler(w, r, msgRqst.Image())

		case request.MSG_TYPE_VIDEO:
			handler.setting.VideoRequestHandler(w, r, msgRqst.Video())

		default: // unknown request message type
			// 因为 msgRqstBody 底层需要缓存, 所以这里需要一个副本
			msgRqstBodyCopy := make([]byte, len(msgRqstBody))
			copy(msgRqstBodyCopy, msgRqstBody)
			handler.setting.UnknownRequestHandler(w, r, msgRqstBodyCopy)
		}

	case "GET": // 首次验证 ======================================================
		var urlValues url.Values
		var signature, timestamp, nonce, echostr string
		var err error

		if r.URL == nil {
			handler.setting.InvalidRequestHandler(w, r, errors.New("r.URL == nil"))
			return
		}
		if urlValues, err = url.ParseQuery(r.URL.RawQuery); err != nil {
			handler.setting.InvalidRequestHandler(w, r, err)
			return
		}

		if signature = urlValues.Get("signature"); signature == "" {
			handler.setting.InvalidRequestHandler(w, r, errors.New("signature is empty"))
			return
		}
		if timestamp = urlValues.Get("timestamp"); timestamp == "" {
			handler.setting.InvalidRequestHandler(w, r, errors.New("timestamp is empty"))
			return
		}
		if nonce = urlValues.Get("nonce"); nonce == "" {
			handler.setting.InvalidRequestHandler(w, r, errors.New("nonce is empty"))
			return
		}
		if echostr = urlValues.Get("echostr"); echostr == "" {
			handler.setting.InvalidRequestHandler(w, r, errors.New("echostr is empty"))
			return
		}

		bufferUnit := handler.getBufferUnitFromPool() // *bufferUnit
		defer handler.putBufferUnitToPool(bufferUnit) // important!

		if !checkSignature(signature, timestamp, nonce, handler.setting.Token, bufferUnit.signatureBuf[:]) {
			handler.setting.InvalidRequestHandler(w, r, errors.New("check signature failed"))
			return
		}

		io.WriteString(w, echostr)
	}
}
