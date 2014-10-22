// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"crypto/sha1"
	"crypto/subtle"
	"encoding/hex"
	"fmt"
	"hash"
	"strconv"
)

// 用户在成功完成支付后，微信后台通知商户服务器（notify_url）支付结果。
// 商户可以使用notify_url 的通知结果进行个性化页面的展示。
//
// 对后台通知交互时，如果微信收到商户的应答不是success 或超时，微信认为通知失败，
// 微信会通过一定的策略（如30 分钟共8 次）定期重新发起通知，尽可能提高通知的成功率，
// 但微信不保证通知最终能成功。
// 由于存在重新发送后台通知的情况，因此同样的通知可能会多次发送给商户系统。商户
// 系统必须能够正确处理重复的通知。
//
// 微信后台通过 notify_url 通知商户，商户做业务处理后，需要以字符串的形式反馈处理
// 结果，内容如下：
// success 处理成功，微信系统收到此结果后不再进行后续通知
// fail 或其它字符处理不成功，微信收到此结果或者没有收到任何结果，系统通过补单机制再次通知
//
// 这是支付成功后通知消息 post 部分的数据结构.
type OrderNotifyPostData struct {
	XMLName struct{} `xml:"xml" json:"-"`

	AppId     string `xml:"AppId"     json:"AppId"`     // 公众号id
	NonceStr  string `xml:"NonceStr"  json:"NonceStr"`  // 随机字符串
	TimeStamp int64  `xml:"TimeStamp" json:"TimeStamp"` // 时间戳, unixtime

	OpenId      string `xml:"OpenId"      json:"OpenId"`      // 支付该笔订单的用户ID，商户可通过公众号其他接口为付款用户服务。
	IsSubscribe int    `xml:"IsSubscribe" json:"IsSubscribe"` // 标记用户是否关注了公众号。1 为关注，0 为未关注。

	Signature  string `xml:"AppSignature" json:"AppSignature"` // 签名
	SignMethod string `xml:"SignMethod"   json:"SignMethod"`   // 签名方式, 目前只支持"sha1"
}

// 检查 data *OrderNotifyPostData 的签名是否正确, 正确时返回 nil, 否则返回错误信息.
//  appKey: 即 paySignKey, 公众号支付请求中用于加密的密钥 Key
func (data *OrderNotifyPostData) CheckSignature(appKey string) (err error) {
	var Hash hash.Hash
	var Signature []byte

	switch data.SignMethod {
	case "sha1", "SHA1":
		if len(data.Signature) != sha1.Size*2 {
			err = fmt.Errorf(`不正确的签名: %q, 长度不对, have: %d, want: %d`,
				data.Signature, len(data.Signature), sha1.Size*2)
			return
		}

		Hash = sha1.New()
		Signature = make([]byte, sha1.Size*2)

	default:
		err = fmt.Errorf(`unknown sign method: %q`, data.SignMethod)
		return
	}

	// 字典序
	// appid
	// appkey
	// issubscribe
	// noncestr
	// openid
	// timestamp
	Hash.Write([]byte("appid="))
	Hash.Write([]byte(data.AppId))
	Hash.Write([]byte("&appkey="))
	Hash.Write([]byte(appKey))
	Hash.Write([]byte("&issubscribe="))
	Hash.Write([]byte(strconv.FormatInt(int64(data.IsSubscribe), 10)))
	Hash.Write([]byte("&noncestr="))
	Hash.Write([]byte(data.NonceStr))
	Hash.Write([]byte("&openid="))
	Hash.Write([]byte(data.OpenId))
	Hash.Write([]byte("&timestamp="))
	Hash.Write([]byte(strconv.FormatInt(data.TimeStamp, 10)))

	hex.Encode(Signature, Hash.Sum(nil))

	if subtle.ConstantTimeCompare(Signature, []byte(data.Signature)) != 1 {
		err = fmt.Errorf("不正确的签名, \r\nhave: %q, \r\nwant: %q", Signature, data.Signature)
		return
	}
	return
}
