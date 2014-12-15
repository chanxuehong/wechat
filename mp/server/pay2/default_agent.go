// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"github.com/chanxuehong/wechat/mp/pay/feedback"
	"github.com/chanxuehong/wechat/mp/pay/pay2"
)

var _ Agent = new(DefaultAgent)

type DefaultAgent struct {
	AppId      string
	AppKey     string
	PartnerId  string
	PartnerKey string
}

func (this *DefaultAgent) GetAppId() string {
	return this.AppId
}
func (this *DefaultAgent) GetAppKey() string {
	return this.AppKey
}
func (this *DefaultAgent) GetPartnerId() string {
	return this.PartnerId
}
func (this *DefaultAgent) GetPartnerKey() string {
	return this.PartnerKey
}

func (this *DefaultAgent) ServeUnknownMsg(para *RequestParameters) {
}
func (this *DefaultAgent) ServePayPackageRequest(req pay2.PayPackageRequest, para *RequestParameters) {
}
func (this *DefaultAgent) ServeOrderNotification(urlData pay2.OrderNotifyURLData, postData pay2.OrderNotifyPostData, para *RequestParameters) {
}
func (this *DefaultAgent) ServeAlarmNotification(data pay2.AlarmNotifyPostData, para *RequestParameters) {
}
func (this *DefaultAgent) ServeFeedbackComplaint(req *feedback.Complaint, para *RequestParameters) {
}
func (this *DefaultAgent) ServeFeedbackConfirmation(req *feedback.Confirmation, para *RequestParameters) {
}
func (this *DefaultAgent) ServeFeedbackRejection(req *feedback.Rejection, para *RequestParameters) {
}
