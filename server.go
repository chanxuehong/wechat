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
type UnknownRequestHandlerFunc func(w http.ResponseWriter, r *http.Request)

// 正常的从微信服务器推送过来的消息处理函数
type TextRequestHandlerFunc func(http.ResponseWriter, *http.Request, *message.TextRequest)
type ImageRequestHandlerFunc func(http.ResponseWriter, *http.Request, *message.ImageRequest)
type VoiceRequestHandlerFunc func(http.ResponseWriter, *http.Request, *message.VoiceRequest)
type VoiceRecognitionRequestHandlerFunc func(http.ResponseWriter, *http.Request, *message.VoiceRecognitionRequest)
type VideoRequestHandlerFunc func(http.ResponseWriter, *http.Request, *message.VideoRequest)
type LocationRequestHandlerFunc func(http.ResponseWriter, *http.Request, *message.LocationRequest)
type LinkRequestHandlerFunc func(http.ResponseWriter, *http.Request, *message.LinkRequest)
type SubscribeEventHandlerFunc func(http.ResponseWriter, *http.Request, *message.SubscribeEvent)
type UnsubscribeEventHandlerFunc func(http.ResponseWriter, *http.Request, *message.UnsubscribeEvent)
type SubscribeByScanEventHandlerFunc func(http.ResponseWriter, *http.Request, *message.SubscribeByScanEvent)
type ScanEventHandlerFunc func(http.ResponseWriter, *http.Request, *message.ScanEvent)
type LocationEventHandlerFunc func(http.ResponseWriter, *http.Request, *message.LocationEvent)
type MenuClickEventHandlerFunc func(http.ResponseWriter, *http.Request, *message.MenuClickEvent)
type MenuViewEventHandlerFunc func(http.ResponseWriter, *http.Request, *message.MenuViewEvent)
type MassSendJobFinishEventHandlerFunc func(http.ResponseWriter, *http.Request, *message.MassSendJobFinishEvent)
type MerchantOrderEventHandlerFunc func(http.ResponseWriter, *http.Request, *message.MerchantOrderEvent)

// 默认的消息处理函数是什么都不做
func defaultInvalidRequestHandler(http.ResponseWriter, *http.Request, error) {}
func defaultUnknownRequestHandler(http.ResponseWriter, *http.Request)        {}

func defaultTextRequestHandler(http.ResponseWriter, *http.Request, *message.TextRequest)   {}
func defaultImageRequestHandler(http.ResponseWriter, *http.Request, *message.ImageRequest) {}
func defaultVoiceRequestHandler(http.ResponseWriter, *http.Request, *message.VoiceRequest) {}
func defaultVoiceRecognitionRequestHandler(http.ResponseWriter, *http.Request, *message.VoiceRecognitionRequest) {
}
func defaultVideoRequestHandler(http.ResponseWriter, *http.Request, *message.VideoRequest)         {}
func defaultLocationRequestHandler(http.ResponseWriter, *http.Request, *message.LocationRequest)   {}
func defaultLinkRequestHandler(http.ResponseWriter, *http.Request, *message.LinkRequest)           {}
func defaultSubscribeEventHandler(http.ResponseWriter, *http.Request, *message.SubscribeEvent)     {}
func defaultUnsubscribeEventHandler(http.ResponseWriter, *http.Request, *message.UnsubscribeEvent) {}
func defaultSubscribeByScanEventHandler(http.ResponseWriter, *http.Request, *message.SubscribeByScanEvent) {
}
func defaultScanEventHandler(http.ResponseWriter, *http.Request, *message.ScanEvent)           {}
func defaultLocationEventHandler(http.ResponseWriter, *http.Request, *message.LocationEvent)   {}
func defaultMenuClickEventHandler(http.ResponseWriter, *http.Request, *message.MenuClickEvent) {}
func defaultMenuViewEventHandler(http.ResponseWriter, *http.Request, *message.MenuViewEvent)   {}
func defaultMassSendJobFinishEventHandler(http.ResponseWriter, *http.Request, *message.MassSendJobFinishEvent) {
}
func defaultMerchantOrderEventHandler(http.ResponseWriter, *http.Request, *message.MerchantOrderEvent) {
}

