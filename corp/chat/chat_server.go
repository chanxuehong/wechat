// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package chat

import (
	"bytes"
	"errors"
	"sync"
)

type ChatServer interface {
	CorpId() string // 企业号Id, 用于约束消息的 CorpId, 如果为空表示不约束
	Token() string  // 获取应用的Token

	CurrentAESKey() [32]byte                // 获取当前有效的 AES 加密 Key
	LastAESKey() (key [32]byte, valid bool) // 获取上一个有效的 AES 加密 Key

	MessageHandler() MessageHandler // 获取 MessageHandler
}

var _ ChatServer = (*DefaultChatServer)(nil)

type DefaultChatServer struct {
	corpId            string
	token             string

	rwmutex           sync.RWMutex
	currentAESKey     [32]byte // 当前的 AES Key
	lastAESKey        [32]byte // 最后一个 AES Key
	isLastAESKeyValid bool     // lastAESKey 是否有效, 如果 lastAESKey 是 zero 则无效

	messageHandler    MessageHandler
}

// NewDefaultChatServer 创建一个新的 DefaultChatServer.
func NewDefaultChatServer(corpId string, token string, aesKey []byte, handler MessageHandler) (srv *DefaultChatServer) {
	if len(aesKey) != 32 {
		panic("the length of aesKey must equal to 32")
	}
	if handler == nil {
		panic("nil MessageHandler")
	}

	srv = &DefaultChatServer{
		corpId:         corpId,
		token:          token,
		messageHandler: handler,
	}
	copy(srv.currentAESKey[:], aesKey)
	return
}

func (srv *DefaultChatServer) CorpId() string {
	return srv.corpId
}


func (srv *DefaultChatServer) Token() string {
	return srv.token
}
func (srv *DefaultChatServer) MessageHandler() MessageHandler {
	return srv.messageHandler
}
func (srv *DefaultChatServer) CurrentAESKey() (key [32]byte) {
	srv.rwmutex.RLock()
	key = srv.currentAESKey
	srv.rwmutex.RUnlock()
	return
}
func (srv *DefaultChatServer) LastAESKey() (key [32]byte, valid bool) {
	srv.rwmutex.RLock()
	key = srv.lastAESKey
	valid = srv.isLastAESKeyValid
	srv.rwmutex.RUnlock()
	return
}

func (srv *DefaultChatServer) UpdateAESKey(aesKey []byte) (err error) {
	if len(aesKey) != 32 {
		return errors.New("the length of aesKey must equal to 32")
	}

	srv.rwmutex.Lock()
	defer srv.rwmutex.Unlock()

	if bytes.Equal(aesKey, srv.currentAESKey[:]) {
		return
	}

	srv.isLastAESKeyValid = true
	srv.lastAESKey = srv.currentAESKey
	copy(srv.currentAESKey[:], aesKey)
	return
}
