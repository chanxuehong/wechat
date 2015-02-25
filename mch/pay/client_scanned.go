// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

// 提交被扫支付API.
func (clt *Client) MicroPay(req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML("https://api.mch.weixin.qq.com/pay/micropay", req)
}

// 撤销支付API.
//  NOTE: 请求需要双向证书.
func (clt *Client) Reverse(req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML("https://api.mch.weixin.qq.com/secapi/pay/reverse", req)
}
