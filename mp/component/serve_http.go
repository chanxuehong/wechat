// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package component

import (
	"bytes"
	"crypto/subtle"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/util"
)

var zeroAESKey [32]byte

// 微信服务器请求 http body
type RequestHttpBody struct {
	XMLName        struct{} `xml:"xml" json:"-"`
	ComponentAppId string   `xml:"AppId"`
	EncryptedMsg   string   `xml:"Encrypt"`
}

// ServeHTTP 处理 http 消息请求
//  NOTE: 调用者保证所有参数有效
func ServeHTTP(w http.ResponseWriter, r *http.Request, urlValues url.Values,
	componentServer ComponentServer, invalidRequestHandler mp.InvalidRequestHandler) {

	switch r.Method {
	case "POST": // 消息处理
		timestampStr, nonce, encryptType, msgSignature1, err := parsePostURLQuery(urlValues)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		switch encryptType {
		case "aes":
			break // 目前只有 aes, 所以可以這麼寫
		default:
			err := errors.New("unknown encrypt_type: " + encryptType)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		// 首先判断签名长度是否合法
		if len(msgSignature1) != 40 {
			err = fmt.Errorf("the length of msg_signature mismatch, have: %d, want: 40", len(msgSignature1))
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			err = errors.New("can not parse timestamp to int64: " + timestampStr)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		// 解析 RequestHttpBody
		var requestHttpBody RequestHttpBody
		if err := xml.NewDecoder(r.Body).Decode(&requestHttpBody); err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		haveComponentAppId := requestHttpBody.ComponentAppId
		wantComponentAppId := componentServer.ComponentAppId()
		if len(haveComponentAppId) != len(wantComponentAppId) {
			err = fmt.Errorf("the RequestHttpBody's AppId mismatch, have: %s, want: %s", haveComponentAppId, wantComponentAppId)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}
		if subtle.ConstantTimeCompare([]byte(haveComponentAppId), []byte(wantComponentAppId)) != 1 {
			err = fmt.Errorf("the RequestHttpBody's AppId mismatch, have: %s, want: %s", haveComponentAppId, wantComponentAppId)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		componentToken := componentServer.ComponentToken()

		// 验证签名
		msgSignature2 := util.MsgSign(componentToken, timestampStr, nonce, requestHttpBody.EncryptedMsg)
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

		AESKey := componentServer.CurrentAESKey()
		Random, RawMsgXML, err := util.AESDecryptMsg(EncryptedMsgBytes, wantComponentAppId, AESKey)
		if err != nil {
			// 尝试用上一次的 AESKey 来解密
			LastAESKey := componentServer.LastAESKey()
			if bytes.Equal(AESKey[:], LastAESKey[:]) || bytes.Equal(zeroAESKey[:], LastAESKey[:]) {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			AESKey = LastAESKey // NOTE
			Random, RawMsgXML, err = util.AESDecryptMsg(EncryptedMsgBytes, wantComponentAppId, AESKey)
			if err != nil {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}
		}

		// 解密成功, 解析 MixedMessage
		var MixedMsg MixedComponentMessage
		if err = xml.Unmarshal(RawMsgXML, &MixedMsg); err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		// 安全考虑再次验证
		if haveComponentAppId != MixedMsg.ComponentAppId {
			err = fmt.Errorf("the RequestHttpBody's AppId(==%s) mismatch the MixedMessage's ComponentAppId(==%s)", haveComponentAppId, MixedMsg.ComponentAppId)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		// 成功, 交给 ComponentMessageHandler
		r := &Request{
			HttpRequest: r,

			QueryValues:  urlValues,
			MsgSignature: msgSignature1,
			EncryptType:  encryptType,
			TimeStamp:    timestamp,
			Nonce:        nonce,

			RawMsgXML: RawMsgXML,
			MixedMsg:  &MixedMsg,

			AESKey: AESKey,
			Random: Random,

			ComponentAppId: haveComponentAppId,
			ComponentToken: componentToken,
		}
		componentServer.ComponentMessageHandler().ServeComponentMessage(w, r)

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

		signature2 := util.Sign(componentServer.ComponentToken(), timestamp, nonce)
		if subtle.ConstantTimeCompare([]byte(signature1), []byte(signature2)) != 1 {
			err = fmt.Errorf("check signature failed, input: %s, local: %s", signature1, signature2)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		io.WriteString(w, echostr)
	}
}
