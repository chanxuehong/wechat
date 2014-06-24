package wechat

import (
	"errors"
	"github.com/chanxuehong/wechat/merchant/product"
)

// 增加商品
func (c *Client) MerchantProductAdd(_product *product.Product) (productId string, err error) {
	if _product == nil {
		return "", errors.New("_product == nil")
	}

	_product.Id = "" // 这个时候还没有 product id

	token, err := c.Token()
	if err != nil {
		return "", err
	}
	_url := clientMerchantProductAddURL(token)

	var result struct {
		Error
		ProductId string `json:"product_id"`
	}
	if err = c.postJSON(_url, _product, &result); err != nil {
		return "", err
	}

	if result.ErrCode != 0 {
		return "", &result.Error
	}

	return result.ProductId, nil
}

// 删除商品
func (c *Client) MerchantProductDelete(productId string) error {
	if productId == "" {
		return errors.New(`productId == ""`)
	}

	token, err := c.Token()
	if err != nil {
		return err
	}
	_url := clientMerchantProductDeleteURL(token)

	var request = struct {
		ProductId string `json:"product_id"`
	}{
		ProductId: productId,
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

// 修改商品
//  NOTE:
//  1. 需要指定 _product.Id 字段
//  2. 从未上架的商品所有信息均可修改，否则商品的名称(name)、商品分类(category)、
//  商品属性(property)这三个字段*不可修改*。
func (c *Client) MerchantProductUpdate(_product *product.Product) error {
	if _product == nil {
		return errors.New("_product == nil")
	}
	if _product.Id == "" {
		return errors.New("product id is not set")
	}

	token, err := c.Token()
	if err != nil {
		return err
	}
	_url := clientMerchantProductUpdateURL(token)

	var result Error
	if err = c.postJSON(_url, _product, &result); err != nil {
		return err
	}

	if result.ErrCode != 0 {
		return &result
	}

	return nil
}

// 查询商品
func (c *Client) MerchantProductGet(productId string) (*product.Product, error) {
	if productId == "" {
		return nil, errors.New(`productId == ""`)
	}

	token, err := c.Token()
	if err != nil {
		return nil, err
	}
	_url := clientMerchantProductGetURL(token)

	var request = struct {
		ProductId string `json:"product_id"`
	}{
		ProductId: productId,
	}

	var result struct {
		Error
		ProductInfo product.Product `json:"product_info"`
	}
	if err = c.postJSON(_url, request, &result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}

	return &result.ProductInfo, nil
}

// 获取指定状态的所有商品.
// 0-所有商品, 1-上架商品, 2-下架商品
func (c *Client) MerchantProductGetByStatus(status int) ([]*product.Product, error) {
	switch status {
	case product.PRODUCT_STATUS_ALL,
		product.PRODUCT_STATUS_ONSHELF,
		product.PRODUCT_STATUS_OFFSHELF:
	default:
		return nil, errors.New("invalid status")
	}

	token, err := c.Token()
	if err != nil {
		return nil, err
	}
	_url := clientMerchantProductGetByStatusURL(token)

	var request = struct {
		Status int `json:"status"`
	}{
		Status: status,
	}

	var result struct {
		Error
		ProductsInfo []*product.Product `json:"products_info"`
	}
	if err = c.postJSON(_url, request, &result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}

	return result.ProductsInfo, nil
}

// 修改商品状态.
// status: 商品上下架标识(0-下架, 1-上架)
func (c *Client) merchantProductModifyStatus(productId string, status int) error {
	if productId == "" {
		return errors.New(`productId == ""`)
	}
	switch status {
	case 0, 1:
	default:
		return errors.New("invalid status")
	}

	token, err := c.Token()
	if err != nil {
		return err
	}
	_url := clientMerchantProductModifyStatusURL(token)

	var request = struct {
		ProductId string `json:"product_id"`
		Status    int    `json:"status"`
	}{
		ProductId: productId,
		Status:    status,
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

// 修改商品到上架状态
func (c *Client) MerchantProductModifyStatusOnShelf(productId string) error {
	return c.merchantProductModifyStatus(productId, 1)
}

// 修改商品到下架状态
func (c *Client) MerchantProductModifyStatusOffShelf(productId string) error {
	return c.merchantProductModifyStatus(productId, 0)
}
