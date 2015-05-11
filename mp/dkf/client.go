// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package dkf

import (
	"net/http"

	"github.com/chanxuehong/wechat/mp"
)

type Client struct {
	mp.WechatClient
}

// 创建一个新的 Client.
//  如果 HttpClient == nil 则默认用 http.DefaultClient
func NewClient(AccessTokenServer mp.AccessTokenServer, HttpClient *http.Client) *Client {
	if AccessTokenServer == nil {
		panic("AccessTokenServer == nil")
	}
	if HttpClient == nil {
		HttpClient = http.DefaultClient
	}

	return &Client{
		WechatClient: mp.WechatClient{
			AccessTokenServer: AccessTokenServer,
			HttpClient:        HttpClient,
		},
	}
}
