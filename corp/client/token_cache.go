// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"errors"
	"sync"
)

var ErrCacheMiss = errors.New("token cache miss")

type TokenCache interface {
	// 从 cache 里获取 token
	//  NOTE: 如果没有找到返回 ErrCacheMiss
	Token() (token string, err error)

	// 添加或者重置 token
	PutToken(token string) (err error)
}

var _ TokenCache = new(DefaultTokenCache)

// 一个简单的 TokenCache 的实现
type DefaultTokenCache struct {
	rwmutex sync.RWMutex
	token   string
}

func (this *DefaultTokenCache) Token() (token string, err error) {
	this.rwmutex.RLock()
	if len(this.token) == 0 {
		err = ErrCacheMiss
	} else {
		token = this.token
	}
	this.rwmutex.RUnlock()
	return
}

func (this *DefaultTokenCache) PutToken(token string) (err error) {
	if len(token) == 0 {
		return errors.New("token is empty")
	}
	this.rwmutex.Lock()
	this.token = token
	this.rwmutex.Unlock()
	return
}
