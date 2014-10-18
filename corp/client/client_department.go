// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"errors"
	"github.com/chanxuehong/wechat/corp/addresslist"
)

// 创建部门
func (c *Client) DepartmentCreate(para *addresslist.DepartmentCreateParameters) (id int64, err error) {
	if para == nil {
		err = errors.New("para == nil")
		return
	}

	var result struct {
		Error
		Id int64 `json:"id"`
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	if err = c.postJSON(_DepartmentCreateURL(token), para, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		id = result.Id
		return

	case errCodeTimeout:
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

// 更新部门
func (c *Client) DepartmentUpdate(para addresslist.DepartmentUpdateParameters) (err error) {
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
	if err = c.postJSON(_DepartmentUpdateURL(token), para, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return

	case errCodeTimeout:
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

// 删除部门
func (c *Client) DepartmentDelete(id int64) (err error) {
	var result Error

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	if err = c.getJSON(_DepartmentDeleteURL(token, id), &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return

	case errCodeTimeout:
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

// 获取部门列表
func (c *Client) DepartmentList() (departments []addresslist.Department, err error) {
	var result struct {
		Error
		Departments []addresslist.Department `json:"department"`
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	if err = c.getJSON(_DepartmentListURL(token), &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		departments = result.Departments
		return

	case errCodeTimeout:
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
