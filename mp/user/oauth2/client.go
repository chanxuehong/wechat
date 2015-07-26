// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

// +build !wechatdebug

package oauth2

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TokenStorage interface {
	Get() (*Token, error)
	Put(*Token) error
}

type Client struct {
	Config Config

	// TokenStorage, Token 两个字段正常情况下只用指定一个, 如果两个同时被指定了, 优先使用 TokenStorage;
	// Client 会自动将最新的 Token 更新到 Client.Token 字段, 不管 Token 字段一开始是否被指定!!!
	TokenStorage TokenStorage
	Token        *Token

	HttpClient *http.Client // 如果 HttpClient == nil 则默认用 http.DefaultClient
}

func (clt *Client) httpClient() *http.Client {
	if clt.HttpClient != nil {
		return clt.HttpClient
	}
	return http.DefaultClient
}

func (clt *Client) getJSON(url string, response interface{}) (err error) {
	httpResp, err := clt.httpClient().Get(url)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}

	return json.NewDecoder(httpResp.Body).Decode(response)
}
