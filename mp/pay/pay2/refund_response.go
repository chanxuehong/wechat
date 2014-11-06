// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"crypto/subtle"
	"fmt"

	"github.com/chanxuehong/wechat/mp/pay"
)

// 退款请求的回复
type RefundResponse map[string]string

func (resp RefundResponse) SignMethod() string {
	str := resp["sign_type"]
	if str != "" {
		return str
	}
	return "MD5"
}
func (resp RefundResponse) Charset() string {
	str := resp["input_charset"]
	if str != "" {
		return str
	}
	return "GBK"
}
func (resp RefundResponse) Signature() string {
	return resp["sign"]
}
func (resp RefundResponse) RetCode() string {
	return resp["retcode"]
}
func (resp RefundResponse) RetMsg() string {
	return resp["retmsg"]
}
func (resp RefundResponse) PartnerId() string {
	return resp["partner"]
}
func (resp RefundResponse) TransactionId() string {
	return resp["transaction_id"]
}
func (resp RefundResponse) OutTradeNo() string {
	return resp["out_trade_no"]
}
func (resp RefundResponse) OutRefundNo() string {
	return resp["out_refund_no"]
}
func (resp RefundResponse) RefundId() string {
	return resp["refund_id"]
}
func (resp RefundResponse) RefundChannel() string {
	return resp["refund_channel"]
}
func (resp RefundResponse) RefundFee() string {
	return resp["refund_fee"]
}
func (resp RefundResponse) RefundStatus() string {
	return resp["refund_status"]
}
func (resp RefundResponse) RecvUserId() string {
	return resp["recv_user_id"]
}
func (resp RefundResponse) RecvUserName() string {
	return resp["reccv_user_name"]
}

// 检查 RefundResponse 的签名是否合法, 合法返回 nil, 否则返回错误信息.
//  Key: 商户支付密钥Key
func (resp RefundResponse) CheckSignature(Key string) (err error) {
	Signature1 := resp.Signature()
	SignMethod := resp.SignMethod()

	switch SignMethod {
	case "MD5", "md5":
		if len(Signature1) != 32 {
			err = fmt.Errorf(`不正确的签名: %q, 长度不对, have: %d, want: %d`,
				Signature1, len(Signature1), 32)
			return
		}

		Signature2 := pay.TenpayMD5Sign(resp, Key)

		if subtle.ConstantTimeCompare([]byte(Signature2), []byte(Signature1)) != 1 {
			err = fmt.Errorf("签名不匹配, \r\nlocal: %q, \r\ninput: %q", Signature2, Signature1)
			return
		}
		return

	default:
		err = fmt.Errorf(`unknown sign method: %q`, SignMethod)
		return
	}
}
