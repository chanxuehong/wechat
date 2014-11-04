// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"testing"
)

func TestMakePayPackage(t *testing.T) {
	para := make(PayPackageParameters, 16)
	para.SetBankType(BANK_TYPE_WX)
	para.SetFeeType(FEE_TYPE_RMB)
	para.SetBody("支付测试")
	para.SetCharset(CHARSET_UTF8)
	para.SetPartnerId("1900000109")
	para.SetTotalFee(1)
	para.SetBillCreateIP("196.168.1.1")
	para.SetOutTradeNo("7240b65810859cbf2a8d9f76a638c0a3")
	para.SetNotifyURL("http://weixin.qq.com")

	have := string(MakePayPackage(para, "8934e7d15453e97507ef794cf7b0519d"))
	want := "bank_type=WX&body=%E6%94%AF%E4%BB%98%E6%B5%8B%E8%AF%95&fee_type=1&input_charset=UTF-8&notify_url=http%3A%2F%2Fweixin.qq.com&out_trade_no=7240b65810859cbf2a8d9f76a638c0a3&partner=1900000109&spbill_create_ip=196.168.1.1&total_fee=1&sign=7F77B507B755B3262884291517E380F8"

	if have != want {
		t.Errorf("failed, \r\nhave %#s, \r\nwant %#s", have, want)
		return
	}
}
