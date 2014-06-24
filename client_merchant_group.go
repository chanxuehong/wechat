package wechat

import (
	"errors"
	"github.com/chanxuehong/wechat/merchant/group"
)

// 增加分组
func (c *Client) MerchantGroupAdd(_group *group.GroupExt) (groupId int64, err error) {
	if _group == nil {
		err = errors.New("_group == nil")
		return
	}

	_group.Id = 0

	token, err := c.Token()
	if err != nil {
		return
	}
	_url := clientMerchantGroupAddURL(token)

	var request = struct {
		GroupDetail *group.GroupExt `json:"group_detail"`
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
	_url := clientMerchantGroupDeleteURL(token)

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
	_url := clientMerchantGroupPropertyModifyURL(token)

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
	_url := clientMerchantGroupProductModifyURL(token)

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
func (c *Client) MerchantGroupGetAll() ([]*group.Group, error) {
	token, err := c.Token()
	if err != nil {
		return nil, err
	}
	_url := clientMerchantGroupGetAllURL(token)

	var result struct {
		Error
		GroupsDetail []*group.Group `json:"groups_detail"`
	}
	if err = c.getJSON(_url, &result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}

	return result.GroupsDetail, nil
}

// 根据分组ID获取分组信息
func (c *Client) MerchantGroupGetById(groupId int64) (*group.GroupExt, error) {
	token, err := c.Token()
	if err != nil {
		return nil, err
	}
	_url := clientMerchantGroupGetByIdURL(token)

	var request = struct {
		GroupId int64 `json:"group_id"`
	}{
		GroupId: groupId,
	}

	var result struct {
		Error
		GroupDetail group.GroupExt `json:"group_detail"`
	}
	if err = c.postJSON(_url, request, &result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}

	return &result.GroupDetail, nil
}
