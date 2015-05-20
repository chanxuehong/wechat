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
	// 根据 SuiteId 获取套件当前的 SuiteTicket, 如果没有找到返回 ErrNotFound
	GetSuiteTicket(SuiteId string) (ticket string, err error)

	// 沒有實際意義, 接口標識而已
	TagEF8503CCFE9811E4959AA4DB30FED8E1()
}

var _ SuiteTicketGetter = (*SuiteTicketCache)(nil)

type SuiteTicketCache struct {
	rwmutex sync.RWMutex
	ticket  string
}

func (cache *SuiteTicketCache) TagEF8503CCFE9811E4959AA4DB30FED8E1() {}

func (cache *SuiteTicketCache) SetSuiteTicket(suiteId string, ticket string) (err error) {
	//if suiteId == "" {
	//	return errors.New("empty suiteId")
	//}
	if ticket == "" {
		return errors.New("empty SuiteTicket")
	}

	cache.rwmutex.Lock()
	cache.ticket = ticket
	cache.rwmutex.Unlock()
	return
}

func (cache *SuiteTicketCache) GetSuiteTicket(suiteId string) (ticket string, err error) {
	cache.rwmutex.RLock()
	ticket = cache.ticket
	if ticket == "" {
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

func (cache *SuiteTicketCache2) TagEF8503CCFE9811E4959AA4DB30FED8E1() {}

func (cache *SuiteTicketCache2) SetSuiteTicket(suiteId string, ticket string) (err error) {
	if suiteId == "" {
		return errors.New("empty suiteId")
	}
	if ticket == "" {
		return errors.New("empty SuiteTicket")
	}

	cache.rwmutex.Lock()
	cache.m[suiteId] = ticket
	cache.rwmutex.Unlock()
	return
}

func (cache *SuiteTicketCache2) GetSuiteTicket(suiteId string) (ticket string, err error) {
	cache.rwmutex.RLock()
	ticket = cache.m[suiteId]
	if ticket == "" {
		err = ErrNotFound
	}
	cache.rwmutex.RUnlock()
	return
}
