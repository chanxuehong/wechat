// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

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
type OrderNotifyPostData map[string]string

func (data OrderNotifyPostData) AppId() string {
	return data["AppId"]
}
func (data OrderNotifyPostData) NonceStr() string {
	return data["NonceStr"]
}
func (data OrderNotifyPostData) TimeStamp() string {
	return data["TimeStamp"]
}
func (data OrderNotifyPostData) OpenId() string {
	return data["OpenId"]
}
func (data OrderNotifyPostData) IsSubscribe() string {
	return data["IsSubscribe"]
}
func (data OrderNotifyPostData) Signature() string {
	return data["AppSignature"]
}
func (data OrderNotifyPostData) SignMethod() string {
	return data["SignMethod"]
}

// 检查 OrderNotifyPostData 的签名是否合法, 合法返回 nil, 否则返回错误信息.
//  appKey: 即 paySignKey, 公众号支付请求中用于加密的密钥 Key
func (data OrderNotifyPostData) CheckSignature(appKey string) (err error) {
	return AlarmNotifyPostData(data).CheckSignature(appKey)
}
