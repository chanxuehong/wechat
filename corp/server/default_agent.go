// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import "github.com/chanxuehong/wechat/corp/message/passive/request"

var _ Agent = new(DefaultAgent)

type DefaultAgent struct {
	CorpId  string
	AgentId int64
	Token   string
	AESKey  [32]byte
}

func (this *DefaultAgent) Init(CorpId string, AgentId int64, Token string, AESKey []byte) {
	if len(AESKey) != 32 {
		panic("the length of AESKey must equal to 32")
	}
	this.CorpId = CorpId
	this.AgentId = AgentId
	this.Token = Token
	copy(this.AESKey[:], AESKey)
}

func (this *DefaultAgent) GetCorpId() string {
	return this.CorpId
}
func (this *DefaultAgent) GetAgentId() int64 {
	return this.AgentId
}
func (this *DefaultAgent) GetToken() string {
	return this.Token
}
func (this *DefaultAgent) GetAESKey() [32]byte {
	return this.AESKey
}

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

func (this *DefaultAgent) ServeSubscribeEvent(event *request.SubscribeEvent, para *RequestParameters) {
}
func (this *DefaultAgent) ServeUnsubscribeEvent(event *request.UnsubscribeEvent, para *RequestParameters) {
}
func (this *DefaultAgent) ServeLocationEvent(event *request.LocationEvent, para *RequestParameters) {
}
func (this *DefaultAgent) ServeMenuClickEvent(event *request.MenuClickEvent, para *RequestParameters) {
}
func (this *DefaultAgent) ServeMenuViewEvent(event *request.MenuViewEvent, para *RequestParameters) {
}
func (this *DefaultAgent) ServeMenuScanCodePushEvent(event *request.MenuScanCodePushEvent, para *RequestParameters) {
}
func (this *DefaultAgent) ServeMenuScanCodeWaitMsgEvent(event *request.MenuScanCodeWaitMsgEvent, para *RequestParameters) {
}
func (this *DefaultAgent) ServeMenuPicSysPhotoEvent(event *request.MenuPicSysPhotoEvent, para *RequestParameters) {
}
func (this *DefaultAgent) ServeMenuPicPhotoOrAlbumEvent(event *request.MenuPicPhotoOrAlbumEvent, para *RequestParameters) {
}
func (this *DefaultAgent) ServeMenuPicWeixinEvent(event *request.MenuPicWeixinEvent, para *RequestParameters) {
}
func (this *DefaultAgent) ServeMenuLocationSelectEvent(event *request.MenuLocationSelectEvent, para *RequestParameters) {
}
