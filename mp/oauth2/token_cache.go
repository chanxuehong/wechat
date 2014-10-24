// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package oauth2

import (
	"errors"
	"sync"
)

var ErrCacheMiss = errors.New("oauth2: cache miss")

type TokenCache interface {
	// 从缓存中读取 OAuth2Token, 如果没有找到则返回错误 ErrCacheMiss
	// 如果成功 tok != nil && err == nil, 否则 tok == nil && err != nil
	Token() (tok *OAuth2Token, err error)

	// 把 OAuth2Token 存入到缓存, 如果原来有 OAuth2Token 则覆盖原来的
	PutToken(tok *OAuth2Token) (err error)
}

var _ TokenCache = new(DefaultTokenCache)

// 一个简单的 TokenCache 的实现
type DefaultTokenCache struct {
	rwmutex sync.RWMutex
	token   *OAuth2Token
}

func (this *DefaultTokenCache) Token() (tok *OAuth2Token, err error) {
	this.rwmutex.RLock()
	if this.token == nil {
		err = ErrCacheMiss
	} else {
		tok = new(OAuth2Token)
		*tok = *this.token
	}
	this.rwmutex.RUnlock()
	return
}

func (this *DefaultTokenCache) PutToken(tok *OAuth2Token) (err error) {
	if tok == nil {
		return errors.New("input OAuth2Token is nil")
	}

	this.rwmutex.Lock()
	if this.token == nil {
		this.token = new(OAuth2Token)
	}
	*this.token = *tok
	this.rwmutex.Unlock()
	return
}
