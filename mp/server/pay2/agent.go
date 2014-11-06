// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"github.com/chanxuehong/wechat/mp/pay/feedback"
	"github.com/chanxuehong/wechat/mp/pay/pay2"
	"net/http"
)

// 微信支付消息处理接口
type Agent interface {
	GetAppId() string      // 公众号身份的唯一标识
	GetAppKey() string     // 公众号支付请求中用于加密的密钥Key，可验证商户唯一身份
	GetPartnerId() string  // 财付通商户身份的标识
	GetPartnerKey() string // 财付通商户权限密钥Key

	// 未知类型的消息处理方法
	//  postRawXMLMsg 是 xml 消息体
	ServeUnknownMsg(w http.ResponseWriter, r *http.Request, postRawXMLMsg []byte)

	// Native（原生）支付回调商户后台获取package
	//  postRawXMLMsg 是 xml 消息体
	ServePayPackageRequest(w http.ResponseWriter, r *http.Request, req pay2.PayPackageRequest, postRawXMLMsg []byte)

	// 用户在成功完成支付后，微信后台通知（POST）商户服务器（notify_url）支付结果。
	//  postRawXMLMsg 是 postData 的原始 xml 消息体
	ServeOrderNotification(w http.ResponseWriter, r *http.Request, urlData pay2.OrderNotifyURLData, postData pay2.OrderNotifyPostData, postRawXMLMsg []byte)

	// 微信后台向商户推送告警通知的处理方法
	//  postRawXMLMsg 是 xml 消息体
	ServeAlarmNotification(w http.ResponseWriter, r *http.Request, data pay2.AlarmNotifyPostData, postRawXMLMsg []byte)

	// 用户维权系统接口的 投诉 处理方法
	//  postRawXMLMsg 是 xml 消息体
	ServeFeedbackComplaint(w http.ResponseWriter, r *http.Request, req *feedback.Complaint, postRawXMLMsg []byte)
	// 用户维权系统接口的 确认消除投诉 的处理方法
	//  postRawXMLMsg 是 xml 消息体
	ServeFeedbackConfirmation(w http.ResponseWriter, r *http.Request, req *feedback.Confirmation, postRawXMLMsg []byte)
	// 用户维权系统接口的 拒绝消除投诉 的处理方法
	//  postRawXMLMsg 是 xml 消息体
	ServeFeedbackRejection(w http.ResponseWriter, r *http.Request, req *feedback.Rejection, postRawXMLMsg []byte)
}
