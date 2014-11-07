// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package tokencache

import "errors"

var ErrCacheMiss = errors.New("token cache miss")

type TokenCache interface {
	// 从 cache 里获取 token
	//  NOTE: 如果没有找到返回 ErrCacheMiss
	Token() (token string, err error)

	// 添加或者重置 token
	PutToken(token string) (err error)
}
