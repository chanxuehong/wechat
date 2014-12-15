// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"sync"

	"github.com/chanxuehong/wechat/mp/message/passive/request"
)

var _ Agent = new(DefaultAgent)

type DefaultAgent struct {
	Id    string // 公众号原始ID, 等于后台中的 公众号设置-->帐号详情-->原始ID
	Token string // 公众号的 Token, 和后台中的设置相等
	AppId string // 貌似需要认证才会有的???

	RWMutex       sync.RWMutex
	LastAESKey    [32]byte // 最后一个 AES Key
	CurrentAESKey [32]byte // 当前的 AES Key
}

// 初始化 DefaultAgent
//  如果不知道自己的 AppId 是多少, 可以先随便填入一个字符串,
//  这样正常情况下会出现 AppId mismatch 错误, 错误中 have 后面的就是正确的 AppId
func (this *DefaultAgent) Init(Id, Token, AppId string, AESKey []byte) {
	if len(AESKey) != 32 {
		panic("the length of AESKey must equal to 32")
	}
	this.Id = Id
	this.Token = Token
	this.AppId = AppId
	copy(this.CurrentAESKey[:], AESKey)
}

func (this *DefaultAgent) GetId() string {
	return this.Id
}
func (this *DefaultAgent) GetToken() string {
	return this.Token
}
func (this *DefaultAgent) GetAppId() string {
	return this.AppId
}
func (this *DefaultAgent) GetLastAESKey() (key [32]byte) {
	this.RWMutex.RLock()
	key = this.LastAESKey
	this.RWMutex.RUnlock()
	return
}
func (this *DefaultAgent) GetCurrentAESKey() (key [32]byte) {
	this.RWMutex.RLock()
	key = this.CurrentAESKey
	this.RWMutex.RUnlock()
	return
}
func (this *DefaultAgent) UpdateAESKey(AESKey [32]byte) {
	this.RWMutex.Lock()
	this.LastAESKey = this.CurrentAESKey
	this.CurrentAESKey = AESKey
	this.RWMutex.Unlock()
	return
}

// 明文模式 ======================================================================================================================================================================

func (this *DefaultAgent) ServeUnknownMsg(para *RequestParameters) {
}

func (this *DefaultAgent) ServeTextMsg(msg *request.Text, para *RequestParameters) {
}
func (this *DefaultAgent) ServeImageMsg(msg *request.Image, para *RequestParameters) {
}
func (this *DefaultAgent) ServeVoiceMsg(msg *request.Voice, para *RequestParameters) {
}
func (this *DefaultAgent) ServeVideoMsg(msg *request.Video, para *RequestParameters) {
}
func (this *DefaultAgent) ServeLocationMsg(msg *request.Location, para *RequestParameters) {
}
func (this *DefaultAgent) ServeLinkMsg(msg *request.Link, para *RequestParameters) {
}

func (this *DefaultAgent) ServeSubscribeEvent(msg *request.SubscribeEvent, para *RequestParameters) {
}
func (this *DefaultAgent) ServeUnsubscribeEvent(msg *request.UnsubscribeEvent, para *RequestParameters) {
}
func (this *DefaultAgent) ServeSubscribeByScanEvent(msg *request.SubscribeByScanEvent, para *RequestParameters) {
}
func (this *DefaultAgent) ServeScanEvent(msg *request.ScanEvent, para *RequestParameters) {
}
func (this *DefaultAgent) ServeLocationEvent(msg *request.LocationEvent, para *RequestParameters) {
}
func (this *DefaultAgent) ServeMenuClickEvent(msg *request.MenuClickEvent, para *RequestParameters) {
}
func (this *DefaultAgent) ServeMenuViewEvent(msg *request.MenuViewEvent, para *RequestParameters) {
}
func (this *DefaultAgent) ServeMenuScanCodePushEvent(msg *request.MenuScanCodePushEvent, para *RequestParameters) {
}
func (this *DefaultAgent) ServeMenuScanCodeWaitMsgEvent(msg *request.MenuScanCodeWaitMsgEvent, para *RequestParameters) {
}
func (this *DefaultAgent) ServeMenuPicSysPhotoEvent(msg *request.MenuPicSysPhotoEvent, para *RequestParameters) {
}
func (this *DefaultAgent) ServeMenuPicPhotoOrAlbumEvent(msg *request.MenuPicPhotoOrAlbumEvent, para *RequestParameters) {
}
func (this *DefaultAgent) ServeMenuPicWeixinEvent(msg *request.MenuPicWeixinEvent, para *RequestParameters) {
}
func (this *DefaultAgent) ServeMenuLocationSelectEvent(msg *request.MenuLocationSelectEvent, para *RequestParameters) {
}
func (this *DefaultAgent) ServeMassSendJobFinishEvent(msg *request.MassSendJobFinishEvent, para *RequestParameters) {
}
func (this *DefaultAgent) ServeTemplateSendJobFinishEvent(msg *request.TemplateSendJobFinishEvent, para *RequestParameters) {
}
func (this *DefaultAgent) ServeMerchantOrderEvent(msg *request.MerchantOrderEvent, para *RequestParameters) {
}

