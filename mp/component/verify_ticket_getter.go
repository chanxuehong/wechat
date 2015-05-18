// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package component

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("item not found")

// component_verify_ticket 獲取接口
type VerifyTicketGetter interface {
	// 根据 component_appid 获取第三方平台当前的 component_verify_ticket, 如果没有找到返回 ErrNotFound
	GetVerifyTicket(appId string) (verifyTicket string, err error)
}

var _ VerifyTicketGetter = (*VerifyTicketCache)(nil)

type VerifyTicketCache struct {
	rwmutex      sync.RWMutex
	verifyTicket string
}

func (cache *VerifyTicketCache) SetVerifyTicket(appId, verifyTicket string) (err error) {
	//if appId == "" {
	//	return errors.New("empty appId")
	//}
	if verifyTicket == "" {
		return errors.New("empty verifyTicket")
	}

	cache.rwmutex.Lock()
	cache.verifyTicket = verifyTicket
	cache.rwmutex.Unlock()
	return
}

func (cache *VerifyTicketCache) GetVerifyTicket(appId string) (verifyTicket string, err error) {
	cache.rwmutex.RLock()
	verifyTicket = cache.verifyTicket
	if verifyTicket == "" {
		err = ErrNotFound
	}
	cache.rwmutex.RUnlock()
	return
}

var _ VerifyTicketGetter = (*VerifyTicketCache2)(nil)

type VerifyTicketCache2 struct {
	rwmutex sync.RWMutex
	m       map[string]string
}

func NewVerifyTicketCache2() *VerifyTicketCache2 {
	return &VerifyTicketCache2{
		m: make(map[string]string),
	}
}

func (cache *VerifyTicketCache2) SetVerifyTicket(appId, verifyTicket string) (err error) {
	if appId == "" {
		return errors.New("empty appId")
	}
	if verifyTicket == "" {
		return errors.New("empty verifyTicket")
	}

	cache.rwmutex.Lock()
	cache.m[appId] = verifyTicket
	cache.rwmutex.Unlock()
	return
}

func (cache *VerifyTicketCache2) GetVerifyTicket(appId string) (verifyTicket string, err error) {
	cache.rwmutex.RLock()
	verifyTicket = cache.m[appId]
	if verifyTicket == "" {
		err = ErrNotFound
	}
	cache.rwmutex.RUnlock()
	return
}
