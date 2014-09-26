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
	"github.com/chanxuehong/wechat/message/passive/request"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// net/http.Handler 的实现
type HttpHandler struct {
	MsgHandler MsgHandler
}

func (handler HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST": // 处理从微信服务器推送过来的消息(事件) ==============================
		if r.URL == nil {
			handler.MsgHandler.InvalidRequestHandler(w, r, errors.New("r.URL == nil"))
			return
		}

		urlValues, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			handler.MsgHandler.InvalidRequestHandler(w, r, err)
			return
		}

		signature := urlValues.Get("signature")
		if signature == "" {
			handler.MsgHandler.InvalidRequestHandler(w, r, errors.New("msg_signature is empty"))
			return
		}

		const signatureLen = sha1.Size * 2
		if len(signature) != signatureLen {
			handler.MsgHandler.InvalidRequestHandler(w, r, errors.New("check signature failed"))
			return
		}

		timestampStr := urlValues.Get("timestamp")
		if timestampStr == "" {
			handler.MsgHandler.InvalidRequestHandler(w, r, errors.New("timestamp is empty"))
			return
		}

		timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			handler.MsgHandler.InvalidRequestHandler(w, r, err)
			return
		}

		nonce := urlValues.Get("nonce")
		if nonce == "" {
			handler.MsgHandler.InvalidRequestHandler(w, r, errors.New("nonce is empty"))
			return
		}

		signaturex := handler.MsgHandler.Signature(timestampStr, nonce)
		// 采用 subtle.ConstantTimeCompare 是防止 计时攻击!
		if subtle.ConstantTimeCompare([]byte(signature), []byte(signaturex)) != 1 {
			handler.MsgHandler.InvalidRequestHandler(w, r, errors.New("check signature failed"))
			return
		}

		rawXMLMsg, err := ioutil.ReadAll(r.Body)
		if err != nil {
			handler.MsgHandler.InvalidRequestHandler(w, r, err)
			return
		}

		var msgReq request.Request
		if err = xml.Unmarshal(rawXMLMsg, &msgReq); err != nil {
			handler.MsgHandler.InvalidRequestHandler(w, r, err)
			return
		}

		// request router, 可一个根据自己的实际业务调整顺序!
		switch msgReq.MsgType {
		case request.MSG_TYPE_TEXT:
			handler.MsgHandler.TextMsgHandler(w, r, msgReq.Text(), rawXMLMsg, timestamp)

		case request.MSG_TYPE_EVENT:
			// event router
			switch msgReq.Event {
			case request.EVENT_TYPE_LOCATION:
				handler.MsgHandler.LocationEventHandler(w, r, msgReq.LocationEvent(), rawXMLMsg, timestamp)

			case request.EVENT_TYPE_CLICK:
				handler.MsgHandler.MenuClickEventHandler(w, r, msgReq.MenuClickEvent(), rawXMLMsg, timestamp)

			case request.EVENT_TYPE_VIEW:
				handler.MsgHandler.MenuViewEventHandler(w, r, msgReq.MenuViewEvent(), rawXMLMsg, timestamp)

			case request.EVENT_TYPE_SCANCODE_PUSH:
				handler.MsgHandler.MenuScanCodePushEventHandler(w, r, msgReq.MenuScanCodePushEvent(), rawXMLMsg, timestamp)

			case request.EVENT_TYPE_SCANCODE_WAITMSG:
				handler.MsgHandler.MenuScanCodeWaitMsgEventHandler(w, r, msgReq.MenuScanCodeWaitMsgEvent(), rawXMLMsg, timestamp)

			case request.EVENT_TYPE_PIC_SYSPHOTO:
				handler.MsgHandler.MenuPicSysPhotoEventHandler(w, r, msgReq.MenuPicSysPhotoEvent(), rawXMLMsg, timestamp)

			case request.EVENT_TYPE_PIC_PHOTO_OR_ALBUM:
				handler.MsgHandler.MenuPicPhotoOrAlbumEventHandler(w, r, msgReq.MenuPicPhotoOrAlbumEvent(), rawXMLMsg, timestamp)

			case request.EVENT_TYPE_PIC_WEIXIN:
				handler.MsgHandler.MenuPicWeixinEventHandler(w, r, msgReq.MenuPicWeixinEvent(), rawXMLMsg, timestamp)

			case request.EVENT_TYPE_LOCATION_SELECT:
				handler.MsgHandler.MenuLocationSelectEventHandler(w, r, msgReq.MenuLocationSelectEvent(), rawXMLMsg, timestamp)

			case request.EVENT_TYPE_TEMPLATESENDJOBFINISH:
				handler.MsgHandler.TemplateSendJobFinishEventHandler(w, r, msgReq.TemplateSendJobFinishEvent(), rawXMLMsg, timestamp)

			case request.EVENT_TYPE_MASSSENDJOBFINISH:
				handler.MsgHandler.MassSendJobFinishEventHandler(w, r, msgReq.MassSendJobFinishEvent(), rawXMLMsg, timestamp)

			case request.EVENT_TYPE_MERCHANTORDER:
				handler.MsgHandler.MerchantOrderEventHandler(w, r, msgReq.MerchantOrderEvent(), rawXMLMsg, timestamp)

			case request.EVENT_TYPE_SUBSCRIBE:
				if msgReq.Ticket == "" { // 普通订阅
					handler.MsgHandler.SubscribeEventHandler(w, r, msgReq.SubscribeEvent(), rawXMLMsg, timestamp)
				} else { // 扫描二维码订阅
					handler.MsgHandler.SubscribeByScanEventHandler(w, r, msgReq.SubscribeByScanEvent(), rawXMLMsg, timestamp)
				}

			case request.EVENT_TYPE_UNSUBSCRIBE:
				handler.MsgHandler.UnsubscribeEventHandler(w, r, msgReq.UnsubscribeEvent(), rawXMLMsg, timestamp)

			case request.EVENT_TYPE_SCAN:
				handler.MsgHandler.ScanEventHandler(w, r, msgReq.ScanEvent(), rawXMLMsg, timestamp)

			default: // unknown event
				handler.MsgHandler.UnknownMsgHandler(w, r, rawXMLMsg, timestamp)
			}

		case request.MSG_TYPE_LINK:
			handler.MsgHandler.LinkMsgHandler(w, r, msgReq.Link(), rawXMLMsg, timestamp)

		case request.MSG_TYPE_VOICE:
			handler.MsgHandler.VoiceMsgHandler(w, r, msgReq.Voice(), rawXMLMsg, timestamp)

		case request.MSG_TYPE_LOCATION:
			handler.MsgHandler.LocationMsgHandler(w, r, msgReq.Location(), rawXMLMsg, timestamp)

		case request.MSG_TYPE_IMAGE:
			handler.MsgHandler.ImageMsgHandler(w, r, msgReq.Image(), rawXMLMsg, timestamp)

		case request.MSG_TYPE_VIDEO:
			handler.MsgHandler.VideoMsgHandler(w, r, msgReq.Video(), rawXMLMsg, timestamp)

		default: // unknown request message type
			handler.MsgHandler.UnknownMsgHandler(w, r, rawXMLMsg, timestamp)
		}

	case "GET": // 首次验证 ======================================================
		if r.URL == nil {
			handler.MsgHandler.InvalidRequestHandler(w, r, errors.New("r.URL == nil"))
			return
		}

		urlValues, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			handler.MsgHandler.InvalidRequestHandler(w, r, err)
			return
		}

		signature := urlValues.Get("signature")
		if signature == "" {
			handler.MsgHandler.InvalidRequestHandler(w, r, errors.New("msg_signature is empty"))
			return
		}

		const signatureLen = sha1.Size * 2
		if len(signature) != signatureLen {
			handler.MsgHandler.InvalidRequestHandler(w, r, errors.New("check signature failed"))
			return
		}

		timestamp := urlValues.Get("timestamp")
		if timestamp == "" {
			handler.MsgHandler.InvalidRequestHandler(w, r, errors.New("timestamp is empty"))
			return
		}

		nonce := urlValues.Get("nonce")
		if nonce == "" {
			handler.MsgHandler.InvalidRequestHandler(w, r, errors.New("nonce is empty"))
			return
		}

		echostr := urlValues.Get("echostr")
		if echostr == "" {
			handler.MsgHandler.InvalidRequestHandler(w, r, errors.New("echostr is empty"))
			return
		}

		signaturex := handler.MsgHandler.Signature(timestamp, nonce)
		// 采用 subtle.ConstantTimeCompare 是防止 计时攻击!
		if subtle.ConstantTimeCompare([]byte(signature), []byte(signaturex)) != 1 {
			handler.MsgHandler.InvalidRequestHandler(w, r, errors.New("check signature failed"))
			return
		}

		io.WriteString(w, echostr)
	}
}
