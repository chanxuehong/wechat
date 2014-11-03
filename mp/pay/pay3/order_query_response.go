// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay3

// 订单查询接口 返回参数
type OrderQueryResponse map[string]string

// 检查 OrderQueryResponse 的签名是否正确, 正确时返回 nil, 否则返回错误信息.
//  appKey: 商户支付密钥Key
func (resp OrderQueryResponse) CheckSignature(appKey string) (err error) {
	if resp["return_code"] != RET_CODE_SUCCESS {
		return
	}

	return CheckSignature(resp, appKey)
}
