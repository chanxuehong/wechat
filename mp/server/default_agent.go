// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"github.com/chanxuehong/wechat/mp/message/passive/request"
	"net/http"
	"sync"
)

var _ Agent = new(DefaultAgent)

type DefaultAgent struct {
	RWMutex       sync.RWMutex
	Id            string   // 公众号原始ID, 等于后台中的 公众号设置-->帐号详情-->原始ID
	Token         string   // 公众号的 Token, 和后台中的设置相等
	AppId         string   // 貌似需要认证才会有的???
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

// 兼容模式, 安全模式 ==============================================================================================================================================================

func (this *DefaultAgent) ServeAESUnknownMsg(w http.ResponseWriter, r *http.Request, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) {
}

func (this *DefaultAgent) ServeAESTextMsg(w http.ResponseWriter, r *http.Request, msg *request.Text, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) {
}
func (this *DefaultAgent) ServeAESImageMsg(w http.ResponseWriter, r *http.Request, msg *request.Image, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) {
}
func (this *DefaultAgent) ServeAESVoiceMsg(w http.ResponseWriter, r *http.Request, msg *request.Voice, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) {
}
func (this *DefaultAgent) ServeAESVideoMsg(w http.ResponseWriter, r *http.Request, msg *request.Video, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) {
}
func (this *DefaultAgent) ServeAESLocationMsg(w http.ResponseWriter, r *http.Request, msg *request.Location, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) {
}
func (this *DefaultAgent) ServeAESLinkMsg(w http.ResponseWriter, r *http.Request, msg *request.Link, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) {
}

func (this *DefaultAgent) ServeAESSubscribeEvent(w http.ResponseWriter, r *http.Request, msg *request.SubscribeEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) {
}
func (this *DefaultAgent) ServeAESUnsubscribeEvent(w http.ResponseWriter, r *http.Request, msg *request.UnsubscribeEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) {
}
func (this *DefaultAgent) ServeAESSubscribeByScanEvent(w http.ResponseWriter, r *http.Request, msg *request.SubscribeByScanEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) {
}
func (this *DefaultAgent) ServeAESScanEvent(w http.ResponseWriter, r *http.Request, msg *request.ScanEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) {
}
func (this *DefaultAgent) ServeAESLocationEvent(w http.ResponseWriter, r *http.Request, msg *request.LocationEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) {
}
func (this *DefaultAgent) ServeAESMenuClickEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuClickEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) {
}
func (this *DefaultAgent) ServeAESMenuViewEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuViewEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) {
}
func (this *DefaultAgent) ServeAESMenuScanCodePushEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuScanCodePushEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) {
}
func (this *DefaultAgent) ServeAESMenuScanCodeWaitMsgEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuScanCodeWaitMsgEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) {
}
func (this *DefaultAgent) ServeAESMenuPicSysPhotoEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuPicSysPhotoEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) {
}
func (this *DefaultAgent) ServeAESMenuPicPhotoOrAlbumEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuPicPhotoOrAlbumEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) {
}
func (this *DefaultAgent) ServeAESMenuPicWeixinEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuPicWeixinEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) {
}
func (this *DefaultAgent) ServeAESMenuLocationSelectEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuLocationSelectEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) {
}
func (this *DefaultAgent) ServeAESMassSendJobFinishEvent(w http.ResponseWriter, r *http.Request, msg *request.MassSendJobFinishEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) {
}
func (this *DefaultAgent) ServeAESTemplateSendJobFinishEvent(w http.ResponseWriter, r *http.Request, msg *request.TemplateSendJobFinishEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) {
}
func (this *DefaultAgent) ServeAESMerchantOrderEvent(w http.ResponseWriter, r *http.Request, msg *request.MerchantOrderEvent, rawXMLMsg []byte, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) {
}
