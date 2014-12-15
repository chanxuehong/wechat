// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay3

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

func (this *DefaultAgent) ServeUnknownMsg(para *RequestParameters) {
}
func (this *DefaultAgent) ServePayPackageRequest(req map[string]string, para *RequestParameters) {
}
func (this *DefaultAgent) ServeOrderNotification(data map[string]string, para *RequestParameters) {
}
