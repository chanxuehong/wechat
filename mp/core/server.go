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

	"github.com/chanxuehong/wechat/util"
)

// Server 用于处理微信服务器推送过来的消息(事件).
//
// 通常一个 Server 实例处理一个公众号的消息(事件), 此时建议指定 oriId(原始ID), appId 用于约束消息(事件),
// 也可以一个 Server 实例处理多个公众号的消息(事件), 此时 oriId(原始ID), appId 必须设置为 "".
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
//  appId:        可选; 公众号的AppId, 如果设置了值则该Server只能处理 AppId 为该值的公众号的消息(事件);
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
			panic(fmt.Sprintf("Decode base64AESKey:%q failed", base64AESKey))
		}
	}

	if errorHandler == nil {
		errorHandler = DefaultErrorHandler
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

// ServeHTTP ===========================================================================================================

// 安全模式下微信服务器推送过来的消息(事件)的数据结构
type RequestHttpBody struct {
	XMLName struct{} `xml:"xml" json:"-"`

	ToUserName   string `xml:"ToUserName" json:"ToUserName"`
	EncryptedMsg string `xml:"Encrypt"    json:"Encrypt"`
}

func (srv *Server) ServeHTTP(w http.ResponseWriter, r *http.Request, queryParams url.Values) {
	if queryParams == nil {
		queryParams = r.URL.Query()
	}
	errorHandler := srv.errorHandler

	switch r.Method {
	case "POST": // 推送消息(事件)
		encryptType := queryParams.Get("encrypt_type")
		switch encryptType {
		case "aes": // 安全模式, 兼容模式
			signature := queryParams.Get("signature") // 只读取, 不做校验

			haveMsgSignature := queryParams.Get("msg_signature")
			if haveMsgSignature == "" {
				errorHandler.ServeError(w, r, errors.New("msg_signature is empty"))
				return
			}
			timestampString := queryParams.Get("timestamp")
			if timestampString == "" {
				errorHandler.ServeError(w, r, errors.New("timestamp is empty"))
				return
			}
			timestamp, err := strconv.ParseInt(timestampString, 10, 64)
			if err != nil {
				err = errors.New("can not parse timestamp to int64: " + timestampString)
				errorHandler.ServeError(w, r, err)
				return
			}
			nonce := queryParams.Get("nonce")
			if nonce == "" {
				errorHandler.ServeError(w, r, errors.New("nonce is empty"))
				return
			}

			var requestHttpBody RequestHttpBody
			if err := xml.NewDecoder(r.Body).Decode(&requestHttpBody); err != nil {
				errorHandler.ServeError(w, r, err)
				return
			}

			haveToUserName := requestHttpBody.ToUserName
			wantToUserName := srv.oriId
			if wantToUserName != "" && !security.SecureCompareString(haveToUserName, wantToUserName) {
				err := fmt.Errorf("the RequestHttpBody's ToUserName mismatch, have: %s, want: %s", haveToUserName, wantToUserName)
				errorHandler.ServeError(w, r, err)
				return
			}

			wantMsgSignature := util.MsgSign(srv.token, timestampString, nonce, requestHttpBody.EncryptedMsg)
			if !security.SecureCompareString(haveMsgSignature, wantMsgSignature) {
				err := fmt.Errorf("check msg_signature failed, have: %s, want: %s", haveMsgSignature, wantMsgSignature)
				errorHandler.ServeError(w, r, err)
				return
			}

			encryptedMsgBytes, err := base64.StdEncoding.DecodeString(requestHttpBody.EncryptedMsg)
			if err != nil {
				errorHandler.ServeError(w, r, err)
				return
			}
			aesKey := srv.getCurrentAESKey()
			if len(aesKey) == 0 {
				errorHandler.ServeError(w, r, errors.New("not found aes key"))
				return
			}
			random, msgPlaintext, haveAppIdBytes, err := util.AESDecryptMsg(encryptedMsgBytes, aesKey)
			if err != nil {
				lastAESKey := srv.getLastAESKey()
				if len(lastAESKey) == 0 {
					errorHandler.ServeError(w, r, err)
					return
				}
				aesKey = lastAESKey // NOTE

				random, msgPlaintext, haveAppIdBytes, err = util.AESDecryptMsg(encryptedMsgBytes, aesKey)
				if err != nil {
					errorHandler.ServeError(w, r, err)
					return
				}
			}
			haveAppId := string(haveAppIdBytes)
			wantAppId := srv.appId
			if wantAppId != "" && !security.SecureCompareString(haveAppId, wantAppId) {
				err := fmt.Errorf("the message's AppId mismatch, have: %s, want: %s", haveAppId, wantAppId)
				errorHandler.ServeError(w, r, err)
				return
			}

			var mixedMsg MixedMsg
			if err := xml.Unmarshal(msgPlaintext, &mixedMsg); err != nil {
				errorHandler.ServeError(w, r, err)
				return
			}
			if haveToUserName != mixedMsg.ToUserName {
				err := fmt.Errorf("the RequestHttpBody's ToUserName(==%s) mismatch the MixedMsg's ToUserName(==%s)", haveToUserName, mixedMsg.ToUserName)
				errorHandler.ServeError(w, r, err)
				return
			}

			ctx := &Context{
				ResponseWriter: w,
				Request:        r,

				QueryParams:  queryParams,
				EncryptType:  encryptType,
				MsgSignature: haveMsgSignature,
				Signature:    signature,
				Timestamp:    timestamp,
				Nonce:        nonce,

				MsgPlaintext: msgPlaintext,
				MixedMsg:     &mixedMsg,

				Token:  srv.token,
				AESKey: aesKey,
				Random: random,
				AppId:  haveAppId,

				handlerIndex: initHandlerIndex,
			}
			srv.handler.ServeMsg(ctx)

		case "", "raw": // 明文模式
			haveSignature := queryParams.Get("signature")
			if haveSignature == "" {
				errorHandler.ServeError(w, r, errors.New("signature is empty"))
				return
			}
			timestampString := queryParams.Get("timestamp")
			if timestampString == "" {
				errorHandler.ServeError(w, r, errors.New("timestamp is empty"))
				return
			}
			timestamp, err := strconv.ParseInt(timestampString, 10, 64)
			if err != nil {
				err = errors.New("can not parse timestamp to int64: " + timestampString)
				errorHandler.ServeError(w, r, err)
				return
			}
			nonce := queryParams.Get("nonce")
			if nonce == "" {
				errorHandler.ServeError(w, r, errors.New("nonce is empty"))
				return
			}

			wantSignature := util.Sign(srv.token, timestampString, nonce)
			if !security.SecureCompareString(haveSignature, wantSignature) {
				err := fmt.Errorf("check signature failed, have: %s, want: %s", haveSignature, wantSignature)
				errorHandler.ServeError(w, r, err)
				return
			}

			msgPlaintext, err := ioutil.ReadAll(r.Body)
			if err != nil {
				errorHandler.ServeError(w, r, err)
				return
			}

			var mixedMsg MixedMsg
			if err := xml.Unmarshal(msgPlaintext, &mixedMsg); err != nil {
				errorHandler.ServeError(w, r, err)
				return
			}

			haveToUserName := mixedMsg.ToUserName
			wantToUserName := srv.oriId
			if wantToUserName != "" && !security.SecureCompareString(haveToUserName, wantToUserName) {
				err := fmt.Errorf("the Message's ToUserName mismatch, have: %s, want: %s", haveToUserName, wantToUserName)
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

		default: // 未知的加密类型
			err := errors.New("unknown encrypt_type: " + encryptType)
			errorHandler.ServeError(w, r, err)
			return
		}

	case "GET": // 验证回调URL是否有效
		haveSignature := queryParams.Get("signature")
		if haveSignature == "" {
			errorHandler.ServeError(w, r, errors.New("signature is empty"))
			return
		}
		timestamp := queryParams.Get("timestamp")
		if timestamp == "" {
			errorHandler.ServeError(w, r, errors.New("timestamp is empty"))
			return
		}
		nonce := queryParams.Get("nonce")
		if nonce == "" {
			errorHandler.ServeError(w, r, errors.New("nonce is empty"))
			return
		}
		echostr := queryParams.Get("echostr")
		if echostr == "" {
			errorHandler.ServeError(w, r, errors.New("echostr is empty"))
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
