// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/chanxuehong/wechat/mp/pay"
)

// 因为某一方技术的原因，可能导致商户在预期时间内都收不到最终支付通知，此时商户
// 可以通过该API 来查询订单的详细支付状态。
//
// 订单查询的真正数据是放在PostData 中的，格式为json
type OrderQueryRequest map[string]string

func (req OrderQueryRequest) SetAppId(AppId string) {
	req["appid"] = AppId
}
func (req OrderQueryRequest) SetTimeStamp(t time.Time) {
	req["timestamp"] = strconv.FormatInt(t.Unix(), 10)
}
func (req OrderQueryRequest) SetPackage(OutTradeNo, PartnerId, PartnerKey string) {
	const keysLen1 = len(`out_trade_no=&partner=&key=`)
	n1 := keysLen1 + len(OutTradeNo) + len(PartnerId) + len(PartnerKey)

	escapedOutTradeNo := pay.URLEscape(OutTradeNo)
	escapedPartnerId := pay.URLEscape(PartnerId)
	const keysLen2 = len(`out_trade_no=&partner=&sign=`)
	n2 := keysLen2 + len(escapedOutTradeNo) + len(escapedPartnerId) + 32 // md5sum

	var buf []byte
	if n1 >= n2 {
		buf = make([]byte, n1)
	} else {
		buf = make([]byte, n2)
	}
	string1 := buf[:0]
	string2 := buf[:0]

	// 字典序
	// out_trade_no
	// partner
	string1 = append(string1, "out_trade_no="...)
	string1 = append(string1, OutTradeNo...)
	string1 = append(string1, "&partner="...)
	string1 = append(string1, PartnerId...)
	string1 = append(string1, "&key="...)
	string1 = append(string1, PartnerKey...)

	hashsum := md5.Sum(string1)
	signature := make([]byte, 32)
	hex.Encode(signature, hashsum[:])
	signature = bytes.ToUpper(signature)

	// 字典序
	// out_trade_no
	// partner
	string2 = append(string2, "out_trade_no="...)
	string2 = append(string2, escapedOutTradeNo...)
	string2 = append(string2, "&partner="...)
	string2 = append(string2, escapedPartnerId...)
	string2 = append(string2, "&sign="...)
	string2 = append(string2, signature...)

	req["package"] = string(string2)
}
func (req OrderQueryRequest) SetSignMethod(SignMethod string) {
	req["sign_method"] = SignMethod
}

// 设置签名字段.
//  appKey: 即 paySignKey, 公众号支付请求中用于加密的密钥 Key
//
//  NOTE: 要求在 OrderQueryRequest 其他字段设置完毕后才能调用这个函数, 否则签名就不正确.
func (req OrderQueryRequest) SetSignature(appKey string) (err error) {
	SignMethod := req["sign_method"]

	switch SignMethod {
	case "sha1", "SHA1":
		req["app_signature"] = pay.WXSHA1Sign1(req, appKey, []string{"app_signature", "sign_method"})
		return

	default:
		err = fmt.Errorf(`unknown sign method: %q`, SignMethod)
		return
	}
}
