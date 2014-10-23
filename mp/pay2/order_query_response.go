// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

// 因为某一方技术的原因，可能导致商户在预期时间内都收不到最终支付通知，此时商户
// 可以通过该API 来查询订单的详细支付状态。
//
// 这是订单查询成功时返回的数据结构
type OrderQueryResponse struct {
	RetCode       int    `json:"ret_code"`             // 查询结果状态码, 0表明成功, 其他表明错误
	RetMsg        string `json:"ret_msg"`              // 查询结果出错信息
	Charset       string `json:"input_charset"`        // 返回信息中的编码方式
	TradeMode     int    `json:"trade_mode,string"`    // 订单状态, 0为成功, 其他为失败
	TradeState    int    `json:"trade_state,string"`   // 交易模式, 1为即时到帐, 其他保留
	PartnerId     string `json:"partner"`              // 财付通商户号
	BankType      string `json:"bank_type"`            // 银行类型
	BankBillNo    string `json:"bank_billno"`          // 银行订单号
	TotalFee      int    `json:"total_fee,string"`     // 总金额, 单位为分
	FeeType       int    `json:"fee_type,string"`      // 币种, 1为人民币
	TransactionId string `json:"transaction_id"`       // 财付通订单号
	OutTradeNo    string `json:"out_trade_no"`         // 第三方订单号
	IsSplit       bool   `json:"is_split,string"`      // 表明是否分账, false为无分账, true为有分账
	IsRefund      bool   `json:"is_refund,string"`     // 表明是否退款, false为无退款, ture为退款
	Attach        string `json:"attach"`               // 商户数据包, 即生成订单package时商户填入的attach
	TimeEnd       string `json:"time_end"`             // 支付完成时间
	TransportFee  int    `json:"transport_fee,string"` // 物流费用, 单位为分
	ProductFee    int    `json:"product_fee,string"`   // 物品费用, 单位为分
	Discount      int    `json:"discount,string"`      // 折扣价格, 单位为分
	RMBTotalFee   string `json:"rmb_total_fee"`        // 换算成人民币之后的总金额, 单位为分, 一般看total_fee即可
}
