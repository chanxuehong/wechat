// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"github.com/chanxuehong/wechat/pay"
	"github.com/chanxuehong/wechat/pay/native"
	"net/http"
)

// 非法请求的处理函数, 比如签名认证不通过, 等等.
//  @err: 具体的错误信息
type InvalidRequestHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)

// 微信请求获取订单详情的处理函数.
//  NOTE: 参数 *native.BillRequest 已经经过验证了, 无需再次认证签名!
type BillRequestHandlerFunc func(http.ResponseWriter, *http.Request, *native.BillRequest)

// 支付成功后, 微信服务器会通知支付结果, 该函数就是处理这个通知的.
//  NOTE: 参数 *pay.OrderNotifyPostData, *pay.OrderNotifyURLDataVer1 已经经过验证了, 是合法的通知消息!
type OrderNotifyHandlerFuncVer1 func(http.ResponseWriter, *http.Request, *pay.OrderNotifyPostData, *pay.OrderNotifyURLDataVer1)
