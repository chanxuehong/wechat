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
	ComponentAppId() string // 获取第三方平台AppId
	ComponentToken() string // 获取第三方平台的Token

	CurrentAESKey() [32]byte // 获取当前有效的 AES 加密 Key
	LastAESKey() [32]byte    // 获取最后一个有效的 AES 加密 Key

	ComponentMessageHandler() ComponentMessageHandler // 获取 ComponentMessageHandler
}

var _ ComponentServer = (*DefaultComponentServer)(nil)

type DefaultComponentServer struct {
	componentAppId string
	componentToken string

	rwmutex           sync.RWMutex
	currentAESKey     [32]byte // 当前的 AES Key
	lastAESKey        [32]byte // 最后一个 AES Key
	isLastAESKeyValid bool     // lastAESKey 是否有效, 如果 lastAESKey 是 zero 则无效

	componentMessageHandler ComponentMessageHandler
}

// NewDefaultComponentServer 创建一个新的 DefaultComponentServer.
func NewDefaultComponentServer(componentAppId, componentToken string, AESKey []byte,
	componentMessageHandler ComponentMessageHandler) (srv *DefaultComponentServer) {

	if len(AESKey) != 32 {
		panic("the length of AESKey must equal to 32")
	}
	if componentMessageHandler == nil {
		panic("nil ComponentMessageHandler")
	}

	srv = &DefaultComponentServer{
		componentAppId:          componentAppId,
		componentToken:          componentToken,
		componentMessageHandler: componentMessageHandler,
	}
	copy(srv.currentAESKey[:], AESKey)
	return
}

func (srv *DefaultComponentServer) ComponentAppId() string {
	return srv.componentAppId
}
func (srv *DefaultComponentServer) ComponentToken() string {
	return srv.componentToken
}
func (srv *DefaultComponentServer) ComponentMessageHandler() ComponentMessageHandler {
	return srv.componentMessageHandler
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
