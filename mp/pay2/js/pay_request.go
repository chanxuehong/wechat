// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package js

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"hash"
	"strconv"
)

// js api 微信支付接口 getBrandWCPayRequest 的参数.
//
//  在前端 js 中这样调用:
//
//  WeixinJSBridge.invoke(
//      'getBrandWCPayRequest',
//      {
//          "appId" : "wxf8b4f85f3a794e77",
//          "timeStamp" : "189026618",
//          "nonceStr" : "adssdasssd13d",
//          "package" : "bank_type=WX&body=XXX&fee_type=1&input_charset=GBK&notify_url=http%3a%2f%2fwww.qq.com&out_trade_no=16642817866003386000&partner=1900000109&spbill_create_ip=127.0.0.1&total_fee=1&sign=BEEF37AD19575D92E191C1E4B1474CA9",
//          "signType" : "SHA1",
//          "paySign" : "7717231c335a05165b1874658306fa431fe9a0de"
//      },
//      function(res){
//          // 返回 res.err_msg,取值
//          // get_brand_wcpay_request:cancel 用户取消
//          // get_brand_wcpay_request:fail 发送失败
//          // get_brand_wcpay_request:ok 发送成功
//          WeixinJSBridge.log(res.err_msg);
//          alert(res.err_code+res.err_desc);
//      }
//  );
//
type PayRequestParameters struct {
	AppId     string `json:"appId"`     // 必须, 公众号身份的唯一标识
	NonceStr  string `json:"nonceStr"`  // 必须, 商户生成的随机字符串, 32个字符以内
	TimeStamp string `json:"timeStamp"` // 必须, unixtime, 商户生成

	Package string `json:"package"` // 必须, 订单详情组合成的字符串, 4096个字符以内, see ../PayPackage.Package

	Signature  string `json:"paySign"`  // 必须, 该 PayRequestParameters 自身的签名. see PayRequestParameters.SetSignature
	SignMethod string `json:"signType"` // 必须, 签名方式, 目前仅支持 SHA1
}

func (this *PayRequestParameters) GetTimeStamp() (timestamp int64, err error) {
	return strconv.ParseInt(this.TimeStamp, 10, 64)
}
func (this *PayRequestParameters) SetTimeStamp(timestamp int64) {
	this.TimeStamp = strconv.FormatInt(timestamp, 10)
}

// 设置签名字段.
//  appKey: 即 paySignKey, 公众号支付请求中用于加密的密钥 Key
//
//  NOTE: 要求在 para *PayRequestParameters 其他字段设置完毕后才能调用这个函数, 否则签名就不正确.
func (para *PayRequestParameters) SetSignature(appKey string) (err error) {
	var Hash hash.Hash

	switch para.SignMethod {
	case "sha1", "SHA1":
		Hash = sha1.New()

	default:
		err = fmt.Errorf(`unknown sign method: %q`, para.SignMethod)
		return
	}

	// 字典序
	// appid
	// appkey
	// noncestr
	// package
	// timestamp
	Hash.Write([]byte("appid="))
	Hash.Write([]byte(para.AppId))
	Hash.Write([]byte("&appkey="))
	Hash.Write([]byte(appKey))
	Hash.Write([]byte("&noncestr="))
	Hash.Write([]byte(para.NonceStr))
	Hash.Write([]byte("&package="))
	Hash.Write([]byte(para.Package))
	Hash.Write([]byte("&timestamp="))
	Hash.Write([]byte(para.TimeStamp))

	para.Signature = hex.EncodeToString(Hash.Sum(nil))
	return
}
