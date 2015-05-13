// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package menu

import (
	"net/http"
	"strconv"

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

// 创建自定义菜单.
func (clt Client) CreateMenu(agentId int64, menu Menu) (err error) {
	var result corp.Error

	incompleteURL := "https://qyapi.weixin.qq.com/cgi-bin/menu/create?agentid=" +
		strconv.FormatInt(agentId, 10) + "&access_token="
	if err = clt.PostJSON(incompleteURL, &menu, &result); err != nil {
		return
	}

	if result.ErrCode != corp.ErrCodeOK {
		err = &result
		return
	}
	return
}

// 删除自定义菜单
func (clt Client) DeleteMenu(agentId int64) (err error) {
	var result corp.Error

	incompleteURL := "https://qyapi.weixin.qq.com/cgi-bin/menu/delete?agentid=" +
		strconv.FormatInt(agentId, 10) + "&access_token="
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != corp.ErrCodeOK {
		err = &result
		return
	}
	return
}

// 获取自定义菜单
func (clt Client) GetMenu(agentId int64) (menu Menu, err error) {
	var result struct {
		corp.Error
		Menu Menu `json:"menu"`
	}

	incompleteURL := "https://qyapi.weixin.qq.com/cgi-bin/menu/get?agentid=" +
		strconv.FormatInt(agentId, 10) + "&access_token="
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != corp.ErrCodeOK {
		err = &result.Error
		return
	}
	menu = result.Menu
	return
}
