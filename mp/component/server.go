// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package component

import (
	"errors"
	"sync"
)

// 公众号第三方平台服务器接口
type Server interface {
	AppId() string // 获取第三方平台AppId
	Token() string // 获取第三方平台的Token

	CurrentAESKey() [32]byte                // 获取当前有效的 AES 加密 Key
	LastAESKey() (key [32]byte, valid bool) // 获取上一个有效的 AES 加密 Key

	MessageHandler() MessageHandler // 获取 MessageHandler

	RequestSizeLimit() int64 // 消息請求的 http body 大小限制, 如果 <= 0 則不做限制
}

var _ Server = (*DefaultServer)(nil)

type DefaultServer struct {
	appId string
	token string

	rwmutex           sync.RWMutex
	currentAESKey     [32]byte // 当前的 AES Key
	lastAESKey        [32]byte // 最后一个 AES Key
	isLastAESKeyValid bool     // lastAESKey 是否有效, 如果 lastAESKey 是 zero 则无效

	messageHandler MessageHandler

	requestSizeLimit int64
}

func NewDefaultServer(appId, token string, AESKey []byte, handler MessageHandler, requestSizeLimit int64) (srv *DefaultServer) {
	if len(AESKey) != 32 {
		panic("the length of AESKey must equal to 32")
	}
	if handler == nil {
		panic("nil MessageHandler")
	}

	srv = &DefaultServer{
		appId:            appId,
		token:            token,
		messageHandler:   handler,
		requestSizeLimit: requestSizeLimit,
	}
	copy(srv.currentAESKey[:], AESKey)
	return
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
func (srv *DefaultServer) RequestSizeLimit() int64 {
	return srv.requestSizeLimit
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