type ServerSetting struct {
	Token string

	// Invalid request handler
	InvalidRequestHandler InvalidRequestHandlerFunc

	// unknown request handler
	UnknownRequestHandler UnknownRequestHandlerFunc

	// request handler
	TextRequestHandler             TextRequestHandlerFunc
	ImageRequestHandler            ImageRequestHandlerFunc
	VoiceRequestHandler            VoiceRequestHandlerFunc
	VoiceRecognitionRequestHandler VoiceRecognitionRequestHandlerFunc
	VideoRequestHandler            VideoRequestHandlerFunc
	LocationRequestHandler         LocationRequestHandlerFunc
	LinkRequestHandler             LinkRequestHandlerFunc
	// event handler
	SubscribeEventHandler         SubscribeEventHandlerFunc
	UnsubscribeEventHandler       UnsubscribeEventHandlerFunc
	SubscribeByScanEventHandler   SubscribeByScanEventHandlerFunc
	ScanEventHandler              ScanEventHandlerFunc
	LocationEventHandler          LocationEventHandlerFunc
	MenuClickEventHandler         MenuClickEventHandlerFunc
	MenuViewEventHandler          MenuViewEventHandlerFunc
	MassSendJobFinishEventHandler MassSendJobFinishEventHandlerFunc
	MerchantOrderEventHandler     MerchantOrderEventHandlerFunc
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
		ss.TextRequestHandler = defaultTextRequestHandler
	}
	if setting.ImageRequestHandler != nil {
		ss.ImageRequestHandler = setting.ImageRequestHandler
	} else {
		ss.ImageRequestHandler = defaultImageRequestHandler
	}
	if setting.VoiceRequestHandler != nil {
		ss.VoiceRequestHandler = setting.VoiceRequestHandler
	} else {
		ss.VoiceRequestHandler = defaultVoiceRequestHandler
	}
	if setting.VoiceRecognitionRequestHandler != nil {
		ss.VoiceRecognitionRequestHandler = setting.VoiceRecognitionRequestHandler
	} else {
		ss.VoiceRecognitionRequestHandler = defaultVoiceRecognitionRequestHandler
	}
	if setting.VideoRequestHandler != nil {
		ss.VideoRequestHandler = setting.VideoRequestHandler
	} else {
		ss.VideoRequestHandler = defaultVideoRequestHandler
	}
	if setting.LocationRequestHandler != nil {
		ss.LocationRequestHandler = setting.LocationRequestHandler
	} else {
		ss.LocationRequestHandler = defaultLocationRequestHandler
	}
	if setting.LinkRequestHandler != nil {
		ss.LinkRequestHandler = setting.LinkRequestHandler
	} else {
		ss.LinkRequestHandler = defaultLinkRequestHandler
	}

	// event handler
	if setting.SubscribeEventHandler != nil {
		ss.SubscribeEventHandler = setting.SubscribeEventHandler
	} else {
		ss.SubscribeEventHandler = defaultSubscribeEventHandler
	}
	if setting.UnsubscribeEventHandler != nil {
		ss.UnsubscribeEventHandler = setting.UnsubscribeEventHandler
	} else {
		ss.UnsubscribeEventHandler = defaultUnsubscribeEventHandler
	}
	if setting.SubscribeByScanEventHandler != nil {
		ss.SubscribeByScanEventHandler = setting.SubscribeByScanEventHandler
	} else {
		ss.SubscribeByScanEventHandler = defaultSubscribeByScanEventHandler
	}
	if setting.ScanEventHandler != nil {
		ss.ScanEventHandler = setting.ScanEventHandler
	} else {
		ss.ScanEventHandler = defaultScanEventHandler
	}
	if setting.LocationEventHandler != nil {
		ss.LocationEventHandler = setting.LocationEventHandler
	} else {
		ss.LocationEventHandler = defaultLocationEventHandler
	}
	if setting.MenuClickEventHandler != nil {
		ss.MenuClickEventHandler = setting.MenuClickEventHandler
	} else {
		ss.MenuClickEventHandler = defaultMenuClickEventHandler
	}
	if setting.MenuViewEventHandler != nil {
		ss.MenuViewEventHandler = setting.MenuViewEventHandler
	} else {
		ss.MenuViewEventHandler = defaultMenuViewEventHandler
	}
	if setting.MassSendJobFinishEventHandler != nil {
		ss.MassSendJobFinishEventHandler = setting.MassSendJobFinishEventHandler
	} else {
		ss.MassSendJobFinishEventHandler = defaultMassSendJobFinishEventHandler
	}
	if setting.MerchantOrderEventHandler != nil {
		ss.MerchantOrderEventHandler = setting.MerchantOrderEventHandler
	} else {
		ss.MerchantOrderEventHandler = defaultMerchantOrderEventHandler
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
	bufferUnitPool *sync.Pool
}

