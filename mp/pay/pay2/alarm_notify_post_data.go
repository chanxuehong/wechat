// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"crypto/subtle"
	"fmt"

	"github.com/chanxuehong/wechat/mp/pay"
)

// 为了及时通知商户异常，提高商户在微信平台的服务质量。微信后台会向商户推送告警
// 通知，包括发货延迟、调用失败、通知失败等情况，通知的地址是商户在申请支付时填写的
// 告警通知URL，在“公众平台-服务-服务中心-商户功能-商户基本资料-告警通知URL”可
// 以查看。商户接收到告警通知后请尽快修复其中提到的问题，以免影响线上经营。（发货时
// 间要求请参考5.3.1）
// 商户收到告警通知后，需要成功返回success。在通过功能发布检测时，请保证已调通。
//
// 这是告警通知URL 接收的postData 中的 xml 数据
type AlarmNotifyPostData map[string]string

func (data AlarmNotifyPostData) AppId() string {
	return data["AppId"]
}
func (data AlarmNotifyPostData) TimeStamp() string {
	return data["TimeStamp"]
}
func (data AlarmNotifyPostData) ErrorType() string {
	return data["ErrorType"]
}
func (data AlarmNotifyPostData) Content() string {
	return data["AlarmContent"]
}
func (data AlarmNotifyPostData) Description() string {
	return data["Description"]
}
func (data AlarmNotifyPostData) Signature() string {
	return data["AppSignature"]
}
func (data AlarmNotifyPostData) SignMethod() string {
	return data["SignMethod"]
}

// 检查 AlarmNotifyPostData 的签名是否合法, 合法返回 nil, 否则返回错误信息.
//  appKey: 即 paySignKey, 公众号支付请求中用于加密的密钥 Key
func (data AlarmNotifyPostData) CheckSignature(appKey string) (err error) {
	Signature1 := data["AppSignature"]
	SignMethod := data["SignMethod"]

	switch SignMethod {
	case "sha1", "SHA1":
		if len(Signature1) != 40 {
			err = fmt.Errorf(`不正确的签名: %q, 长度不对, have: %d, want: %d`,
				Signature1, len(Signature1), 40)
			return
		}

		Signature2 := pay.WXSHA1Sign1(data, appKey, []string{"AppSignature", "SignMethod"})

		if subtle.ConstantTimeCompare([]byte(Signature2), []byte(Signature1)) != 1 {
			err = fmt.Errorf("签名不匹配, \r\nlocal: %q, \r\ninput: %q", Signature2, Signature1)
			return
		}
		return

	default:
		err = fmt.Errorf(`unknown sign method: %q`, SignMethod)
		return
	}
}
