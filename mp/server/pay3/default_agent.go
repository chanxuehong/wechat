// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay3

import (
	"net/http"
)

var _ Agent = new(DefaultAgent)

type DefaultAgent struct {
	AppId  string
	MchId  string
	AppKey string
}

func (this *DefaultAgent) GetAppId() string {
	return this.AppId
}
func (this *DefaultAgent) GetMchId() string {
	return this.MchId
}
func (this *DefaultAgent) GetAppKey() string {
	return this.AppKey
}

func (this *DefaultAgent) ServeUnknownMsg(w http.ResponseWriter, r *http.Request, postRawXMLMsg []byte) {
}
func (this *DefaultAgent) ServePayPackageRequest(w http.ResponseWriter, r *http.Request, req map[string]string, postRawXMLMsg []byte) {
}
func (this *DefaultAgent) ServeOrderNotification(w http.ResponseWriter, r *http.Request, data map[string]string, postRawXMLMsg []byte) {
}
