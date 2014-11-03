// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"crypto/sha1"
	"encoding/hex"
)

// 生成 native 支付 URL.
//  AppId:      必须, 公众号身份的唯一标识
//  AppKey:     必须, 即 paySignKey, 公众号支付请求中用于加密的密钥 Key
//  NonceStr:   必须, 32个字符以内, 商户生成的随机字符串
//  Timestamp:  必须, unixtime, 商户生成
//  ProductId:  必须, 32个字符以内, 商户需要定义并维护自己的商品id, 这个id与一张订单等价,
//              微信后台凭借该id通过POST商户后台获取交易必须信息;
//
//  NOTE: 该函数没有做 url escape, 因为正常情况下根本不需要做 url escape
func NativeURL(AppId, AppKey, NonceStr, Timestamp, ProductId string) string {
	Hash := sha1.New()
	hashsum := make([]byte, sha1.Size*2)

	// 字典序
	// appid
	// appkey
	// noncestr
	// productid
	// timestamp
	Hash.Write([]byte("appid="))
	Hash.Write([]byte(AppId))
	Hash.Write([]byte("&appkey="))
	Hash.Write([]byte(AppKey))
	Hash.Write([]byte("&noncestr="))
	Hash.Write([]byte(NonceStr))
	Hash.Write([]byte("&productid="))
	Hash.Write([]byte(ProductId))
	Hash.Write([]byte("&timestamp="))
	Hash.Write([]byte(Timestamp))

	hex.Encode(hashsum, Hash.Sum(nil))
	signature := string(hashsum)

	// weixin://wxpay/bizpayurl?sign=XXXXX&appid=XXXXXX&productid=XXXXXX
	// &timestamp=XXXXXX&noncestr=XXXXXX
	return "weixin://wxpay/bizpayurl?sign=" + signature +
		"&appid=" + AppId +
		"&productid=" + ProductId +
		"&timestamp=" + Timestamp +
		"&noncestr=" + NonceStr
}
