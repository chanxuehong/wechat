// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package mp

import (
	"bytes"
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

var zeroAESKey [32]byte

// 安全模式 和 兼容模式, 微信服务器推送过来的 http body
type RequestHttpBody struct {
	XMLName      struct{} `xml:"xml" json:"-"`
	MixedMessage          // ToUserName 始终有效
	EncryptedMsg string   `xml:"Encrypt" json:"Encrypt"`
}

// ServeHTTP 处理 http 消息请求
//  NOTE: 调用者保证所有参数有效
func ServeHTTP(w http.ResponseWriter, r *http.Request, urlValues url.Values,
	wechatServer WechatServer, invalidRequestHandler InvalidRequestHandler) {

	switch r.Method {
	case "POST": // 消息处理
		signature1, timestampStr, nonce, encryptType, msgSignature1, err := parsePostURLQuery(urlValues)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			err = errors.New("can not parse timestamp to int64: " + timestampStr)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		switch encryptType {
		case "aes": // 兼容模式, 安全模式
			//if len(signature1) != 40 {
			//	err = fmt.Errorf("the length of signature mismatch, have: %d, want: 40", len(signature1))
			//	invalidRequestHandler.ServeInvalidRequest(w, r, err)
			//	return
			//}

			//WechatToken := wechatServer.Token()
			//signature2 := util.Sign(WechatToken, timestampStr, nonce)
			//if subtle.ConstantTimeCompare([]byte(signature1), []byte(signature2)) != 1 {
			//	err = fmt.Errorf("check signature failed, input: %s, local: %s", signature1, signature2)
			//	invalidRequestHandler.ServeInvalidRequest(w, r, err)
			//	return
			//}

			// 首先验证密文签名长度
			if len(msgSignature1) != 40 {
				err = fmt.Errorf("the length of msg_signature mismatch, have: %d, want: 40", len(msgSignature1))
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			var requestHttpBody RequestHttpBody
			if err := xml.NewDecoder(r.Body).Decode(&requestHttpBody); err != nil {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			// 安全考虑验证下 ToUserName
			haveToUserName := requestHttpBody.ToUserName
			wantToUserName := wechatServer.WechatId()
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
				err = fmt.Errorf("check signature failed, input: %s, local: %s", msgSignature1, msgSignature2)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			// 解密
			EncryptedMsgBytes, err := base64.StdEncoding.DecodeString(requestHttpBody.EncryptedMsg)
			if err != nil {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			WechatAppId := wechatServer.AppId()
			AESKey := wechatServer.CurrentAESKey()
			Random, RawMsgXML, err := util.AESDecryptMsg(EncryptedMsgBytes, WechatAppId, AESKey)
			if err != nil {
				// 尝试用上一次的 AESKey 来解密
				LastAESKey := wechatServer.LastAESKey()
				if bytes.Equal(zeroAESKey[:], LastAESKey[:]) || bytes.Equal(AESKey[:], LastAESKey[:]) {
					invalidRequestHandler.ServeInvalidRequest(w, r, err)
					return
				}

				AESKey = LastAESKey // NOTE
				Random, RawMsgXML, err = util.AESDecryptMsg(EncryptedMsgBytes, WechatAppId, AESKey)
				if err != nil {
					invalidRequestHandler.ServeInvalidRequest(w, r, err)
					return
				}
			}

			// 解密成功, 解析 MixedMessage
			var MixedMsg MixedMessage
			if err = xml.Unmarshal(RawMsgXML, &MixedMsg); err != nil {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			// 安全考虑再次验证 ToUserName
			if haveToUserName != MixedMsg.ToUserName {
				err = fmt.Errorf("the RequestHttpBody's ToUserName(==%s) mismatch the MixedMessage's ToUserName(==%s)", haveToUserName, MixedMsg.ToUserName)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			// 成功, 交给 MessageHandler
			r := &Request{
				HttpRequest: r,

				Signature: signature1,
				TimeStamp: timestamp,
				Nonce:     nonce,
				RawMsgXML: RawMsgXML,
				MixedMsg:  &MixedMsg,

				MsgSignature: msgSignature1,
				EncryptType:  encryptType,
				AESKey:       AESKey,
				Random:       Random,

				WechatId:    wantToUserName,
				WechatToken: wechatToken,
				WechatAppId: WechatAppId,
			}
			wechatServer.MessageHandler().ServeMessage(w, r)

		case "", "raw": // 明文模式
			// 首先验证签名
			if len(signature1) != 40 {
				err = fmt.Errorf("the length of signature mismatch, have: %d, want: 40", len(signature1))
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			WechatToken := wechatServer.Token()
			signature2 := util.Sign(WechatToken, timestampStr, nonce)
			if subtle.ConstantTimeCompare([]byte(signature1), []byte(signature2)) != 1 {
				err = fmt.Errorf("check signature failed, input: %s, local: %s", signature1, signature2)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			// 验证签名成功, 解析 MixedMessage
			RawMsgXML, err := ioutil.ReadAll(r.Body)
			if err != nil {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			var MixedMsg MixedMessage
			if err := xml.Unmarshal(RawMsgXML, &MixedMsg); err != nil {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			// 安全考虑验证 ToUserName
			haveToUserName := MixedMsg.ToUserName
			wantToUserName := wechatServer.WechatId()
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

				Signature: signature1,
				TimeStamp: timestamp,
				Nonce:     nonce,
				RawMsgXML: RawMsgXML,
				MixedMsg:  &MixedMsg,

				WechatId:    wantToUserName,
				WechatToken: WechatToken,
				WechatAppId: wechatServer.AppId(),
			}
			wechatServer.MessageHandler().ServeMessage(w, r)

		default: // 未知的加密类型
			err := errors.New("unknown encrypt_type: " + encryptType)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

	case "GET": // 首次验证
		signature1, timestamp, nonce, echostr, err := parseGetURLQuery(urlValues)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		if len(signature1) != 40 {
			err = fmt.Errorf("the length of signature mismatch, have: %d, want: 40", len(signature1))
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		signature2 := util.Sign(wechatServer.Token(), timestamp, nonce)
		if subtle.ConstantTimeCompare([]byte(signature1), []byte(signature2)) != 1 {
			err = fmt.Errorf("check signature failed, input: %s, local: %s", signature1, signature2)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		io.WriteString(w, echostr)
	}
}
