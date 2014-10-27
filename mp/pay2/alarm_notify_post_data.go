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

// 为了及时通知商户异常，提高商户在微信平台的服务质量。微信后台会向商户推送告警
// 通知，包括发货延迟、调用失败、通知失败等情况，通知的地址是商户在申请支付时填写的
// 告警通知URL，在“公众平台-服务-服务中心-商户功能-商户基本资料-告警通知URL”可
// 以查看。商户接收到告警通知后请尽快修复其中提到的问题，以免影响线上经营。（发货时
// 间要求请参考5.3.1）
// 商户收到告警通知后，需要成功返回success。在通过功能发布检测时，请保证已调通。
//
// 这是告警通知URL 接收的postData 中的 xml 数据
type AlarmNotifyPostData struct {
	XMLName struct{} `xml:"xml" json:"-"`

	AppId     string `xml:"AppId"     json:"AppId"`
	TimeStamp int64  `xml:"TimeStamp" json:"TimeStamp"`

	ErrorType   int    `xml:"ErrorType"    json:"ErrorType"` // 1001:发货超时
	Description string `xml:"Description"  json:"Description"`
	Content     string `xml:"AlarmContent" json:"AlarmContent"`

	Signature  string `xml:"AppSignature" json:"AppSignature"`
	SignMethod string `xml:"SignMethod"   json:"SignMethod"`
}

// 检查 data *AlarmNotifyPostData 的签名是否合法, 合法返回 nil, 否则返回错误信息.
//  appKey: 即 paySignKey, 公众号支付请求中用于加密的密钥 Key
func (data *AlarmNotifyPostData) CheckSignature(appKey string) (err error) {
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
	// alarmcontent
	// appid
	// appkey
	// description
	// errortype
	// timestamp
	Hash.Write([]byte("alarmcontent="))
	Hash.Write([]byte(data.Content))
	Hash.Write([]byte("&appid="))
	Hash.Write([]byte(data.AppId))
	Hash.Write([]byte("&appkey="))
	Hash.Write([]byte(appKey))
	Hash.Write([]byte("&description="))
	Hash.Write([]byte(data.Description))
	Hash.Write([]byte("&errortype="))
	Hash.Write([]byte(strconv.FormatInt(int64(data.ErrorType), 10)))
	Hash.Write([]byte("&timestamp="))
	Hash.Write([]byte(strconv.FormatInt(data.TimeStamp, 10)))

	hex.Encode(Signature, Hash.Sum(nil))

	if subtle.ConstantTimeCompare(Signature, []byte(data.Signature)) != 1 {
		err = fmt.Errorf("不正确的签名, \r\nhave: %q, \r\nwant: %q", Signature, data.Signature)
		return
	}
	return
}
