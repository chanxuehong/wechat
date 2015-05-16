// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package component

import (
	"errors"
	"sync"
)

type ComponentServer interface {
	AppId() string // 获取第三方平台AppId
	Token() string // 获取第三方平台的Token

	CurrentAESKey() [32]byte // 获取当前有效的 AES 加密 Key
	LastAESKey() [32]byte    // 获取最后一个有效的 AES 加密 Key

	MessageHandler() ComponentMessageHandler // 获取 ComponentMessageHandler
}

var _ ComponentServer = (*DefaultComponentServer)(nil)

type DefaultComponentServer struct {
	appId string
	token string

	rwmutex           sync.RWMutex
	currentAESKey     [32]byte // 当前的 AES Key
	lastAESKey        [32]byte // 最后一个 AES Key
	isLastAESKeyValid bool     // lastAESKey 是否有效, 如果 lastAESKey 是 zero 则无效

	messageHandler ComponentMessageHandler
}

// NewDefaultComponentServer 创建一个新的 DefaultComponentServer.
func NewDefaultComponentServer(componentAppId, componentToken string, AESKey []byte, messageHandler ComponentMessageHandler) (srv *DefaultComponentServer) {
	if len(AESKey) != 32 {
		panic("the length of AESKey must equal to 32")
	}
	if messageHandler == nil {
		panic("nil ComponentMessageHandler")
	}

	srv = &DefaultComponentServer{
		appId:          componentAppId,
		token:          componentToken,
		messageHandler: messageHandler,
	}
	copy(srv.currentAESKey[:], AESKey)
	return
}

func (srv *DefaultComponentServer) AppId() string {
	return srv.appId
}
func (srv *DefaultComponentServer) Token() string {
	return srv.token
}
func (srv *DefaultComponentServer) MessageHandler() ComponentMessageHandler {
	return srv.messageHandler
}
func (srv *DefaultComponentServer) CurrentAESKey() (key [32]byte) {
	srv.rwmutex.RLock()
	key = srv.currentAESKey
	srv.rwmutex.RUnlock()
	return
}
func (srv *DefaultComponentServer) LastAESKey() (key [32]byte) {
	srv.rwmutex.RLock()
	if srv.isLastAESKeyValid {
		key = srv.lastAESKey
	} else {
		key = srv.currentAESKey
	}
	srv.rwmutex.RUnlock()
	return
}
func (srv *DefaultComponentServer) UpdateAESKey(AESKey []byte) (err error) {
	if len(AESKey) != 32 {
		return errors.New("the length of AESKey must equal to 32")
	}

	srv.rwmutex.Lock()
	srv.isLastAESKeyValid = true
	srv.lastAESKey = srv.currentAESKey
	copy(srv.currentAESKey[:], AESKey)
	srv.rwmutex.Unlock()
	return
}
