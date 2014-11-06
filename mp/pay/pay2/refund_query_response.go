// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"strconv"
)

// 退款明细查询 回复参数
type NormalRefundQueryResponse map[string]string

func (resp NormalRefundQueryResponse) SignMethod() string {
	str := resp["sign_type"]
	if str != "" {
		return str
	}
	return "MD5"
}
func (resp NormalRefundQueryResponse) Charset() string {
	str := resp["input_charset"]
	if str != "" {
		return str
	}
	return "GBK"
}
func (resp NormalRefundQueryResponse) Signature() string {
	return resp["sign"]
}
func (resp NormalRefundQueryResponse) RetCode() string {
	return resp["retcode"]
}
func (resp NormalRefundQueryResponse) RetMsg() string {
	return resp["retmsg"]
}
func (resp NormalRefundQueryResponse) PartnerId() string {
	return resp["partner"]
}
func (resp NormalRefundQueryResponse) OutTradeNo() string {
	return resp["out_trade_no"]
}
func (resp NormalRefundQueryResponse) TransactionId() string {
	return resp["transaction_id"]
}
func (resp NormalRefundQueryResponse) RefundCount() string {
	return resp["refund_count"]
}
func (resp NormalRefundQueryResponse) OutRefundNo(n int) string {
	return resp["out_refund_no_"+strconv.FormatInt(int64(n), 10)]
}
func (resp NormalRefundQueryResponse) RefundId(n int) string {
	return resp["refund_id_"+strconv.FormatInt(int64(n), 10)]
}
func (resp NormalRefundQueryResponse) RefundChannel(n int) string {
	return resp["refund_channel_"+strconv.FormatInt(int64(n), 10)]
}
func (resp NormalRefundQueryResponse) RefundFee(n int) string {
	return resp["refund_fee_"+strconv.FormatInt(int64(n), 10)]
}
func (resp NormalRefundQueryResponse) RefundState(n int) string {
	return resp["refund_state_"+strconv.FormatInt(int64(n), 10)]
}
func (resp NormalRefundQueryResponse) RecvUserId(n int) string {
	return resp["recv_user_id_"+strconv.FormatInt(int64(n), 10)]
}
func (resp NormalRefundQueryResponse) RecvUserName(n int) string {
	return resp["reccv_user_name_"+strconv.FormatInt(int64(n), 10)]
}

// 检查 NormalRefundQueryResponse 的签名是否合法, 合法返回 nil, 否则返回错误信息.
//  Key: 商户支付密钥Key
func (resp NormalRefundQueryResponse) CheckSignature(Key string) (err error) {
	return RefundResponse(resp).CheckSignature(Key)
}
