package wxfrontend

import (
	"encoding/xml"
	"github.com/chanxuehong/wechat/message"
	"io"
	"io/ioutil"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var signature, timestamp, nonce string

	if err = r.ParseForm(); err != nil {
		//TODO: 增加相应的处理
		io.WriteString(w, "")
		return
	}

	if signature = r.FormValue("signature"); signature == "" {
		//TODO: 增加相应的处理
		io.WriteString(w, "")
		return
	}
	if timestamp = r.FormValue("timestamp"); timestamp == "" {
		//TODO: 增加相应的处理
		io.WriteString(w, "")
		return
	}
	if nonce = r.FormValue("nonce"); nonce == "" {
		//TODO: 增加相应的处理
		io.WriteString(w, "")
		return
	}

	if !message.CheckSignature(signature, timestamp, nonce, wechatToken) {
		//TODO: 增加相应的处理
		io.WriteString(w, "")
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//TODO: 增加相应的处理
		io.WriteString(w, "")
		return
	}

	rqstMsg := getRequestMsg()   // *message.RequestMsg
	defer putRequestMsg(rqstMsg) // important!

	if err = xml.Unmarshal(b, rqstMsg); err != nil {
		//TODO: 增加相应的处理
		io.WriteString(w, "")
		return
	}

	// request router
	switch rqstMsg.MsgType {

	case message.RQST_MSG_TYPE_TEXT:
		textRequestHandler(w, r, rqstMsg)

	case message.RQST_MSG_TYPE_IMAGE:
		imageRequestHandler(w, r, rqstMsg)

	case message.RQST_MSG_TYPE_VOICE:
		if rqstMsg.Recognition == "" { // 普通的语音请求
			voiceRequestHandler(w, r, rqstMsg)
		} else { // 语音识别请求
			voiceRecognitionRequestHandler(w, r, rqstMsg)
		}

	case message.RQST_MSG_TYPE_VIDEO:
		videoRequestHandler(w, r, rqstMsg)

	case message.RQST_MSG_TYPE_LOCATION:
		locationRequestHandler(w, r, rqstMsg)

	case message.RQST_MSG_TYPE_LINK:
		linkRequestHandler(w, r, rqstMsg)

	case message.RQST_MSG_TYPE_EVENT:
		// event router
		switch rqstMsg.Event {

		case message.RQST_EVENT_TYPE_SUBSCRIBE:
			if rqstMsg.Ticket == "" {
				subscribeEventRequestHandler(w, r, rqstMsg)
			} else { // 扫描二维码订阅
				subscribeEventByScanRequestHandler(w, r, rqstMsg)
			}

		case message.RQST_EVENT_TYPE_UNSUBSCRIBE:
			unsubscribeEventRequestHandler(w, r, rqstMsg)

		case message.RQST_EVENT_TYPE_SCAN:
			scanEventRequestHandler(w, r, rqstMsg)

		case message.RQST_EVENT_TYPE_LOCATION:
			locationEventRequestHandler(w, r, rqstMsg)

		case message.RQST_EVENT_TYPE_CLICK:
			clickEventRequestHandler(w, r, rqstMsg)

		case message.RQST_EVENT_TYPE_VIEW:
			viewEventRequestHandler(w, r, rqstMsg)

		case message.RQST_EVENT_TYPE_MASSSENDJOBFINISH:
			masssendjobfinishEventRequestHandler(w, r, rqstMsg)

		default: // unknown event
			//TODO: 增加相应的处理
			io.WriteString(w, "")
		}

	default: // unknown request message type
		//TODO: 增加相应的处理
		io.WriteString(w, "")
	}
}
