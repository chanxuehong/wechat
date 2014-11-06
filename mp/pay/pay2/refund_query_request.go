// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import "strconv"

// 退款明细查询 请求参数
type NormalRefundQueryRequest map[string]string

func (req NormalRefundQueryRequest) SetSignMethod(str string) {
	req["sign_type"] = str
}
func (req NormalRefundQueryRequest) SetCharset(str string) {
	req["input_charset"] = str
}
func (req NormalRefundQueryRequest) SetSignKeyIndex(n int) {
	req["sign_key_index"] = strconv.FormatInt(int64(n), 10)
}

func (req NormalRefundQueryRequest) SetPartnerId(str string) {
	req["partner"] = str
}
func (req NormalRefundQueryRequest) SetOutTradeNo(str string) {
	req["out_trade_no"] = str
}
func (req NormalRefundQueryRequest) SetTransactionId(str string) {
	req["transaction_id"] = str
}
func (req NormalRefundQueryRequest) SetOutRefundNo(str string) {
	req["out_refund_no"] = str
}
func (req NormalRefundQueryRequest) SetRefundId(str string) {
	req["refund_id"] = str
}
func (req NormalRefundQueryRequest) SetUseSPBillNoFlag(n int) {
	req["use_spbill_no_flag"] = strconv.FormatInt(int64(n), 10)
}

// 设置签名字段.
//  Key: 商户支付密钥Key
//
//  NOTE: 要求在 NormalRefundQueryRequest 其他字段设置完毕后才能调用这个函数, 否则签名就不正确.
func (req NormalRefundQueryRequest) SetSignature(Key string) (err error) {
	return RefundRequest(req).SetSignature(Key)
}