func NewServer(setting *ServerSetting) *Server {
	if setting == nil {
		panic("error, wechat.NewServer: setting == nil")
	}

	var srv Server
	srv.setting.initialize(setting)
	srv.bufferUnitPool = &sync.Pool{
		New: newServerBufferUnit,
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

		bufferUnit := s.getBufferUnitFromPool() // *serverBufferUnit
		defer s.putBufferUnitToPool(bufferUnit) // important!

		if !CheckSignatureEx(signature, timestamp, nonce, s.setting.Token, bufferUnit.buf) {
			s.setting.InvalidRequestHandler(w, r, errors.New("check signature failed"))
			return
		}

		msgRqst := &bufferUnit.msgRequest
		if err = xml.NewDecoder(r.Body).Decode(msgRqst); err != nil {
			s.setting.InvalidRequestHandler(w, r, err)
			return
		}

		// request router, 可一个根据自己的实际业务调整顺序!
		switch msgRqst.MsgType {
		case message.RQST_MSG_TYPE_TEXT:
			s.setting.TextRequestHandler(w, r, msgRqst.TextRequest())

		case message.RQST_MSG_TYPE_EVENT:
			// event router
			switch msgRqst.Event {
			case message.RQST_EVENT_TYPE_CLICK:
				s.setting.MenuClickEventHandler(w, r, msgRqst.MenuClickEvent())

			case message.RQST_EVENT_TYPE_VIEW:
				s.setting.MenuViewEventHandler(w, r, msgRqst.MenuViewEvent())

			case message.RQST_EVENT_TYPE_LOCATION:
				s.setting.LocationEventHandler(w, r, msgRqst.LocationEvent())

			case message.RQST_EVENT_TYPE_MERCHANTORDER:
				s.setting.MerchantOrderEventHandler(w, r, msgRqst.MerchantOrderEvent())

			case message.RQST_EVENT_TYPE_SUBSCRIBE:
				if msgRqst.Ticket == "" { // 普通订阅
					s.setting.SubscribeEventHandler(w, r, msgRqst.SubscribeEvent())
				} else { // 扫描二维码订阅
					s.setting.SubscribeByScanEventHandler(w, r, msgRqst.SubscribeByScanEvent())
				}

			case message.RQST_EVENT_TYPE_UNSUBSCRIBE:
				s.setting.UnsubscribeEventHandler(w, r, msgRqst.UnsubscribeEvent())

			case message.RQST_EVENT_TYPE_SCAN:
				s.setting.ScanEventHandler(w, r, msgRqst.ScanEvent())

			case message.RQST_EVENT_TYPE_MASSSENDJOBFINISH:
				s.setting.MassSendJobFinishEventHandler(w, r, msgRqst.MassSendJobFinishEvent())

			default: // unknown event
				s.setting.UnknownRequestHandler(w, r)
			}

		case message.RQST_MSG_TYPE_LINK:
			s.setting.LinkRequestHandler(w, r, msgRqst.LinkRequest())

		case message.RQST_MSG_TYPE_VOICE:
			if msgRqst.Recognition == "" { // 普通的语音请求
				s.setting.VoiceRequestHandler(w, r, msgRqst.VoiceRequest())
			} else { // 语音识别请求
				s.setting.VoiceRecognitionRequestHandler(w, r, msgRqst.VoiceRecognitionRequest())
			}

		case message.RQST_MSG_TYPE_LOCATION:
			s.setting.LocationRequestHandler(w, r, msgRqst.LocationRequest())

		case message.RQST_MSG_TYPE_IMAGE:
			s.setting.ImageRequestHandler(w, r, msgRqst.ImageRequest())

		case message.RQST_MSG_TYPE_VIDEO:
			s.setting.VideoRequestHandler(w, r, msgRqst.VideoRequest())

		default: // unknown request message type
			s.setting.UnknownRequestHandler(w, r)
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
