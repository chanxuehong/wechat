// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
)

// 告警通知.
// 为了及时通知商户异常, 提高商户在微信平台的服务质量. 微信后台会向商户推送告警通知,
// 包括发货延迟, 调用失败, 通知失败等情况, 通知的地址是商户在申请支付时填写的告警通知 URL,
// 在"公众平台-服务-服务中心-商户功能-商户基本资料-告警通知URL"可以查看.
// 商户接收到告警通知后请尽快修复其中提到的问题, 以免影响线上经营
//
// 商户收到告警通知后, 需要成功返回 success. 在通过功能发布检测时, 请保证已调通.
//
// 这是告警通知URL接收的postData的xml数据结构.
type AlarmNotifyData struct {
	XMLName struct{} `xml:"xml" json:"-"`

	AppId     string `xml:"AppId"`
	TimeStamp int64  `xml:"TimeStamp"`

	ErrCode     int    `xml:"ErrorType"`
	Description string `xml:"Description"`
	Content     string `xml:"AlarmContent"`

	Signature  string `xml:"AppSignature"`
	SignMethod string `xml:"SignMethod"`
}

// 检查 data *AlarmNotifyData 是否合法(包括签名的检查), 合法返回 nil, 否则返回错误信息.
//  @paySignKey: 公众号支付请求中用于加密的密钥 Key, 对应于支付场景中的 appKey
func (data *AlarmNotifyData) Check(paySignKey string) (err error) {
	// 检查签名
	var hashSumLen, twoHashSumLen int
	var sumFunc hashSumFunc

	switch data.SignMethod {
	case "sha1", "SHA1":
		hashSumLen = 40
		twoHashSumLen = 80
		sumFunc = sha1Sum

	default:
		err = fmt.Errorf(`not implement for "%s" sign method`, data.SignMethod)
		return
	}

	if len(data.Signature) != hashSumLen {
		err = errors.New("签名不正确")
		return
	}

	ErrCodeStr := strconv.FormatInt(int64(data.ErrCode), 10)
	TimeStampStr := strconv.FormatInt(data.TimeStamp, 10)

	const keysLen = len(`alarmcontent=&appid=&appkey=&description=&errortype=&timestamp=`)

	n := twoHashSumLen + keysLen + len(data.Content) + len(data.AppId) + len(paySignKey) +
		len(data.Description) + len(ErrCodeStr) + len(TimeStampStr)

	// buf[:hashSumLen] 保存参数 data.Signature,
	// buf[hashSumLen:twoHashSumLen] 保存生成的签名
	// buf[twoHashSumLen:] 按照字典序列保存 string1
	buf := make([]byte, n)
	dataSignature := buf[:hashSumLen]
	signature := buf[hashSumLen:twoHashSumLen]
	string1 := buf[twoHashSumLen:twoHashSumLen]

	copy(dataSignature, data.Signature)

	// 字典序
	// alarmcontent
	// appid
	// appkey
	// description
	// errortype
	// timestamp
	string1 = append(string1, "alarmcontent="...)
	string1 = append(string1, data.Content...)
	string1 = append(string1, "&appid="...)
	string1 = append(string1, data.AppId...)
	string1 = append(string1, "&appkey="...)
	string1 = append(string1, paySignKey...)
	string1 = append(string1, "&description="...)
	string1 = append(string1, data.Description...)
	string1 = append(string1, "&errortype="...)
	string1 = append(string1, ErrCodeStr...)
	string1 = append(string1, "&timestamp="...)
	string1 = append(string1, TimeStampStr...)

	hex.Encode(signature, sumFunc(string1))

	// 采用 subtle.ConstantTimeCompare 是防止 计时攻击!
	if subtle.ConstantTimeCompare(dataSignature, signature) != 1 {
		err = errors.New("签名不正确")
		return
	}
	return
}
