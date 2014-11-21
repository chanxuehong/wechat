// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

// copy from github.com/chanxuehong/wechat/mp/client/client.go

package merchant

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	wechatjson "github.com/chanxuehong/wechat/json"
	"github.com/chanxuehong/wechat/mp/tokenservice"
)

type Client struct {
	tokenService tokenservice.TokenService
	httpClient   *http.Client
}

// 创建一个新的 Client.
//  如果 httpClient == nil 则默认用 http.DefaultClient,
func NewClient(tokenService tokenservice.TokenService, httpClient *http.Client) (clt *Client) {
	if tokenService == nil {
		panic("tokenService == nil")
	}
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	clt = &Client{
		tokenService: tokenService,
		httpClient:   httpClient,
	}
	return
}

// 获取 access token
// 正常情况下 token != "" && err == nil, 否则 token == "" && err != nil
func (c *Client) Token() (token string, err error) {
	return c.tokenService.Token()
}

// 从微信服务器获取新的 access token.
// 正常情况下 token != "" && err == nil, 否则 token == "" && err != nil
//  NOTE:
//  1. 一般情况下无需调用该函数, 请使用 Token() 获取 access token.
//  2. 即使 access token 过期(错误代码 40001, 正常情况下不会出现),
//     也请谨慎调用 TokenRefresh, 建议直接返回错误! 因为很有可能高并发情况下造成雪崩效应!
//  3. 再次强调, 调用这个函数你应该知道发生了什么!!!
func (c *Client) TokenRefresh() (token string, err error) {
	return c.tokenService.TokenRefresh()
}

// Client 通用的 json post 请求
func (c *Client) postJSON(url_ string, request interface{}, response interface{}) (err error) {
	buf := textBufferPool.Get().(*bytes.Buffer) // io.ReadWriter
	buf.Reset()                                 // important
	defer textBufferPool.Put(buf)

	if err = wechatjson.NewEncoder(buf).Encode(request); err != nil {
		return
	}

	resp, err := c.httpClient.Post(url_, "application/json; charset=utf-8", buf)
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
func (c *Client) getJSON(url_ string, response interface{}) (err error) {
	resp, err := c.httpClient.Get(url_)
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
