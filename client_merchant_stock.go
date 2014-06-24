package wechat

import (
	"errors"
)

// 增加库存
// @productId: 商品ID;
// @skuInfo:   sku信息,格式"id1:vid1;id2:vid2",如商品为统一规格，则此处赋值为空字符串即可;
// @quantity:  增加的库存数量.
func (c *Client) MerchantStockAdd(productId string, skuInfo string,
	quantity int) error {

	if productId == "" {
		return errors.New(`productId == ""`)
	}
	if skuInfo == "" {
		return errors.New(`skuInfo == ""`)
	}
	if quantity <= 0 {
		return errors.New(`quantity <= 0`)
	}

	token, err := c.Token()
	if err != nil {
		return err
	}
	_url := clientMerchantStockAddURL(token)

	var request = struct {
		ProductId string `json:"product_id"`
		SkuInfo   string `json:"sku_info"`
		Quantity  int    `json:"quantity"`
	}{
		ProductId: productId,
		SkuInfo:   skuInfo,
		Quantity:  quantity,
	}

	var result Error
	if err = c.postJSON(_url, &request, &result); err != nil {
		return err
	}

	if result.ErrCode != 0 {
		return &result
	}

	return nil
}

// 减少库存
// @productId: 商品ID;
// @skuInfo:   sku信息,格式"id1:vid1;id2:vid2",如商品为统一规格，则此处赋值为空字符串即可;
// @quantity:  增加的库存数量.
func (c *Client) MerchantStockReduce(productId string, skuInfo string,
	quantity int) error {

	if productId == "" {
		return errors.New(`productId == ""`)
	}
	if skuInfo == "" {
		return errors.New(`skuInfo == ""`)
	}
	if quantity <= 0 {
		return errors.New(`quantity <= 0`)
	}

	token, err := c.Token()
	if err != nil {
		return err
	}
	_url := clientMerchantStockReduceURL(token)

	var request = struct {
		ProductId string `json:"product_id"`
		SkuInfo   string `json:"sku_info"`
		Quantity  int    `json:"quantity"`
	}{
		ProductId: productId,
		SkuInfo:   skuInfo,
		Quantity:  quantity,
	}

	var result Error
	if err = c.postJSON(_url, &request, &result); err != nil {
		return err
	}

	if result.ErrCode != 0 {
		return &result
	}

	return nil
}
