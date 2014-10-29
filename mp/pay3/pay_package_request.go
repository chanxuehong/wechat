// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay3

import (
	"bytes"
	"crypto/md5"
	"crypto/subtle"
	"encoding/hex"
	"fmt"
)

// 模式一下请求商户订单信息，微信会主动调用。用户扫码后，微信主动调用把下列信息
// 发送到该链接，商户在返回前先调用统一支付接口，提交订单后返回prepayid，再将prepayid
// 返回给微信。注意：接收该信息前需要先验证签名。
type PayPackageRequest struct {
	XMLName struct{} `xml:"xml" json:"-"`

	AppId       string `xml:"appid"        json:"appid"`        // 必须, 微信分配的公众账号ID
	MerchantId  string `xml:"mch_id"       json:"mch_id"`       // 必须, 微信支付分配的商户号
	OpenId      string `xml:"openid"       json:"openid"`       // 必须, 用户在商户appid 下的唯一标识
	IsSubscribe string `xml:"is_subscribe" json:"is_subscribe"` // 必须, 用户是否关注公众账号，Y-关注，N-未关注，仅在公众账号类型支付有效
	ProductId   string `xml:"product_id"   json:"product_id"`   // 必须, 商户需要定义并维护自己的商品id，这个id 与一张订单等价，微信后台凭借该id 通过POST 商户后台获取交易必须信息；
	NonceStr    string `xml:"nonce_str"    json:"nonce_str"`    // 必须, 随机字符串，不长于32 位
	Signature   string `xml:"sign"         json:"sign"`         // 必须, 签名
}

// 检查 req *PayPackageRequest 的签名是否正确, 正确时返回 nil, 否则返回错误信息.
//  appKey: 商户支付密钥Key
func (req *PayPackageRequest) CheckSignature(appKey string) (err error) {
	if len(req.Signature) != md5.Size*2 {
		err = fmt.Errorf(`不正确的签名: %q, 长度不对, have: %d, want: %d`,
			req.Signature, len(req.Signature), md5.Size*2)
		return
	}

	Hash := md5.New()
	hashsum := make([]byte, md5.Size*2)

	// 字典序
	// appid
	// is_subscribe
	// mch_id
	// nonce_str
	// openid
	// product_id
	if len(req.AppId) > 0 {
		Hash.Write([]byte("appid="))
		Hash.Write([]byte(req.AppId))
		Hash.Write([]byte{'&'})
	}
	if len(req.IsSubscribe) > 0 {
		Hash.Write([]byte("is_subscribe="))
		Hash.Write([]byte(req.IsSubscribe))
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
	if len(req.OpenId) > 0 {
		Hash.Write([]byte("openid="))
		Hash.Write([]byte(req.OpenId))
		Hash.Write([]byte{'&'})
	}
	if len(req.ProductId) > 0 {
		Hash.Write([]byte("product_id="))
		Hash.Write([]byte(req.ProductId))
		Hash.Write([]byte{'&'})
	}

	Hash.Write([]byte("key="))
	Hash.Write([]byte(appKey))

	hex.Encode(hashsum, Hash.Sum(nil))
	hashsum = bytes.ToUpper(hashsum)

	if subtle.ConstantTimeCompare(hashsum, []byte(req.Signature)) != 1 {
		err = fmt.Errorf("签名不匹配, \r\nlocal: %q, \r\ninput: %q", hashsum, req.Signature)
		return
	}
	return
}
