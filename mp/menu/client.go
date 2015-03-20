// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package menu

import (
	"net/http"

	"github.com/chanxuehong/wechat/mp"
)

type Client struct {
	mp.WechatClient
}

// 创建一个新的 Client.
//  如果 HttpClient == nil 则默认用 http.DefaultClient
func NewClient(TokenServer mp.TokenServer, HttpClient *http.Client) *Client {
	if TokenServer == nil {
		panic("TokenServer == nil")
	}
	if HttpClient == nil {
		HttpClient = http.DefaultClient
	}

	return &Client{
		WechatClient: mp.WechatClient{
			TokenServer: TokenServer,
			HttpClient:  HttpClient,
		},
	}
}

// 创建自定义菜单.
func (clt *Client) CreateMenu(menu Menu) (err error) {
	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/menu/create?access_token="
	if err = clt.PostJSON(incompleteURL, &menu, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}

// 删除自定义菜单
func (clt *Client) DeleteMenu() (err error) {
	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token="
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}

// 获取自定义菜单
func (clt *Client) GetMenu() (menu Menu, err error) {
	var result struct {
		mp.Error
		Menu Menu `json:"menu"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/menu/get?access_token="
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	menu = result.Menu
	return
}
