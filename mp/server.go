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
type WechatServer interface {
	WechatId() string // 获取公众号的原始ID, 等于后台中的 公众号设置-->帐号详情-->原始ID
	Token() string    // 获取公众号的Token, 和后台中的设置相等
	AppId() string    // 获取公众号的 AppId

	CurrentAESKey() [32]byte // 获取当前有效的 AES 加密 Key
	LastAESKey() [32]byte    // 获取最后一个有效的 AES 加密 Key

	MessageHandler() MessageHandler // 获取 MessageHandler
}

var _ WechatServer = new(DefaultWechatServer)

type DefaultWechatServer struct {
	wechatId string
	token    string
	appId    string

	rwmutex           sync.RWMutex
	currentAESKey     [32]byte // 当前的 AES Key
	lastAESKey        [32]byte // 最后一个 AES Key
	isLastAESKeyValid bool     // lastAESKey 是否有效, 如果 lastAESKey 是 zero 则无效

	messageHandler MessageHandler
}

// NewDefaultWechatServer 创建一个新的 DefaultWechatServer.
//  如果不知道自己的 AppId 是多少, 可以先随便填入一个字符串,
//  这样正常情况下会出现 AppId mismatch 错误, 错误的 have 后面的就是正确的 AppId.
func NewDefaultWechatServer(wechatId, token, appId string, AESKey []byte,
	messageHandler MessageHandler) (srv *DefaultWechatServer) {

	if len(AESKey) != 32 {
		panic("mp: the length of AESKey must equal to 32")
	}
	if messageHandler == nil {
		panic("mp: nil messageHandler")
	}

	srv = &DefaultWechatServer{
		wechatId:       wechatId,
		token:          token,
		appId:          appId,
		messageHandler: messageHandler,
	}
	copy(srv.currentAESKey[:], AESKey)
	return
}

func (srv *DefaultWechatServer) WechatId() string {
	return srv.wechatId
}
func (srv *DefaultWechatServer) Token() string {
	return srv.token
}
func (srv *DefaultWechatServer) AppId() string {
	return srv.appId
}
func (srv *DefaultWechatServer) MessageHandler() MessageHandler {
	return srv.messageHandler
}
func (srv *DefaultWechatServer) CurrentAESKey() (key [32]byte) {
	srv.rwmutex.RLock()
	key = srv.currentAESKey
	srv.rwmutex.RUnlock()
	return
}
func (srv *DefaultWechatServer) LastAESKey() (key [32]byte) {
	srv.rwmutex.RLock()
	if srv.isLastAESKeyValid {
		key = srv.lastAESKey
	} else {
		key = srv.currentAESKey
	}
	srv.rwmutex.RUnlock()
	return
}
func (srv *DefaultWechatServer) UpdateAESKey(AESKey []byte) (err error) {
	if len(AESKey) != 32 {
		return errors.New("the length of AESKey must equal to 32")
	}

	srv.rwmutex.Lock()
	srv.lastAESKey = srv.currentAESKey
	srv.isLastAESKeyValid = true
	copy(srv.currentAESKey[:], AESKey)
	srv.rwmutex.Unlock()
	return
}
