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
	"fmt"
	"net/url"
	"sort"
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
type OrderNotifyURLData url.Values

func (data OrderNotifyURLData) Signature() string {
	return url.Values(data).Get("sign")
}
func (data OrderNotifyURLData) SignMethod() string {
	str := url.Values(data).Get("sign_type")
	if str == "" {
		return "MD5"
	}
	return str
}
func (data OrderNotifyURLData) Charset() string {
	str := url.Values(data).Get("input_charset")
	if str == "" {
		return "GBK"
	}
	return str
}
func (data OrderNotifyURLData) NotifyId() string {
	return url.Values(data).Get("notify_id")
}
func (data OrderNotifyURLData) TradeMode() string {
	return url.Values(data).Get("trade_mode")
}
func (data OrderNotifyURLData) TradeState() string {
	return url.Values(data).Get("trade_state")
}
func (data OrderNotifyURLData) BankBillNo() string {
	return url.Values(data).Get("bank_billno")
}
func (data OrderNotifyURLData) TransactionId() string {
	return url.Values(data).Get("transaction_id")
}
func (data OrderNotifyURLData) TimeEnd() string {
	return url.Values(data).Get("time_end")
}
func (data OrderNotifyURLData) BankType() string {
	return url.Values(data).Get("bank_type")
}
func (data OrderNotifyURLData) PartnerId() string {
	return url.Values(data).Get("partner")
}
func (data OrderNotifyURLData) OutTradeNo() string {
	return url.Values(data).Get("out_trade_no")
}
func (data OrderNotifyURLData) Attach() string {
	return url.Values(data).Get("attach")
}
func (data OrderNotifyURLData) TotalFee() string {
	return url.Values(data).Get("total_fee")
}
func (data OrderNotifyURLData) Discount() string {
	return url.Values(data).Get("discount")
}
func (data OrderNotifyURLData) TransportFee() string {
	return url.Values(data).Get("transport_fee")
}
func (data OrderNotifyURLData) ProductFee() string {
	return url.Values(data).Get("product_fee")
}
func (data OrderNotifyURLData) FeeType() string {
	return url.Values(data).Get("fee_type")
}

// 检查 OrderNotifyURLData 的签名是否合法, 合法返回 nil, 否则返回错误信息.
//  partnerKey: 财付通商户权限密钥Key
func (data OrderNotifyURLData) CheckSignature(partnerKey string) (err error) {
	Signature := data.Signature()
	SignMethod := data.SignMethod()

	switch SignMethod {
	case "MD5", "md5":
		if len(Signature) != 32 {
			err = fmt.Errorf(`不正确的签名: %q, 长度不对, have: %d, want: %d`,
				Signature, len(Signature), 32)
			return
		}

		keys := make([]string, 0, len(data))
		for key := range data {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		Hash := md5.New()
		hashsum := make([]byte, 32)

		for _, key := range keys {
			if key == "sign" {
				continue
			}

			value := data[key][0] // len(data[key]) > 0
			if value == "" {
				continue
			}

			Hash.Write([]byte(key))
			Hash.Write([]byte{'='})
			Hash.Write([]byte(value))
			Hash.Write([]byte{'&'})
		}
		Hash.Write([]byte("key="))
		Hash.Write([]byte(partnerKey))

		hex.Encode(hashsum, Hash.Sum(nil))
		hashsum = bytes.ToUpper(hashsum)

		if subtle.ConstantTimeCompare(hashsum, []byte(Signature)) != 1 {
			err = fmt.Errorf("签名不匹配, \r\nlocal: %q, \r\ninput: %q", hashsum, Signature)
			return
		}
		return

	default:
		err = fmt.Errorf(`unknown sign method: %q`, SignMethod)
		return
	}
}
