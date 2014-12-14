// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"io"
	"sync"
)

var _ Agent = new(DefaultAgent)

type DefaultAgent struct {
	SuiteId string
	Token   string

	RWMutex       sync.RWMutex
	LastAESKey    [32]byte // 最后一个 AES Key
	CurrentAESKey [32]byte // 当前的 AES Key
}

func (this *DefaultAgent) Init(SuiteId, Token string, AESKey []byte) {
	if len(AESKey) != 32 {
		panic("the length of AESKey must equal to 32")
	}
	this.SuiteId = SuiteId
	this.Token = Token
	copy(this.CurrentAESKey[:], AESKey)
}

func (this *DefaultAgent) GetSuiteId() string {
	return this.SuiteId
}
func (this *DefaultAgent) GetToken() string {
	return this.Token
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

func (this *DefaultAgent) ServeUnknownMsg(para *InputParameters) {
	io.WriteString(para.w, ResponseSuccess)
}

func (this *DefaultAgent) ServeSuiteTicketMsg(para *InputParameters, msg *SuiteTicket) {
	io.WriteString(para.w, ResponseSuccess)
}
func (this *DefaultAgent) ServeChangeAuthMsg(para *InputParameters, msg *ChangeAuth) {
	io.WriteString(para.w, ResponseSuccess)
}
func (this *DefaultAgent) ServeCancelAuthMsg(para *InputParameters, msg *CancelAuth) {
	io.WriteString(para.w, ResponseSuccess)
}
