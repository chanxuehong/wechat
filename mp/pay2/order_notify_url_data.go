// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"bytes"
	"crypto/md5"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/chanxuehong/wechat/util"
	"net/url"
	"sort"
	"strconv"
	"time"
)

// 用户在成功完成支付后，微信后台通知商户服务器（notify_url）支付结果。
// 商户可以使用notify_url 的通知结果进行个性化页面的展示。
//
// 对后台通知交互时，如果微信收到商户的应答不是success 或超时，微信认为通知失败，
// 微信会通过一定的策略（如30 分钟共8 次）定期重新发起通知，尽可能提高通知的成功率，
// 但微信不保证通知最终能成功。
// 由于存在重新发送后台通知的情况，因此同样的通知可能会多次发送给商户系统。商户
// 系统必须能够正确处理重复的通知。
//
// 微信后台通过 notify_url 通知商户，商户做业务处理后，需要以字符串的形式反馈处理
// 结果，内容如下：
// success 处理成功，微信系统收到此结果后不再进行后续通知
// fail 或其它字符处理不成功，微信收到此结果或者没有收到任何结果，系统通过补单机制再次通知
//
// 这是支付成功后通知消息 url query 部分的数据结构
type OrderNotifyURLData struct {
	// 协议参数 ==================================================================

	//ServiceVersion string // 可选, service_version, 版本号, 默认为 1.0
	//SignKeyIndex   int    // 可选, sign_key_index, 多密钥支持的密钥序号, 默认为 1
	Signature  string // 必须, sign, 签名
	SignMethod string // 可选, sign_type, 签名类型，取值：MD5、RSA，默认：MD5
	Charset    string // 可选, input_charset, 字符编码,取值：GBK、UTF-8，默认：GBK。

	// 业务参数 ==================================================================

	NotifyId string // 必须, notify_id, 支付结果通知id，对于某些特定商户，只返回通知id，要求商户据此查询交易结果

	TradeMode     int       // 必须, trade_mode, 交易模式, 1-即时到账, 其他保留
	TradeState    int       // 必须, trade_state, 交易状态(支付结果), 0-成功, 其他保留
	BankBillNo    string    // 可选, bank_billno, 银行订单号
	TransactionId string    // 必须, transaction_id, 交易号，28 位长的数值，其中前10位为商户号，之后8 位为订单产生的日期，如20090415，最后10 位是流水号。
	TimeEnd       time.Time // 必须, time_end, 支付完成时间
	//PayInfo       string    // 可选, pay_info, 支付结果信息, 支付成功时为 ""
	//BuyerAlias    string    // 可选, buyer_alias , 买家别名, 对应买家账号的一个加密串

	// 下面这 4 个字段和支付账单 PayPackage 里的同名字段内容相同
	BankType   string // 必须, bank_type, 银行类型, 微信中固定为 WX
	PartnerId  string // 必须, partner, 商户号，也即之前步骤的partnerid,由微信统一分配的10 位正整数(120XXXXXXX)号
	OutTradeNo string // 必须, out_trade_no, 商户系统的订单号，与请求一致。
	Attach     string // 可选, attach, 商户数据包，原样返回，空参数不传递

	TotalFee     int // 必须, total_fee, 支付金额，单位为分，如果 Discount 有值，通知的 TotalFee + Discount == 请求的 TotalFee
	Discount     int // 可选, discount, 折扣价格，单位分，如果有值，通知的 TotalFee + Discount == 请求的 TotalFee
	TransportFee int // 可选, transport_fee, 物流费用，单位分，默认0。如果有值， 必须保证 TransportFee + ProductFee == TotalFee
	ProductFee   int // 可选, product_fee, 物品费用，单位分。如果有值，必须保证 TransportFee + ProductFee == TotalFee
	FeeType      int // 必须, fee_type, 币种, 目前只支持人民币, 默认值是 1-人民币
}

