// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"bytes"
	"crypto/md5"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"time"
)

// 多密钥支持的情况下, 根据密钥 index 获取 partnerKey, 找不到合法的密钥返回空值 ""
type GetPartnerKey func(keyIndex int) string

// 用户在成功完成支付后，微信后台通知（POST）商户服务器（notify_url）支付结果。
// 商户可以使用 notify_url 的通知结果进行个性化页面的展示。
//
// 这是支付成功后通知消息 url query string 部分, 1.0 版本
type OrderNotifyURLDataVer1 struct {
	// 协议参数 ==================================================================

	ServiceVersion string // 必须, 版本号
	Charset        string // 必须, 字符编码, 取值: GBK, UTF-8
	Signature      string // 必须, 签名
	SignMethod     string // 必须, 签名类型, 取值: MD5, RSA
	SignKeyIndex   int    // 必须, 多密钥支持的密钥序号

	// 业务参数 ==================================================================

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

// 根据 values url.Values(来自对 notify url query string 的解析) 来初始化 data *OrderNotifyURLDataVer1.
// 如果 values url.Values 里的参数不合法(包括签名不正确) 则返回错误信息, 否则返回 nil.
func (data *OrderNotifyURLDataVer1) CheckAndInit(values url.Values, getPartnerKey GetPartnerKey) (err error) {
	if values == nil {
		return errors.New("values == nil")
	}
	if getPartnerKey == nil {
		return errors.New("getPartnerKey ==  nil")
	}

	// 先检查签名是否正确 =========================================================

	var signMethod string
	if signMethods := values["sign_type"]; len(signMethods) > 0 && len(signMethods[0]) > 0 {
		signMethod = signMethods[0]
	} else {
		signMethod = "MD5"
	}

	var charset string
	if charsets := values["input_charset"]; len(charsets) > 0 && len(charsets[0]) > 0 {
		charset = charsets[0]
	} else {
		charset = "GBK"
	}

	var signature string
	if signatures := values["sign"]; len(signatures) > 0 && len(signatures[0]) > 0 {
		signature = signatures[0]
		values.Del("sign")
	} else {
		return errors.New("sign is empty")
	}

	var signKeyIndex int
	if signKeyIndexes := values["sign_key_index"]; len(signKeyIndexes) > 0 && len(signKeyIndexes[0]) > 0 {
		index, err := strconv.ParseInt(signKeyIndexes[0], 10, 64)
		if err != nil {
			return fmt.Errorf("解析密钥 %s 出错: %s", signKeyIndexes[0], err.Error())
		}
		signKeyIndex = int(index)
	} else {
		signKeyIndex = 1
	}

	partnerKey := getPartnerKey(signKeyIndex)
	if partnerKey == "" {
		return fmt.Errorf("获取编号为 %d 的加密密钥失败", signKeyIndex)
	}

	switch {
	case signMethod == "MD5" && charset == "UTF-8":
		keys := make([]string, 0, len(values))
		for k := range values {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		var buf bytes.Buffer

		// 原始串拼接
		for _, k := range keys {
			vs := values[k]
			for _, v := range vs {
				if buf.Len() > 0 {
					buf.WriteByte('&')
				}
				buf.WriteString(k)
				buf.WriteByte('=')
				buf.WriteString(v)
			}
		}
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString("key=")
		buf.WriteString(partnerKey)

		string1 := buf.Bytes()
		hashSumArray := md5.Sum(string1)
		hashSumHexBytes := make([]byte, hex.EncodedLen(len(hashSumArray)))
		hex.Encode(hashSumHexBytes, hashSumArray[:])
		copy(hashSumHexBytes, bytes.ToUpper(hashSumHexBytes))

		if subtle.ConstantTimeCompare([]byte(signature), hashSumHexBytes) != 1 {
			return errors.New("签名检验不通过")
		}

		// 初始化 ===================================================================

		// 在上面已经判断了
		data.ServiceVersion = "1.0"
		data.SignMethod = signMethod
		data.Charset = charset
		data.Signature = signature
		data.SignKeyIndex = signKeyIndex

		if notifyIds := values["notify_id"]; len(notifyIds) > 0 && len(notifyIds[0]) > 0 {
			data.NotifyId = notifyIds[0]
		} else {
			return errors.New("notify_id is empty")
		}

		if tradeModes := values["trade_mode"]; len(tradeModes) > 0 && len(tradeModes[0]) > 0 {
			mode, err := strconv.ParseInt(tradeModes[0], 10, 64)
			if err != nil {
				return err
			}
			data.TradeMode = int(mode)
		} else {
			return errors.New("trade_mode is empty")
		}

		if tradeStates := values["trade_state"]; len(tradeStates) > 0 && len(tradeStates[0]) > 0 {
			state, err := strconv.ParseInt(tradeStates[0], 10, 64)
			if err != nil {
				return err
			}
			data.TradeState = int(state)
		} else {
			return errors.New("trade_state is empty")
		}

		if payInfo := values["pay_info"]; len(payInfo) > 0 {
			data.PayInfo = payInfo[0]
		}

		if bankBillNo := values["bank_billno"]; len(bankBillNo) > 0 {
			data.BankBillNo = bankBillNo[0]
		}

		if transactionIds := values["transaction_id"]; len(transactionIds) > 0 && len(transactionIds[0]) > 0 {
			data.TransactionId = transactionIds[0]
		} else {
			return errors.New("transaction_id is empty")
		}

		if timeEnds := values["time_end"]; len(timeEnds) > 0 && len(timeEnds[0]) > 0 {
			t, err := ParseTime(timeEnds[0])
			if err != nil {
				return err
			}
			data.TimeEnd = t
		} else {
			return errors.New("time_end is empty")
		}

		if bankTypes := values["bank_type"]; len(bankTypes) > 0 && len(bankTypes[0]) > 0 {
			data.BankType = bankTypes[0]
		} else {
			return errors.New("bank_type is empty")
		}

		if partnerIds := values["partner"]; len(partnerIds) > 0 && len(partnerIds[0]) > 0 {
			data.PartnerId = partnerIds[0]
		} else {
			return errors.New("partner is empty")
		}

		if outTradeNo := values["out_trade_no"]; len(outTradeNo) > 0 && len(outTradeNo[0]) > 0 {
			data.OutTradeNo = outTradeNo[0]
		} else {
			return errors.New("out_trade_no is empty")
		}

		if attaches := values["attach"]; len(attaches) > 0 {
			data.Attach = attaches[0]
		}

		if totalFees := values["total_fee"]; len(totalFees) > 0 && len(totalFees[0]) > 0 {
			fee, err := strconv.ParseInt(totalFees[0], 10, 64)
			if err != nil {
				return err
			}
			data.TotalFee = int(fee)
		} else {
			return errors.New("total_fee is empty")
		}

		if discounts := values["discount"]; len(discounts) > 0 && len(discounts[0]) > 0 {
			discount, err := strconv.ParseInt(discounts[0], 10, 64)
			if err != nil {
				return err
			}
			data.Discount = int(discount)
		}

		if transportFees := values["transport_fee"]; len(transportFees) > 0 && len(transportFees[0]) > 0 {
			fee, err := strconv.ParseInt(transportFees[0], 10, 64)
			if err != nil {
				return err
			}
			data.TransportFee = int(fee)
		}

		if productFees := values["product_fee"]; len(productFees) > 0 && len(productFees[0]) > 0 {
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

		if feeTypes := values["fee_type"]; len(feeTypes) > 0 && len(feeTypes[0]) > 0 {
			feeType, err := strconv.ParseInt(feeTypes[0], 10, 64)
			if err != nil {
				return err
			}
			data.FeeType = int(feeType)
		} else {
			return errors.New("fee_type is empty")
		}

		if buyerAliases := values["buyer_alias"]; len(buyerAliases) > 0 {
			data.BuyerAlias = buyerAliases[0]
		}

		return

	default:
		err = fmt.Errorf("没有实现对 签名方法(%s) 或者 字符编码(%s) 的支持", signMethod, charset)
		return
	}
}
