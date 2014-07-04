// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"github.com/chanxuehong/wechat/merchant/category"
)

// 获取指定分类的所有子分类.
//  @categoryId: 大分类ID(根节点分类id为1)
func (c *Client) MerchantCategoryGetSub(categoryId int64) ([]category.Category, error) {
	token, err := c.Token()
	if err != nil {
		return nil, err
	}
	_url := merchantCategoryGetSubURL(token)

	var request = struct {
		CategoryId int64 `json:"cate_id"`
	}{
		CategoryId: categoryId,
	}

	var result struct {
		Error
		Categories []category.Category `json:"cate_list"`
	}
	result.Categories = make([]category.Category, 0, 64)
	if err = c.postJSON(_url, request, &result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}

	return result.Categories, nil
}

// 获取指定子分类的所有SKU
func (c *Client) MerchantCategoryGetSKU(categoryId int64) ([]category.SKU, error) {
	token, err := c.Token()
	if err != nil {
		return nil, err
	}
	_url := merchantCategoryGetSKUURL(token)

	var request = struct {
		CategoryId int64 `json:"cate_id"`
	}{
		CategoryId: categoryId,
	}

	var result struct {
		Error
		SKUs []category.SKU `json:"sku_table"`
	}
	result.SKUs = make([]category.SKU, 0, 256)
	if err = c.postJSON(_url, request, &result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}

	return result.SKUs, nil
}

// 获取指定分类的所有属性
func (c *Client) MerchantCategoryGetProperty(categoryId int64) ([]category.Property, error) {
	token, err := c.Token()
	if err != nil {
		return nil, err
	}
	_url := merchantCategoryGetPropertyURL(token)

	var request = struct {
		CategoryId int64 `json:"cate_id"`
	}{
		CategoryId: categoryId,
	}

	var result struct {
		Error
		Properties []category.Property `json:"properties"`
	}
	result.Properties = make([]category.Property, 0, 64)
	if err = c.postJSON(_url, request, &result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}

	return result.Properties, nil
}
