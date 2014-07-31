// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

var _test_handler = func() *Handler {
	handler := NewHandler(&HandlerSetting{})
	// 预热
	unit := handler.getBufferUnitFromPool()
	handler.putBufferUnitToPool(unit)

	return handler
}()
