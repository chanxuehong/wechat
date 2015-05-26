// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package mp

import (
	"errors"
	"sync"
)

// 公众号服务端接口, 处理单个公众号的消息(事件)请求.
type Server interface {
	OriId() string // 获取公众号的 原始ID, 如果爲空則不檢查消息的 ToUserName
	AppId() string // 获取公众号的 AppId
	Token() string // 获取公众号的 Token

	CurrentAESKey() [32]byte                // 获取当前有效的 AES 加密 Key
	LastAESKey() (key [32]byte, valid bool) // 获取上一个有效的 AES 加密 Key

	MessageHandler() MessageHandler // 获取 MessageHandler
}

var _ Server = (*DefaultServer)(nil)

type DefaultServer struct {
	oriId string
	appId string
	token string

	rwmutex           sync.RWMutex
	currentAESKey     [32]byte // 当前的 AES Key
	lastAESKey        [32]byte // 最后一个 AES Key
	isLastAESKeyValid bool     // lastAESKey 是否有效, 如果 lastAESKey 是 zero 则无效

	messageHandler MessageHandler
}

func NewDefaultServer(oriId, token, appId string, AESKey []byte, handler MessageHandler) (srv *DefaultServer) {
	if len(AESKey) != 32 {
		panic("the length of AESKey must equal to 32")
	}
	if handler == nil {
		panic("nil MessageHandler")
	}

	srv = &DefaultServer{
		oriId:          oriId,
		appId:          appId,
		token:          token,
		messageHandler: handler,
	}
	copy(srv.currentAESKey[:], AESKey)
	return
}

func (srv *DefaultServer) OriId() string {
	return srv.oriId
}
func (srv *DefaultServer) AppId() string {
	return srv.appId
}
func (srv *DefaultServer) Token() string {
	return srv.token
}
func (srv *DefaultServer) MessageHandler() MessageHandler {
	return srv.messageHandler
}
func (srv *DefaultServer) CurrentAESKey() (key [32]byte) {
	srv.rwmutex.RLock()
	key = srv.currentAESKey
	srv.rwmutex.RUnlock()
	return
}
func (srv *DefaultServer) LastAESKey() (key [32]byte, valid bool) {
	srv.rwmutex.RLock()
	key = srv.lastAESKey
	valid = srv.isLastAESKeyValid
	srv.rwmutex.RUnlock()
	return
}

// 更新當前的 aesKey
func (srv *DefaultServer) UpdateAESKey(aesKey []byte) (err error) {
	if len(aesKey) != 32 {
		return errors.New("the length of aesKey must equal to 32")
	}

	srv.rwmutex.Lock()
	srv.isLastAESKeyValid = true
	srv.lastAESKey = srv.currentAESKey
	copy(srv.currentAESKey[:], aesKey)
	srv.rwmutex.Unlock()
	return
}
