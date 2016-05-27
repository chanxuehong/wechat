package core

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/chanxuehong/util"
	"github.com/chanxuehong/util/security"
	"github.com/chanxuehong/wechat.v2/internal/debug/mch/callback"
)

type Server struct {
	appId  string
	mchId  string
	apiKey string

	handler      Handler
	errorHandler ErrorHandler
}

// NewServer 创建一个新的 Server.
//  appId:        可选; 公众号的 appid, 如果设置了值则该 Server 只能处理 appid 为该值的消息(事件)
//  mchId:        可选; 商户号 mch_id, 如果设置了值则该 Server 只能处理 mch_id 为该值的消息(事件)
//  apiKey:       必须; 商户的签名 key
//  handler:      必须; 处理微信服务器推送过来的消息(事件)的 Handler
//  errorHandler: 可选; 用于处理 Server 在处理消息(事件)过程中产生的错误, 如果没有设置则默认使用 DefaultErrorHandler
func NewServer(appId, mchId, apiKey string, handler Handler, errorHandler ErrorHandler) *Server {
	if apiKey == "" {
		panic("empty apiKey")
	}
	if handler == nil {
		panic("nil Handler")
	}
	if errorHandler == nil {
		errorHandler = DefaultErrorHandler
	}

	return &Server{
		appId:        appId,
		mchId:        mchId,
		apiKey:       apiKey,
		handler:      handler,
		errorHandler: errorHandler,
	}
}

func (srv *Server) AppId() string {
	return srv.appId
}
func (srv *Server) MchId() string {
	return srv.mchId
}
func (srv *Server) ApiKey() string {
	return srv.apiKey
}

// ServeHTTP 处理微信服务器的回调请求, queryParams 参数可以为 nil.
func (srv *Server) ServeHTTP(w http.ResponseWriter, r *http.Request, queryParams url.Values) {
	callback.DebugPrintRequest(r)
	errorHandler := srv.errorHandler

	switch r.Method {
	case "POST":
		requestBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			errorHandler.ServeError(w, r, err)
			return
		}
		callback.DebugPrintRequestMessage(requestBody)

		msg, err := util.DecodeXMLToMap(bytes.NewReader(requestBody))
		if err != nil {
			errorHandler.ServeError(w, r, err)
			return
		}

		returnCode, ok := msg["return_code"]
		if returnCode == ReturnCodeSuccess || !ok {
			haveAppId := msg["appid"]
			wantAppId := srv.appId
			if haveAppId != "" && wantAppId != "" && !security.SecureCompareString(haveAppId, wantAppId) {
				err = fmt.Errorf("appid mismatch, have: %s, want: %s", haveAppId, wantAppId)
				errorHandler.ServeError(w, r, err)
				return
			}

			haveMchId := msg["mch_id"]
			wantMchId := srv.mchId
			if haveMchId != "" && wantMchId != "" && !security.SecureCompareString(haveMchId, wantMchId) {
				err = fmt.Errorf("mch_id mismatch, have: %s, want: %s", haveMchId, wantMchId)
				errorHandler.ServeError(w, r, err)
				return
			}

			// 认证签名
			haveSignature, ok := msg["sign"]
			if !ok {
				err = ErrNotFoundSign
				errorHandler.ServeError(w, r, err)
				return
			}
			wantSignature := Sign(msg, srv.apiKey, nil)
			if !security.SecureCompareString(haveSignature, wantSignature) {
				err = fmt.Errorf("sign mismatch,\nhave: %s,\nwant: %s", haveSignature, wantSignature)
				errorHandler.ServeError(w, r, err)
				return
			}
		}

		ctx := &Context{
			Server: srv,

			ResponseWriter: w,
			Request:        r,

			RequestBody: requestBody,
			Msg:         msg,

			handlerIndex: initHandlerIndex,
		}
		srv.handler.ServeMsg(ctx)
	default:
		errorHandler.ServeError(w, r, errors.New("Unexpected HTTP Method: "+r.Method))
	}
}
