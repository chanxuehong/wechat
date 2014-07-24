// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

// 创建订单查询的 package 数据
func MakeOrderQueryRequestPackage(
	OutTradeNo string, // 第三方唯一订单号
	PartnerId string, // 财付通商户身份标识
	PartnerKey string, // 财付通商户权限密钥

) string {

	// 签名的源
	const keysLen1 = len(`out_trade_no=&partner=&key=`)
	n1 := keysLen1 + len(OutTradeNo) + len(PartnerId) + len(PartnerKey)

	// 最终结果
	const keysLen2 = len(`out_trade_no=&partner=&sign=`)
	n2 := keysLen2 + len(OutTradeNo) + len(PartnerId) + 32 // md5sum

	var n int
	if n1 >= n2 {
		n = n1
	} else {
		n = n2
	}
	buf := make([]byte, n)

	string1 := buf[:0]

	// 字典序
	// out_trade_no
	// partner
	string1 = append(string1, "out_trade_no="...)
	string1 = append(string1, OutTradeNo...)
	string1 = append(string1, "&partner="...)
	string1 = append(string1, PartnerId...)

	string2 := string1 // 到目前为止两者相同

	string1 = append(string1, "&key="...)
	string1 = append(string1, PartnerKey...)

	hashSumArray := md5.Sum(string1)

	string2 = append(string2, "&sign="...)
	signature := buf[len(string2) : len(string2)+32]
	hex.Encode(signature, hashSumArray[:])
	string2 = append(string2, bytes.ToUpper(signature)...)

	return string(string2)
}

// 因为某一方技术的原因，可能导致商户在预期时间内都收不到最终支付通知，
// 此时商户可以通过API来查询订单的详细支付状态。
//
// 订单查询的请求数据
type OrderQueryRequest struct {
	AppId      string `json:"appid"`            // 公众平台账户的 AppId
	Package    string `json:"package"`          // 查询订单的关键信息数据, see MakeOrderQueryRequestPackage
	TimeStamp  int64  `json:"timestamp,string"` // 时间戳, unixtime
	Signature  string `json:"app_signature"`    // 签名
	SignMethod string `json:"sign_method"`      // 签名方法
}

// 设置签名字段.
//  @paySignKey: 公众号支付请求中用于加密的密钥 Key, 对应于支付场景中的 appKey
//  NOTE: 要求在 req *OrderQueryRequest 其他字段设置完毕后才能调用这个函数, 否则签名就不正确.
func (req *OrderQueryRequest) SetSignature(paySignKey string) (err error) {
	var SumFunc func([]byte) []byte // hash 签名函数

	switch {
	case strings.ToLower(req.SignMethod) == "sha1":
		req.SignMethod = "sha1"
		SumFunc = func(src []byte) (hashsum []byte) {
			hashsumArray := sha1.Sum(src)
			hashsum = hashsumArray[:]
			return
		}

	default:
		err = fmt.Errorf(`not implement for "%s" sign method`, req.SignMethod)
		return
	}

	TimeStampStr := strconv.FormatInt(req.TimeStamp, 10)

	const keysLen = len(`appid=&appkey=&package=&timestamp=`)
	n := keysLen + len(req.AppId) + len(paySignKey) + len(req.Package) + len(TimeStampStr)

	string1 := make([]byte, 0, n)

	// 字典序
	// appid
	// appkey
	// package
	// timestamp
	string1 = append(string1, "appid="...)
	string1 = append(string1, req.AppId...)
	string1 = append(string1, "&appkey="...)
	string1 = append(string1, paySignKey...)
	string1 = append(string1, "&package="...)
	string1 = append(string1, req.Package...)
	string1 = append(string1, "&timestamp="...)
	string1 = append(string1, TimeStampStr...)

	req.Signature = hex.EncodeToString(SumFunc(string1))
	return
}
