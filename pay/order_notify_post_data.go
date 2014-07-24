// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"crypto/sha1"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
)

// 用户在成功完成支付后，微信后台通知（POST）商户服务器（notify_url）支付结果。
// 商户可以使用 notify_url 的通知结果进行个性化页面的展示。
//
// 这是支付成功后通知消息 post 部分
type OrderNotifyPostData struct {
	XMLName struct{} `xml:"xml" json:"-"`

	// 下面这三个字段和之前传过去的一样? 可以比对下, 确保安全?
	AppId     string `xml:"AppId"`     // 必须, 公众号 id
	NonceStr  string `xml:"NonceStr"`  // 必须, 随机字符串
	TimeStamp int64  `xml:"TimeStamp"` // 必须, 时间戳, unixtime

	OpenId      string `xml:"OpenId"`      // 必须, 支付该笔订单的用户 OpenId
	IsSubscribe int    `xml:"IsSubscribe"` // 必须, 标记用户是否订阅该公众帐号, 1为关注, 0为未关注

	Signature  string `xml:"AppSignature"` // 必须, 参数的加密签名
	SignMethod string `xml:"SignMethod"`   // 必须, 签名方式, 目前只支持"sha1"
}

// 检查 data *OrderNotifyPostData 是否合法(包括签名的检查), 合法返回 nil, 否则返回错误信息.
//  @paySignKey: 公众号支付请求中用于加密的密钥 Key, 对应于支付场景中的 appKey
func (data *OrderNotifyPostData) Check(paySignKey string) (err error) {
	// 检查签名
	var hashSumLen, twoHashSumLen int
	var SumFunc func([]byte) []byte // hash 签名函数

	switch data.SignMethod {
	case "sha1", "SHA1":
		hashSumLen = 40
		twoHashSumLen = 80
		SumFunc = func(src []byte) (hashsum []byte) {
			hashsumArray := sha1.Sum(src)
			hashsum = hashsumArray[:]
			return
		}

	default:
		err = fmt.Errorf(`not implement for "%s" sign method`, data.SignMethod)
		return
	}

	if len(data.Signature) != hashSumLen {
		err = errors.New("签名不正确")
		return
	}

	TimeStampStr := strconv.FormatInt(data.TimeStamp, 10)
	IsSubscribeStr := strconv.FormatInt(int64(data.IsSubscribe), 10)

	const keysLen = len(`appid=&appkey=&issubscribe=&noncestr=&openid=&timestamp=`)

	n := twoHashSumLen + keysLen + len(data.AppId) + len(paySignKey) + len(IsSubscribeStr) +
		len(data.NonceStr) + len(data.OpenId) + len(TimeStampStr)

	// buf[:hashSumLen] 保存参数 data.Signature,
	// buf[hashSumLen:twoHashSumLen] 保存生成的签名
	// buf[twoHashSumLen:] 按照字典序列保存 string1
	buf := make([]byte, n)
	dataSignature := buf[:hashSumLen]
	signature := buf[hashSumLen:twoHashSumLen]
	string1 := buf[twoHashSumLen:twoHashSumLen]

	copy(dataSignature, data.Signature)

	// 字典序
	// appid
	// appkey
	// issubscribe
	// noncestr
	// openid
	// timestamp
	string1 = append(string1, "appid="...)
	string1 = append(string1, data.AppId...)
	string1 = append(string1, "&appkey="...)
	string1 = append(string1, paySignKey...)
	string1 = append(string1, "&issubscribe="...)
	string1 = append(string1, IsSubscribeStr...)
	string1 = append(string1, "&noncestr="...)
	string1 = append(string1, data.NonceStr...)
	string1 = append(string1, "&openid="...)
	string1 = append(string1, data.OpenId...)
	string1 = append(string1, "&timestamp="...)
	string1 = append(string1, TimeStampStr...)

	hex.Encode(signature, SumFunc(string1))

	// 采用 subtle.ConstantTimeCompare 是防止 计时攻击!
	if subtle.ConstantTimeCompare(dataSignature, signature) != 1 {
		err = errors.New("签名不正确")
		return
	}
	return
}
