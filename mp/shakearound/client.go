// @description wechat
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com) Harry Rong(harrykobe@gmail.com)
package shakearound

import (
    "net/http"

    "github.com/chanxuehong/wechat/mp"
)

type Client struct {
    *mp.WechatClient
}


func NewClient(AccessTokenServer mp.AccessTokenServer, httpClient *http.Client) Client {
    return Client{
        WechatClient: mp.NewWechatClient(AccessTokenServer, httpClient),
    }
}
