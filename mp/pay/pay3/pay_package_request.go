// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay3

// Native支付——请求商家获取商品信息 请求参数
type PayPackageRequest map[string]string

// 检查 PayPackageRequest 的签名是否正确, 正确时返回 nil, 否则返回错误信息.
//  appKey: 商户支付密钥Key
func (req PayPackageRequest) CheckSignature(appKey string) (err error) {
	return CheckSignature(req, appKey)
}
