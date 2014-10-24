// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay3

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
)

// 订单查询接口 请求参数
type OrderQueryRequest struct {
	XMLName struct{} `xml:"xml" json:"-"`

	AppId         string `xml:"appid"                    json:"appid"`                    // 必须, 微信分配的公众账号ID
	MerchantId    string `xml:"mch_id"                   json:"mch_id"`                   // 必须, 微信支付分配的商户号
	TransactionId string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"` // 可选, 微信的订单号，优先使用
	OutTradeNo    string `xml:"out_trade_no,omitempty"   json:"out_trade_no,omitempty"`   // 必须, 商户系统内部的订单号,transaction_id、out_trade_no 二选一，如果同时存在优先级: transaction_id > out_trade_no
	NonceStr      string `xml:"nonce_str"                json:"nonce_str"`                // 必须, 随机字符串，不长于32 位
	Signature     string `xml:"sign"                     json:"sign"`                     // 必须, 签名
}

// 设置签名字段.
//  appKey: 商户支付密钥Key
//
//  NOTE: 要求在 req *OrderQueryRequest 其他字段设置完毕后才能调用这个函数, 否则签名就不正确.
func (req *OrderQueryRequest) SetSignature(appKey string) (err error) {
	Hash := md5.New()
	Signature := make([]byte, md5.Size*2)

	// 字典序
	// appid
	// mch_id
	// nonce_str
	// out_trade_no
	// transaction_id
	if len(req.AppId) > 0 {
		Hash.Write([]byte("appid="))
		Hash.Write([]byte(req.AppId))
		Hash.Write([]byte{'&'})
	}
	if len(req.MerchantId) > 0 {
		Hash.Write([]byte("mch_id="))
		Hash.Write([]byte(req.MerchantId))
		Hash.Write([]byte{'&'})
	}
	if len(req.NonceStr) > 0 {
		Hash.Write([]byte("nonce_str="))
		Hash.Write([]byte(req.NonceStr))
		Hash.Write([]byte{'&'})
	}
	if len(req.OutTradeNo) > 0 {
		Hash.Write([]byte("out_trade_no="))
		Hash.Write([]byte(req.OutTradeNo))
		Hash.Write([]byte{'&'})
	}
	if len(req.TransactionId) > 0 {
		Hash.Write([]byte("transaction_id="))
		Hash.Write([]byte(req.TransactionId))
		Hash.Write([]byte{'&'})
	}

	Hash.Write([]byte("key="))
	Hash.Write([]byte(appKey))

	hex.Encode(Signature, Hash.Sum(nil))
	Signature = bytes.ToUpper(Signature)

	req.Signature = string(Signature)
	return
}
