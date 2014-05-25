package wechat

import (
	"encoding/xml"
	"errors"
	"github.com/chanxuehong/util/pool"
	"github.com/chanxuehong/wechat/message"
	"io"
	"io/ioutil"
	"net/http"
)

type Server struct {
	token string

	// TODO: go1.3有了新的实现(sync.Pool), 目前 GAE 还不支持,
	// 如果你的环境是 go1.3+, 你可以自己更改.
	messageRequestPool *pool.Pool

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

// 非法请求的处理函数
type InvalidRequestHandlerFunc func(http.ResponseWriter, *http.Request, error)

// 正常的从微信服务器推送过来的消息处理函数
//  NOTE: *message.Request 这个对象系统会自动池化的, 所以需要这个对象里的数据要深拷贝
type RequestHandlerFunc func(http.ResponseWriter, *http.Request, *message.Request)

// 目前不能识别的从微信服务器推送过来的消息处理函数
//  NOTE: *message.Request 这个对象系统会自动池化的, 所以需要这个对象里的数据要深拷贝
type UnknownRequestHandlerFunc func(http.ResponseWriter, *http.Request, *message.Request)

// 默认的消息处理函数是什么都不做
func defaultInvalidRequestHandler(w http.ResponseWriter, r *http.Request, err error)                {}
func defaultUnknownRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.Request) {}
func defaultRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.Request)        {}

func NewServer(token string, requestPoolSize int) *Server {
	var srv Server

	srv.token = token
	srv.messageRequestPool = pool.New(newMessageRequest, requestPoolSize)

	// 注册默认的处理函数
	srv.invalidRequestHandler = defaultInvalidRequestHandler
	srv.unknownRequestHandler = defaultUnknownRequestHandler

	srv.textRequestHandler = defaultRequestHandler
	srv.imageRequestHandler = defaultRequestHandler
	srv.voiceRequestHandler = defaultRequestHandler
	srv.voiceRecognitionRequestHandler = defaultRequestHandler
	srv.videoRequestHandler = defaultRequestHandler
	srv.locationRequestHandler = defaultRequestHandler
	srv.linkRequestHandler = defaultRequestHandler
	srv.subscribeEventRequestHandler = defaultRequestHandler
	srv.subscribeEventByScanRequestHandler = defaultRequestHandler
	srv.unsubscribeEventRequestHandler = defaultRequestHandler
	srv.scanEventRequestHandler = defaultRequestHandler
	srv.locationEventRequestHandler = defaultRequestHandler
	srv.clickEventRequestHandler = defaultRequestHandler
	srv.viewEventRequestHandler = defaultRequestHandler
	srv.masssendjobfinishEventRequestHandler = defaultRequestHandler

	return &srv
}

// Server 实现 http.Handler 接口
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	// 处理从微信服务器推送过来的消息 ==============================================
	case "POST":
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

		rqstMsg := s.getRequestEntity()   // *message.Request
		defer s.putRequestEntity(rqstMsg) // important!

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

	// 首次验证 =================================================================
	case "GET":
		var err error
		var signature, timestamp, nonce, echostr string

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
		if echostr = r.FormValue("echostr"); echostr == "" {
			s.invalidRequestHandler(w, r, errors.New("echostr is empty"))
			return
		}

		if !CheckSignature(signature, timestamp, nonce, s.token) {
			s.invalidRequestHandler(w, r, errors.New("check signature failed"))
			return
		}

		io.WriteString(w, echostr)
	}
}
