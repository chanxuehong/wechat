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

// Client 封装了主动请求功能
type Client struct {
	corpId     string
	corpSecret string

	tokenCache TokenCache
	httpClient *http.Client
}

// 创建一个新的 Client.
//  如果 httpClient == nil 则默认用 http.DefaultClient,
//  see github.com/chanxuehong/wechat/CommonHttpClient 和
//      github.com/chanxuehong/wechat/MediaHttpClient
func NewClient(corpId, corpSecret string,
	tokenCache TokenCache, httpClient *http.Client) (clt *Client) {

	if tokenCache == nil {
		panic("tokenCache == nil")
	}

	clt = &Client{
		corpId:     corpId,
		corpSecret: corpSecret,
		tokenCache: tokenCache,
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
