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

type Client struct {
	Config Config

	// TokenStorage, Token 正常情况下只需要指定一个; 如果两个都指定了, 优先使用 TokenStorage;
	// 程序会自动更新最新的 Token 到 Client.Token, 不管一开始是否已经赋值.
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
