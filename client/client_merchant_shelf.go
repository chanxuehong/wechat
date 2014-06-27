// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong@gmail.com

package client

import (
	"errors"
	"github.com/chanxuehong/wechat/merchant/shelf"
)

// 增加货架
func (c *Client) MerchantShelfAdd(_shelf *shelf.Shelf) (shelfId int64, err error) {
	if _shelf == nil {
		err = errors.New("_shelf == nil")
		return
	}

	_shelf.Id = 0

	token, err := c.Token()
	if err != nil {
		return
	}
	_url := merchantShelfAddURL(token)

	var result struct {
		Error
		ShelfId int64 `json:"shelf_id"`
	}
	if err = c.postJSON(_url, _shelf, &result); err != nil {
		return
	}

	if result.ErrCode != 0 {
		err = &result.Error
		return
	}

	shelfId = result.ShelfId
	return
}

// 删除货架
func (c *Client) MerchantShelfDelete(shelfId int64) error {
	token, err := c.Token()
	if err != nil {
		return err
	}
	_url := merchantShelfDeleteURL(token)

	var request = struct {
		ShelfId int64 `json:"shelf_id"`
	}{
		ShelfId: shelfId,
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

// 修改货架
func (c *Client) MerchantShelfModify(_shelf *shelf.Shelf) error {
	if _shelf == nil {
		return errors.New("_shelf == nil")
	}

	token, err := c.Token()
	if err != nil {
		return err
	}
	_url := merchantShelfModifyURL(token)

	var result Error
	if err = c.postJSON(_url, _shelf, &result); err != nil {
		return err
	}

	if result.ErrCode != 0 {
		return &result
	}

	return nil
}

// 获取所有货架
func (c *Client) MerchantShelfGetAll() ([]shelf.ShelfX, error) {
	token, err := c.Token()
	if err != nil {
		return nil, err
	}
	_url := merchantShelfGetAllURL(token)

	var result = struct {
		Error
		Shelves []shelf.ShelfX `json:"shelves"`
	}{
		Shelves: make([]shelf.ShelfX, 0, 16),
	}
	if err = c.getJSON(_url, &result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}

	return result.Shelves, nil
}

// 根据货架ID获取货架信息
func (c *Client) MerchantShelfGetById(shelfId int64) (*shelf.ShelfX, error) {
	token, err := c.Token()
	if err != nil {
		return nil, err
	}
	_url := merchantShelfGetByIdURL(token)

	var request = struct {
		ShelfId int64 `json:"shelf_id"`
	}{
		ShelfId: shelfId,
	}

	var result struct {
		Error
		shelf.ShelfX
	}
	if err = c.postJSON(_url, request, &result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}

	return &result.ShelfX, nil
}
