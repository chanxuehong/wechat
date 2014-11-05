// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"fmt"
	"strconv"
	"time"

	"github.com/chanxuehong/wechat/mp/pay"
)

// 公众平台接到用户点击 Native 支付 URL 之后, 会调用注册时填写的商户获取订单 Package 的回调 URL.
// 这是获取订单详情 package 的回复消息数据结构, xml 格式.
type PayPackageResponse map[string]string

func (resp PayPackageResponse) SetAppId(AppId string) {
	resp["AppId"] = AppId
}
func (resp PayPackageResponse) SetNonceStr(NonceStr string) {
	resp["NonceStr"] = NonceStr
}
func (resp PayPackageResponse) SetTimeStamp(t time.Time) {
	resp["TimeStamp"] = strconv.FormatInt(t.Unix(), 10)
}
func (resp PayPackageResponse) SetPackage(Package string) {
	resp["Package"] = Package
}
func (resp PayPackageResponse) SetRetCode(RetCode int) {
	resp["RetCode"] = strconv.FormatInt(int64(RetCode), 10)
}
func (resp PayPackageResponse) SetRetMsg(RetMsg string) {
	resp["RetErrMsg"] = RetMsg
}
func (resp PayPackageResponse) SetSignMethod(SignMethod string) {
	resp["SignMethod"] = SignMethod
}

// 设置签名字段.
//  appKey: 即 paySignKey, 公众号支付请求中用于加密的密钥 Key
//
//  NOTE: 要求在 PayPackageResponse 其他字段设置完毕后才能调用这个函数, 否则签名就不正确.
func (resp PayPackageResponse) SetSignature(appKey string) (err error) {
	SignMethod := resp["SignMethod"]

	switch SignMethod {
	case "sha1", "SHA1":
		resp["AppSignature"] = pay.WXSHA1Sign1(resp, appKey, []string{"AppSignature", "SignMethod"})
		return

	default:
		return fmt.Errorf(`unknown sign method: %q`, SignMethod)
	}
}
