// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"bytes"
	"crypto/subtle"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/chanxuehong/util"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request, urlValues url.Values,
	messageServer MessageServer, invalidRequestHandler InvalidRequestHandler) {

	switch r.Method {
	case "POST":
		RawMsgXML, err := ioutil.ReadAll(r.Body)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		msg, err := util.ParseXMLToMap(bytes.NewReader(RawMsgXML))
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		ReturnCode, ok := msg["return_code"]
		if !ok || ReturnCode == ReturnCodeSuccess {
			haveAppId := msg["appid"]
			wantAppId := messageServer.AppId()
			if len(haveAppId) != len(wantAppId) {
				err = fmt.Errorf("the message's appid mismatch, have: %s, want: %s", haveAppId, wantAppId)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}
			if subtle.ConstantTimeCompare([]byte(haveAppId), []byte(wantAppId)) != 1 {
				err = fmt.Errorf("the message's appid mismatch, have: %s, want: %s", haveAppId, wantAppId)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			haveMchId := msg["mch_id"]
			wantMchId := messageServer.MchId()
			if len(haveMchId) != len(wantMchId) {
				err = fmt.Errorf("the message's mch_id mismatch, have: %s, want: %s", haveMchId, wantMchId)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}
			if subtle.ConstantTimeCompare([]byte(haveMchId), []byte(wantMchId)) != 1 {
				err = fmt.Errorf("the message's mch_id mismatch, have: %s, want: %s", haveMchId, wantMchId)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			// 认证签名
			signature1, ok := msg["sign"]
			if !ok {
				err = errors.New("no sign parameter")
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}
			signature2 := Sign(msg, messageServer.APIKey(), nil)
			if len(signature1) != len(signature2) {
				err = fmt.Errorf("check signature failed, \r\ninput: %q, \r\nlocal: %q", signature1, signature2)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}
			if subtle.ConstantTimeCompare([]byte(signature1), []byte(signature2)) != 1 {
				err = fmt.Errorf("check signature failed, \r\ninput: %q, \r\nlocal: %q", signature1, signature2)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}
		}

		req := &Request{
			HttpRequest: r,

			RawMsgXML: RawMsgXML,
			Msg:       msg,
		}
		messageServer.MessageHandler().ServeMessage(w, req)

	default:
		invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("Request.Method: "+r.Method))
	}
}