// 兼容模式, 安全模式 ==============================================================================================================================================================

func (this *DefaultAgent) ServeAESUnknownMsg(para *AESRequestParameters) {
}

func (this *DefaultAgent) ServeAESTextMsg(msg *request.Text, para *AESRequestParameters) {
}
func (this *DefaultAgent) ServeAESImageMsg(msg *request.Image, para *AESRequestParameters) {
}
func (this *DefaultAgent) ServeAESVoiceMsg(msg *request.Voice, para *AESRequestParameters) {
}
func (this *DefaultAgent) ServeAESVideoMsg(msg *request.Video, para *AESRequestParameters) {
}
func (this *DefaultAgent) ServeAESLocationMsg(msg *request.Location, para *AESRequestParameters) {
}
func (this *DefaultAgent) ServeAESLinkMsg(msg *request.Link, para *AESRequestParameters) {
}

func (this *DefaultAgent) ServeAESSubscribeEvent(msg *request.SubscribeEvent, para *AESRequestParameters) {
}
func (this *DefaultAgent) ServeAESUnsubscribeEvent(msg *request.UnsubscribeEvent, para *AESRequestParameters) {
}
func (this *DefaultAgent) ServeAESSubscribeByScanEvent(msg *request.SubscribeByScanEvent, para *AESRequestParameters) {
}
func (this *DefaultAgent) ServeAESScanEvent(msg *request.ScanEvent, para *AESRequestParameters) {
}
func (this *DefaultAgent) ServeAESLocationEvent(msg *request.LocationEvent, para *AESRequestParameters) {
}
func (this *DefaultAgent) ServeAESMenuClickEvent(msg *request.MenuClickEvent, para *AESRequestParameters) {
}
func (this *DefaultAgent) ServeAESMenuViewEvent(msg *request.MenuViewEvent, para *AESRequestParameters) {
}
func (this *DefaultAgent) ServeAESMenuScanCodePushEvent(msg *request.MenuScanCodePushEvent, para *AESRequestParameters) {
}
func (this *DefaultAgent) ServeAESMenuScanCodeWaitMsgEvent(msg *request.MenuScanCodeWaitMsgEvent, para *AESRequestParameters) {
}
func (this *DefaultAgent) ServeAESMenuPicSysPhotoEvent(msg *request.MenuPicSysPhotoEvent, para *AESRequestParameters) {
}
func (this *DefaultAgent) ServeAESMenuPicPhotoOrAlbumEvent(msg *request.MenuPicPhotoOrAlbumEvent, para *AESRequestParameters) {
}
func (this *DefaultAgent) ServeAESMenuPicWeixinEvent(msg *request.MenuPicWeixinEvent, para *AESRequestParameters) {
}
func (this *DefaultAgent) ServeAESMenuLocationSelectEvent(msg *request.MenuLocationSelectEvent, para *AESRequestParameters) {
}
func (this *DefaultAgent) ServeAESMassSendJobFinishEvent(msg *request.MassSendJobFinishEvent, para *AESRequestParameters) {
}
func (this *DefaultAgent) ServeAESTemplateSendJobFinishEvent(msg *request.TemplateSendJobFinishEvent, para *AESRequestParameters) {
}
func (this *DefaultAgent) ServeAESMerchantOrderEvent(msg *request.MerchantOrderEvent, para *AESRequestParameters) {
}
