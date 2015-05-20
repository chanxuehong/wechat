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

type ComponentVerifyTicket string

// component_verify_ticket 獲取接口
type VerifyTicketGetter interface {
	// 根据 component_appid 获取第三方平台当前的 component_verify_ticket, 如果没有找到返回 ErrNotFound
	GetVerifyTicket(appId string) (ticket ComponentVerifyTicket, err error)
}

var _ VerifyTicketGetter = (*VerifyTicketCache)(nil)

type VerifyTicketCache struct {
	rwmutex sync.RWMutex
	ticket  ComponentVerifyTicket
}

func (cache *VerifyTicketCache) SetVerifyTicket(appId string, ticket ComponentVerifyTicket) (err error) {
	//if appId == "" {
	//	return errors.New("empty ComponentAppId")
	//}
	if ticket == "" {
		return errors.New("empty ComponentVerifyTicket")
	}

	cache.rwmutex.Lock()
	cache.ticket = ticket
	cache.rwmutex.Unlock()
	return
}

func (cache *VerifyTicketCache) GetVerifyTicket(appId string) (ticket ComponentVerifyTicket, err error) {
	cache.rwmutex.RLock()
	ticket = cache.ticket
	if ticket == "" {
		err = ErrNotFound
	}
	cache.rwmutex.RUnlock()
	return
}

var _ VerifyTicketGetter = (*VerifyTicketCache2)(nil)

type VerifyTicketCache2 struct {
	rwmutex sync.RWMutex
	m       map[string]ComponentVerifyTicket
}

func NewVerifyTicketCache2() *VerifyTicketCache2 {
	return &VerifyTicketCache2{
		m: make(map[string]ComponentVerifyTicket),
	}
}

func (cache *VerifyTicketCache2) SetVerifyTicket(appId string, ticket ComponentVerifyTicket) (err error) {
	if appId == "" {
		return errors.New("empty ComponentAppId")
	}
	if ticket == "" {
		return errors.New("empty ComponentVerifyTicket")
	}

	cache.rwmutex.Lock()
	cache.m[appId] = ticket
	cache.rwmutex.Unlock()
	return
}

func (cache *VerifyTicketCache2) GetVerifyTicket(appId string) (ticket ComponentVerifyTicket, err error) {
	cache.rwmutex.RLock()
	ticket = cache.m[appId]
	if ticket == "" {
		err = ErrNotFound
	}
	cache.rwmutex.RUnlock()
	return
}
