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
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/chanxuehong/wechat/util"
)

var zeroAESKey [32]byte

// ServeHTTP 处理 http 消息请求
//  NOTE: 确保所有参数合法, r.Body 能正确读取数据
func ServeHTTP(w http.ResponseWriter, r *http.Request,
	urlValues url.Values, agent Agent, invalidRequestHandler InvalidRequestHandler) {

	switch r.Method {
	case "POST": // 消息处理
		msgSignature1, timestampStr, nonce, err := parsePostURLQuery(urlValues)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		if len(msgSignature1) != 40 {
			err = fmt.Errorf("the length of msg_signature mismatch, have: %d, want: 40", len(msgSignature1))
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			err = fmt.Errorf("can not parse timestamp(==%q) to int64, error: %s", timestampStr, err.Error())
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		EncryptedMsg, err := ioutil.ReadAll(r.Body)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		msgSignature2 := util.MsgSign(agent.GetToken(), timestampStr, nonce, string(EncryptedMsg))
		if subtle.ConstantTimeCompare([]byte(msgSignature1), []byte(msgSignature2)) != 1 {
			err = fmt.Errorf("check signature failed, input: %s, local: %s", msgSignature1, msgSignature2)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		EncryptedMsgBytes := make([]byte, base64.StdEncoding.DecodedLen(len(EncryptedMsg)))
		n, err := base64.StdEncoding.Decode(EncryptedMsgBytes, EncryptedMsg)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}
		EncryptedMsgBytes = EncryptedMsgBytes[:n]

		AESKey := agent.GetCurrentAESKey()

		random, rawXMLMsg, err := util.AESDecryptMsg(EncryptedMsgBytes, agent.GetSuiteId(), AESKey)
		if err != nil {
			// 尝试上一个 AESKey
			LastAESKey := agent.GetLastAESKey()
			if bytes.Equal(zeroAESKey[:], LastAESKey[:]) || bytes.Equal(AESKey[:], LastAESKey[:]) {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			AESKey = LastAESKey // !!!

			random, rawXMLMsg, err = util.AESDecryptMsg(EncryptedMsgBytes, agent.GetSuiteId(), AESKey)
			if err != nil {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}
		}

		var mixedReq mixedRequest
		if err := xml.Unmarshal(rawXMLMsg, &mixedReq); err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		haveSuiteId := mixedReq.SuiteId
		wantSuiteId := agent.GetSuiteId()
		if len(haveSuiteId) != len(wantSuiteId) {
			err = fmt.Errorf("the request SuiteId mismatch, have: %s, want: %s", haveSuiteId, wantSuiteId)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}
		if subtle.ConstantTimeCompare([]byte(haveSuiteId), []byte(wantSuiteId)) != 1 {
			err = fmt.Errorf("the request SuiteId mismatch, have: %s, want: %s", haveSuiteId, wantSuiteId)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		inputPara := &InputParameters{
			w:         w,
			r:         r,
			rawXMLMsg: rawXMLMsg,
			timestamp: timestamp,
			nonce:     nonce,
			AESKey:    AESKey,
			random:    random,
		}
		msgDispatch(inputPara, &mixedReq, agent)

	case "GET": // 首次验证
		msgSignature1, timestamp, nonce, encryptedMsg, err := parseGetURLQuery(urlValues)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		if len(msgSignature1) != 40 {
			err = fmt.Errorf("the length of msg_signature mismatch, have: %d, want: 40", len(msgSignature1))
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		msgSignature2 := util.MsgSign(agent.GetToken(), timestamp, nonce, encryptedMsg)
		if subtle.ConstantTimeCompare([]byte(msgSignature1), []byte(msgSignature2)) != 1 {
			err = fmt.Errorf("check signature failed, input: %s, local: %s", msgSignature1, msgSignature2)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		EncryptedMsgBytes, err := base64.StdEncoding.DecodeString(encryptedMsg)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		AESKey := agent.GetCurrentAESKey()

		_, echostr, err := util.AESDecryptMsg(EncryptedMsgBytes, agent.GetSuiteId(), AESKey)
		if err != nil {
			// 尝试上一个 AESKey
			LastAESKey := agent.GetLastAESKey()
			if bytes.Equal(zeroAESKey[:], LastAESKey[:]) || bytes.Equal(AESKey[:], LastAESKey[:]) {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			AESKey = LastAESKey // !!!

			_, echostr, err = util.AESDecryptMsg(EncryptedMsgBytes, agent.GetSuiteId(), AESKey)
			if err != nil {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}
		}

		w.Write(echostr)
	}
}
