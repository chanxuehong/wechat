// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"crypto/sha1"
	"crypto/subtle"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/chanxuehong/wechat/message/passive/request"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// net/http.Handler 的实现
//  NOTE: 非并发安全, 要求注册到 URL 之前全部设置好, 注册之后不能再更改设置了.
type NCSHttpHandler struct {
	invalidRequestHandler InvalidRequestHandler
	msgHandlerMap         map[string]MsgHandler
}

// 设置 InvalidRequestHandler, 如果 handler == nil 则使用默认的 DefaultInvalidRequestHandlerFunc
func (this *NCSHttpHandler) SetInvalidRequestHandler(handler InvalidRequestHandler) {
	if handler == nil {
		this.invalidRequestHandler = InvalidRequestHandlerFunc(DefaultInvalidRequestHandlerFunc)
	} else {
		this.invalidRequestHandler = handler
	}
}

// 添加或设置 WechatMPId 对应的 MsgHandler, 如果 handler == nil 则不做任何操作
func (this *NCSHttpHandler) SetMsgHandler(WechatMPId string, handler MsgHandler) {
	if handler == nil {
		return
	}

	if this.msgHandlerMap == nil {
		this.msgHandlerMap = make(map[string]MsgHandler)
	}
	this.msgHandlerMap[WechatMPId] = handler
}

// 删除 WechatMPId 对应的 MsgHandler
func (this *NCSHttpHandler) DeleteMsgHandler(WechatMPId string) {
	delete(this.msgHandlerMap, WechatMPId)
}

// 清除所有的 MsgHandler
func (this *NCSHttpHandler) ClearMsgHandler() {
	this.msgHandlerMap = nil
}

