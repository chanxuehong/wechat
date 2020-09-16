package core

import (
	"net/http"

	"github.com/chanxuehong/wechat/mp/core"
)

type Client struct {
	*core.Client
	AppId          string
	Token          string
	EncodingAESKey string
}

// NewClient 创建一个新的 Client.
//  如果 clt == nil 则默认用 util.DefaultHttpClient
func NewClient(appId string, token string, aesKey string, clt *http.Client) *Client {
	return &Client{
		Client:         core.NewClientWithToken(token, clt),
		AppId:          appId,
		Token:          token,
		EncodingAESKey: aesKey,
	}
}
