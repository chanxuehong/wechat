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
	*corp.CorpClient
}

// 兼容保留, 建議實際項目全局維護一個 *corp.CorpClient
func NewClient(AccessTokenServer corp.AccessTokenServer, httpClient *http.Client) Client {
	return Client{
		CorpClient: corp.NewCorpClient(AccessTokenServer, httpClient),
	}
}
