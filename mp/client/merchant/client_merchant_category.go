// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package merchant

import (
	"github.com/chanxuehong/wechat/mp/merchant/category"
)

// 获取指定分类的所有子分类.
//  @categoryId: 大分类ID(根节点分类id为1)
func (c *Client) MerchantCategoryGetSub(categoryId int64) (categories []category.Category, err error) {
	var request = struct {
		CategoryId int64 `json:"cate_id"`
	}{
		CategoryId: categoryId,
	}

	var result struct {
		Error
		Categories []category.Category `json:"cate_list"`
	}
	result.Categories = make([]category.Category, 0, 16)

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	url_ := merchantCategoryGetSubURL(token)
	if err = c.postJSON(url_, request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		categories = result.Categories
		return

	case errCodeInvalidCredential, errCodeTimeout:
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

// 获取指定子分类的所有SKU
func (c *Client) MerchantCategoryGetSKU(categoryId int64) (skus []category.SKU, err error) {
	var request = struct {
		CategoryId int64 `json:"cate_id"`
	}{
		CategoryId: categoryId,
	}

	var result struct {
		Error
		SKUs []category.SKU `json:"sku_table"`
	}
	result.SKUs = make([]category.SKU, 0, 16)

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	url_ := merchantCategoryGetSKUURL(token)
	if err = c.postJSON(url_, request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		skus = result.SKUs
		return

	case errCodeInvalidCredential, errCodeTimeout:
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

// 获取指定分类的所有属性
func (c *Client) MerchantCategoryGetProperty(categoryId int64) (properties []category.Property, err error) {
	var request = struct {
		CategoryId int64 `json:"cate_id"`
	}{
		CategoryId: categoryId,
	}

	var result struct {
		Error
		Properties []category.Property `json:"properties"`
	}
	result.Properties = make([]category.Property, 0, 16)

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	url_ := merchantCategoryGetPropertyURL(token)
	if err = c.postJSON(url_, request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		properties = result.Properties
		return

	case errCodeInvalidCredential, errCodeTimeout:
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
