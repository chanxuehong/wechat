// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package mp

import (
	"crypto/subtle"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/chanxuehong/wechat/util"
)

// 安全模式, 微信服务器推送过来的 http body
type RequestHttpBody struct {
	XMLName struct{} `xml:"xml" json:"-"`

	ToUserName   string `xml:"ToUserName" json:"ToUserName"`
	EncryptedMsg string `xml:"Encrypt"    json:"Encrypt"`
}

// ServeHTTP 处理 http 消息请求
//  NOTE: 调用者保证所有参数有效
func ServeHTTP(w http.ResponseWriter, r *http.Request, queryValues url.Values, wechatServer WechatServer, invalidRequestHandler InvalidRequestHandler) {
	switch r.Method {
	case "POST": // 消息处理
		switch encryptType := queryValues.Get("encrypt_type"); encryptType {
		case "aes": // 安全模式, 兼容模式
			signature := queryValues.Get("signature") // 只讀取, 不驗證了

			msgSignature1 := queryValues.Get("msg_signature")
			if msgSignature1 == "" {
				invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("msg_signature is empty"))
				return
			}
			if len(msgSignature1) != 40 { // sha1
				err := fmt.Errorf("the length of msg_signature mismatch, have: %d, want: 40", len(msgSignature1))
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			timestampStr := queryValues.Get("timestamp")
			if timestampStr == "" {
				invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("timestamp is empty"))
				return
			}

			timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
			if err != nil {
				err = errors.New("can not parse timestamp to int64: " + timestampStr)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			nonce := queryValues.Get("nonce")
			if nonce == "" {
				invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("nonce is empty"))
				return
			}

			var requestHttpBody RequestHttpBody
			if err := xml.NewDecoder(r.Body).Decode(&requestHttpBody); err != nil {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			// 安全考虑验证下 ToUserName
			haveToUserName := requestHttpBody.ToUserName
			wantToUserName := wechatServer.OriId()
			if len(haveToUserName) != len(wantToUserName) {
				err = fmt.Errorf("the RequestHttpBody's ToUserName mismatch, have: %s, want: %s", haveToUserName, wantToUserName)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}
			if subtle.ConstantTimeCompare([]byte(haveToUserName), []byte(wantToUserName)) != 1 {
				err = fmt.Errorf("the RequestHttpBody's ToUserName mismatch, have: %s, want: %s", haveToUserName, wantToUserName)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			wechatToken := wechatServer.Token()

			// 验证签名
			msgSignature2 := util.MsgSign(wechatToken, timestampStr, nonce, requestHttpBody.EncryptedMsg)
			if subtle.ConstantTimeCompare([]byte(msgSignature1), []byte(msgSignature2)) != 1 {
				err = fmt.Errorf("check msg_signature failed, input: %s, local: %s", msgSignature1, msgSignature2)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			// 解密
			encryptedMsgBytes, err := base64.StdEncoding.DecodeString(requestHttpBody.EncryptedMsg)
			if err != nil {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			wechatAppId := wechatServer.AppId()
			aesKey := wechatServer.CurrentAESKey()

			random, rawMsgXML, err := util.AESDecryptMsg(encryptedMsgBytes, wechatAppId, aesKey)
			if err != nil {
				// 尝试用上一次的 AESKey 来解密
				lastAESKey, isLastAESKeyValid := wechatServer.LastAESKey()
				if !isLastAESKeyValid {
					invalidRequestHandler.ServeInvalidRequest(w, r, err)
					return
				}

				aesKey = lastAESKey // NOTE

				random, rawMsgXML, err = util.AESDecryptMsg(encryptedMsgBytes, wechatAppId, aesKey)
				if err != nil {
					invalidRequestHandler.ServeInvalidRequest(w, r, err)
					return
				}
			}

			// 解密成功, 解析 MixedMessage
			var mixedMsg MixedMessage
			if err := xml.Unmarshal(rawMsgXML, &mixedMsg); err != nil {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			// 安全考虑再次验证 ToUserName
			if haveToUserName != mixedMsg.ToUserName {
				err = fmt.Errorf("the RequestHttpBody's ToUserName(==%s) mismatch the MixedMessage's ToUserName(==%s)", haveToUserName, mixedMsg.ToUserName)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			// 成功, 交给 MessageHandler
			r := &Request{
				HttpRequest: r,

				QueryValues: queryValues,
				Signature:   signature,
				TimeStamp:   timestamp,
				Nonce:       nonce,

				RawMsgXML: rawMsgXML,
				MixedMsg:  &mixedMsg,

				MsgSignature: msgSignature1,
				EncryptType:  encryptType,
				AESKey:       aesKey,
				Random:       random,

				WechatOriId: haveToUserName,
				WechatToken: wechatToken,
				WechatAppId: wechatAppId,
			}
			wechatServer.MessageHandler().ServeMessage(w, r)

		case "", "raw": // 明文模式
			signature1 := queryValues.Get("signature")
			if signature1 == "" {
				invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("signature is empty"))
				return
			}
			if len(signature1) != 40 { // sha1
				err := fmt.Errorf("the length of signature mismatch, have: %d, want: 40", len(signature1))
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			timestampStr := queryValues.Get("timestamp")
			if timestampStr == "" {
				invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("timestamp is empty"))
				return
			}

			timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
			if err != nil {
				err = errors.New("can not parse timestamp to int64: " + timestampStr)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			nonce := queryValues.Get("nonce")
			if nonce == "" {
				invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("nonce is empty"))
				return
			}

			wechatToken := wechatServer.Token()

			signature2 := util.Sign(wechatToken, timestampStr, nonce)
			if subtle.ConstantTimeCompare([]byte(signature1), []byte(signature2)) != 1 {
				err = fmt.Errorf("check signature failed, input: %s, local: %s", signature1, signature2)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			// 验证签名成功, 解析 MixedMessage
			rawMsgXML, err := ioutil.ReadAll(r.Body)
			if err != nil {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			var mixedMsg MixedMessage
			if err := xml.Unmarshal(rawMsgXML, &mixedMsg); err != nil {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			// 安全考虑验证 ToUserName
			haveToUserName := mixedMsg.ToUserName
			wantToUserName := wechatServer.OriId()
			if len(haveToUserName) != len(wantToUserName) {
				err = fmt.Errorf("the message's ToUserName mismatch, have: %s, want: %s", haveToUserName, wantToUserName)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}
			if subtle.ConstantTimeCompare([]byte(haveToUserName), []byte(wantToUserName)) != 1 {
				err = fmt.Errorf("the message's ToUserName mismatch, have: %s, want: %s", haveToUserName, wantToUserName)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			// 成功, 交给 MessageHandler
			r := &Request{
				HttpRequest: r,

				QueryValues: queryValues,
				Signature:   signature1,
				TimeStamp:   timestamp,
				Nonce:       nonce,

				RawMsgXML: rawMsgXML,
				MixedMsg:  &mixedMsg,

				EncryptType: encryptType,

				WechatOriId: haveToUserName,
				WechatAppId: wechatServer.AppId(),
				WechatToken: wechatToken,
			}
			wechatServer.MessageHandler().ServeMessage(w, r)

		default: // 未知的加密类型
			err := errors.New("unknown encrypt_type: " + encryptType)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

	case "GET": // 首次验证
		signature1 := queryValues.Get("signature")
		if signature1 == "" {
			invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("signature is empty"))
			return
		}
		if len(signature1) != 40 { // sha1
			err := fmt.Errorf("the length of signature mismatch, have: %d, want: 40", len(signature1))
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		timestamp := queryValues.Get("timestamp")
		if timestamp == "" {
			invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("timestamp is empty"))
			return
		}

		nonce := queryValues.Get("nonce")
		if nonce == "" {
			invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("nonce is empty"))
			return
		}

		echostr := queryValues.Get("echostr")
		if echostr == "" {
			invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("echostr is empty"))
			return
		}

		signature2 := util.Sign(wechatServer.Token(), timestamp, nonce)
		if subtle.ConstantTimeCompare([]byte(signature1), []byte(signature2)) != 1 {
			err := fmt.Errorf("check signature failed, input: %s, local: %s", signature1, signature2)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		io.WriteString(w, echostr)
	}
}
