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
	para.AppId = "wxf8b4f85f3a794e77"
	para.NonceStr = "adssdasssd13d"
	para.Package = "bank_type=WX&body=XXX&fee_type=1&input_charset=GBK&notify_url=http%3a%2f%2fwww.qq.com&out_trade_no=16642817866003386000&partner=1900000109&spbill_create_ip=127.0.0.1&total_fee=1&sign=BEEF37AD19575D92E191C1E4B1474CA9"
	para.SignMethod = SIGN_METHOD_SHA1
	para.TimeStamp = 189026618

	appKey := "2Wozy2aksie1puXUBpWD8oZxiD1DfQuEaiC7KcRATv1Ino3mdopKaPGQQ7TtkNySuAmCaDCrw4xhPY5qKTBl7Fzm0RgR3c0WaVYIXZARsxzHV2x7iwPPzOz94dnwPWSn"

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

	str := string(bs)
	want := `{"appId":"wxf8b4f85f3a794e77","nonceStr":"adssdasssd13d","timeStamp":"189026618","package":"bank_type=WX\u0026body=XXX\u0026fee_type=1\u0026input_charset=GBK\u0026notify_url=http%3a%2f%2fwww.qq.com\u0026out_trade_no=16642817866003386000\u0026partner=1900000109\u0026spbill_create_ip=127.0.0.1\u0026total_fee=1\u0026sign=BEEF37AD19575D92E191C1E4B1474CA9","paySign":"7717231c335a05165b1874658306fa431fe9a0de","signType":"SHA1"}`
	if str != want {
		t.Errorf("failed, have %#s, want %#s", str, want)
		return
	}
}
