// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"testing"
)

func TestTenpayMD5Sign(t *testing.T) {
	m := make(map[string]string)
	m["bank_type"] = "WX"
	m["body"] = "支付测试"
	m["fee_type"] = "1"
	m["input_charset"] = "UTF-8"
	m["notify_url"] = "http://weixin.qq.com"
	m["out_trade_no"] = "7240b65810859cbf2a8d9f76a638c0a3"
	m["partner"] = "1900000109"
	m["spbill_create_ip"] = "196.168.1.1"
	m["total_fee"] = "1"

	have := TenpayMD5Sign(m, "8934e7d15453e97507ef794cf7b0519d")
	want := "7F77B507B755B3262884291517E380F8"

	if have != want {
		t.Errorf("failed, \r\nhave %#s, \r\nwant %#s", have, want)
		return
	}
}

func TestWXSHA1Sign1(t *testing.T) {
	m := make(map[string]string)
	m["appid"] = "wxd930ea5d5a258f4f"
	m["noncestr"] = "e7d161ac8d8a76529d39d9f5b4249ccb"
	m["package"] = "bank_type=WX&body=%E6%94%AF%E4%BB%98%E6%B5%8B%E8%AF%95&fee_type=1&input_charset=UTF-8&notify_url=http%3A%2F%2Fweixin.qq.com&out_trade_no=7240b65810859cbf2a8d9f76a638c0a3&partner=1900000109&spbill_create_ip=196.168.1.1&total_fee=1&sign=7F77B507B755B3262884291517E380F8"
	m["timestamp"] = "1399514976"
	m["traceid"] = "test_1399514976"

	appKey := "L8LrMqqeGRxST5reouB0K66CaYAWpqhAVsq7ggKkxHCOastWksvuX1uvmvQclxaHoYd3ElNBrNO2DHnnzgfVG9Qs473M3DTOZug5er46FhuGofumV8H2FVR9qkjSlC5K"

	have := WXSHA1Sign1(m, appKey, nil)
	want := "8893870b9004ead28691b60db97a8d2c80dbfdc6"

	if have != want {
		t.Errorf("failed, \r\nhave %#s, \r\nwant %#s", have, want)
		return
	}
}

func TestWXSHA1Sign2(t *testing.T) {
	m := make(map[string]string)
	m["appid"] = "wxd930ea5d5a258f4f"
	m["noncestr"] = "e7d161ac8d8a76529d39d9f5b4249ccb"
	m["package"] = "bank_type=WX&body=%E6%94%AF%E4%BB%98%E6%B5%8B%E8%AF%95&fee_type=1&input_charset=UTF-8&notify_url=http%3A%2F%2Fweixin.qq.com&out_trade_no=7240b65810859cbf2a8d9f76a638c0a3&partner=1900000109&spbill_create_ip=196.168.1.1&total_fee=1&sign=7F77B507B755B3262884291517E380F8"
	m["timestamp"] = "1399514976"
	m["traceid"] = "test_1399514976"

	appKey := "L8LrMqqeGRxST5reouB0K66CaYAWpqhAVsq7ggKkxHCOastWksvuX1uvmvQclxaHoYd3ElNBrNO2DHnnzgfVG9Qs473M3DTOZug5er46FhuGofumV8H2FVR9qkjSlC5K"

	have := WXSHA1Sign2(m, appKey, []string{"appid", "noncestr", "package", "timestamp", "traceid"})
	want := "8893870b9004ead28691b60db97a8d2c80dbfdc6"

	if have != want {
		t.Errorf("failed, \r\nhave %#s, \r\nwant %#s", have, want)
		return
	}
}
