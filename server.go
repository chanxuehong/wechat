package wechat

import (
	"encoding/xml"
	"errors"
	"github.com/chanxuehong/util/pool"
	"github.com/chanxuehong/wechat/message"
	"io/ioutil"
	"net/http"
)

type InvalidRequestHandlerFunc func(http.ResponseWriter, *http.Request, error)
type RequestHandlerFunc func(http.ResponseWriter, *http.Request, *message.Request)
type UnknownRequestHandlerFunc func(http.ResponseWriter, *http.Request, *message.Request)

type Server struct {
	token string

	messageRequestPool *pool.Pool // go1.3有了新的实现(sync.Pool), 目前 GAE 还不支持

	// Invalid or unknown request handler
	invalidRequestHandler InvalidRequestHandlerFunc
	unknownRequestHandler UnknownRequestHandlerFunc

	// request handler
	textRequestHandler                   RequestHandlerFunc
	imageRequestHandler                  RequestHandlerFunc
	voiceRequestHandler                  RequestHandlerFunc
	voiceRecognitionRequestHandler       RequestHandlerFunc
	videoRequestHandler                  RequestHandlerFunc
	locationRequestHandler               RequestHandlerFunc
	linkRequestHandler                   RequestHandlerFunc
	subscribeEventRequestHandler         RequestHandlerFunc
	subscribeEventByScanRequestHandler   RequestHandlerFunc
	unsubscribeEventRequestHandler       RequestHandlerFunc
	scanEventRequestHandler              RequestHandlerFunc
	locationEventRequestHandler          RequestHandlerFunc
	clickEventRequestHandler             RequestHandlerFunc
	viewEventRequestHandler              RequestHandlerFunc
	masssendjobfinishEventRequestHandler RequestHandlerFunc
}

func NewServer(token string, requestPoolSize int) *Server {
	var srv Server

	srv.token = token
	srv.messageRequestPool = pool.New(newMessageRequest, requestPoolSize)

	// 注册默认的处理函数

	srv.invalidRequestHandler = invalidRequestHandler
	srv.unknownRequestHandler = unknownRequestHandler

	srv.textRequestHandler = textRequestHandler
	srv.imageRequestHandler = imageRequestHandler
	srv.voiceRequestHandler = voiceRequestHandler
	srv.voiceRecognitionRequestHandler = voiceRecognitionRequestHandler
	srv.videoRequestHandler = videoRequestHandler
	srv.locationRequestHandler = locationRequestHandler
	srv.linkRequestHandler = linkRequestHandler
	srv.subscribeEventRequestHandler = subscribeEventRequestHandler
	srv.subscribeEventByScanRequestHandler = subscribeEventByScanRequestHandler
	srv.unsubscribeEventRequestHandler = unsubscribeEventRequestHandler
	srv.scanEventRequestHandler = scanEventRequestHandler
	srv.locationEventRequestHandler = locationEventRequestHandler
	srv.clickEventRequestHandler = clickEventRequestHandler
	srv.viewEventRequestHandler = viewEventRequestHandler
	srv.masssendjobfinishEventRequestHandler = masssendjobfinishEventRequestHandler

	return &srv
}

// Server 实现 http.Handler 接口
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	var signature, timestamp, nonce string

	if err = r.ParseForm(); err != nil {
		s.invalidRequestHandler(w, r, err)
		return
	}

	if signature = r.FormValue("signature"); signature == "" {
		s.invalidRequestHandler(w, r, errors.New("signature is empty"))
		return
	}
	if timestamp = r.FormValue("timestamp"); timestamp == "" {
		s.invalidRequestHandler(w, r, errors.New("timestamp is empty"))
		return
	}
	if nonce = r.FormValue("nonce"); nonce == "" {
		s.invalidRequestHandler(w, r, errors.New("nonce is empty"))
		return
	}

	if !CheckSignature(signature, timestamp, nonce, s.token) {
		s.invalidRequestHandler(w, r, errors.New("check signature failed"))
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		s.invalidRequestHandler(w, r, err)
		return
	}

	rqstMsg := s.getRequestFromPool() // *message.Request
	defer s.putRequestToPool(rqstMsg) // important!

	if err = xml.Unmarshal(b, rqstMsg); err != nil {
		s.invalidRequestHandler(w, r, err)
		return
	}

	// request router
	switch rqstMsg.MsgType {

	case message.RQST_MSG_TYPE_TEXT:
		s.textRequestHandler(w, r, rqstMsg)

	case message.RQST_MSG_TYPE_VOICE:
		if rqstMsg.Recognition == "" { // 普通的语音请求
			s.voiceRequestHandler(w, r, rqstMsg)
		} else { // 语音识别请求
			s.voiceRecognitionRequestHandler(w, r, rqstMsg)
		}

	case message.RQST_MSG_TYPE_LOCATION:
		s.locationRequestHandler(w, r, rqstMsg)

	case message.RQST_MSG_TYPE_LINK:
		s.linkRequestHandler(w, r, rqstMsg)

	case message.RQST_MSG_TYPE_IMAGE:
		s.imageRequestHandler(w, r, rqstMsg)

	case message.RQST_MSG_TYPE_VIDEO:
		s.videoRequestHandler(w, r, rqstMsg)

	case message.RQST_MSG_TYPE_EVENT:
		// event router
		switch rqstMsg.Event {

		case message.RQST_EVENT_TYPE_SUBSCRIBE:
			if rqstMsg.Ticket == "" {
				s.subscribeEventRequestHandler(w, r, rqstMsg)
			} else { // 扫描二维码订阅
				s.subscribeEventByScanRequestHandler(w, r, rqstMsg)
			}

		case message.RQST_EVENT_TYPE_UNSUBSCRIBE:
			s.unsubscribeEventRequestHandler(w, r, rqstMsg)

		case message.RQST_EVENT_TYPE_SCAN:
			s.scanEventRequestHandler(w, r, rqstMsg)

		case message.RQST_EVENT_TYPE_LOCATION:
			s.locationEventRequestHandler(w, r, rqstMsg)

		case message.RQST_EVENT_TYPE_CLICK:
			s.clickEventRequestHandler(w, r, rqstMsg)

		case message.RQST_EVENT_TYPE_VIEW:
			s.viewEventRequestHandler(w, r, rqstMsg)

		case message.RQST_EVENT_TYPE_MASSSENDJOBFINISH:
			s.masssendjobfinishEventRequestHandler(w, r, rqstMsg)

		default: // unknown event
			s.unknownRequestHandler(w, r, rqstMsg)
		}

	default: // unknown request message type
		s.unknownRequestHandler(w, r, rqstMsg)
	}
}
