package wechat

import (
	"encoding/xml"
	"errors"
	"github.com/chanxuehong/wechat/message"
	"io"
	"net/http"
	"sync"
)

// 非法请求的处理函数
type InvalidRequestHandlerFunc func(http.ResponseWriter, *http.Request, error)

// 目前不能识别的从微信服务器推送过来的消息处理函数
//  NOTE: *message.Request 这个对象系统会自动池化的, 所以需要这个对象里的数据要深拷贝
type UnknownRequestHandlerFunc func(http.ResponseWriter, *http.Request, *message.Request)

// 正常的从微信服务器推送过来的消息处理函数
//  NOTE: *message.Request 这个对象系统会自动池化的, 所以需要这个对象里的数据要深拷贝
type RequestHandlerFunc func(http.ResponseWriter, *http.Request, *message.Request)

// 默认的消息处理函数是什么都不做
func defaultInvalidRequestHandler(w http.ResponseWriter, r *http.Request, err error)                {}
func defaultUnknownRequestHandler(w http.ResponseWriter, r *http.Request, msgRqst *message.Request) {}
func defaultRequestHandler(w http.ResponseWriter, r *http.Request, msgRqst *message.Request)        {}

type ServerSetting struct {
	Token string

	// Invalid request handler
	InvalidRequestHandler InvalidRequestHandlerFunc

	// unknown request handler
	UnknownRequestHandler UnknownRequestHandlerFunc

	// request handler
	TextRequestHandler                   RequestHandlerFunc
	ImageRequestHandler                  RequestHandlerFunc
	VoiceRequestHandler                  RequestHandlerFunc
	VoiceRecognitionRequestHandler       RequestHandlerFunc
	VideoRequestHandler                  RequestHandlerFunc
	LocationRequestHandler               RequestHandlerFunc
	LinkRequestHandler                   RequestHandlerFunc
	SubscribeEventRequestHandler         RequestHandlerFunc
	SubscribeEventByScanRequestHandler   RequestHandlerFunc
	UnsubscribeEventRequestHandler       RequestHandlerFunc
	ScanEventRequestHandler              RequestHandlerFunc
	LocationEventRequestHandler          RequestHandlerFunc
	ClickEventRequestHandler             RequestHandlerFunc
	ViewEventRequestHandler              RequestHandlerFunc
	MassSendJobFinishEventRequestHandler RequestHandlerFunc
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

	// request handler
	if setting.TextRequestHandler != nil {
		ss.TextRequestHandler = setting.TextRequestHandler
	} else {
		ss.TextRequestHandler = defaultRequestHandler
	}
	if setting.ImageRequestHandler != nil {
		ss.ImageRequestHandler = setting.ImageRequestHandler
	} else {
		ss.ImageRequestHandler = defaultRequestHandler
	}
	if setting.VoiceRequestHandler != nil {
		ss.VoiceRequestHandler = setting.VoiceRequestHandler
	} else {
		ss.VoiceRequestHandler = defaultRequestHandler
	}
	if setting.VoiceRecognitionRequestHandler != nil {
		ss.VoiceRecognitionRequestHandler = setting.VoiceRecognitionRequestHandler
	} else {
		ss.VoiceRecognitionRequestHandler = defaultRequestHandler
	}
	if setting.VideoRequestHandler != nil {
		ss.VideoRequestHandler = setting.VideoRequestHandler
	} else {
		ss.VideoRequestHandler = defaultRequestHandler
	}
	if setting.LocationRequestHandler != nil {
		ss.LocationRequestHandler = setting.LocationRequestHandler
	} else {
		ss.LocationRequestHandler = defaultRequestHandler
	}
	if setting.LinkRequestHandler != nil {
		ss.LinkRequestHandler = setting.LinkRequestHandler
	} else {
		ss.LinkRequestHandler = defaultRequestHandler
	}
	if setting.SubscribeEventRequestHandler != nil {
		ss.SubscribeEventRequestHandler = setting.SubscribeEventRequestHandler
	} else {
		ss.SubscribeEventRequestHandler = defaultRequestHandler
	}
	if setting.SubscribeEventByScanRequestHandler != nil {
		ss.SubscribeEventByScanRequestHandler = setting.SubscribeEventByScanRequestHandler
	} else {
		ss.SubscribeEventByScanRequestHandler = defaultRequestHandler
	}
	if setting.UnsubscribeEventRequestHandler != nil {
		ss.UnsubscribeEventRequestHandler = setting.UnsubscribeEventRequestHandler
	} else {
		ss.UnsubscribeEventRequestHandler = defaultRequestHandler
	}
	if setting.ScanEventRequestHandler != nil {
		ss.ScanEventRequestHandler = setting.ScanEventRequestHandler
	} else {
		ss.ScanEventRequestHandler = defaultRequestHandler
	}
	if setting.LocationEventRequestHandler != nil {
		ss.LocationEventRequestHandler = setting.LocationEventRequestHandler
	} else {
		ss.LocationEventRequestHandler = defaultRequestHandler
	}
	if setting.ClickEventRequestHandler != nil {
		ss.ClickEventRequestHandler = setting.ClickEventRequestHandler
	} else {
		ss.ClickEventRequestHandler = defaultRequestHandler
	}
	if setting.ViewEventRequestHandler != nil {
		ss.ViewEventRequestHandler = setting.ViewEventRequestHandler
	} else {
		ss.ViewEventRequestHandler = defaultRequestHandler
	}
	if setting.MassSendJobFinishEventRequestHandler != nil {
		ss.MassSendJobFinishEventRequestHandler = setting.MassSendJobFinishEventRequestHandler
	} else {
		ss.MassSendJobFinishEventRequestHandler = defaultRequestHandler
	}
}

