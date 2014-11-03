// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay3

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash"
	"strconv"
	"time"
)

type JSAPIPayRequestParameters struct {
	AppId     string `json:"appId"`     // 必须, 公众号身份的唯一标识
	NonceStr  string `json:"nonceStr"`  // 必须, 商户生成的随机字符串, 32个字符以内
	TimeStamp string `json:"timeStamp"` // 必须, unixtime, 商户生成

	Package string `json:"package"` // 必须, 统一支付接口返回的 prepay_id 参数值，提交格式如：prepay_id=***

	Signature  string `json:"paySign"`  // 必须, see JSAPIPayRequestParameters.SetSignature
	SignMethod string `json:"signType"` // 必须, 签名方式, 按照文档中所示填入MD5
}

func (this *JSAPIPayRequestParameters) SetTimeStamp(t time.Time) {
	this.TimeStamp = strconv.FormatInt(t.Unix(), 10)
}

// 设置签名字段.
//  appKey: 商户支付密钥Key
//
//  NOTE: 要求在 JSAPIPayRequestParameters 其他字段设置完毕后才能调用这个函数, 否则签名就不正确.
func (para *JSAPIPayRequestParameters) SetSignature(appKey string) (err error) {
	var Hash hash.Hash
	var hashsum []byte

	switch para.SignMethod {
	case "md5", "MD5":
		Hash = md5.New()
		hashsum = make([]byte, md5.Size*2)

	default:
		err = fmt.Errorf(`unknown sign method: %q`, para.SignMethod)
		return
	}

	// 字典序
	// appId
	// nonceStr
	// package
	// signType
	// timeStamp
	if len(para.AppId) > 0 {
		Hash.Write([]byte("appId="))
		Hash.Write([]byte(para.AppId))
		Hash.Write([]byte{'&'})
	}
	if len(para.NonceStr) > 0 {
		Hash.Write([]byte("nonceStr="))
		Hash.Write([]byte(para.NonceStr))
		Hash.Write([]byte{'&'})
	}
	if len(para.Package) > 0 {
		Hash.Write([]byte("package="))
		Hash.Write([]byte(para.Package))
		Hash.Write([]byte{'&'})
	}
	if len(para.SignMethod) > 0 {
		Hash.Write([]byte("signType="))
		Hash.Write([]byte(para.SignMethod))
		Hash.Write([]byte{'&'})
	}
	if len(para.TimeStamp) > 0 {
		Hash.Write([]byte("timeStamp="))
		Hash.Write([]byte(para.TimeStamp))
		Hash.Write([]byte{'&'})
	}
	Hash.Write([]byte("key="))
	Hash.Write([]byte(appKey))

	hex.Encode(hashsum, Hash.Sum(nil))
	hashsum = bytes.ToUpper(hashsum)

	para.Signature = string(hashsum)
	return
}
