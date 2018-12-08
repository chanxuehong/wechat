package core

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"sync/atomic"
	"unicode"
	"unsafe"

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

	tokenBucketPtrMutex sync.Mutex     // used only by writers
	tokenBucketPtr      unsafe.Pointer // *tokenBucket

	aesKeyBucketPtrMutex sync.Mutex     // used only by writers
	aesKeyBucketPtr      unsafe.Pointer // *aesKeyBucket

	handler      Handler
	errorHandler ErrorHandler
}

func (srv *Server) OriId() string {
	return srv.oriId
}
func (srv *Server) AppId() string {
	return srv.appId
}

type tokenBucket struct {
	currentToken string
	lastToken    string
}

type aesKeyBucket struct {
	currentAESKey []byte
	lastAESKey    []byte
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
		oriId:           oriId,
		appId:           appId,
		tokenBucketPtr:  unsafe.Pointer(&tokenBucket{currentToken: token}),
		aesKeyBucketPtr: unsafe.Pointer(&aesKeyBucket{currentAESKey: aesKey}),
		handler:         handler,
		errorHandler:    errorHandler,
	}
}

func (srv *Server) getToken() (currentToken, lastToken string) {
	if p := (*tokenBucket)(atomic.LoadPointer(&srv.tokenBucketPtr)); p != nil {
		return p.currentToken, p.lastToken
	}
	return
}

// SetToken 设置签名token.
func (srv *Server) SetToken(token string) (err error) {
	if token == "" {
		return errors.New("empty token")
	}

	srv.tokenBucketPtrMutex.Lock()
	defer srv.tokenBucketPtrMutex.Unlock()

	currentToken, _ := srv.getToken()
	if token == currentToken {
		return
	}

	bucket := tokenBucket{
		currentToken: token,
		lastToken:    currentToken,
	}
	atomic.StorePointer(&srv.tokenBucketPtr, unsafe.Pointer(&bucket))
	return
}

func (srv *Server) removeLastToken(lastToken string) {
	srv.tokenBucketPtrMutex.Lock()
	defer srv.tokenBucketPtrMutex.Unlock()

	currentToken2, lastToken2 := srv.getToken()
	if lastToken != lastToken2 {
		return
	}

	bucket := tokenBucket{
		currentToken: currentToken2,
	}
	atomic.StorePointer(&srv.tokenBucketPtr, unsafe.Pointer(&bucket))
	return
}

func (srv *Server) getAESKey() (currentAESKey, lastAESKey []byte) {
	if p := (*aesKeyBucket)(atomic.LoadPointer(&srv.aesKeyBucketPtr)); p != nil {
		return p.currentAESKey, p.lastAESKey
	}
	return
}

// SetAESKey 设置aes加密解密key.
//  base64AESKey: aes加密解密key, 43字节长(base64编码, 去掉了尾部的'=').
func (srv *Server) SetAESKey(base64AESKey string) (err error) {
	if len(base64AESKey) != 43 {
		return errors.New("the length of base64AESKey must equal to 43")
	}
	aesKey, err := base64.StdEncoding.DecodeString(base64AESKey + "=")
	if err != nil {
		return
	}

	srv.aesKeyBucketPtrMutex.Lock()
	defer srv.aesKeyBucketPtrMutex.Unlock()

	currentAESKey, _ := srv.getAESKey()
	if bytes.Equal(aesKey, currentAESKey) {
		return
	}

	bucket := aesKeyBucket{
		currentAESKey: aesKey,
		lastAESKey:    currentAESKey,
	}
	atomic.StorePointer(&srv.aesKeyBucketPtr, unsafe.Pointer(&bucket))
	return
}

func (srv *Server) removeLastAESKey(lastAESKey []byte) {
	srv.aesKeyBucketPtrMutex.Lock()
	defer srv.aesKeyBucketPtrMutex.Unlock()

	currentAESKey2, lastAESKey2 := srv.getAESKey()
	if !bytes.Equal(lastAESKey, lastAESKey2) {
		return
	}

	bucket := aesKeyBucket{
		currentAESKey: currentAESKey2,
	}
	atomic.StorePointer(&srv.aesKeyBucketPtr, unsafe.Pointer(&bucket))
	return
}

