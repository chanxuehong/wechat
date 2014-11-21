// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"errors"

	"github.com/chanxuehong/wechat/corp/addresslist"
)

// 创建成员
func (c *Client) UserCreate(para *addresslist.UserCreateParameters) (err error) {
	if para == nil {
		return errors.New("para == nil")
	}

	var result Error

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	if err = c.postJSON(_UserCreateURL(token), para, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return

	case errCodeInvalidCredential:
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

// 更新成员
func (c *Client) UserUpdate(para addresslist.UserUpdateParameters) (err error) {
	if para == nil {
		return errors.New("para == nil")
	}

	var result Error

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	if err = c.postJSON(_UserUpdateURL(token), para, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return

	case errCodeInvalidCredential:
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

// 删除成员
func (c *Client) UserDelete(userid string) (err error) {
	if len(userid) == 0 {
		return errors.New(`userid == ""`)
	}

	var result Error

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	if err = c.getJSON(_UserDeleteURL(token, userid), &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return

	case errCodeInvalidCredential:
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

// 获取指定成员信息
func (c *Client) UserInfo(userid string) (info *addresslist.UserInfo, err error) {
	if len(userid) == 0 {
		err = errors.New(`userid == ""`)
		return
	}

	var result struct {
		Error
		addresslist.UserInfo
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	if err = c.getJSON(_UserGetURL(token, userid), &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		info = &result.UserInfo
		return

	case errCodeInvalidCredential:
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

// 获取部门成员
//  departmentId: 获取的部门id
//  fetchChild:   是否递归获取子部门下面的成员
//  status:       0 获取全部员工，1 获取已关注成员列表，2 获取禁用成员列表，4 获取未关注成员列表。
//                status可叠加（可用逻辑运算符 | 来叠加, 一般都是后面 3 个叠加）。
func (c *Client) UserSimpleList(departmentId int64,
	fetchChild bool, status int) (userList []addresslist.UserInfoBase, err error) {

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
	if err = c.getJSON(_UserSimpleListURL(token, departmentId, fetchChild, status), &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		userList = result.UserList
		return

	case errCodeInvalidCredential:
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
