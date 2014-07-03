// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

// 增加库存
// @productId: 商品ID;
// @skuInfo:   sku信息,格式"id1:vid1;id2:vid2",如商品为统一规格，则此处赋值为空字符串即可;
// @quantity:  增加的库存数量.
func (c *Client) MerchantStockAdd(productId string, skuInfo string, quantity int) error {
	token, err := c.Token()
	if err != nil {
		return err
	}
	_url := merchantStockAddURL(token)

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
func (c *Client) MerchantStockReduce(productId string, skuInfo string, quantity int) error {
	token, err := c.Token()
	if err != nil {
		return err
	}
	_url := merchantStockReduceURL(token)

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
