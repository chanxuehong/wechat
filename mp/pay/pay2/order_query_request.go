// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"hash"
	"strconv"
)

// 因为某一方技术的原因，可能导致商户在预期时间内都收不到最终支付通知，此时商户
// 可以通过该API 来查询订单的详细支付状态。
//
// 订单查询的真正数据是放在PostData 中的，格式为json
type OrderQueryRequest struct {
	AppId      string `json:"appid"`         // 公众平台账户的 AppId
	Package    string `json:"package"`       // 查询订单的关键信息数据, see MakeOrderQueryRequestPackage
	TimeStamp  string `json:"timestamp"`     // 时间戳, unixtime
	Signature  string `json:"app_signature"` // 签名
	SignMethod string `json:"sign_method"`   // 签名方法
}

func (this *OrderQueryRequest) SetTimeStamp(timestamp int64) {
	this.TimeStamp = strconv.FormatInt(timestamp, 10)
}

// 设置签名字段.
//  appKey: 即 paySignKey, 公众号支付请求中用于加密的密钥 Key
//
//  NOTE: 要求在 req *OrderQueryRequest 其他字段设置完毕后才能调用这个函数, 否则签名就不正确.
func (req *OrderQueryRequest) SetSignature(appKey string) (err error) {
	var Hash hash.Hash

	switch req.SignMethod {
	case "sha1", "SHA1":
		Hash = sha1.New()

	default:
		err = fmt.Errorf(`unknown sign method: %q`, req.SignMethod)
		return
	}

	// 字典序
	// appid
	// appkey
	// package
	// timestamp
	Hash.Write([]byte("appid="))
	Hash.Write([]byte(req.AppId))
	Hash.Write([]byte("&appkey="))
	Hash.Write([]byte(appKey))
	Hash.Write([]byte("&package="))
	Hash.Write([]byte(req.Package))
	Hash.Write([]byte("&timestamp="))
	Hash.Write([]byte(req.TimeStamp))

	req.Signature = hex.EncodeToString(Hash.Sum(nil))
	return
}

// 创建订单查询的 package 数据
func MakeOrderQueryRequestPackage(
	OutTradeNo string, // 第三方唯一订单号
	PartnerId string, // 财付通商户身份标识
	PartnerKey string, // 财付通商户权限密钥
) string {

	const keysLen1 = len(`out_trade_no=&partner=&key=`)
	n1 := keysLen1 + len(OutTradeNo) + len(PartnerId) + len(PartnerKey)

	const keysLen2 = len(`out_trade_no=&partner=&sign=`)
	n2 := keysLen2 + len(OutTradeNo) + len(PartnerId) + md5.Size*2 // md5sum

	// 一次性分配需要的内存
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
	signature := buf[len(string2) : len(string2)+md5.Size*2]
	hex.Encode(signature, hashSumArray[:])
	string2 = append(string2, bytes.ToUpper(signature)...)

	return string(string2)
}