// 根据 URL RawQuery 来初始化 data *OrderNotifyURLData.
// 如果 RawQuery 里的参数不合法(包括签名不正确) 则返回错误信息, 否则返回 nil.
//  partnerKey: 财付通商户权限密钥Key
func (data *OrderNotifyURLData) CheckAndInit(RawQuery string, partnerKey string) (err error) {
	urlValues, err := url.ParseQuery(RawQuery)
	if err != nil {
		return
	}

	var signature string
	if vs := urlValues["sign"]; len(vs) > 0 && len(vs[0]) > 0 {
		signature = vs[0]
	} else {
		return errors.New("sign is empty")
	}

	var signMethod string
	if vs := urlValues["sign_type"]; len(vs) > 0 && len(vs[0]) > 0 {
		signMethod = vs[0]
	} else {
		signMethod = "MD5"
	}

	// 验证签名是否正确 ===========================================================

	switch signMethod {
	case "MD5", "md5":
		if len(signature) != md5.Size*2 {
			err = fmt.Errorf(`不正确的签名: %q, 长度不对, have: %d, want: %d`,
				signature, len(signature), md5.Size*2)
			return
		}

		urlValues.Del("sign") // sign 不参与签名

		keys := make([]string, 0, len(urlValues))
		for key := range urlValues {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		Hash := md5.New()
		for _, key := range keys {
			// len(urlValues[key]) == 1, 都是单值
			value := urlValues[key][0]
			if len(value) == 0 {
				continue
			}

			Hash.Write([]byte(key))
			Hash.Write([]byte{'='})
			Hash.Write([]byte(value))
			Hash.Write([]byte{'&'})
		}
		Hash.Write([]byte("key="))
		Hash.Write([]byte(partnerKey))

		SignatureHave := make([]byte, md5.Size*2)
		hex.Encode(SignatureHave, Hash.Sum(nil))
		copy(SignatureHave, bytes.ToUpper(SignatureHave))

		if subtle.ConstantTimeCompare(SignatureHave, []byte(signature)) != 1 {
			err = fmt.Errorf("不正确的签名, \r\nhave: %q, \r\nwant: %q", SignatureHave, signature)
			return
		}

	default:
		return fmt.Errorf("没有实现对签名方法 %q 的支持", signMethod)
	}

	// 初始化 ===================================================================

	data.Signature = signature
	data.SignMethod = signMethod

	if vs := urlValues["input_charset"]; len(vs) > 0 && len(vs[0]) > 0 {
		data.Charset = vs[0]
	} else {
		data.Charset = CHARSET_GBK
	}

	if vs := urlValues["notify_id"]; len(vs) > 0 && len(vs[0]) > 0 {
		data.NotifyId = vs[0]
	} else {
		return errors.New("notify_id is empty")
	}

	if vs := urlValues["trade_mode"]; len(vs) > 0 && len(vs[0]) > 0 {
		v0, err := strconv.ParseInt(vs[0], 10, 64)
		if err != nil {
			return err
		}
		data.TradeMode = int(v0)
	} else {
		return errors.New("trade_mode is empty")
	}

	if vs := urlValues["trade_state"]; len(vs) > 0 && len(vs[0]) > 0 {
		v0, err := strconv.ParseInt(vs[0], 10, 64)
		if err != nil {
			return err
		}
		data.TradeState = int(v0)
	} else {
		return errors.New("trade_state is empty")
	}

	if vs := urlValues["bank_billno"]; len(vs) > 0 {
		data.BankBillNo = vs[0]
	}

	if vs := urlValues["transaction_id"]; len(vs) > 0 && len(vs[0]) > 0 {
		data.TransactionId = vs[0]
	} else {
		return errors.New("transaction_id is empty")
	}

	if vs := urlValues["time_end"]; len(vs) > 0 && len(vs[0]) > 0 {
		v0, err := util.ParseTime(vs[0])
		if err != nil {
			return err
		}
		data.TimeEnd = v0
	} else {
		return errors.New("time_end is empty")
	}

	if vs := urlValues["bank_type"]; len(vs) > 0 && len(vs[0]) > 0 {
		data.BankType = vs[0]
	} else {
		return errors.New("bank_type is empty")
	}

	if vs := urlValues["partner"]; len(vs) > 0 && len(vs[0]) > 0 {
		data.PartnerId = vs[0]
	} else {
		return errors.New("partner is empty")
	}

	if vs := urlValues["out_trade_no"]; len(vs) > 0 && len(vs[0]) > 0 {
		data.OutTradeNo = vs[0]
	} else {
		return errors.New("out_trade_no is empty")
	}

	if vs := urlValues["attach"]; len(vs) > 0 {
		data.Attach = vs[0]
	}

	if vs := urlValues["total_fee"]; len(vs) > 0 && len(vs[0]) > 0 {
		v0, err := strconv.ParseInt(vs[0], 10, 64)
		if err != nil {
			return err
		}
		data.TotalFee = int(v0)
	} else {
		return errors.New("total_fee is empty")
	}

	if vs := urlValues["discount"]; len(vs) > 0 && len(vs[0]) > 0 {
		v0, err := strconv.ParseInt(vs[0], 10, 64)
		if err != nil {
			return err
		}
		data.Discount = int(v0)
	}

	if vs := urlValues["transport_fee"]; len(vs) > 0 && len(vs[0]) > 0 {
		v0, err := strconv.ParseInt(vs[0], 10, 64)
		if err != nil {
			return err
		}
		data.TransportFee = int(v0)
	}

	if vs := urlValues["product_fee"]; len(vs) > 0 && len(vs[0]) > 0 {
		v0, err := strconv.ParseInt(vs[0], 10, 64)
		if err != nil {
			return err
		}
		data.ProductFee = int(v0)
	}

	if vs := urlValues["fee_type"]; len(vs) > 0 && len(vs[0]) > 0 {
		v0, err := strconv.ParseInt(vs[0], 10, 64)
		if err != nil {
			return err
		}
		data.FeeType = int(v0)
	} else {
		return errors.New("fee_type is empty")
	}

	return
}
