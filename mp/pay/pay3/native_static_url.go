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

// 生成 native 支付 静态链接 URL.
//  AppId:      必须, 微信分配的公众账号ID
//  AppKey:     必须, 即 paySignKey, 公众号支付请求中用于加密的密钥 Key
//  NonceStr:   必须, 32个字符以内, 商户生成的随机字符串
//  Timestamp:  必须, unixtime, 商户生成
//  ProductId:  必须, 32个字符以内, 商户需要定义并维护自己的商品id，这个id 与一张订单等价，
//                   微信后台凭借该id 通过POST 商户后台获取交易必须信息；传此参数必须在
//                   申请的时候配置了Package 请求回调地址；
//  MerchantId: 必须, 微信支付分配的商户号
//
//  NOTE: 该函数没有做 url escape, 因为正常情况下根本不需要做 url escape
func NativeURL(AppId, AppKey, NonceStr, Timestamp, ProductId, MerchantId string) string {
	Hash := md5.New()
	hashsum := make([]byte, md5.Size*2)

	// 字典序
	// appid
	// mch_id
	// nonce_str
	// product_id
	// time_stamp
	if len(AppId) > 0 {
		Hash.Write([]byte("appid="))
		Hash.Write([]byte(AppId))
		Hash.Write([]byte{'&'})
	}
	if len(MerchantId) > 0 {
		Hash.Write([]byte("mch_id="))
		Hash.Write([]byte(MerchantId))
		Hash.Write([]byte{'&'})
	}
	if len(NonceStr) > 0 {
		Hash.Write([]byte("nonce_str="))
		Hash.Write([]byte(NonceStr))
		Hash.Write([]byte{'&'})
	}
	if len(ProductId) > 0 {
		Hash.Write([]byte("product_id="))
		Hash.Write([]byte(ProductId))
		Hash.Write([]byte{'&'})
	}
	if len(Timestamp) > 0 {
		Hash.Write([]byte("time_stamp="))
		Hash.Write([]byte(Timestamp))
		Hash.Write([]byte{'&'})
	}
	Hash.Write([]byte("key="))
	Hash.Write([]byte(AppKey))

	hex.Encode(hashsum, Hash.Sum(nil))
	hashsum = bytes.ToUpper(hashsum)

	// weixin://wxpay/bizpayurl?sign=XXXXX&appid=XXXXX&mch_id=XXXXX
	//          &product_id=XXXXXX&time_stamp=XXXXXX&nonce_str=XXXXX
	return "weixin://wxpay/bizpayurl?sign=" + string(hashsum) +
		"&appid=" + AppId +
		"&mch_id=" + MerchantId +
		"&product_id=" + ProductId +
		"&time_stamp=" + Timestamp +
		"&nonce_str=" + NonceStr
}
