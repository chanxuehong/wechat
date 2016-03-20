package core

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	"github.com/chanxuehong/util/security"

	"github.com/chanxuehong/wechat/internal/debug/callback"
	"github.com/chanxuehong/wechat/internal/util"
)

// Server 用于处理微信服务器的回调请求, 并发安全!
//  通常情况下一个 Server 实例用于处理一个公众号的消息(事件), 此时建议指定 oriId(原始ID) 和 appId(明文模式下无需指定) 用于约束消息(事件);
//  特殊情况下也可以一个 Server 实例用于处理多个公众号的消息(事件), 此时要求这些公众号的 token 是一样的, 并且 oriId 和 appId 必须设置为 "".
type Server struct {
	oriId string
	appId string
	token string

	aesRWMutex    sync.RWMutex
	currentAESKey []byte
	lastAESKey    []byte

	handler      Handler
	errorHandler ErrorHandler
}

// NewServer 创建一个新的 Server.
//  oriId:        可选; 公众号的原始ID(微信公众号管理后台查看), 如果设置了值则该Server只能处理 ToUserName 为该值的公众号的消息(事件);
//  appId:        可选; 公众号的AppId, 如果设置了值则安全模式时该Server只能处理 AppId 为该值的公众号的消息(事件);
//  token:        必须; 公众号用于验证签名的token;
//  base64AESKey: 可选; aes加密解密key, 43字节长(base64编码, 去掉了尾部的'='), 安全模式必须设置;
//  handler:      必须; 处理微信服务器推送过来的消息(事件)的Handler;
//  errorHandler: 可选; 用于处理Server在处理消息(事件)过程中产生的错误, 如果没有设置则默认使用 DefaultErrorHandler.
func NewServer(oriId, appId, token, base64AESKey string, handler Handler, errorHandler ErrorHandler) (srv *Server) {
	if token == "" {
		panic("empty token")
	}
	if handler == nil {
		panic("nil Handler")
	}
	if errorHandler == nil {
		errorHandler = DefaultErrorHandler
	}

	var (
		aesKey []byte
		err    error
	)
	if base64AESKey != "" {
		if len(base64AESKey) != 43 {
			panic("the length of base64AESKey must equal to 43")
		}
		aesKey, err = base64.StdEncoding.DecodeString(base64AESKey + "=")
		if err != nil {
			panic(fmt.Sprintf("Decode base64AESKey:%s failed", base64AESKey))
		}
	}

	return &Server{
		oriId:         oriId,
		appId:         appId,
		token:         token,
		currentAESKey: aesKey,
		handler:       handler,
		errorHandler:  errorHandler,
	}
}

func (srv *Server) getCurrentAESKey() (key []byte) {
	srv.aesRWMutex.RLock()
	key = srv.currentAESKey
	srv.aesRWMutex.RUnlock()
	return
}

func (srv *Server) getLastAESKey() (key []byte) {
	srv.aesRWMutex.RLock()
	key = srv.lastAESKey
	srv.aesRWMutex.RUnlock()
	return
}

// SetAESKey 设置新的aes加密解密key.
//  base64AESKey: aes加密解密key, 43字节长(base64编码, 去掉了尾部的'=').
func (srv *Server) SetAESKey(base64AESKey string) (err error) {
	if len(base64AESKey) != 43 {
		return errors.New("the length of base64AESKey must equal to 43")
	}
	aesKey, err := base64.StdEncoding.DecodeString(base64AESKey + "=")
	if err != nil {
		return
	}

	srv.aesRWMutex.Lock()
	defer srv.aesRWMutex.Unlock()

	if bytes.Equal(aesKey, srv.currentAESKey) {
		return
	}
	srv.lastAESKey = srv.currentAESKey
	srv.currentAESKey = aesKey
	return
}

