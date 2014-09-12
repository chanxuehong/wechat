// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	wechatjson "github.com/chanxuehong/wechat/json"
	"net/http"
)

// Client 封装了主动请求功能, 比如创建菜单, 回复客服消息等等
type Client struct {
	tokenCache   TokenCache
	tokenService TokenService
	httpClient   *http.Client
}

// 创建一个新的 Client.
//  NOTE:
//  1. tokenCache 必须实现, tokenService, httpClient 均可为 nil.
//  2. *DefaultTokenCacheService 同时实现了 TokenCache, TokenService 接口,
//     请注意:
//     A. 如果 tokenCache 的类型是 *DefaultTokenCacheService, 则 tokenService 要么为 nil,
//        要么也必须是 *DefaultTokenCacheService
//     B. 如果 tokenCache 的类型不是 *DefaultTokenCacheService, 则 tokenService 的
//        类型也不能是 *DefaultTokenCacheService, 需要自己实现!
//  3. 如果 httpClient == nil 则默认用 http.DefaultClient, see ../CommonHttpClient 和 ../MediaHttpClient.
func NewClient(tokenCache TokenCache, tokenService TokenService, httpClient *http.Client) (clt *Client) {
	if tokenCache == nil {
		panic("tokenCache == nil")
	}

	clt = &Client{
		tokenCache:   tokenCache,
		tokenService: tokenService,
	}

	if httpClient == nil {
		clt.httpClient = http.DefaultClient
	} else {
		clt.httpClient = httpClient
	}

	return
}

// Client 通用的 json post 请求
func (c *Client) postJSON(_url string, request interface{}, response interface{}) (err error) {
	buf := textBufferPool.Get().(*bytes.Buffer) // io.ReadWriter
	buf.Reset()                                 // important
	defer textBufferPool.Put(buf)

	if err = wechatjson.NewEncoder(buf).Encode(request); err != nil {
		return
	}

	resp, err := c.httpClient.Post(_url, "application/json; charset=utf-8", buf)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", resp.Status)
	}

	if err = json.NewDecoder(resp.Body).Decode(response); err != nil {
		return
	}

	return
}

// Client 通用的 json get 请求
func (c *Client) getJSON(_url string, response interface{}) (err error) {
	resp, err := c.httpClient.Get(_url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", resp.Status)
	}

	if err = json.NewDecoder(resp.Body).Decode(response); err != nil {
		return
	}

	return
}
