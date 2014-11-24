// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"github.com/chanxuehong/wechat/corp/menu"
)

// 创建自定义菜单
func (c *Client) MenuCreate(menu_ menu.Menu, agentId int64) (err error) {
	var result Error

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	if err = c.postJSON(_MenuCreateURL(token, agentId), menu_, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return
	case errCodeTimeout, errCodeInvalidCredential:
		if !hasRetry {
			hasRetry = true

			if token, err = c.TokenRefresh(); err != nil {
				return
			}
			goto RETRY
		}
		fallthrough
	default:
		err = &result
		return
	}
}

// 删除自定义菜单
func (c *Client) MenuDelete(agentId int64) (err error) {
	var result Error

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	if err = c.getJSON(_MenuDeleteURL(token, agentId), &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return
	case errCodeTimeout, errCodeInvalidCredential:
		if !hasRetry {
			hasRetry = true

			if token, err = c.TokenRefresh(); err != nil {
				return
			}
			goto RETRY
		}
		fallthrough
	default:
		err = &result
		return
	}
}

// 获取自定义菜单
func (c *Client) MenuGet(agentId int64) (menu_ menu.Menu, err error) {
	var result struct {
		Error
		Menu menu.Menu `json:"menu"`
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	if err = c.getJSON(_MenuGetURL(token, agentId), &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		menu_ = result.Menu
		return
	case errCodeTimeout, errCodeInvalidCredential:
		if !hasRetry {
			hasRetry = true

			if token, err = c.TokenRefresh(); err != nil {
				return
			}
			goto RETRY
		}
		fallthrough
	default:
		err = &result.Error
		return
	}
}
