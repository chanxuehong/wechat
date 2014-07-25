// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"errors"
)

// access token 存储在分布式缓存中的基本单元
type Token struct {
	Value     string // access token 的值
	ExpiresAt int64  // 过期时间戳; 因为是绝对时间, 所以要求服务器之间时间要同步!
}

// TokenCache.Token() 获取 access token 不存在时必须返回这个错误
var ErrTokenCacheMiss = errors.New("token cache miss")

// 对于分布式应用, access token 可以放在一个分布式缓存中; TokenCache 就是缓存的接口.
//   NOTE: 要求分布式服务器之间 !!!时间同步!!!
type TokenCache interface {
	// 从服务器获取 access token, 如果找不到则返回 ErrTokenCacheMiss
	Token() (tk Token, err error)

	// 设置 access token 到服务器
	SetToken(tk Token) (err error)
}
