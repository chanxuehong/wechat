// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"testing"
)

func TestPayPackagePackage(t *testing.T) {
	var payPackage PayPackage
	payPackage.BankType = BANK_TYPE_WX
	payPackage.FeeType = FEE_TYPE_RMB
	payPackage.Body = "支付测试"
	payPackage.Charset = CHARSET_UTF8
	payPackage.PartnerId = "1900000109"
	payPackage.TotalFee = 1
	payPackage.BillCreateIP = "196.168.1.1"
	payPackage.OutTradeNo = "7240b65810859cbf2a8d9f76a638c0a3"
	payPackage.NotifyURL = "http://weixin.qq.com"

	Package := payPackage.Package("8934e7d15453e97507ef794cf7b0519d")

	have := string(Package)
	want := `bank_type=WX&body=%E6%94%AF%E4%BB%98%E6%B5%8B%E8%AF%95&fee_type=1&input_charset=UTF-8&notify_url=http%3A%2F%2Fweixin.qq.com&out_trade_no=7240b65810859cbf2a8d9f76a638c0a3&partner=1900000109&spbill_create_ip=196.168.1.1&total_fee=1&sign=7F77B507B755B3262884291517E380F8`
	if have != want {
		t.Errorf("failed, \r\nhave %#s, \r\nwant %#s", have, want)
		return
	}
}
