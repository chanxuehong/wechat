// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"github.com/chanxuehong/wechat/pay"
	"github.com/chanxuehong/wechat/pay/feedback"
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

// 为了及时通知商户异常，提高商户在微信平台的服务质量。
// 微信后台会向商户推送告警通知，包括发货延迟、调用失败、通知失败等情况
//  NOTE: 参数 *pay.AlarmNotifyData 已经经过验证了, 无需再次认证签名!
type AlarmNotifyHandlerFunc func(http.ResponseWriter, *http.Request, *pay.AlarmNotifyData)

// 处理维权接口用户投诉消息
//  NOTE: 参数 *feedback.Request 已经经过验证了, 无需再次认证签名!
type FeedbackRequestHandlerFunc func(http.ResponseWriter, *http.Request, *feedback.Request)

// 处理维权接口用户确认消除投诉消息
//  NOTE: 参数 *feedback.Confirm 已经经过验证了, 无需再次认证签名!
type FeedbackConfirmHandlerFunc func(http.ResponseWriter, *http.Request, *feedback.Confirm)

// 处理维权接口用户拒绝消除投诉消息
//  NOTE: 参数 *feedback.Reject 已经经过验证了, 无需再次认证签名!
type FeedbackRejectHandlerFunc func(http.ResponseWriter, *http.Request, *feedback.Reject)
