// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay3

// 支付完成后，微信会把相关支付和用户信息发送到该 notify URL，商户需要接收处理信息。
// 对后台通知交互时，如果微信收到商户的应答不是成功或超时，微信认为通知失败，微
// 信会通过一定的策略（如30 分钟共8 次）定期重新发起通知，尽可能提高通知的成功率，
// 但微信不保证通知最终能成功。
// 由于存在重新发送后台通知的情况，因此同样的通知可能会多次发送给商户系统。商户
// 系统必须能够正确处理重复的通知。
//
// 这是通知的 post 参数的数据结构
type OrderNotifyRequest map[string]string

// 检查 OrderNotifyRequest 的签名是否正确, 正确时返回 nil, 否则返回错误信息.
//  appKey: 商户支付密钥Key
func (req OrderNotifyRequest) CheckSignature(appKey string) (err error) {
	if req["return_code"] != RET_CODE_SUCCESS {
		return
	}

	return CheckSignature(req, appKey)
}
