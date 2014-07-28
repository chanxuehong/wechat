// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package js

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

// js api 编辑并获取收货地址 editAddress 的参数.
//
//  在前端 js 中这样调用:
//  WeixinJSBridge.invoke(
//      'editAddress',
//      {
//          "appId" : getAppId(),
//          "scope" : "jsapi_address",
//          "signType" : "sha1",
//          "addrSign" : "xxxxx",
//          "timeStamp" : "12345",
//          "nonceStr" : "10000",
//      },
//      function(res) {
//          // 若res 中所带的返回值不为空，则表示用户选择该返回值作为收货地址。
//          // 否则若返回空，则表示用户取消了这一次编辑收货地址。
//          document.form1.address1.value = res.proviceFirstStageName;
//          document.form1.address2.value = res.addressCitySecondStageName;
//          document.form1.address3.value = res.addressCountiesThirdStageName;
//          document.form1.detail.value = res.addressDetailInfo;
//          document.form1.phone.value = res.telNumber;
//      }
//  });
//
type EditAddressParameters struct {
	AppId      string `json:"appId"`            // 公众号 id
	NonceStr   string `json:"nonceStr"`         // 随机字符串
	TimeStamp  int64  `json:"timeStamp,string"` // 时间戳, unixtime
	Scope      string `json:"scope"`            // 填写"jsapi_address", 获得编辑地址权限
	Signature  string `json:"addrSign"`         // 签名
	SignMethod string `json:"signType"`         // 签名方式, 目前仅支持SHA1
}

// 设置签名字段.
//  @url:               当前网页的 URL
//  @oauth2AccessToken: oauth2 用户授权凭证
//  NOTE: 要求在 para *EditAddressParameters 其他字段设置完毕后才能调用这个函数, 否则签名就不正确.
func (para *EditAddressParameters) SetSignature(url, oauth2AccessToken string) (err error) {
	var sumFunc hashSumFunc

	switch {
	case para.SignMethod == SIGN_METHOD_SHA1:
		fallthrough

	case strings.ToLower(para.SignMethod) == "sha1":
		para.SignMethod = "sha1"
		sumFunc = sha1Sum

	default:
		err = fmt.Errorf(`not implement for "%s" sign method`, para.SignMethod)
		return
	}

	TimeStampStr := strconv.FormatInt(para.TimeStamp, 10)

	const keysLen = len(`accesstoken=&appid=&noncestr=&timestamp=&url=`)
	n := keysLen + len(oauth2AccessToken) + len(para.AppId) + len(para.NonceStr) +
		len(TimeStampStr) + len(url)

	string1 := make([]byte, 0, n)

	// 字典序
	// accesstoken
	// appid
	// noncestr
	// timestamp
	// url
	string1 = append(string1, "accesstoken="...)
	string1 = append(string1, oauth2AccessToken...)
	string1 = append(string1, "&appid="...)
	string1 = append(string1, para.AppId...)
	string1 = append(string1, "&noncestr="...)
	string1 = append(string1, para.NonceStr...)
	string1 = append(string1, "&timestamp="...)
	string1 = append(string1, TimeStampStr...)
	string1 = append(string1, "&url="...)
	string1 = append(string1, url...)

	para.Signature = hex.EncodeToString(sumFunc(string1))
	return
}
