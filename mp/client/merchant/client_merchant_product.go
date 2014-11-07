// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package merchant

import (
	"errors"

	"github.com/chanxuehong/wechat/mp/merchant/product"
)

// 增加商品
//  NOTE: 无需指定 Id 和 Status 字段
func (c *Client) MerchantProductAdd(_product *product.Product) (productId string, err error) {
	if _product == nil {
		err = errors.New("_product == nil")
		return
	}

	// 无需指定 Id 和 Status 字段
	_product.Id = ""
	_product.Status = 0

	var result struct {
		Error
		ProductId string `json:"product_id"`
	}

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := merchantProductAddURL(token)
	if err = c.postJSON(_url, _product, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		productId = result.ProductId
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = &result.Error
		return
	}
}

// 删除商品
func (c *Client) MerchantProductDelete(productId string) (err error) {
	if productId == "" {
		return errors.New(`productId == ""`)
	}

	var request = struct {
		ProductId string `json:"product_id"`
	}{
		ProductId: productId,
	}

	var result Error

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := merchantProductDeleteURL(token)
	if err = c.postJSON(_url, request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = &result
		return
	}
}

// 修改商品
//  NOTE:
//  1. 需要指定 _product.Id 字段
//  2. 从未上架的商品所有信息均可修改，否则商品的名称(name)、商品分类(category)、
//  商品属性(property)这三个字段*不可修改*。
func (c *Client) MerchantProductUpdate(_product *product.Product) (err error) {
	if _product == nil {
		return errors.New("_product == nil")
	}
	if _product.Id == "" {
		return errors.New("product id is not set")
	}

	var result Error

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := merchantProductUpdateURL(token)
	if err = c.postJSON(_url, _product, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = &result
		return
	}
}

// 查询商品
func (c *Client) MerchantProductGet(productId string) (_product *product.Product, err error) {
	var request = struct {
		ProductId string `json:"product_id"`
	}{
		ProductId: productId,
	}

	var result struct {
		Error
		ProductInfo product.Product `json:"product_info"`
	}

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := merchantProductGetURL(token)
	if err = c.postJSON(_url, request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		_product = &result.ProductInfo
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = &result.Error
		return
	}
}

// 获取所有商品，包括上架商品 和 下架商品
func (c *Client) MerchantProductGetAll() ([]product.Product, error) {
	return c.merchantProductGetByStatus(0)
}

// 获取所有上架商品
func (c *Client) MerchantProductGetAllOnShelf() ([]product.Product, error) {
	return c.merchantProductGetByStatus(1)
}

// 获取所有下架商品
func (c *Client) MerchantProductGetAllOffShelf() ([]product.Product, error) {
	return c.merchantProductGetByStatus(2)
}

// 获取指定状态的所有商品.
// 0-所有商品, 1-上架商品, 2-下架商品
func (c *Client) merchantProductGetByStatus(status int) (products []product.Product, err error) {
	var request = struct {
		Status int `json:"status"`
	}{
		Status: status,
	}

	var result struct {
		Error
		ProductsInfo []product.Product `json:"products_info"`
	}
	result.ProductsInfo = make([]product.Product, 0, 64)

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := merchantProductGetByStatusURL(token)
	if err = c.postJSON(_url, request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		products = result.ProductsInfo
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = &result.Error
		return
	}
}

// 修改商品到上架状态
func (c *Client) MerchantProductModifyStatusOnShelf(productId string) error {
	return c.merchantProductModifyStatus(productId, 1)
}

// 修改商品到下架状态
func (c *Client) MerchantProductModifyStatusOffShelf(productId string) error {
	return c.merchantProductModifyStatus(productId, 0)
}

// 修改商品状态.
// status: 商品上下架标识(0-下架, 1-上架)
func (c *Client) merchantProductModifyStatus(productId string, status int) (err error) {
	var request = struct {
		ProductId string `json:"product_id"`
		Status    int    `json:"status"`
	}{
		ProductId: productId,
		Status:    status,
	}

	var result Error

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := merchantProductModifyStatusURL(token)
	if err = c.postJSON(_url, request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = &result
		return
	}
}
