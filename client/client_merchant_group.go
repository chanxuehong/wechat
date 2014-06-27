// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong@gmail.com

package client

import (
	"errors"
	"github.com/chanxuehong/wechat/merchant/group"
)

// 增加分组
func (c *Client) MerchantGroupAdd(_group *group.GroupEx) (groupId int64, err error) {
	if _group == nil {
		err = errors.New("_group == nil")
		return
	}

	_group.Id = 0

	token, err := c.Token()
	if err != nil {
		return
	}
	_url := merchantGroupAddURL(token)

	var request = struct {
		GroupDetail *group.GroupEx `json:"group_detail"`
	}{
		GroupDetail: _group,
	}

	var result struct {
		Error
		GroupId int64 `json:"group_id"`
	}
	if err = c.postJSON(_url, request, &result); err != nil {
		return
	}

	if result.ErrCode != 0 {
		err = &result.Error
		return
	}

	groupId = result.GroupId
	return
}

// 删除分组
func (c *Client) MerchantGroupDelete(groupId int64) error {
	token, err := c.Token()
	if err != nil {
		return err
	}
	_url := merchantGroupDeleteURL(token)

	var request = struct {
		GroupId int64 `json:"group_id"`
	}{
		GroupId: groupId,
	}

	var result Error
	if err = c.postJSON(_url, request, &result); err != nil {
		return err
	}

	if result.ErrCode != 0 {
		return &result
	}

	return nil
}

// 修改分组名称
func (c *Client) MerchantGroupRename(groupId int64, newName string) error {
	if newName == "" {
		return errors.New(`newName == ""`)
	}

	token, err := c.Token()
	if err != nil {
		return err
	}
	_url := merchantGroupPropertyModifyURL(token)

	var request = struct {
		GroupId   int64  `json:"group_id"`
		GroupName string `json:"group_name"`
	}{
		GroupId:   groupId,
		GroupName: newName,
	}

	var result Error
	if err = c.postJSON(_url, request, &result); err != nil {
		return err
	}

	if result.ErrCode != 0 {
		return &result
	}

	return nil
}

// 修改分组商品
func (c *Client) MerchantGroupProductModify(modifyRequest *group.GroupProductModifyRequest) error {
	if modifyRequest == nil {
		return errors.New("modifyRequest == nil")
	}

	token, err := c.Token()
	if err != nil {
		return err
	}
	_url := merchantGroupProductModifyURL(token)

	var result Error
	if err = c.postJSON(_url, modifyRequest, &result); err != nil {
		return err
	}

	if result.ErrCode != 0 {
		return &result
	}

	return nil
}

// 获取所有分组
func (c *Client) MerchantGroupGetAll() ([]group.Group, error) {
	token, err := c.Token()
	if err != nil {
		return nil, err
	}
	_url := merchantGroupGetAllURL(token)

	var result struct {
		Error
		GroupsDetail []group.Group `json:"groups_detail"`
	}
	result.GroupsDetail = make([]group.Group, 0, 64)
	if err = c.getJSON(_url, &result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}

	return result.GroupsDetail, nil
}

// 根据分组ID获取分组信息
func (c *Client) MerchantGroupGetById(groupId int64) (*group.GroupEx, error) {
	token, err := c.Token()
	if err != nil {
		return nil, err
	}
	_url := merchantGroupGetByIdURL(token)

	var request = struct {
		GroupId int64 `json:"group_id"`
	}{
		GroupId: groupId,
	}

	var result struct {
		Error
		GroupDetail group.GroupEx `json:"group_detail"`
	}
	if err = c.postJSON(_url, request, &result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}

	return &result.GroupDetail, nil
}