// 对于微信服务器推送过来的消息或者事件, 公众号服务程序就相当于服务器.
// 被动回复和处理各种事件功能都封装在这个结构里; Server 并发安全.
//  NOTE: 必须调用 NewServer() 创建对象!
type Server struct {
	setting ServerSetting

	// 对于微信服务器推送过来的请求, 基本都是中间处理下就丢弃, 所以可以缓存起来.
	//  NOTE: require go1.3+ , 如果你的环境不满足这个条件, 可以自己实现一个简单的 Pool,
	//        see github.com/chanxuehong/util/pool 或者直接用 sync.Pool.patch 目录下的文件;
	messageRequestPool *sync.Pool
}

func NewServer(setting *ServerSetting) *Server {
	if setting == nil {
		panic("error, wechat.NewServer: setting == nil")
	}

	var srv Server
	srv.setting.initialize(setting)
	srv.messageRequestPool = &sync.Pool{
		New: serverNewMessageRequest,
	}

	return &srv
}

// Server 实现 http.Handler 接口
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST": // 处理从微信服务器推送过来的消息 ===================================
		var err error
		var signature, timestamp, nonce string

		if err = r.ParseForm(); err != nil {
			s.setting.InvalidRequestHandler(w, r, err)
			return
		}

		if signature = r.FormValue("signature"); signature == "" {
			s.setting.InvalidRequestHandler(w, r, errors.New("signature is empty"))
			return
		}
		if timestamp = r.FormValue("timestamp"); timestamp == "" {
			s.setting.InvalidRequestHandler(w, r, errors.New("timestamp is empty"))
			return
		}
		if nonce = r.FormValue("nonce"); nonce == "" {
			s.setting.InvalidRequestHandler(w, r, errors.New("nonce is empty"))
			return
		}

		if !CheckSignature(signature, timestamp, nonce, s.setting.Token) {
			s.setting.InvalidRequestHandler(w, r, errors.New("check signature failed"))
			return
		}

		msgRqst := s.getMessageRequestFromPool() // *message.Request
		defer s.putMessageRequestToPool(msgRqst) // important!

		if err = xml.NewDecoder(r.Body).Decode(msgRqst); err != nil {
			s.setting.InvalidRequestHandler(w, r, err)
			return
		}

		// request router, 可一个根据自己的实际业务调整顺序!
		switch msgRqst.MsgType {
		case message.RQST_MSG_TYPE_TEXT:
			s.setting.TextRequestHandler(w, r, msgRqst)

		case message.RQST_MSG_TYPE_EVENT:
			// event router
			switch msgRqst.Event {
			case message.RQST_EVENT_TYPE_CLICK:
				s.setting.ClickEventRequestHandler(w, r, msgRqst)

			case message.RQST_EVENT_TYPE_VIEW:
				s.setting.ViewEventRequestHandler(w, r, msgRqst)

			case message.RQST_EVENT_TYPE_LOCATION:
				s.setting.LocationEventRequestHandler(w, r, msgRqst)

			case message.RQST_EVENT_TYPE_SUBSCRIBE:
				if msgRqst.Ticket == "" { // 普通订阅
					s.setting.SubscribeEventRequestHandler(w, r, msgRqst)
				} else { // 扫描二维码订阅
					s.setting.SubscribeEventByScanRequestHandler(w, r, msgRqst)
				}

			case message.RQST_EVENT_TYPE_UNSUBSCRIBE:
				s.setting.UnsubscribeEventRequestHandler(w, r, msgRqst)

			case message.RQST_EVENT_TYPE_SCAN:
				s.setting.ScanEventRequestHandler(w, r, msgRqst)

			case message.RQST_EVENT_TYPE_MASSSENDJOBFINISH:
				s.setting.MassSendJobFinishEventRequestHandler(w, r, msgRqst)

			default: // unknown event
				s.setting.UnknownRequestHandler(w, r, msgRqst)
			}

		case message.RQST_MSG_TYPE_LINK:
			s.setting.LinkRequestHandler(w, r, msgRqst)

		case message.RQST_MSG_TYPE_VOICE:
			if msgRqst.Recognition == "" { // 普通的语音请求
				s.setting.VoiceRequestHandler(w, r, msgRqst)
			} else { // 语音识别请求
				s.setting.VoiceRecognitionRequestHandler(w, r, msgRqst)
			}

		case message.RQST_MSG_TYPE_LOCATION:
			s.setting.LocationRequestHandler(w, r, msgRqst)

		case message.RQST_MSG_TYPE_IMAGE:
			s.setting.ImageRequestHandler(w, r, msgRqst)

		case message.RQST_MSG_TYPE_VIDEO:
			s.setting.VideoRequestHandler(w, r, msgRqst)

		default: // unknown request message type
			s.setting.UnknownRequestHandler(w, r, msgRqst)
		}

	case "GET": // 首次验证 ======================================================
		var err error
		var signature, timestamp, nonce, echostr string

		if err = r.ParseForm(); err != nil {
			s.setting.InvalidRequestHandler(w, r, err)
			return
		}

		if signature = r.FormValue("signature"); signature == "" {
			s.setting.InvalidRequestHandler(w, r, errors.New("signature is empty"))
			return
		}
		if timestamp = r.FormValue("timestamp"); timestamp == "" {
			s.setting.InvalidRequestHandler(w, r, errors.New("timestamp is empty"))
			return
		}
		if nonce = r.FormValue("nonce"); nonce == "" {
			s.setting.InvalidRequestHandler(w, r, errors.New("nonce is empty"))
			return
		}
		if echostr = r.FormValue("echostr"); echostr == "" {
			s.setting.InvalidRequestHandler(w, r, errors.New("echostr is empty"))
			return
		}

		if !CheckSignature(signature, timestamp, nonce, s.setting.Token) {
			s.setting.InvalidRequestHandler(w, r, errors.New("check signature failed"))
			return
		}

		io.WriteString(w, echostr)
	}
}
