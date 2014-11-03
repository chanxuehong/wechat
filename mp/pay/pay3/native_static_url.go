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
//  appId:      必须, 微信分配的公众账号ID
//  appKey:     必须, 即 paySignKey, 公众号支付请求中用于加密的密钥 Key
//  nonceStr:   必须, 32个字符以内, 商户生成的随机字符串
//  timestamp:  必须, unixtime, 商户生成
//  productId:  必须, 32个字符以内, 商户需要定义并维护自己的商品id，这个id 与一张订单等价，
//                   微信后台凭借该id 通过POST 商户后台获取交易必须信息；传此参数必须在
//                   申请的时候配置了Package 请求回调地址；
//  merchantId: 必须, 微信支付分配的商户号
//
//  NOTE: 该函数没有做 url escape, 因为正常情况下根本不需要做 url escape
func NativeURL(appId, appKey, nonceStr, timestamp, productId, merchantId string) string {
	Hash := md5.New()
	hashsum := make([]byte, md5.Size*2)

	// 字典序
	// appid
	// mch_id
	// nonce_str
	// product_id
	// time_stamp
	if len(appId) > 0 {
		Hash.Write([]byte("appid="))
		Hash.Write([]byte(appId))
		Hash.Write([]byte{'&'})
	}
	if len(merchantId) > 0 {
		Hash.Write([]byte("mch_id="))
		Hash.Write([]byte(merchantId))
		Hash.Write([]byte{'&'})
	}
	if len(nonceStr) > 0 {
		Hash.Write([]byte("nonce_str="))
		Hash.Write([]byte(nonceStr))
		Hash.Write([]byte{'&'})
	}
	if len(productId) > 0 {
		Hash.Write([]byte("product_id="))
		Hash.Write([]byte(productId))
		Hash.Write([]byte{'&'})
	}
	if len(timestamp) > 0 {
		Hash.Write([]byte("time_stamp="))
		Hash.Write([]byte(timestamp))
		Hash.Write([]byte{'&'})
	}
	Hash.Write([]byte("key="))
	Hash.Write([]byte(appKey))

	hex.Encode(hashsum, Hash.Sum(nil))
	hashsum = bytes.ToUpper(hashsum)

	// weixin://wxpay/bizpayurl?sign=XXXXX&appid=XXXXX&mch_id=XXXXX
	//          &product_id=XXXXXX&time_stamp=XXXXXX&nonce_str=XXXXX
	return "weixin://wxpay/bizpayurl?sign=" + string(hashsum) +
		"&appid=" + appId +
		"&mch_id=" + merchantId +
		"&product_id=" + productId +
		"&time_stamp=" + timestamp +
		"&nonce_str=" + nonceStr
}
