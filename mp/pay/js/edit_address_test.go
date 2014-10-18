// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package js

import (
	"encoding/json"
	"testing"
)

func TestEditAddressParametersSetSignatureAndMarshal(t *testing.T) {
	var para EditAddressParameters
	para.AppId = "wx17ef1eaef46752cb"
	para.NonceStr = "123456"
	para.TimeStamp = 1384841012
	para.Scope = "jsapi_address"
	para.SignMethod = SIGN_METHOD_SHA1

	accessToken := "OezXcEiiBSKSxW0eoylIeBFk1b8VbNtfWALJ5g6aMgZHaqZwK4euEskSn78Qd5pLsfQtuMdgmhajVM5QDm24W8X3tJ18kz5mhmkUcI3RoLm7qGgh1cEnCHejWQo8s5L3VvsFAdawhFxUuLmgh5FRA"

	err := para.SetSignature("http://open.weixin.qq.com/", accessToken)
	if err != nil {
		t.Error(err)
		return
	}

	bs, err := json.Marshal(&para)
	if err != nil {
		t.Error(err)
		return
	}

	str := string(bs)
	want := `{"appId":"wx17ef1eaef46752cb","nonceStr":"123456","timeStamp":"1384841012","scope":"jsapi_address","addrSign":"b43e348ff8e0a8d04b7b4274ce1bd6c4db00b1a4","signType":"sha1"}`
	if str != want {
		t.Errorf("failed, have %#s, want %#s", str, want)
		return
	}
}
