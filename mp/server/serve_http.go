// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

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

	"github.com/chanxuehong/wechat/mp/message/passive/request"
	"github.com/chanxuehong/wechat/util"
)

var zeroAESKey [32]byte

// ServeHTTP 处理 http 消息请求
//  NOTE: 确保所有参数合法, r.Body 能正确读取数据
func ServeHTTP(w http.ResponseWriter, r *http.Request,
	urlValues url.Values, agent Agent, invalidRequestHandler InvalidRequestHandler) {

	switch r.Method {
	case "POST": // 消息处理
		signature1, timestampStr, nonce, encryptType, msgSignature1, err := parsePostURLQuery(urlValues)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			err = fmt.Errorf("can not parse timestamp(==%q) to int64, error: %s", timestampStr, err.Error())
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

			//signature2 := util.Sign(agent.GetToken(), timestampStr, nonce)
			//if subtle.ConstantTimeCompare([]byte(signature1), []byte(signature2)) != 1 {
			//	err = fmt.Errorf("check signature failed, input: %s, local: %s", signature1, signature2)
			//	invalidRequestHandler.ServeInvalidRequest(w, r, err)
			//	return
			//}

			if len(msgSignature1) != 40 {
				err = fmt.Errorf("the length of msg_signature mismatch, have: %d, want: 40", len(msgSignature1))
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			var requestHttpBody request.RequestHttpBody
			if err := xml.NewDecoder(r.Body).Decode(&requestHttpBody); err != nil {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			haveToUserName := requestHttpBody.ToUserName
			wantToUserName := agent.GetId()
			if len(haveToUserName) != len(wantToUserName) {
				err = fmt.Errorf("the message RequestHttpBody's ToUserName mismatch, have: %s, want: %s", haveToUserName, wantToUserName)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}
			if subtle.ConstantTimeCompare([]byte(haveToUserName), []byte(wantToUserName)) != 1 {
				err = fmt.Errorf("the message RequestHttpBody's ToUserName mismatch, have: %s, want: %s", haveToUserName, wantToUserName)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			msgSignature2 := util.MsgSign(agent.GetToken(), timestampStr, nonce, requestHttpBody.EncryptedMsg)
			if subtle.ConstantTimeCompare([]byte(msgSignature1), []byte(msgSignature2)) != 1 {
				err = fmt.Errorf("check signature failed, input: %s, local: %s", msgSignature1, msgSignature2)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			EncryptedMsgBytes, err := base64.StdEncoding.DecodeString(requestHttpBody.EncryptedMsg)
			if err != nil {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			AESKey := agent.GetCurrentAESKey()

			random, rawXMLMsg, err := util.AESDecryptMsg(EncryptedMsgBytes, agent.GetAppId(), AESKey)
			if err != nil {
				// 尝试上一个 AESKey
				LastAESKey := agent.GetLastAESKey()
				if bytes.Equal(zeroAESKey[:], LastAESKey[:]) || bytes.Equal(AESKey[:], LastAESKey[:]) {
					invalidRequestHandler.ServeInvalidRequest(w, r, err)
					return
				}

				AESKey = LastAESKey // !!!

				random, rawXMLMsg, err = util.AESDecryptMsg(EncryptedMsgBytes, agent.GetAppId(), AESKey)
				if err != nil {
					invalidRequestHandler.ServeInvalidRequest(w, r, err)
					return
				}
			}

			var msgReq request.Request
			if err = xml.Unmarshal(rawXMLMsg, &msgReq); err != nil {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			if haveToUserName != msgReq.ToUserName {
				err = fmt.Errorf("the RequestHttpBody's ToUserName(==%s) mismatch the Request's ToUserName(==%s)", haveToUserName, msgReq.ToUserName)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			aesMsgDispatch(w, r, &msgReq, rawXMLMsg, timestamp, nonce, AESKey, random, agent)

		case "", "raw": // 明文模式
			if len(signature1) != 40 {
				err = fmt.Errorf("the length of signature mismatch, have: %d, want: 40", len(signature1))
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			signature2 := util.Sign(agent.GetToken(), timestampStr, nonce)
			if subtle.ConstantTimeCompare([]byte(signature1), []byte(signature2)) != 1 {
				err = fmt.Errorf("check signature failed, input: %s, local: %s", signature1, signature2)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			rawXMLMsg, err := ioutil.ReadAll(r.Body)
			if err != nil {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			var msgReq request.Request
			if err = xml.Unmarshal(rawXMLMsg, &msgReq); err != nil {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			wantToUserName := agent.GetId()
			if len(msgReq.ToUserName) != len(wantToUserName) {
				err = fmt.Errorf("the message Request's ToUserName mismatch, have: %s, want: %s", msgReq.ToUserName, wantToUserName)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}
			if subtle.ConstantTimeCompare([]byte(msgReq.ToUserName), []byte(wantToUserName)) != 1 {
				err = fmt.Errorf("the message Request's ToUserName mismatch, have: %s, want: %s", msgReq.ToUserName, wantToUserName)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			rawMsgDispatch(w, r, &msgReq, rawXMLMsg, timestamp, agent)

		default: // 未知的加密类型
			invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("unknown encrypt_type"))
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

		signature2 := util.Sign(agent.GetToken(), timestamp, nonce)
		if subtle.ConstantTimeCompare([]byte(signature1), []byte(signature2)) != 1 {
			err = fmt.Errorf("check signature failed, input: %s, local: %s", signature1, signature2)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		io.WriteString(w, echostr)
	}
}
