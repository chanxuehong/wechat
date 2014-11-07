// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"fmt"
	"strconv"

	"github.com/chanxuehong/wechat/mp/pay"
)

// 退款请求 请求参数
type RefundRequest map[string]string

func (req RefundRequest) SetSignMethod(str string) {
	req["sign_type"] = str
}
func (req RefundRequest) SetCharset(str string) {
	req["input_charset"] = str
}
func (req RefundRequest) SetSignKeyIndex(n int) {
	req["sign_key_index"] = strconv.FormatInt(int64(n), 10)
}
func (req RefundRequest) SetServiceVersion(str string) {
	req["service_version"] = str
}
func (req RefundRequest) SetPartnerId(str string) {
	req["partner"] = str
}
func (req RefundRequest) SetOutTradeNo(str string) {
	req["out_trade_no"] = str
}
func (req RefundRequest) SetTransactionId(str string) {
	req["transaction_id"] = str
}
func (req RefundRequest) SetOutRefundNo(str string) {
	req["out_refund_no"] = str
}
func (req RefundRequest) SetTotalFee(n int) {
	req["total_fee"] = strconv.FormatInt(int64(n), 10)
}
func (req RefundRequest) SetRefundFee(n int) {
	req["refund_fee"] = strconv.FormatInt(int64(n), 10)
}
func (req RefundRequest) SetOperUserId(id int) {
	req["op_user_id"] = strconv.FormatInt(int64(id), 10)
}
func (req RefundRequest) SetOperUserPwd(str string) {
	req["op_user_passwd"] = str
}
func (req RefundRequest) SetRecvUserId(str string) {
	req["recv_user_id"] = str
}
func (req RefundRequest) SetRecvUserName(str string) {
	req["reccv_user_name"] = str
}
func (req RefundRequest) SetUseSPBillNoFlag(n int) {
	req["use_spbill_no_flag"] = strconv.FormatInt(int64(n), 10)
}
func (req RefundRequest) SetRefundType(n int) {
	req["refund_type"] = strconv.FormatInt(int64(n), 10)
}

// 设置签名字段.
//  Key: 商户支付密钥Key
//
//  NOTE: 要求在 RefundRequest 其他字段设置完毕后才能调用这个函数, 否则签名就不正确.
func (req RefundRequest) SetSignature(Key string) (err error) {
	SignMethod := req["sign_type"]
	if SignMethod == "" {
		SignMethod = "MD5"
	}

	switch SignMethod {
	case "MD5", "md5":
		req["sign"] = pay.TenpayMD5Sign(req, Key)
		return

	default:
		return fmt.Errorf(`unknown sign method: %q`, SignMethod)
	}
}
