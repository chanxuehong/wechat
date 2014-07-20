// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package js

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
)

// js api 微信支付接口 getBrandWCPayRequest 的参数
type Parameters struct {
	AppId     string `json:"appId"`            // 必须, 公众号 id, 商户注册具有支付权限的公众号成功后即可获得
	TimeStamp int64  `json:"timeStamp,string"` // 必须, unixtime, 商户生成
	NonceStr  string `json:"nonceStr"`         // 必须, 商户生成的随机字符串, 32个字符以内

	Package    string `json:"package"`  // 必须, 订单详情组合成的字符串, 4096个字符以内, see ../Bill.Package
	SignMethod string `json:"signType"` // 必须, 签名方式, 目前仅支持 SHA1
}

// 获取 js api 微信支付接口 getBrandWCPayRequest 的参数 package, JSON 格式.
//  @paySignKey: 公众号支付请求中用于加密的密钥 Key, 对应于支付场景中的 appKey
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
func (para *Parameters) MarshalToJSON(paySignKey string) (jsonBytes []byte, err error) {
	var dst = struct {
		*Parameters
		PaySignature string `json:"paySign"`
	}{
		Parameters: para,
	}

	timestamp := strconv.FormatInt(para.TimeStamp, 10)

	const keysLen = len(`appid=&appkey=&noncestr=&package=&timestamp=`)
	n := keysLen + len(para.AppId) + len(paySignKey) + len(para.NonceStr) +
		len(para.Package) + len(timestamp)

	string1 := make([]byte, 0, n)

	// 字典序
	// appid
	// appkey
	// noncestr
	// package
	// timestamp
	string1 = append(string1, "appid="...)
	string1 = append(string1, para.AppId...)
	string1 = append(string1, "&appkey="...)
	string1 = append(string1, paySignKey...)
	string1 = append(string1, "&noncestr="...)
	string1 = append(string1, para.NonceStr...)
	string1 = append(string1, "&package="...)
	string1 = append(string1, para.Package...)
	string1 = append(string1, "&timestamp="...)
	string1 = append(string1, timestamp...)

	switch para.SignMethod {
	case PARAMETERS_SIGN_MOTHOD_SHA1:
		hashsumArray := sha1.Sum(string1)
		dst.PaySignature = hex.EncodeToString(hashsumArray[:])
	default:
		err = fmt.Errorf("not implement for %s sign method", para.SignMethod)
		return
	}

	jsonBytes, err = json.Marshal(dst)
	return
}
