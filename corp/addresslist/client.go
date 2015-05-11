// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package addresslist

import (
	"net/http"

	"github.com/chanxuehong/wechat/corp"
)

type Client struct {
	corp.CorpClient
}

// 创建一个新的 Client.
//  如果 HttpClient == nil 则默认用 http.DefaultClient
func NewClient(AccessTokenServer corp.AccessTokenServer, HttpClient *http.Client) *Client {
	if AccessTokenServer == nil {
		panic("AccessTokenServer == nil")
	}
	if HttpClient == nil {
		HttpClient = http.DefaultClient
	}

	return &Client{
		CorpClient: corp.CorpClient{
			AccessTokenServer: AccessTokenServer,
			HttpClient:        HttpClient,
		},
	}
}