// ServeHTTP 处理微信服务器的回调请求, queryParams 参数可以为 nil.
func (srv *Server) ServeHTTP(w http.ResponseWriter, r *http.Request, queryParams url.Values) {
	callback.DebugPrintRequest(r)
	if queryParams == nil {
		queryParams = r.URL.Query()
	}
	errorHandler := srv.errorHandler

	switch r.Method {
	case "POST": // 推送消息(事件)
		switch encryptType := queryParams.Get("encrypt_type"); encryptType {
		case "aes":
			haveSignature := queryParams.Get("signature")
			if haveSignature == "" {
				errorHandler.ServeError(w, r, errors.New("not found signature query parameter"))
				return
			}
			haveMsgSignature := queryParams.Get("msg_signature")
			if haveMsgSignature == "" {
				errorHandler.ServeError(w, r, errors.New("not found msg_signature query parameter"))
				return
			}
			timestampString := queryParams.Get("timestamp")
			if timestampString == "" {
				errorHandler.ServeError(w, r, errors.New("not found timestamp query parameter"))
				return
			}
			timestamp, err := strconv.ParseInt(timestampString, 10, 64)
			if err != nil {
				err = fmt.Errorf("can not parse timestamp query parameter %q to int64", timestampString)
				errorHandler.ServeError(w, r, err)
				return
			}
			nonce := queryParams.Get("nonce")
			if nonce == "" {
				errorHandler.ServeError(w, r, errors.New("not found nonce query parameter"))
				return
			}

			wantSignature := util.Sign(srv.token, timestampString, nonce)
			if !security.SecureCompareString(haveSignature, wantSignature) {
				err = fmt.Errorf("check signature failed, have: %s, want: %s", haveSignature, wantSignature)
				errorHandler.ServeError(w, r, err)
				return
			}

			var requestHttpBody struct {
				XMLName            struct{} `xml:"xml"`
				ToUserName         string   `xml:"ToUserName"`
				Base64EncryptedMsg []byte   `xml:"Encrypt"`
			}
			if err = xml.NewDecoder(r.Body).Decode(&requestHttpBody); err != nil {
				errorHandler.ServeError(w, r, err)
				return
			}

			haveToUserName := requestHttpBody.ToUserName
			wantToUserName := srv.oriId
			if wantToUserName != "" && !security.SecureCompareString(haveToUserName, wantToUserName) {
				err = fmt.Errorf("the message ToUserName mismatch, have: %s, want: %s",
					haveToUserName, wantToUserName)
				errorHandler.ServeError(w, r, err)
				return
			}

			wantMsgSignature := util.MsgSign(srv.token, timestampString, nonce, string(requestHttpBody.Base64EncryptedMsg))
			if !security.SecureCompareString(haveMsgSignature, wantMsgSignature) {
				err = fmt.Errorf("check msg_signature failed, have: %s, want: %s", haveMsgSignature, wantMsgSignature)
				errorHandler.ServeError(w, r, err)
				return
			}

			encryptedMsg := make([]byte, base64.StdEncoding.DecodedLen(len(requestHttpBody.Base64EncryptedMsg)))
			encryptedMsgLen, err := base64.StdEncoding.Decode(encryptedMsg, requestHttpBody.Base64EncryptedMsg)
			if err != nil {
				errorHandler.ServeError(w, r, err)
				return
			}
			encryptedMsg = encryptedMsg[:encryptedMsgLen]

			aesKey := srv.getCurrentAESKey()
			if aesKey == nil {
				err = errors.New("aes key was not set for Server, see NewServer function or Server.SetAESKey method")
				errorHandler.ServeError(w, r, err)
				return
			}
			random, msgPlaintext, haveAppIdBytes, err := util.AESDecryptMsg(encryptedMsg, aesKey)
			if err != nil {
				lastAESKey := srv.getLastAESKey()
				if lastAESKey == nil {
					errorHandler.ServeError(w, r, err)
					return
				}
				aesKey = lastAESKey // NOTE
				random, msgPlaintext, haveAppIdBytes, err = util.AESDecryptMsg(encryptedMsg, aesKey)
				if err != nil {
					errorHandler.ServeError(w, r, err)
					return
				}
			}
			callback.DebugPrintPlainRequestMessage(msgPlaintext)

			haveAppId := string(haveAppIdBytes)
			wantAppId := srv.appId
			if wantAppId != "" && !security.SecureCompareString(haveAppId, wantAppId) {
				err = fmt.Errorf("the message AppId mismatch, have: %s, want: %s", haveAppId, wantAppId)
				errorHandler.ServeError(w, r, err)
				return
			}

			var mixedMsg MixedMsg
			if err = xml.Unmarshal(msgPlaintext, &mixedMsg); err != nil {
				errorHandler.ServeError(w, r, err)
				return
			}
			if haveToUserName != mixedMsg.ToUserName {
				err = fmt.Errorf("the message ToUserName mismatch between ciphertext and plaintext, %q != %q",
					haveToUserName, mixedMsg.ToUserName)
				errorHandler.ServeError(w, r, err)
				return
			}

			ctx := &Context{
				ResponseWriter: w,
				Request:        r,

				QueryParams:  queryParams,
				EncryptType:  encryptType,
				MsgSignature: haveMsgSignature,
				Signature:    haveSignature,
				Timestamp:    timestamp,
				Nonce:        nonce,

				MsgCiphertext: requestHttpBody.Base64EncryptedMsg,
				MsgPlaintext:  msgPlaintext,
				MixedMsg:      &mixedMsg,

				Token:  srv.token,
				AESKey: aesKey,
				Random: random,
				AppId:  haveAppId,

				handlerIndex: initHandlerIndex,
			}
			srv.handler.ServeMsg(ctx)

		case "", "raw":
			haveSignature := queryParams.Get("signature")
			if haveSignature == "" {
				errorHandler.ServeError(w, r, errors.New("not found signature query parameter"))
				return
			}
			timestampString := queryParams.Get("timestamp")
			if timestampString == "" {
				errorHandler.ServeError(w, r, errors.New("not found timestamp query parameter"))
				return
			}
			timestamp, err := strconv.ParseInt(timestampString, 10, 64)
			if err != nil {
				err = fmt.Errorf("can not parse timestamp query parameter %q to int64", timestampString)
				errorHandler.ServeError(w, r, err)
				return
			}
			nonce := queryParams.Get("nonce")
			if nonce == "" {
				errorHandler.ServeError(w, r, errors.New("not found nonce query parameter"))
				return
			}

			wantSignature := util.Sign(srv.token, timestampString, nonce)
			if !security.SecureCompareString(haveSignature, wantSignature) {
				err = fmt.Errorf("check signature failed, have: %s, want: %s", haveSignature, wantSignature)
				errorHandler.ServeError(w, r, err)
				return
			}

			msgPlaintext, err := ioutil.ReadAll(r.Body)
			if err != nil {
				errorHandler.ServeError(w, r, err)
				return
			}
			callback.DebugPrintPlainRequestMessage(msgPlaintext)

			var mixedMsg MixedMsg
			if err = xml.Unmarshal(msgPlaintext, &mixedMsg); err != nil {
				errorHandler.ServeError(w, r, err)
				return
			}

			haveToUserName := mixedMsg.ToUserName
			wantToUserName := srv.oriId
			if wantToUserName != "" && !security.SecureCompareString(haveToUserName, wantToUserName) {
				err = fmt.Errorf("the message ToUserName mismatch, have: %s, want: %s",
					haveToUserName, wantToUserName)
				errorHandler.ServeError(w, r, err)
				return
			}

			ctx := &Context{
				ResponseWriter: w,
				Request:        r,

				QueryParams: queryParams,
				EncryptType: encryptType,
				Signature:   haveSignature,
				Timestamp:   timestamp,
				Nonce:       nonce,

				MsgPlaintext: msgPlaintext,
				MixedMsg:     &mixedMsg,

				Token: srv.token,

				handlerIndex: initHandlerIndex,
			}
			srv.handler.ServeMsg(ctx)

		default:
			errorHandler.ServeError(w, r, errors.New("unknown encrypt_type: "+encryptType))
		}

	case "GET": // 验证回调URL是否有效
		haveSignature := queryParams.Get("signature")
		if haveSignature == "" {
			errorHandler.ServeError(w, r, errors.New("not found signature query parameter"))
			return
		}
		timestamp := queryParams.Get("timestamp")
		if timestamp == "" {
			errorHandler.ServeError(w, r, errors.New("not found timestamp query parameter"))
			return
		}
		nonce := queryParams.Get("nonce")
		if nonce == "" {
			errorHandler.ServeError(w, r, errors.New("not found nonce query parameter"))
			return
		}
		echostr := queryParams.Get("echostr")
		if echostr == "" {
			errorHandler.ServeError(w, r, errors.New("not found echostr query parameter"))
			return
		}

		wantSignature := util.Sign(srv.token, timestamp, nonce)
		if !security.SecureCompareString(haveSignature, wantSignature) {
			err := fmt.Errorf("check signature failed, have: %s, want: %s", haveSignature, wantSignature)
			errorHandler.ServeError(w, r, err)
			return
		}
		io.WriteString(w, echostr)
	}
}
