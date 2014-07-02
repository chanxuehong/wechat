// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

var _test_client = func() *Client {
	// 填入正确的 appid, appsecret
	clt := NewClient("appid", "appsecret", nil)

	// 预热
	buf := clt.getBufferFromPool()
	clt.putBufferToPool(buf)

	return clt
}()
