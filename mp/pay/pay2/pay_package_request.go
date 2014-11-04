// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

// 公众平台接到用户点击 Native 支付 URL 之后, 会调用注册时填写的商户获取订单 Package 的回调 URL.
// 微信公众平台调用时会使用POST方式, 这是推送的 xml 格式的数据结构.
type PayPackageRequest map[string]string

func (req PayPackageRequest) AppId() string {
	return req["AppId"]
}
func (req PayPackageRequest) NonceStr() string {
	return req["NonceStr"]
}
func (req PayPackageRequest) TimeStamp() string {
	return req["TimeStamp"]
}
func (req PayPackageRequest) OpenId() string {
	return req["OpenId"]
}
func (req PayPackageRequest) IsSubscribe() string {
	return req["IsSubscribe"]
}
func (req PayPackageRequest) ProductId() string {
	return req["ProductId"]
}
func (req PayPackageRequest) Signature() string {
	return req["AppSignature"]
}
func (req PayPackageRequest) SignMethod() string {
	return req["SignMethod"]
}

// 检查 PayPackageRequest 的签名是否合法, 合法返回 nil, 否则返回错误信息.
//  appKey: 即 paySignKey, 公众号支付请求中用于加密的密钥 Key
func (req PayPackageRequest) CheckSignature(appKey string) (err error) {
	return AlarmNotifyPostData(req).CheckSignature(appKey)
}
