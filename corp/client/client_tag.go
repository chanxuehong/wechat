// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"errors"
	"strings"

	"github.com/chanxuehong/wechat/corp/addresslist"
)

// 创建标签
func (c *Client) TagCreate(name string) (id int64, err error) {
	var request = struct {
		Name string `json:"tagname"`
	}{
		Name: name,
	}

	var result struct {
		Error
		Id int64 `json:"tagid"`
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	if err = c.postJSON(_TagCreateURL(token), request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		id = result.Id
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

// 更新标签名字
func (c *Client) TagUpdate(id int64, name string) (err error) {
	var request = struct {
		Id   int64  `json:"tagid"`
		Name string `json:"tagname"`
	}{
		Id:   id,
		Name: name,
	}

	var result Error

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	if err = c.postJSON(_TagUpdateURL(token), &request, &result); err != nil {
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

// 删除标签
func (c *Client) TagDelete(id int64) (err error) {
	var result Error

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	if err = c.getJSON(_TagDeleteURL(token, id), &result); err != nil {
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

// 获取标签成员
func (c *Client) TagUserList(tagId int64) (userList []addresslist.UserInfoBase, err error) {
	var result struct {
		Error
		UserList []addresslist.UserInfoBase `json:"userlist"`
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	if err = c.getJSON(_TagUserListURL(token, tagId), &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		userList = result.UserList
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

// 增加标签成员
//  NOTE: 如果 err != nil 则不用考虑 invalidUsers, 否则还要考虑 invalidUsers
func (c *Client) TagUserAdd(tagId int64, users []string) (invalidUsers []string, err error) {
	if len(users) == 0 {
		err = errors.New("users is empty")
		return
	}

	var request = struct {
		TagId int64    `json:"tagid"`
		Users []string `json:"userlist"`
	}{
		TagId: tagId,
		Users: users,
	}

	var result struct {
		Error
		InvalidList string `json:"invalidlist"`
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	if err = c.postJSON(_TagUserAddURL(token), &request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		if len(result.InvalidList) == 0 { // 全部成功
			return
		} else { // 部分userid非法
			invalidUsers = strings.Split(result.InvalidList, "|")
			return
		}

	case 40070: // userid全部非法
		invalidUsers = users
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

// 删除标签成员
//  NOTE: 如果 err != nil 则不用考虑 invalidUsers, 否则还要考虑 invalidUsers
func (c *Client) TagUserDel(tagId int64, users []string) (invalidUsers []string, err error) {
	if len(users) == 0 {
		err = errors.New("users is empty")
		return
	}

	var request = struct {
		TagId int64    `json:"tagid"`
		Users []string `json:"userlist"`
	}{
		TagId: tagId,
		Users: users,
	}

	var result struct {
		Error
		InvalidList string `json:"invalidlist"`
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	if err = c.postJSON(_TagUserDeleteURL(token), &request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		if len(result.InvalidList) == 0 { // 全部成功
			return
		} else { // 部分userid非法
			invalidUsers = strings.Split(result.InvalidList, "|")
			return
		}

	case 40031: // userid全部非法
		invalidUsers = users
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