// ServeHTTP 处理微信服务器的回调请求, query 参数可以为 nil.
func (srv *Server) ServeHTTP(w http.ResponseWriter, r *http.Request, query url.Values) {
	callback.DebugPrintRequest(r)
	if query == nil {
		query = r.URL.Query()
	}
	errorHandler := srv.errorHandler

	switch r.Method {
	case "POST": // 推送消息(事件)
		switch encryptType := query.Get("encrypt_type"); encryptType {
		case "aes":
			haveSignature := query.Get("signature")
			if haveSignature == "" {
				errorHandler.ServeError(w, r, errors.New("not found signature query parameter"))
				return
			}
			haveMsgSignature := query.Get("msg_signature")
			if haveMsgSignature == "" {
				errorHandler.ServeError(w, r, errors.New("not found msg_signature query parameter"))
				return
			}
			timestampString := query.Get("timestamp")
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
			nonce := query.Get("nonce")
			if nonce == "" {
				errorHandler.ServeError(w, r, errors.New("not found nonce query parameter"))
				return
			}

			var token string
			currentToken, lastToken := srv.getToken()
			if currentToken == "" {
				err = errors.New("token was not set for Server, see NewServer function or Server.SetToken method")
				errorHandler.ServeError(w, r, err)
				return
			}
			token = currentToken
			wantSignature := util.Sign(token, timestampString, nonce)
			if !security.SecureCompareString(haveSignature, wantSignature) {
				if lastToken == "" {
					err = fmt.Errorf("check signature failed, have: %s, want: %s", haveSignature, wantSignature)
					errorHandler.ServeError(w, r, err)
					return
				}
				token = lastToken
				wantSignature = util.Sign(token, timestampString, nonce)
				if !security.SecureCompareString(haveSignature, wantSignature) {
					err = fmt.Errorf("check signature failed, have: %s, want: %s", haveSignature, wantSignature)
					errorHandler.ServeError(w, r, err)
					return
				}
			} else {
				if lastToken != "" {
					srv.removeLastToken(lastToken)
				}
			}

			buffer := textBufferPool.Get().(*bytes.Buffer)
			buffer.Reset()
			defer textBufferPool.Put(buffer)

			if _, err = buffer.ReadFrom(r.Body); err != nil {
				errorHandler.ServeError(w, r, err)
				return
			}
			requestBodyBytes := buffer.Bytes()

			var requestHttpBody cipherRequestHttpBody
			if err = xmlUnmarshal(requestBodyBytes, &requestHttpBody); err != nil {
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

			wantMsgSignature := util.MsgSign(token, timestampString, nonce, string(requestHttpBody.Base64EncryptedMsg))
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

			var aesKey []byte
			currentAESKey, lastAESKey := srv.getAESKey()
			if currentAESKey == nil {
				err = errors.New("aes key was not set for Server, see NewServer function or Server.SetAESKey method")
				errorHandler.ServeError(w, r, err)
				return
			}
			aesKey = currentAESKey
			random, msgPlaintext, haveAppIdBytes, err := util.AESDecryptMsg(encryptedMsg, aesKey)
			if err != nil {
				if lastAESKey == nil {
					errorHandler.ServeError(w, r, err)
					return
				}
				aesKey = lastAESKey
				random, msgPlaintext, haveAppIdBytes, err = util.AESDecryptMsg(encryptedMsg, aesKey)
				if err != nil {
					errorHandler.ServeError(w, r, err)
					return
				}
			} else {
				if lastAESKey != nil {
					srv.removeLastAESKey(lastAESKey)
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

				QueryParams:  query,
				EncryptType:  encryptType,
				MsgSignature: haveMsgSignature,
				Signature:    haveSignature,
				Timestamp:    timestamp,
				Nonce:        nonce,

				MsgCiphertext: requestHttpBody.Base64EncryptedMsg,
				MsgPlaintext:  msgPlaintext,
				MixedMsg:      &mixedMsg,

				Token:  token,
				AESKey: aesKey,
				Random: random,
				AppId:  haveAppId,

				handlerIndex: initHandlerIndex,
			}
			srv.handler.ServeMsg(ctx)

		case "", "raw":
			haveSignature := query.Get("signature")
			if haveSignature == "" {
				errorHandler.ServeError(w, r, errors.New("not found signature query parameter"))
				return
			}
			timestampString := query.Get("timestamp")
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
			nonce := query.Get("nonce")
			if nonce == "" {
				errorHandler.ServeError(w, r, errors.New("not found nonce query parameter"))
				return
			}

			var token string
			currentToken, lastToken := srv.getToken()
			if currentToken == "" {
				err = errors.New("token was not set for Server, see NewServer function or Server.SetToken method")
				errorHandler.ServeError(w, r, err)
				return
			}
			token = currentToken
			wantSignature := util.Sign(token, timestampString, nonce)
			if !security.SecureCompareString(haveSignature, wantSignature) {
				if lastToken == "" {
					err = fmt.Errorf("check signature failed, have: %s, want: %s", haveSignature, wantSignature)
					errorHandler.ServeError(w, r, err)
					return
				}
				token = lastToken
				wantSignature = util.Sign(token, timestampString, nonce)
				if !security.SecureCompareString(haveSignature, wantSignature) {
					err = fmt.Errorf("check signature failed, have: %s, want: %s", haveSignature, wantSignature)
					errorHandler.ServeError(w, r, err)
					return
				}
			} else {
				if lastToken != "" {
					srv.removeLastToken(lastToken)
				}
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

				QueryParams: query,
				EncryptType: encryptType,
				Signature:   haveSignature,
				Timestamp:   timestamp,
				Nonce:       nonce,

				MsgPlaintext: msgPlaintext,
				MixedMsg:     &mixedMsg,

				Token: token,

				handlerIndex: initHandlerIndex,
			}
			srv.handler.ServeMsg(ctx)

		default:
			errorHandler.ServeError(w, r, errors.New("unknown encrypt_type: "+encryptType))
		}

	case "GET": // 验证回调URL是否有效
		haveSignature := query.Get("signature")
		if haveSignature == "" {
			errorHandler.ServeError(w, r, errors.New("not found signature query parameter"))
			return
		}
		timestamp := query.Get("timestamp")
		if timestamp == "" {
			errorHandler.ServeError(w, r, errors.New("not found timestamp query parameter"))
			return
		}
		nonce := query.Get("nonce")
		if nonce == "" {
			errorHandler.ServeError(w, r, errors.New("not found nonce query parameter"))
			return
		}
		echostr := query.Get("echostr")
		if echostr == "" {
			errorHandler.ServeError(w, r, errors.New("not found echostr query parameter"))
			return
		}

		var token string
		currentToken, lastToken := srv.getToken()
		if currentToken == "" {
			err := errors.New("token was not set for Server, see NewServer function or Server.SetToken method")
			errorHandler.ServeError(w, r, err)
			return
		}
		token = currentToken
		wantSignature := util.Sign(token, timestamp, nonce)
		if !security.SecureCompareString(haveSignature, wantSignature) {
			if lastToken == "" {
				err := fmt.Errorf("check signature failed, have: %s, want: %s", haveSignature, wantSignature)
				errorHandler.ServeError(w, r, err)
				return
			}
			token = lastToken
			wantSignature = util.Sign(token, timestamp, nonce)
			if !security.SecureCompareString(haveSignature, wantSignature) {
				err := fmt.Errorf("check signature failed, have: %s, want: %s", haveSignature, wantSignature)
				errorHandler.ServeError(w, r, err)
				return
			}
		} else {
			if lastToken != "" {
				srv.removeLastToken(lastToken)
			}
		}

		io.WriteString(w, echostr)
	}
}

// =====================================================================================================================

type cipherRequestHttpBody struct {
	XMLName            struct{} `xml:"xml"`
	ToUserName         string   `xml:"ToUserName"`
	Base64EncryptedMsg []byte   `xml:"Encrypt"`
}

var (
	msgStartElementLiteral = []byte("<xml>")
	msgEndElementLiteral   = []byte("</xml>")

	msgToUserNameStartElementLiteral = []byte("<ToUserName>")
	msgToUserNameEndElementLiteral   = []byte("</ToUserName>")

	msgEncryptStartElementLiteral = []byte("<Encrypt>")
	msgEncryptEndElementLiteral   = []byte("</Encrypt>")

	cdataStartLiteral = []byte("<![CDATA[")
	cdataEndLiteral   = []byte("]]>")
)

//<xml>
//    <ToUserName><![CDATA[gh_b1eb3f8bd6c6]]></ToUserName>
//    <Encrypt><![CDATA[DlCGq+lWQuyjNNK+vDaO0zUltpdUW3u4V00WCzsdNzmZGEhrU7TPxG52viOKCWYPwTMbCzgbCtakZHyNxr5hjoZJ7ORAUYoIAGQy/LDWtAnYgDO+ppKLp0rDq+67Dv3yt+vatMQTh99NII6x9SEGpY3O2h8RpG99+NYevQiOLVKqiQYzan21sX/jE4Y3wZaeudsb4QVjqzRAPaCJ5nS3T31uIR9fjSRgHTDRDOzjQ1cHchge+t6faUhniN5VQVTE+wIYtmnejc55BmHYPfBnTkYah9+cTYnI3diUPJRRiyVocJyHlb+XOZN22dsx9yzKHBAyagaoDIV8Yyb/PahcUbsqGv5wziOgLJQIa6z93/VY7d2Kq2C2oBS+Qb+FI9jLhgc3RvCi+Yno2X3cWoqbsRwoovYdyg6jme/H7nMZn77PSxOGRt/dYiWx2NuBAF7fNFigmbRiive3DyOumNCMvA==]]></Encrypt>
//</xml>
func xmlUnmarshal(data []byte, p *cipherRequestHttpBody) error {
	data = bytes.TrimSpace(data)
	if !bytes.HasPrefix(data, msgStartElementLiteral) || !bytes.HasSuffix(data, msgEndElementLiteral) {
		log.Printf("[WARNING] xmlUnmarshal failed, data:\n%s\n", data)
		return xml.Unmarshal(data, p)
	}
	data2 := data[len(msgStartElementLiteral) : len(data)-len(msgEndElementLiteral)]

	// ToUserName
	ToUserNameElementBytes := data2
	i := bytes.Index(ToUserNameElementBytes, msgToUserNameStartElementLiteral)
	if i == -1 {
		log.Printf("[WARNING] xmlUnmarshal failed, data:\n%s\n", data)
		return xml.Unmarshal(data, p)
	}
	ToUserNameElementBytes = ToUserNameElementBytes[i+len(msgToUserNameStartElementLiteral):]
	ToUserNameElementBytes = bytes.TrimLeftFunc(ToUserNameElementBytes, unicode.IsSpace)
	if !bytes.HasPrefix(ToUserNameElementBytes, cdataStartLiteral) {
		log.Printf("[WARNING] xmlUnmarshal failed, data:\n%s\n", data)
		return xml.Unmarshal(data, p)
	}
	ToUserNameElementBytes = ToUserNameElementBytes[len(cdataStartLiteral):]
	i = bytes.Index(ToUserNameElementBytes, cdataEndLiteral)
	if i == -1 {
		log.Printf("[WARNING] xmlUnmarshal failed, data:\n%s\n", data)
		return xml.Unmarshal(data, p)
	}
	ToUserName := ToUserNameElementBytes[:i]
	ToUserNameElementBytes = ToUserNameElementBytes[i+len(cdataEndLiteral):]
	ToUserNameElementBytes = bytes.TrimLeftFunc(ToUserNameElementBytes, unicode.IsSpace)
	if !bytes.HasPrefix(ToUserNameElementBytes, msgToUserNameEndElementLiteral) {
		log.Printf("[WARNING] xmlUnmarshal failed, data:\n%s\n", data)
		return xml.Unmarshal(data, p)
	}
	ToUserNameElementBytes = ToUserNameElementBytes[len(msgToUserNameEndElementLiteral):]

	// Encrypt
	EncryptElementBytes := ToUserNameElementBytes
	i = bytes.Index(EncryptElementBytes, msgEncryptStartElementLiteral)
	if i == -1 {
		EncryptElementBytes = data2
		i = bytes.Index(EncryptElementBytes, msgEncryptStartElementLiteral)
		if i == -1 {
			log.Printf("[WARNING] xmlUnmarshal failed, data:\n%s\n", data)
			return xml.Unmarshal(data, p)
		}
	}
	EncryptElementBytes = EncryptElementBytes[i+len(msgEncryptStartElementLiteral):]
	EncryptElementBytes = bytes.TrimLeftFunc(EncryptElementBytes, unicode.IsSpace)
	if !bytes.HasPrefix(EncryptElementBytes, cdataStartLiteral) {
		log.Printf("[WARNING] xmlUnmarshal failed, data:\n%s\n", data)
		return xml.Unmarshal(data, p)
	}
	EncryptElementBytes = EncryptElementBytes[len(cdataStartLiteral):]
	i = bytes.Index(EncryptElementBytes, cdataEndLiteral)
	if i == -1 {
		log.Printf("[WARNING] xmlUnmarshal failed, data:\n%s\n", data)
		return xml.Unmarshal(data, p)
	}
	Encrypt := EncryptElementBytes[:i]
	EncryptElementBytes = EncryptElementBytes[i+len(cdataEndLiteral):]
	EncryptElementBytes = bytes.TrimLeftFunc(EncryptElementBytes, unicode.IsSpace)
	if !bytes.HasPrefix(EncryptElementBytes, msgEncryptEndElementLiteral) {
		log.Printf("[WARNING] xmlUnmarshal failed, data:\n%s\n", data)
		return xml.Unmarshal(data, p)
	}

	p.ToUserName = string(ToUserName)
	p.Base64EncryptedMsg = Encrypt
	return nil
}
