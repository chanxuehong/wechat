// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"errors"
)

// 获取 access token
func (c *Client) Token() (token string, err error) {
	return c.tokenCache.Token()
}

// 从微信服务器获取新的 access token.
// 正常情况下 token != "" && err == nil, 否则 token == "" && err != nil
//
//  NOTE:
//  1. 一般情况下无需调用该函数, 请使用 Token() 获取 access token.
//  2. 即使函数调用过程中返回了 access token 过期错误(错误代码 42001, 正常情况下不会出现),
//     也请谨慎调用 TokenRefresh, 建议直接返回错误! 因为很有可能造成雪崩效应!
//  3. 再次强调, 调用这个函数你应该知道发生了什么!!!
func (c *Client) TokenRefresh() (token string, err error) {
	if c.tokenService == nil {
		err = errors.New("Client 没有指定有效的 TokenService")
		return
	}
	return c.tokenService.TokenRefresh()
}
