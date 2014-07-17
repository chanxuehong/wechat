// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package js

import (
	"testing"
)

func TestOrderInfoPackage(t *testing.T) {
	var oi OrderInfo
	oi.BankType = ORDER_INFO_BANK_TYPE_WX
	oi.FeeType = ORDER_INFO_FEE_TYPE_RMB
	oi.Body = "XXX"
	oi.Charset = ORDER_INFO_CHARSET_GBK
	oi.PartnerId = "1900000109"
	oi.TotalFee = 1
	oi.CreateIP = "127.0.0.1"
	oi.OutTradeNo = "16642817866003386000"
	oi.NotifyURL = "http://www.qq.com"

	bs := oi.Package("8934e7d15453e97507ef794cf7b0519d")

	str := string(bs)
	want := `bank_type=WX&body=XXX&fee_type=1&input_charset=GBK&notify_url=http%3A%2F%2Fwww.qq.com&out_trade_no=16642817866003386000&partner=1900000109&spbill_create_ip=127.0.0.1&total_fee=1&sign=BEEF37AD19575D92E191C1E4B1474CA9`
	if str != want {
		t.Errorf("OrderInfo.Package failed, have %#s\n, want %#s\n", str, want)
		return
	}
}
