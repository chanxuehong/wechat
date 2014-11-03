// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/chanxuehong/wechat/mp/pay"
	"time"
)

// 因为某一方技术的原因，可能导致商户在预期时间内都收不到最终支付通知，此时商户
// 可以通过该API 来查询订单的详细支付状态。
//
// 这是订单查询成功时返回的数据结构（注意 json 格式化的时候要用 *OrderQueryResponse）
type OrderQueryResponse struct {
	RetCode       json.Number     `json:"ret_code"`       // 查询结果状态码, 0表明成功, 其他表明错误
	RetMsg        string          `json:"ret_msg"`        // 查询结果出错信息
	Charset       string          `json:"input_charset"`  // 返回信息中的编码方式
	TradeMode     json.Number     `json:"trade_mode"`     // 订单状态, 0为成功, 其他为失败
	TradeState    json.Number     `json:"trade_state"`    // 交易模式, 1为即时到帐, 其他保留
	PartnerId     string          `json:"partner"`        // 财付通商户号
	BankType      string          `json:"bank_type"`      // 银行类型
	BankBillNo    string          `json:"bank_billno"`    // 银行订单号
	TotalFee      json.Number     `json:"total_fee"`      // 总金额, 单位为分
	FeeType       json.Number     `json:"fee_type"`       // 币种, 1为人民币
	TransactionId string          `json:"transaction_id"` // 财付通订单号
	OutTradeNo    string          `json:"out_trade_no"`   // 第三方订单号
	IsSplitBytes  json.RawMessage `json:"is_split"`       // boolean, 表明是否分账, false为无分账, true为有分账
	IsRefundBytes json.RawMessage `json:"is_refund"`      // boolean, 表明是否退款, false为无退款, ture为退款
	Attach        string          `json:"attach"`         // 商户数据包, 即生成订单package时商户填入的attach
	TimeEnd       string          `json:"time_end"`       // 支付完成时间
	TransportFee  json.Number     `json:"transport_fee"`  // 物流费用, 单位为分
	ProductFee    json.Number     `json:"product_fee"`    // 物品费用, 单位为分
	Discount      json.Number     `json:"discount"`       // 折扣价格, 单位为分
	RMBTotalFee   json.Number     `json:"rmb_total_fee"`  // 换算成人民币之后的总金额, 单位为分, 一般看total_fee即可
}

var (
	json_false_bytes        = []byte(`false`)
	json_false_bytes_quoted = []byte(`"false"`)
	json_true_bytes         = []byte(`true`)
	json_true_bytes_quoted  = []byte(`"true"`)
)

func (this *OrderQueryResponse) IsSplit() (b bool, err error) {
	src := this.IsSplitBytes

	if bytes.Equal(src, json_false_bytes_quoted) ||
		bytes.Equal(src, json_false_bytes) {
		return
	}
	if bytes.Equal(src, json_true_bytes_quoted) ||
		bytes.Equal(src, json_true_bytes) {
		b = true
		return
	}
	err = fmt.Errorf("invalid is_split: %q", src)
	return
}
func (this *OrderQueryResponse) IsRefund() (b bool, err error) {
	src := this.IsRefundBytes

	if bytes.Equal(src, json_false_bytes_quoted) ||
		bytes.Equal(src, json_false_bytes) {
		return
	}
	if bytes.Equal(src, json_true_bytes_quoted) ||
		bytes.Equal(src, json_true_bytes) {
		b = true
		return
	}
	err = fmt.Errorf("invalid is_refund: %q", src)
	return
}
func (this *OrderQueryResponse) GetTimeEnd() (t time.Time, err error) {
	return pay.ParseTime(this.TimeEnd)
}
