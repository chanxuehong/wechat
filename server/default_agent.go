// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"github.com/chanxuehong/wechat/message/passive/request"
	"net/http"
)

var _ Agent = new(DefaultAgent)

type DefaultAgent struct {
	Id    string // 公众号原始ID, 等于后台中的 公众号设置-->帐号详情-->原始ID
	Token string // 公众号的 Token, 和后台中的设置相等
}

func (this *DefaultAgent) GetId() string {
	return this.Id
}
func (this *DefaultAgent) GetToken() string {
	return this.Token
}

func (this *DefaultAgent) ServeUnknownMsg(w http.ResponseWriter, r *http.Request, rawXMLMsg []byte, timestamp int64) {
}

func (this *DefaultAgent) ServeTextMsg(w http.ResponseWriter, r *http.Request, msg *request.Text, rawXMLMsg []byte, timestamp int64) {
}
func (this *DefaultAgent) ServeImageMsg(w http.ResponseWriter, r *http.Request, msg *request.Image, rawXMLMsg []byte, timestamp int64) {
}
func (this *DefaultAgent) ServeVoiceMsg(w http.ResponseWriter, r *http.Request, msg *request.Voice, rawXMLMsg []byte, timestamp int64) {
}
func (this *DefaultAgent) ServeVideoMsg(w http.ResponseWriter, r *http.Request, msg *request.Video, rawXMLMsg []byte, timestamp int64) {
}
func (this *DefaultAgent) ServeLocationMsg(w http.ResponseWriter, r *http.Request, msg *request.Location, rawXMLMsg []byte, timestamp int64) {
}
func (this *DefaultAgent) ServeLinkMsg(w http.ResponseWriter, r *http.Request, msg *request.Link, rawXMLMsg []byte, timestamp int64) {
}

func (this *DefaultAgent) ServeSubscribeEvent(w http.ResponseWriter, r *http.Request, msg *request.SubscribeEvent, rawXMLMsg []byte, timestamp int64) {
}
func (this *DefaultAgent) ServeUnsubscribeEvent(w http.ResponseWriter, r *http.Request, msg *request.UnsubscribeEvent, rawXMLMsg []byte, timestamp int64) {
}
func (this *DefaultAgent) ServeSubscribeByScanEvent(w http.ResponseWriter, r *http.Request, msg *request.SubscribeByScanEvent, rawXMLMsg []byte, timestamp int64) {
}
func (this *DefaultAgent) ServeScanEvent(w http.ResponseWriter, r *http.Request, msg *request.ScanEvent, rawXMLMsg []byte, timestamp int64) {
}
func (this *DefaultAgent) ServeLocationEvent(w http.ResponseWriter, r *http.Request, msg *request.LocationEvent, rawXMLMsg []byte, timestamp int64) {
}
func (this *DefaultAgent) ServeMenuClickEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuClickEvent, rawXMLMsg []byte, timestamp int64) {
}
func (this *DefaultAgent) ServeMenuViewEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuViewEvent, rawXMLMsg []byte, timestamp int64) {
}
func (this *DefaultAgent) ServeMenuScanCodePushEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuScanCodePushEvent, rawXMLMsg []byte, timestamp int64) {
}
func (this *DefaultAgent) ServeMenuScanCodeWaitMsgEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuScanCodeWaitMsgEvent, rawXMLMsg []byte, timestamp int64) {
}
func (this *DefaultAgent) ServeMenuPicSysPhotoEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuPicSysPhotoEvent, rawXMLMsg []byte, timestamp int64) {
}
func (this *DefaultAgent) ServeMenuPicPhotoOrAlbumEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuPicPhotoOrAlbumEvent, rawXMLMsg []byte, timestamp int64) {
}
func (this *DefaultAgent) ServeMenuPicWeixinEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuPicWeixinEvent, rawXMLMsg []byte, timestamp int64) {
}
func (this *DefaultAgent) ServeMenuLocationSelectEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuLocationSelectEvent, rawXMLMsg []byte, timestamp int64) {
}
func (this *DefaultAgent) ServeMassSendJobFinishEvent(w http.ResponseWriter, r *http.Request, msg *request.MassSendJobFinishEvent, rawXMLMsg []byte, timestamp int64) {
}
func (this *DefaultAgent) ServeTemplateSendJobFinishEvent(w http.ResponseWriter, r *http.Request, msg *request.TemplateSendJobFinishEvent, rawXMLMsg []byte, timestamp int64) {
}
func (this *DefaultAgent) ServeMerchantOrderEvent(w http.ResponseWriter, r *http.Request, msg *request.MerchantOrderEvent, rawXMLMsg []byte, timestamp int64) {
}
