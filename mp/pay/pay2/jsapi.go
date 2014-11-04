// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"crypto/sha1"
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

	Package string `json:"package"` // 必须, 订单详情组合成的字符串, 4096个字符以内, MakePayPackage

	Signature  string `json:"paySign"`  // 必须, 该 PayRequestParameters 自身的签名. see PayRequestParameters.SetSignature
	SignMethod string `json:"signType"` // 必须, 签名方式, 目前仅支持 SHA1
}

func (this *JSAPIPayRequestParameters) SetTimeStamp(t time.Time) {
	this.TimeStamp = strconv.FormatInt(t.Unix(), 10)
}

// 设置签名字段.
//  appKey: 即 paySignKey, 公众号支付请求中用于加密的密钥 Key
//
//  NOTE: 要求在 para *PayRequestParameters 其他字段设置完毕后才能调用这个函数, 否则签名就不正确.
func (para *JSAPIPayRequestParameters) SetSignature(appKey string) (err error) {
	var Hash hash.Hash
	var hashsum []byte

	switch para.SignMethod {
	case "SHA1", "sha1":
		Hash = sha1.New()
		hashsum = make([]byte, 40)

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

	hex.Encode(hashsum, Hash.Sum(nil))
	para.Signature = string(hashsum)
	return
}
