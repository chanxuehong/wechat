// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package native

import (
	"crypto/sha1"
	"encoding/hex"
	"strconv"
)

// 生成 native 支付 URL.
//  @appId:      必须, 商户注册具有支付权限的公众号成功后即可获得
//  @paySignKey: 必须, 公众号支付请求中用于加密的密钥 Key
//  @nonceStr:   必须, 32个字符以内, 商户生成的随机字符串
//  @productId:  必须, 32个字符以内, 商户需要定义并维护自己的商品id, 这个id与一张订单等价, 微信后台凭借该id通过POST商户后台获取交易必须信息；
//  @timestamp:  必须, unixtime, 商户生成
func NativeURL(appId, paySignKey, nonceStr, productId string, timestamp int64) string {
	timestampStr := strconv.FormatInt(timestamp, 10)

	// len(`appid=&appkey=&noncestr=&productid=&timestamp=`) == 46
	n := 46 + len(appId) + len(paySignKey) + len(nonceStr) +
		len(productId) + len(timestampStr)

	buf := make([]byte, 0, n)

	// 字典序
	// appid
	// appkey
	// noncestr
	// productid
	// timestamp
	buf = append(buf, "appid="...)
	buf = append(buf, appId...)
	buf = append(buf, "&appkey="...)
	buf = append(buf, paySignKey...)
	buf = append(buf, "&noncestr="...)
	buf = append(buf, nonceStr...)
	buf = append(buf, "&productid="...)
	buf = append(buf, productId...)
	buf = append(buf, "&timestamp="...)
	buf = append(buf, timestampStr...)

	hashsumArray := sha1.Sum(buf)
	hashsumStr := hex.EncodeToString(hashsumArray[:])

	return "weixin://wxpay/bizpayurl?appid=" + appId +
		"&noncestr=" + nonceStr +
		"&productid=" + productId +
		"&sign=" + hashsumStr +
		"&timestamp=" + timestampStr
}