func (this *NCSHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	invalidRequestHandler := this.invalidRequestHandler
	if invalidRequestHandler == nil {
		invalidRequestHandler = InvalidRequestHandlerFunc(DefaultInvalidRequestHandlerFunc)
	}
	if len(this.msgHandlerMap) == 0 {
		invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("no MsgHandler"))
		return
	}

	switch r.Method {
	case "POST": // 处理从微信服务器推送过来的消息(事件) ==============================
		rawXMLMsg, err := ioutil.ReadAll(r.Body)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		var msgReq request.Request
		if err = xml.Unmarshal(rawXMLMsg, &msgReq); err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		msgHandler := this.msgHandlerMap[msgReq.ToUserName]
		if msgHandler == nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, fmt.Errorf("Not found MsgHandler for WechatMPId: %s", msgReq.ToUserName))
			return
		}

		if r.URL == nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("r.URL == nil"))
			return
		}

		urlValues, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		signature := urlValues.Get("signature")
		if signature == "" {
			invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("signature is empty"))
			return
		}

		const signatureLen = sha1.Size * 2
		if len(signature) != signatureLen {
			invalidRequestHandler.ServeInvalidRequest(w, r, fmt.Errorf("the length of signature mismatch, have: %d, want: %d", len(signature), signatureLen))
			return
		}

		timestampStr := urlValues.Get("timestamp")
		if timestampStr == "" {
			invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("timestamp is empty"))
			return
		}

		timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, fmt.Errorf("can not parse timestamp(==%q) to int64, error: %s", timestampStr, err.Error()))
			return
		}

		nonce := urlValues.Get("nonce")
		if nonce == "" {
			invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("nonce is empty"))
			return
		}

		signaturex := msgHandler.Signature(timestampStr, nonce)
		if subtle.ConstantTimeCompare([]byte(signature), []byte(signaturex)) != 1 {
			invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("check signature failed"))
			return
		}

		// request router, 可一个根据自己的实际业务调整顺序!
		switch msgReq.MsgType {
		case request.MSG_TYPE_TEXT:
			msgHandler.TextMsgHandler(w, r, msgReq.Text(), rawXMLMsg, timestamp)

		case request.MSG_TYPE_EVENT:
			// event router
			switch msgReq.Event {
			case request.EVENT_TYPE_LOCATION:
				msgHandler.LocationEventHandler(w, r, msgReq.LocationEvent(), rawXMLMsg, timestamp)

			case request.EVENT_TYPE_CLICK:
				msgHandler.MenuClickEventHandler(w, r, msgReq.MenuClickEvent(), rawXMLMsg, timestamp)

			case request.EVENT_TYPE_VIEW:
				msgHandler.MenuViewEventHandler(w, r, msgReq.MenuViewEvent(), rawXMLMsg, timestamp)

			case request.EVENT_TYPE_SCANCODE_PUSH:
				msgHandler.MenuScanCodePushEventHandler(w, r, msgReq.MenuScanCodePushEvent(), rawXMLMsg, timestamp)

			case request.EVENT_TYPE_SCANCODE_WAITMSG:
				msgHandler.MenuScanCodeWaitMsgEventHandler(w, r, msgReq.MenuScanCodeWaitMsgEvent(), rawXMLMsg, timestamp)

			case request.EVENT_TYPE_PIC_SYSPHOTO:
				msgHandler.MenuPicSysPhotoEventHandler(w, r, msgReq.MenuPicSysPhotoEvent(), rawXMLMsg, timestamp)

			case request.EVENT_TYPE_PIC_PHOTO_OR_ALBUM:
				msgHandler.MenuPicPhotoOrAlbumEventHandler(w, r, msgReq.MenuPicPhotoOrAlbumEvent(), rawXMLMsg, timestamp)

			case request.EVENT_TYPE_PIC_WEIXIN:
				msgHandler.MenuPicWeixinEventHandler(w, r, msgReq.MenuPicWeixinEvent(), rawXMLMsg, timestamp)

			case request.EVENT_TYPE_LOCATION_SELECT:
				msgHandler.MenuLocationSelectEventHandler(w, r, msgReq.MenuLocationSelectEvent(), rawXMLMsg, timestamp)

			case request.EVENT_TYPE_TEMPLATESENDJOBFINISH:
				msgHandler.TemplateSendJobFinishEventHandler(w, r, msgReq.TemplateSendJobFinishEvent(), rawXMLMsg, timestamp)

			case request.EVENT_TYPE_MASSSENDJOBFINISH:
				msgHandler.MassSendJobFinishEventHandler(w, r, msgReq.MassSendJobFinishEvent(), rawXMLMsg, timestamp)

			case request.EVENT_TYPE_MERCHANTORDER:
				msgHandler.MerchantOrderEventHandler(w, r, msgReq.MerchantOrderEvent(), rawXMLMsg, timestamp)

			case request.EVENT_TYPE_SUBSCRIBE:
				if msgReq.Ticket == "" { // 普通订阅
					msgHandler.SubscribeEventHandler(w, r, msgReq.SubscribeEvent(), rawXMLMsg, timestamp)
				} else { // 扫描二维码订阅
					msgHandler.SubscribeByScanEventHandler(w, r, msgReq.SubscribeByScanEvent(), rawXMLMsg, timestamp)
				}

			case request.EVENT_TYPE_UNSUBSCRIBE:
				msgHandler.UnsubscribeEventHandler(w, r, msgReq.UnsubscribeEvent(), rawXMLMsg, timestamp)

			case request.EVENT_TYPE_SCAN:
				msgHandler.ScanEventHandler(w, r, msgReq.ScanEvent(), rawXMLMsg, timestamp)

			default: // unknown event
				msgHandler.UnknownMsgHandler(w, r, rawXMLMsg, timestamp)
			}

		case request.MSG_TYPE_LINK:
			msgHandler.LinkMsgHandler(w, r, msgReq.Link(), rawXMLMsg, timestamp)

		case request.MSG_TYPE_VOICE:
			msgHandler.VoiceMsgHandler(w, r, msgReq.Voice(), rawXMLMsg, timestamp)

		case request.MSG_TYPE_LOCATION:
			msgHandler.LocationMsgHandler(w, r, msgReq.Location(), rawXMLMsg, timestamp)

		case request.MSG_TYPE_IMAGE:
			msgHandler.ImageMsgHandler(w, r, msgReq.Image(), rawXMLMsg, timestamp)

		case request.MSG_TYPE_VIDEO:
			msgHandler.VideoMsgHandler(w, r, msgReq.Video(), rawXMLMsg, timestamp)

		default: // unknown request message type
			msgHandler.UnknownMsgHandler(w, r, rawXMLMsg, timestamp)
		}

	case "GET": // 首次验证 ======================================================
		if r.URL == nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("r.URL == nil"))
			return
		}

		urlValues, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		signature := urlValues.Get("signature")
		if signature == "" {
			invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("signature is empty"))
			return
		}

		const signatureLen = sha1.Size * 2
		if len(signature) != signatureLen {
			invalidRequestHandler.ServeInvalidRequest(w, r, fmt.Errorf("the length of signature mismatch, have: %d, want: %d", len(signature), signatureLen))
			return
		}

		timestamp := urlValues.Get("timestamp")
		if timestamp == "" {
			invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("timestamp is empty"))
			return
		}

		nonce := urlValues.Get("nonce")
		if nonce == "" {
			invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("nonce is empty"))
			return
		}

		echostr := urlValues.Get("echostr")
		if echostr == "" {
			invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("echostr is empty"))
			return
		}

		for _, msgHandler := range this.msgHandlerMap {
			signaturex := msgHandler.Signature(timestamp, nonce)
			if subtle.ConstantTimeCompare([]byte(signature), []byte(signaturex)) != 1 {
				continue
			}

			io.WriteString(w, echostr)
			return
		}

		// 所有的 MsgHandler 都不能验证
		invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("check signature failed"))
	}
}
