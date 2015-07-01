// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package user

import (
	"net/http"

	"github.com/chanxuehong/wechat/mp"
)

type Client struct {
	*mp.Client
}

// 兼容保留, 建议实际项目中全局维护一个 *mp.Client
func NewClient(srv mp.AccessTokenServer, clt *http.Client) Client {
	return Client{
		Client: mp.NewClient(srv, clt),
	}
}
