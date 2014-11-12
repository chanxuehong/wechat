// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"net/http"

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

func (this *DefaultAgent) ServeUnknownMsg(w http.ResponseWriter, r *http.Request, postRawXMLMsg []byte) {
}
func (this *DefaultAgent) ServePayPackageRequest(w http.ResponseWriter, r *http.Request, req pay2.PayPackageRequest, postRawXMLMsg []byte) {
}
func (this *DefaultAgent) ServeOrderNotification(w http.ResponseWriter, r *http.Request, urlData pay2.OrderNotifyURLData, postData pay2.OrderNotifyPostData, postRawXMLMsg []byte) {
}
func (this *DefaultAgent) ServeAlarmNotification(w http.ResponseWriter, r *http.Request, data pay2.AlarmNotifyPostData, postRawXMLMsg []byte) {
}
func (this *DefaultAgent) ServeFeedbackComplaint(w http.ResponseWriter, r *http.Request, req *feedback.Complaint, postRawXMLMsg []byte) {
}
func (this *DefaultAgent) ServeFeedbackConfirmation(w http.ResponseWriter, r *http.Request, req *feedback.Confirmation, postRawXMLMsg []byte) {
}
func (this *DefaultAgent) ServeFeedbackRejection(w http.ResponseWriter, r *http.Request, req *feedback.Rejection, postRawXMLMsg []byte) {
}
