package wechat

import (
	"github.com/chanxuehong/wechat/merchant/product/category"
)

// 获取指定分类的所有子分类.
// @categoryId: 大分类ID(根节点分类id为1)
func (c *Client) MerchantCategoryGetSub(categoryId int) ([]category.Category, error) {
	token, err := c.Token()
	if err != nil {
		return nil, err
	}
	_url := clientMerchantCategoryGetSubURL(token)

	var request = struct {
		CategoryId int `json:"cate_id"`
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
func (c *Client) MerchantCategoryGetSKU(categoryId int) ([]category.SKU, error) {
	token, err := c.Token()
	if err != nil {
		return nil, err
	}
	_url := clientMerchantCategoryGetSKUURL(token)

	var request = struct {
		CategoryId int `json:"cate_id"`
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
func (c *Client) MerchantCategoryGetProperty(categoryId int) ([]category.Property, error) {
	token, err := c.Token()
	if err != nil {
		return nil, err
	}
	_url := clientMerchantCategoryGetPropertyURL(token)

	var request = struct {
		CategoryId int `json:"cate_id"`
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
