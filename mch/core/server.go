package core

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/chanxuehong/util"
	"github.com/chanxuehong/util/security"

	"github.com/chanxuehong/wechat/internal/debug/mch/callback"
)

type Server struct {
	appId  string
	mchId  string
	apiKey string

	subAppId string
	subMchId string

	handler      Handler
	errorHandler ErrorHandler
}

// NewServer 创建一个新的 Server.
//  appId:        可选; 公众号的 appid, 如果设置了值则该 Server 只能处理 appid 为该值的消息(事件)
//  mchId:        可选; 商户号 mch_id, 如果设置了值则该 Server 只能处理 mch_id 为该值的消息(事件)
//  apiKey:       必选; 商户的签名 key
//  handler:      必选; 处理微信服务器推送过来的消息(事件)的 Handler
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

// NewSubMchServer 创建一个新的 Server.
//  appId:        可选; 公众号的 appid, 如果设置了值则该 Server 只能处理 appid 为该值的消息(事件)
//  mchId:        可选; 商户号 mch_id, 如果设置了值则该 Server 只能处理 mch_id 为该值的消息(事件)
//  apiKey:       必选; 商户的签名 key
//  subAppId:     可选; 公众号的 sub_appid, 如果设置了值则该 Server 只能处理 sub_appid 为该值的消息(事件)
//  subMchId:     可选; 商户号 sub_mch_id, 如果设置了值则该 Server 只能处理 sub_mch_id 为该值的消息(事件)
//  handler:      必选; 处理微信服务器推送过来的消息(事件)的 Handler
//  errorHandler: 可选; 用于处理 Server 在处理消息(事件)过程中产生的错误, 如果没有设置则默认使用 DefaultErrorHandler
func NewSubMchServer(appId, mchId, apiKey string, subAppId, subMchId string, handler Handler, errorHandler ErrorHandler) *Server {
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
		subAppId:     subAppId,
		subMchId:     subMchId,
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

func (srv *Server) SubAppId() string {
	return srv.subAppId
}
func (srv *Server) SubMchId() string {
	return srv.subMchId
}

// ServeHTTP 处理微信服务器的回调请求, query 参数可以为 nil.
func (srv *Server) ServeHTTP(w http.ResponseWriter, r *http.Request, query url.Values) {
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

		returnCode := msg["return_code"]
		if returnCode != "" && returnCode != ReturnCodeSuccess {
			err = &Error{
				ReturnCode: returnCode,
				ReturnMsg:  msg["return_msg"],
			}
			errorHandler.ServeError(w, r, err)
			return
		}

		resultCode := msg["result_code"]
		if resultCode != "" && resultCode != ResultCodeSuccess {
			err = &BizError{
				ResultCode:  resultCode,
				ErrCode:     msg["err_code"],
				ErrCodeDesc: msg["err_code_des"],
			}
			errorHandler.ServeError(w, r, err)
			return
		}

		if srv.appId != "" {
			wantAppId := srv.appId
			haveAppId := msg["appid"]
			if haveAppId != "" && !security.SecureCompareString(haveAppId, wantAppId) {
				err = fmt.Errorf("appid mismatch, have: %s, want: %s", haveAppId, wantAppId)
				errorHandler.ServeError(w, r, err)
				return
			}
		}
		if srv.mchId != "" {
			wantMchId := srv.mchId
			haveMchId := msg["mch_id"]
			if haveMchId != "" && !security.SecureCompareString(haveMchId, wantMchId) {
				err = fmt.Errorf("mch_id mismatch, have: %s, want: %s", haveMchId, wantMchId)
				errorHandler.ServeError(w, r, err)
				return
			}
		}

		if srv.subAppId != "" {
			wantSubAppId := srv.subAppId
			haveSubAppId := msg["sub_appid"]
			if haveSubAppId != "" && !security.SecureCompareString(haveSubAppId, wantSubAppId) {
				err = fmt.Errorf("sub_appid mismatch, have: %s, want: %s", haveSubAppId, wantSubAppId)
				errorHandler.ServeError(w, r, err)
				return
			}
		}
		if srv.subMchId != "" {
			wantSubMchId := srv.subMchId
			haveSubMchId := msg["sub_mch_id"]
			if haveSubMchId != "" && !security.SecureCompareString(haveSubMchId, wantSubMchId) {
				err = fmt.Errorf("sub_mch_id mismatch, have: %s, want: %s", haveSubMchId, wantSubMchId)
				errorHandler.ServeError(w, r, err)
				return
			}
		}

		// 认证签名
		if haveSignature := msg["sign"]; haveSignature != "" {
			var wantSignature string
			switch signType := msg["sign_type"]; signType {
			case "", SignType_MD5:
				wantSignature = Sign2(msg, srv.apiKey, md5.New())
			case SignType_HMAC_SHA256:
				wantSignature = Sign2(msg, srv.apiKey, hmac.New(sha256.New, []byte(srv.apiKey)))
			default:
				err = fmt.Errorf("unsupported notification sign_type: %s", signType)
				errorHandler.ServeError(w, r, err)
				return
			}
			if !security.SecureCompareString(haveSignature, wantSignature) {
				err = fmt.Errorf("sign mismatch,\nhave: %s,\nwant: %s", haveSignature, wantSignature)
				errorHandler.ServeError(w, r, err)
				return
			}
		} else {
			if _, ok := msg["req_info"]; !ok { // 退款结果通知没有 sign 字段
				err = ErrNotFoundSign
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
