// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package component

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("github.com/chanxuehong/wechat/mp/component: item not found")

// component_verify_ticket 獲取接口
type ComponentVerifyTicketGetter interface {
	// 根据 component_appid 获取第三方平台当前的 component_verify_ticket, 如果没有找到返回 ErrNotFound
	GetComponentVerifyTicket(componentAppId string) (componentVerifyTicket string, err error)
}

var _ ComponentVerifyTicketGetter = (*ComponentVerifyTicketCache)(nil)

type ComponentVerifyTicketCache struct {
	rwmutex               sync.RWMutex
	componentVerifyTicket string
}

func (cache *ComponentVerifyTicketCache) SetComponentVerifyTicket(componentAppId, componentVerifyTicket string) (err error) {
	//if componentAppId == "" {
	//	return errors.New("empty componentAppId")
	//}
	if componentVerifyTicket == "" {
		return errors.New("empty componentVerifyTicket")
	}

	cache.rwmutex.Lock()
	cache.componentVerifyTicket = componentVerifyTicket
	cache.rwmutex.Unlock()
	return
}

func (cache *ComponentVerifyTicketCache) GetComponentVerifyTicket(componentAppId string) (componentVerifyTicket string, err error) {
	cache.rwmutex.RLock()
	componentVerifyTicket = cache.componentVerifyTicket
	if componentVerifyTicket == "" {
		err = ErrNotFound
	}
	cache.rwmutex.RUnlock()
	return
}

var _ ComponentVerifyTicketGetter = (*ComponentVerifyTicketCache2)(nil)

type ComponentVerifyTicketCache2 struct {
	rwmutex sync.RWMutex
	m       map[string]string
}

func NewComponentVerifyTicketCache2() *ComponentVerifyTicketCache2 {
	return &ComponentVerifyTicketCache2{
		m: make(map[string]string),
	}
}

func (cache *ComponentVerifyTicketCache2) SetComponentVerifyTicket(componentAppId, componentVerifyTicket string) (err error) {
	if componentAppId == "" {
		return errors.New("empty componentAppId")
	}
	if componentVerifyTicket == "" {
		return errors.New("empty componentVerifyTicket")
	}

	cache.rwmutex.Lock()
	cache.m[componentAppId] = componentVerifyTicket
	cache.rwmutex.Unlock()
	return
}

func (cache *ComponentVerifyTicketCache2) GetComponentVerifyTicket(componentAppId string) (componentVerifyTicket string, err error) {
	cache.rwmutex.RLock()
	componentVerifyTicket = cache.m[componentAppId]
	if componentVerifyTicket == "" {
		err = ErrNotFound
	}
	cache.rwmutex.RUnlock()
	return
}
