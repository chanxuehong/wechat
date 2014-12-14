// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chanxuehong/wechat"
	wechatjson "github.com/chanxuehong/wechat/json"
)

// Client 封装了主动请求功能
type Client struct {
	tokenCache  wechat.TokenCache
	tokenGetter wechat.TokenGetter
	httpClient  *http.Client
}

// 创建一个新的 Client.
//  如果 httpClient == nil 则默认用 http.DefaultClient
func NewClient(tokenCache wechat.TokenCache, tokenGetter wechat.TokenGetter,
	httpClient *http.Client) (clt *Client) {

	if tokenCache == nil {
		panic("tokenCache == nil")
	}
	if tokenGetter == nil {
		panic("tokenGetter == nil")
	}
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		tokenCache:  tokenCache,
		tokenGetter: tokenGetter,
		httpClient:  httpClient,
	}
}

// Client 通用的 json post 请求
func (c *Client) postJSON(url_ string, request interface{}, response interface{}) (err error) {
	buf := textBufferPool.Get().(*bytes.Buffer) // io.ReadWriter
	defer textBufferPool.Put(buf)               // important
	buf.Reset()                                 // important

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
