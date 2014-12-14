// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package wechat

import "errors"

var ErrCacheMiss = errors.New("cache miss")

// access token 缓存接口, 一般用于分布式获取 access token 场景, see token_cache.png
type TokenCache interface {
	// 从 cache 里获取 token
	// NOTE: 如果没有找到返回 ErrCacheMiss!!!
	Token() (token string, err error)

	// 添加或者重置 token
	PutToken(token string) (err error)
}

// 从微信服务器获取新的 access token 接口.
// 一般用于分布式获取 access token 场景, see token_cache.png
type TokenGetter interface {
	GetNewToken() (token string, err error)
}
