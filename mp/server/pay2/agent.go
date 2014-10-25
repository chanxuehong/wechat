// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"github.com/chanxuehong/wechat/mp/pay2"
	"github.com/chanxuehong/wechat/mp/pay2/feedback"
	"net/http"
)

type Agent interface {
	GetAppId() string
	GetAppKey() string
	GetPartnerId() string
	GetPartnerKey() string

	// 未知类型的消息处理方法
	//  rawXMLMsg 是 xml 消息体
	ServeUnknownMsg(w http.ResponseWriter, r *http.Request, rawXMLMsg []byte)

	// Native（原生）支付回调商户后台获取package
	//  rawXMLMsg 是 xml 消息体
	ServePayPackageRequest(w http.ResponseWriter, r *http.Request, req *pay2.PayPackageRequest, rawXMLMsg []byte)

	// 用户在成功完成支付后，微信后台通知（POST）商户服务器（notify_url）支付结果。
	//  postRawXMLMsg 是 postData 的原始 xml 消息体
	ServeOrderNotification(w http.ResponseWriter, r *http.Request, urlData *pay2.OrderNotifyURLData, postData *pay2.OrderNotifyPostData, postRawXMLMsg []byte)

	// 微信后台向商户推送告警通知的处理方法
	//  rawXMLMsg 是 xml 消息体
	ServeAlarmNotification(w http.ResponseWriter, r *http.Request, data *pay2.AlarmNotifyPostData, rawXMLMsg []byte)

	// 用户维权系统接口的 投诉 处理方法
	//  rawXMLMsg 是 xml 消息体
	ServeFeedbackComplaint(w http.ResponseWriter, r *http.Request, req *feedback.Complaint, rawXMLMsg []byte)
	// 用户维权系统接口的 确认消除投诉 的处理方法
	//  rawXMLMsg 是 xml 消息体
	ServeFeedbackConfirmation(w http.ResponseWriter, r *http.Request, req *feedback.Confirmation, rawXMLMsg []byte)
	// 用户维权系统接口的 拒绝消除投诉 的处理方法
	//  rawXMLMsg 是 xml 消息体
	ServeFeedbackRejection(w http.ResponseWriter, r *http.Request, req *feedback.Rejection, rawXMLMsg []byte)
}
