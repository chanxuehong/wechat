// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"time"
)

// 支付结果通知消息
type Notification struct {
	Id string // 必须, 该 Notification id, 对于某些特定商户, 只返回通知 id, 要求商户据此查询交易结果

	// 下面这三个字段和之前传过去的一样, 可以比对下, 确保安全
	AppId     string // 必须, 公众号 id
	TimeStamp int64  // 必须, 时间戳, unixtime
	NonceStr  string // 必须, 随机字符串

	OpenId      string // 必须, 支付该笔订单的用户 id, 商户可通过公众号其他接口为付款用户服务.
	IsSubscribe int    // 必须, 用户是否关注了公众号. 1为关注, 0为未关注.
	BuyerAlias  string // 可选, 买家别名 ,对应买家账号的一个加密串

	TradeMode     int       // 必须, 交易模式, 1-即时到账, 其他保留
	TradeState    int       // 必须, 交易状态(支付结果), 0-成功, 其他保留
	PayInfo       string    // 可选, 支付结果信息, 支付成功时为空!
	BankBillNo    string    // 可选, 银行订单号
	TransactionId string    // 必须, 交易号, 28位长的数值, 其中前10位为商户号, 之后8位为订单产生的日期, 如20090415, 最后10位是流水号.
	TimeEnd       time.Time // 必须, 支付完成时间

	// 下面这 4 个字段和支付账单 Bill 里的同名字段内容相同
	BankType   string // 必须, 银行类型, 微信中固定为 WX
	PartnerId  string // 必须, 财付通商户 partnerId
	OutTradeNo string // 必须, 商户系统的订单号
	Attach     string // 可选, 商户数据包

	TotalFee     int // 必须, 支付金额, 单位为分; 如果 discount 有值, 则有 TotalFee + Discount == 支付请求的 Bill.TotalFee
	Discount     int // 可选, 折扣价格, 单位为分; 如果有值, 则有 TotalFee + Discount == 支付请求的 Bill.TotalFee
	TransportFee int // 可选, 物流费用, 单位为分, 默认0; 如果有值, 必须保证 TransportFee + ProductFee == TotalFee
	ProductFee   int // 可选, 物品费用, 单位为分; 如果有值, 必须保证 TransportFee + ProductFee == TotalFee
	FeeType      int // 必须, 币种, 目前只支持人民币, 默认值是 1-人民币
}
