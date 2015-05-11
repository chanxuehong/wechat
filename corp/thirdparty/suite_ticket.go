// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package thirdparty

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("github.com/chanxuehong/wechat/corp/thirdparty: item not found")

type SuiteTicketGetter interface {
	// 根据 suiteId 获取套件当前的 suiteTicket, 如果没有找到返回 ErrNotFound
	GetSuiteTicket(suiteId string) (suiteTicket string, err error)
}

var _ SuiteTicketGetter = (*SuiteTicketCache)(nil)

type SuiteTicketCache struct {
	rwmutex     sync.RWMutex
	suiteTicket string
}

func (cache *SuiteTicketCache) SetSuiteTicket(suiteId, suiteTicket string) (err error) {
	//if suiteId == "" {
	//	return errors.New("empty suiteId")
	//}
	if suiteTicket == "" {
		return errors.New("empty suiteTicket")
	}

	cache.rwmutex.Lock()
	cache.suiteTicket = suiteTicket
	cache.rwmutex.Unlock()
	return
}

func (cache *SuiteTicketCache) GetSuiteTicket(suiteId string) (suiteTicket string, err error) {
	cache.rwmutex.RLock()
	suiteTicket = cache.suiteTicket
	if suiteTicket == "" {
		err = ErrNotFound
	}
	cache.rwmutex.RUnlock()
	return
}

var _ SuiteTicketGetter = (*SuiteTicketCache2)(nil)

type SuiteTicketCache2 struct {
	rwmutex sync.RWMutex
	m       map[string]string
}

func NewSuiteTicketCache2() *SuiteTicketCache2 {
	return &SuiteTicketCache2{
		m: make(map[string]string),
	}
}

func (cache *SuiteTicketCache2) SetSuiteTicket(suiteId, suiteTicket string) (err error) {
	if suiteId == "" {
		return errors.New("empty suiteId")
	}
	if suiteTicket == "" {
		return errors.New("empty suiteTicket")
	}

	cache.rwmutex.Lock()
	cache.m[suiteId] = suiteTicket
	cache.rwmutex.Unlock()
	return
}

func (cache *SuiteTicketCache2) GetSuiteTicket(suiteId string) (suiteTicket string, err error) {
	cache.rwmutex.RLock()
	suiteTicket = cache.m[suiteId]
	if suiteTicket == "" {
		err = ErrNotFound
	}
	cache.rwmutex.RUnlock()
	return
}
