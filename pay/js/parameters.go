// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package js

import (
	"crypto/sha1"
	"encoding/hex"
	"strconv"
)

// jsapi 微信支付接口 getBrandWCPayRequest 的参数集合
//
// WeixinJSBridge.invoke(
//     'getBrandWCPayRequest',
//     {
//         "appId" : "wxf8b4f85f3a794e77",
//         "timeStamp" : "189026618",
//         "nonceStr" : "adssdasssd13d",
//         "package" : "bank_type=WX&body=XXX&fee_type=1&input_charset=GBK&notify_url=http%3a%2f%2fwww.qq.com&out_trade_no=16642817866003386000&partner=1900000109&spbill_create_ip=127.0.0.1&total_fee=1&sign=BEEF37AD19575D92E191C1E4B1474CA9",
//         "signType" : "SHA1",
//         "paySign" : "7717231c335a05165b1874658306fa431fe9a0de"
//     },
//     function(res){
//         // 返回 res.err_msg,取值
//         // get_brand_wcpay_request:cancel 用户取消
//         // get_brand_wcpay_request:fail 发送失败
//         // get_brand_wcpay_request:ok 发送成功
//         WeixinJSBridge.log(res.err_msg);
//         alert(res.err_code+res.err_desc);
//     }
// );
type Parameters struct {
	AppId        string `json:"appId"`            // 必须, 商户注册具有支付权限的公众号成功后即可获得
	TimeStamp    int64  `json:"timeStamp,string"` // 必须, unixtime, 商户生成
	NonceStr     string `json:"nonceStr"`         // 必须, 商户生成的随机字符串
	Package      string `json:"package"`          // 必须, 订单详情组合成的字符串, see OrderInfo.Package
	SignType     string `json:"signType"`         // 必须, 目前仅支持 SHA1
	PaySignature string `json:"paySign"`          // 必须, Parameters 的签名, see Parameters.SetPaySignature 函数
}

// 将 Parameters 做签名, 把签名填入 PaySign 字段.
//  NOTE: 需要在设置了其他字段后才能调用.
//  @paySignKey: 公众号支付请求中用于加密的密钥 Key
func (para *Parameters) SetPaySignature(paySignKey string) {
	timestamp := strconv.FormatInt(para.TimeStamp, 10)

	// appid=&appkey=&noncestr=&package=&timestamp=
	n := 44 + len(para.AppId) + len(paySignKey) + len(para.NonceStr) +
		len(para.Package) + len(timestamp)
	buf := make([]byte, 0, n)

	// 字典序
	// appid
	// appkey
	// noncestr
	// package
	// timestamp
	buf = append(buf, "appid="...)
	buf = append(buf, para.AppId...)
	buf = append(buf, "&appkey="...)
	buf = append(buf, paySignKey...)
	buf = append(buf, "&noncestr="...)
	buf = append(buf, para.NonceStr...)
	buf = append(buf, "&package="...)
	buf = append(buf, para.Package...)
	buf = append(buf, "&timestamp="...)
	buf = append(buf, timestamp...)

	// 目前仅支持 SHA1
	hashsumArray := sha1.Sum(buf)
	para.PaySignature = hex.EncodeToString(hashsumArray[:])
}
