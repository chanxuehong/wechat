// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package native

import (
	"crypto/sha1"
	"crypto/subtle"
	"encoding/hex"
	"fmt"
	"hash"
	"strconv"
)

// 公众平台接到用户点击 Native 支付 URL 之后, 会调用注册时填写的商户获取订单 Package 的回调 URL.
// 微信公众平台调用时会使用POST方式, 这是推送的 xml 格式的数据结构.
type PayPackageRequest struct {
	XMLName struct{} `xml:"xml" json:"-"`

	AppId     string `xml:"AppId"     json:"AppId"`     // 公众帐号的appid
	NonceStr  string `xml:"NonceStr"  json:"NonceStr"`  // 随机串
	TimeStamp int64  `xml:"TimeStamp" json:"TimeStamp"` // 时间戳

	OpenId      string `xml:"OpenId"      json:"OpenId"`      // 点击链接准备购买商品的用户标识
	IsSubscribe int    `xml:"IsSubscribe" json:"IsSubscribe"` // 标记用户是否订阅该公众帐号，1 为关注，0 为未关注

	ProductId string `xml:"ProductId" json:"ProductId"` // 第三方的商品ID 号

	Signature  string `xml:"AppSignature" json:"AppSignature"` // 参数的加密签名
	SignMethod string `xml:"SignMethod"   json:"SignMethod"`   // 签名方式，目前只支持“SHA1”，该字段不参与签名
}

// 检查 req *PayPackageRequest 的签名是否正确, 正确时返回 nil, 否则返回错误信息.
//  appKey: 即 paySignKey, 公众号支付请求中用于加密的密钥 Key
func (req *PayPackageRequest) CheckSignature(appKey string) (err error) {
	var Hash hash.Hash
	var Signature []byte

	switch req.SignMethod {
	case "sha1", "SHA1":
		if len(req.Signature) != sha1.Size*2 {
			err = fmt.Errorf(`不正确的签名: %q, 长度不对, have: %d, want: %d`,
				req.Signature, len(req.Signature), sha1.Size*2)
			return
		}

		Hash = sha1.New()
		Signature = make([]byte, sha1.Size*2)

	default:
		err = fmt.Errorf(`unknown sign method: %q`, req.SignMethod)
		return
	}

	// 字典序
	// appid
	// appkey
	// issubscribe
	// noncestr
	// openid
	// productid
	// timestamp
	Hash.Write([]byte("appid="))
	Hash.Write([]byte(req.AppId))
	Hash.Write([]byte("&appkey="))
	Hash.Write([]byte(appKey))
	Hash.Write([]byte("&issubscribe="))
	Hash.Write([]byte(strconv.FormatInt(int64(req.IsSubscribe), 10)))
	Hash.Write([]byte("&noncestr="))
	Hash.Write([]byte(req.NonceStr))
	Hash.Write([]byte("&openid="))
	Hash.Write([]byte(req.OpenId))
	Hash.Write([]byte("&productid="))
	Hash.Write([]byte(req.ProductId))
	Hash.Write([]byte("&timestamp="))
	Hash.Write([]byte(strconv.FormatInt(req.TimeStamp, 10)))

	hex.Encode(Signature, Hash.Sum(nil))

	if subtle.ConstantTimeCompare(Signature, []byte(req.Signature)) != 1 {
		err = fmt.Errorf("不正确的签名, \r\nhave: %q, \r\nwant: %q", Signature, req.Signature)
		return
	}
	return
}
