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

// js api 编辑并获取收货地址 editAddress 的参数.
//
//  在前端 js 中这样调用:
//
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
	AppId     string `json:"appId"`     // 必须, 公众号身份的唯一标识
	NonceStr  string `json:"nonceStr"`  // 必须, 商户生成的随机字符串, 32个字符以内
	TimeStamp string `json:"timeStamp"` // 必须, unixtime, 商户生成

	Scope string `json:"scope"` // 必须, 填写"jsapi_address", 获得编辑地址权限

	Signature  string `json:"addrSign"` // 必须, 该 EditAddressParameters 自身的签名. see EditAddressParameters.SetSignature
	SignMethod string `json:"signType"` // 必须, 签名方式, 目前仅支持 sha1
}

func (this *EditAddressParameters) GetTimeStamp() (timestamp int64, err error) {
	return strconv.ParseInt(this.TimeStamp, 10, 64)
}
func (this *EditAddressParameters) SetTimeStamp(timestamp int64) {
	this.TimeStamp = strconv.FormatInt(timestamp, 10)
}

// 设置签名字段.
//  url:               当前网页的 URL
//  oauth2AccessToken: oauth2 用户授权凭证
//
//  NOTE: 要求在 para *EditAddressParameters 其他字段设置完毕后才能调用这个函数, 否则签名就不正确.
func (para *EditAddressParameters) SetSignature(url, oauth2AccessToken string) (err error) {
	var Hash hash.Hash

	switch para.SignMethod {
	case "sha1", "SHA1":
		Hash = sha1.New()

	default:
		err = fmt.Errorf(`unknown sign method: %q`, para.SignMethod)
		return
	}

	// 字典序
	// accesstoken
	// appid
	// noncestr
	// timestamp
	// url
	Hash.Write([]byte("accesstoken="))
	Hash.Write([]byte(oauth2AccessToken))
	Hash.Write([]byte("&appid="))
	Hash.Write([]byte(para.AppId))
	Hash.Write([]byte("&noncestr="))
	Hash.Write([]byte(para.NonceStr))
	Hash.Write([]byte("&timestamp="))
	Hash.Write([]byte(para.TimeStamp))
	Hash.Write([]byte("&url="))
	Hash.Write([]byte(url))

	para.Signature = hex.EncodeToString(Hash.Sum(nil))
	return
}
