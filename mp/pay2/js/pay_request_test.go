// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package js

import (
	"encoding/json"
	"testing"
)

func TestPayRequestParametersSetSignatureAndMarshal(t *testing.T) {
	var para PayRequestParameters
	para.AppId = "wxd930ea5d5a258f4f"
	para.NonceStr = "e7d161ac8d8a76529d39d9f5b4249ccb"
	para.Package = "bank_type=WX&body=%E6%94%AF%E4%BB%98%E6%B5%8B%E8%AF%95&fee_type=1&input_charset=UTF-8&notify_url=http%3A%2F%2Fweixin.qq.com&out_trade_no=7240b65810859cbf2a8d9f76a638c0a3&partner=1900000109&spbill_create_ip=196.168.1.1&total_fee=1&sign=7F77B507B755B3262884291517E380F8"
	para.SignMethod = "SHA1"
	para.TimeStamp = 1399514976

	appKey := "L8LrMqqeGRxST5reouB0K66CaYAWpqhAVsq7ggKkxHCOastWksvuX1uvmvQclxaHoYd3ElNBrNO2DHnnzgfVG9Qs473M3DTOZug5er46FhuGofumV8H2FVR9qkjSlC5K"

	err := para.SetSignature(appKey)
	if err != nil {
		t.Error(err)
		return
	}

	bs, err := json.Marshal(&para)
	if err != nil {
		t.Error(err)
		return
	}

	have := string(bs)
	want := `{"appId":"wxd930ea5d5a258f4f","nonceStr":"e7d161ac8d8a76529d39d9f5b4249ccb","timeStamp":"1399514976","package":"bank_type=WX\u0026body=%E6%94%AF%E4%BB%98%E6%B5%8B%E8%AF%95\u0026fee_type=1\u0026input_charset=UTF-8\u0026notify_url=http%3A%2F%2Fweixin.qq.com\u0026out_trade_no=7240b65810859cbf2a8d9f76a638c0a3\u0026partner=1900000109\u0026spbill_create_ip=196.168.1.1\u0026total_fee=1\u0026sign=7F77B507B755B3262884291517E380F8","paySign":"8d794be018f6a2c99e617f74c8d4aca5c6ce25a0","signType":"SHA1"}`
	if have != want {
		t.Errorf("failed, \r\nhave %#s, \r\nwant %#s", have, want)
		return
	}
}
