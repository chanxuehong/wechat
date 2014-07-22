// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"errors"
	"net/url"
	"strconv"
	"time"
)

// 支付成功后通知消息 url query string 部分
type NotifyURLData struct {
	// 协议参数
	ServiceVersion string // 必须, 版本号
	Charset        string // 必须, 字符编码, 取值: GBK, UTF-8
	Signature      string // 必须, 签名
	SignMethod     string // 必须, 签名类型, 取值: MD5, RSA
	SignKeyIndex   int    // 必须, 多密钥支持的密钥序号

	// 业务参数

	NotifyId string // 必须, 支付结果通知 id, 对于某些特定商户, 只返回通知 id, 要求商户据此查询交易结果

	TradeMode     int       // 必须, 交易模式, 1-即时到账, 其他保留
	TradeState    int       // 必须, 交易状态(支付结果), 0-成功, 其他保留
	PayInfo       string    // 可选, 支付结果信息, 支付成功时为 "".
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

	BuyerAlias string // 可选, 买家别名, 对应买家账号的一个加密串
}

func (data *NotifyURLData) Init(values url.Values) (err error) {
	if values == nil {
		return errors.New("values == nil")
	}

	serviceVersions := values["service_version"]
	if len(serviceVersions) > 0 && len(serviceVersions[0]) > 0 {
		data.ServiceVersion = serviceVersions[0]
	} else {
		data.ServiceVersion = "1.0"
	}

	charsets := values["input_charset"]
	if len(charsets) > 0 && len(charsets[0]) > 0 {
		data.Charset = charsets[0]
	} else {
		data.Charset = NOTIFY_URL_DATA_CHARSET_GBK
	}

	signatures := values["sign"]
	if len(signatures) > 0 && len(signatures[0]) > 0 {
		data.Signature = signatures[0]
	} else {
		return errors.New("sign is empty")
	}

	signMethods := values["sign_type"]
	if len(signMethods) > 0 && len(signMethods[0]) > 0 {
		data.SignMethod = signMethods[0]
	} else {
		data.SignMethod = NOTIFY_URL_DATA_SIGN_METHOD_MD5
	}

	signKeyIndexes := values["sign_key_index"]
	if len(signKeyIndexes) > 0 && len(signKeyIndexes[0]) > 0 {
		index, err := strconv.ParseInt(signMethods[0], 10, 64)
		if err != nil {
			return err
		}
		data.SignKeyIndex = int(index)
	} else {
		data.SignKeyIndex = 1
	}

	notifyIds := values["notify_id"]
	if len(notifyIds) > 0 && len(notifyIds[0]) > 0 {
		data.NotifyId = notifyIds[0]
	} else {
		return errors.New("notify_id is empty")
	}

	tradeModes := values["trade_mode"]
	if len(tradeModes) > 0 && len(tradeModes[0]) > 0 {
		mode, err := strconv.ParseInt(tradeModes[0], 10, 64)
		if err != nil {
			return err
		}
		data.TradeMode = int(mode)
	} else {
		return errors.New("trade_mode is empty")
	}

	tradeStates := values["trade_state"]
	if len(tradeStates) > 0 && len(tradeStates[0]) > 0 {
		state, err := strconv.ParseInt(tradeStates[0], 10, 64)
		if err != nil {
			return err
		}
		data.TradeState = int(state)
	} else {
		return errors.New("trade_state is empty")
	}

	payInfo := values["pay_info"]
	if len(payInfo) > 0 {
		data.PayInfo = payInfo[0]
	}

	bankBillNo := values["bank_billno"]
	if len(bankBillNo) > 0 {
		data.BankBillNo = bankBillNo[0]
	}

	transactionIds := values["transaction_id"]
	if len(transactionIds) > 0 && len(transactionIds[0]) > 0 {
		data.TransactionId = transactionIds[0]
	} else {
		return errors.New("transaction_id is empty")
	}

	timeEnds := values["time_end"]
	if len(timeEnds) > 0 && len(timeEnds[0]) > 0 {
		t, err := ParseTime(timeEnds[0])
		if err != nil {
			return err
		}
		data.TimeEnd = t
	} else {
		return errors.New("time_end is empty")
	}

	bankTypes := values["bank_type"]
	if len(bankTypes) > 0 && len(bankTypes[0]) > 0 {
		data.BankType = bankTypes[0]
	} else {
		return errors.New("bank_type is empty")
	}

	partnerIds := values["partner"]
	if len(partnerIds) > 0 && len(partnerIds[0]) > 0 {
		data.PartnerId = partnerIds[0]
	} else {
		return errors.New("partner is empty")
	}

	outTradeNo := values["out_trade_no"]
	if len(outTradeNo) > 0 && len(outTradeNo[0]) > 0 {
		data.OutTradeNo = outTradeNo[0]
	} else {
		return errors.New("out_trade_no is empty")
	}

	attaches := values["attach"]
	if len(attaches) > 0 {
		data.Attach = attaches[0]
	}

	totalFees := values["total_fee"]
	if len(totalFees) > 0 && len(totalFees[0]) > 0 {
		fee, err := strconv.ParseInt(totalFees[0], 10, 64)
		if err != nil {
			return err
		}
		data.TotalFee = int(fee)
	} else {
		return errors.New("total_fee is empty")
	}

	discounts := values["discount"]
	if len(discounts) > 0 && len(discounts[0]) > 0 {
		discount, err := strconv.ParseInt(discounts[0], 10, 64)
		if err != nil {
			return err
		}
		data.Discount = int(discount)
	}

	transportFees := values["transport_fee"]
	if len(transportFees) > 0 && len(transportFees[0]) > 0 {
		fee, err := strconv.ParseInt(transportFees[0], 10, 64)
		if err != nil {
			return err
		}
		data.TransportFee = int(fee)
	}

	productFees := values["product_fee"]
	if len(productFees) > 0 && len(productFees[0]) > 0 {
		fee, err := strconv.ParseInt(productFees[0], 10, 64)
		if err != nil {
			return err
		}
		data.ProductFee = int(fee)
	}

	if data.TransportFee != 0 || data.ProductFee != 0 {
		if data.TransportFee+data.ProductFee != data.TotalFee {
			return errors.New(`transport_fee+product_fee != total_fee`)
		}
	}

	feeTypes := values["fee_type"]
	if len(feeTypes) > 0 && len(feeTypes[0]) > 0 {
		feeType, err := strconv.ParseInt(feeTypes[0], 10, 64)
		if err != nil {
			return err
		}
		data.FeeType = int(feeType)
	} else {
		return errors.New("fee_type is empty")
	}

	buyerAliases := values["buyer_alias"]
	if len(buyerAliases) > 0 {
		data.BuyerAlias = buyerAliases[0]
	}

	return
}
