// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

// +build wechatdebug

package component

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

	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/util"
)

// 微信服务器请求 http body
type RequestHttpBody struct {
	XMLName struct{} `xml:"xml" json:"-"`

	AppId        string `xml:"AppId"`
	EncryptedMsg string `xml:"Encrypt"`
}

// ServeHTTP 处理 http 消息请求
//  NOTE: 调用者保证所有参数有效
func ServeHTTP(w http.ResponseWriter, r *http.Request, queryValues url.Values, srv Server, irh mp.InvalidRequestHandler) {
	mp.LogInfoln("[WECHAT_DEBUG] request uri:", r.RequestURI)
	mp.LogInfoln("[WECHAT_DEBUG] request remote-addr:", r.RemoteAddr)
	mp.LogInfoln("[WECHAT_DEBUG] request user-agent:", r.UserAgent())

	switch r.Method {
	case "POST": // 消息处理
		switch encryptType := queryValues.Get("encrypt_type"); encryptType {
		case "aes":
			msgSignature1 := queryValues.Get("msg_signature")
			if msgSignature1 == "" {
				irh.ServeInvalidRequest(w, r, errors.New("msg_signature is empty"))
				return
			}
			if len(msgSignature1) != 40 { // sha1
				err := fmt.Errorf("the length of msg_signature mismatch, have: %d, want: 40", len(msgSignature1))
				irh.ServeInvalidRequest(w, r, err)
				return
			}

			timestampStr := queryValues.Get("timestamp")
			if timestampStr == "" {
				irh.ServeInvalidRequest(w, r, errors.New("timestamp is empty"))
				return
			}

			timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
			if err != nil {
				err = errors.New("can not parse timestamp to int64: " + timestampStr)
				irh.ServeInvalidRequest(w, r, err)
				return
			}

			nonce := queryValues.Get("nonce")
			if nonce == "" {
				irh.ServeInvalidRequest(w, r, errors.New("nonce is empty"))
				return
			}

			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				irh.ServeInvalidRequest(w, r, err)
				return
			}
			mp.LogInfoln("[WECHAT_DEBUG] request msg http body:\r\n", string(reqBody))

			var requestHttpBody RequestHttpBody
			if err := xml.Unmarshal(reqBody, &requestHttpBody); err != nil {
				irh.ServeInvalidRequest(w, r, err)
				return
			}

			appId := srv.AppId()

			// 安全考虑验证下 AppId
			haveAppId := requestHttpBody.AppId
			if len(haveAppId) != len(appId) {
				err = fmt.Errorf("the RequestHttpBody's AppId mismatch, have: %s, want: %s", haveAppId, appId)
				irh.ServeInvalidRequest(w, r, err)
				return
			}
			if subtle.ConstantTimeCompare([]byte(haveAppId), []byte(appId)) != 1 {
				err = fmt.Errorf("the RequestHttpBody's AppId mismatch, have: %s, want: %s", haveAppId, appId)
				irh.ServeInvalidRequest(w, r, err)
				return
			}

			token := srv.Token()

			// 验证签名
			msgSignature2 := util.MsgSign(token, timestampStr, nonce, requestHttpBody.EncryptedMsg)
			if subtle.ConstantTimeCompare([]byte(msgSignature1), []byte(msgSignature2)) != 1 {
				err = fmt.Errorf("check msg_signature failed, input: %s, local: %s", msgSignature1, msgSignature2)
				irh.ServeInvalidRequest(w, r, err)
				return
			}

			// 解密
			encryptedMsgBytes, err := base64.StdEncoding.DecodeString(requestHttpBody.EncryptedMsg)
			if err != nil {
				irh.ServeInvalidRequest(w, r, err)
				return
			}

			aesKey := srv.CurrentAESKey()

			random, rawMsgXML, err := util.AESDecryptMsg(encryptedMsgBytes, appId, aesKey)
			if err != nil {
				// 尝试用上一次的 AESKey 来解密
				lastAESKey, isLastAESKeyValid := srv.LastAESKey()
				if !isLastAESKeyValid {
					irh.ServeInvalidRequest(w, r, err)
					return
				}

				aesKey = lastAESKey // NOTE

				random, rawMsgXML, err = util.AESDecryptMsg(encryptedMsgBytes, appId, aesKey)
				if err != nil {
					irh.ServeInvalidRequest(w, r, err)
					return
				}
			}

			mp.LogInfoln("[WECHAT_DEBUG] request msg raw xml:\r\n", string(rawMsgXML))

			// 解密成功, 解析 MixedMessage
			var mixedMsg MixedMessage
			if err := xml.Unmarshal(rawMsgXML, &mixedMsg); err != nil {
				irh.ServeInvalidRequest(w, r, err)
				return
			}

			// 安全考虑再次验证 AppId
			if haveAppId != mixedMsg.AppId {
				err = fmt.Errorf("the RequestHttpBody's AppId(==%s) mismatch the MixedMessage's AppId(==%s)", haveAppId, mixedMsg.AppId)
				irh.ServeInvalidRequest(w, r, err)
				return
			}

			// 成功, 交给 MessageHandler
			r := &Request{
				HttpRequest: r,

				QueryValues:  queryValues,
				MsgSignature: msgSignature1,
				EncryptType:  encryptType,
				Timestamp:    timestamp,
				Nonce:        nonce,

				RawMsgXML: rawMsgXML,
				MixedMsg:  &mixedMsg,

				AESKey: aesKey,
				Random: random,

				AppId: appId,
				Token: token,
			}
			srv.MessageHandler().ServeMessage(w, r)

		default: // 未知的加密类型
			err := errors.New("unknown encrypt_type: " + encryptType)
			irh.ServeInvalidRequest(w, r, err)
			return
		}

	case "GET": // 首次验证
		signature1 := queryValues.Get("signature")
		if signature1 == "" {
			irh.ServeInvalidRequest(w, r, errors.New("signature is empty"))
			return
		}
		if len(signature1) != 40 { // sha1
			err := fmt.Errorf("the length of signature mismatch, have: %d, want: 40", len(signature1))
			irh.ServeInvalidRequest(w, r, err)
			return
		}

		timestamp := queryValues.Get("timestamp")
		if timestamp == "" {
			irh.ServeInvalidRequest(w, r, errors.New("timestamp is empty"))
			return
		}

		nonce := queryValues.Get("nonce")
		if nonce == "" {
			irh.ServeInvalidRequest(w, r, errors.New("nonce is empty"))
			return
		}

		echostr := queryValues.Get("echostr")
		if echostr == "" {
			irh.ServeInvalidRequest(w, r, errors.New("echostr is empty"))
			return
		}

		signature2 := util.Sign(srv.Token(), timestamp, nonce)
		if subtle.ConstantTimeCompare([]byte(signature1), []byte(signature2)) != 1 {
			err := fmt.Errorf("check signature failed, input: %s, local: %s", signature1, signature2)
			irh.ServeInvalidRequest(w, r, err)
			return
		}

		io.WriteString(w, echostr)
	}
}
