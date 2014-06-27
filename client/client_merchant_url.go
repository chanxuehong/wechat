// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"net/url"
)

// https://api.weixin.qq.com/merchant/common/upload_img?access_token=ACCESS_TOKEN&filename=test.png
func merchantUploadImageURL(accesstoken, filename string) string {
	return "https://api.weixin.qq.com/merchant/common/upload_img?access_token=" +
		accesstoken +
		"&filename=" +
		url.QueryEscape(filename)
}

// =============================================================================

// https://api.weixin.qq.com/merchant/create?access_token=ACCESS_TOKEN
func merchantProductAddURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/create?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/del?access_token=ACCESS_TOKEN
func merchantProductDeleteURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/del?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/update?access_token=ACCESS_TOKEN
func merchantProductUpdateURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/update?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/get?access_token=ACCESS_TOKEN
func merchantProductGetURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/get?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/getbystatus?access_token=ACCESS_TOKEN
func merchantProductGetByStatusURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/getbystatus?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/modproductstatus?access_token=ACCESS_TOKEN
func merchantProductModifyStatusURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/modproductstatus?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/category/getsub?access_token=ACCESS_TOKEN
func merchantCategoryGetSubURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/category/getsub?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/category/getsku?access_token=ACCESS_TOKEN
func merchantCategoryGetSKUURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/category/getsku?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/category/getproperty?access_token=ACCESS_TOKEN
func merchantCategoryGetPropertyURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/category/getproperty?access_token=" +
		accesstoken
}

// =============================================================================

// https://api.weixin.qq.com/merchant/stock/add?access_token=ACCESS_TOKEN
func merchantStockAddURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/stock/add?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/stock/reduce?access_token=ACCESS_TOKEN
func merchantStockReduceURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/stock/reduce?access_token=" +
		accesstoken
}

// =============================================================================

// https://api.weixin.qq.com/merchant/express/add?access_token=ACCESS_TOKEN
func merchantExpressAddURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/express/add?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/express/del?access_token=ACCESS_TOKEN
func merchantExpressDeleteURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/express/del?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/express/update?access_token=ACCESS_TOKEN
func merchantExpressUpdateURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/express/update?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/express/getbyid?access_token=ACCESS_TOKEN
func merchantExpressGetByIdURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/express/getbyid?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/express/getall?access_token=ACCESS_TOKEN
func merchantExpressGetAllURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/express/getall?access_token=" +
		accesstoken
}

// =============================================================================

// https://api.weixin.qq.com/merchant/group/add?access_token=ACCESS_TOKEN
func merchantGroupAddURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/group/add?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/group/del?access_token=ACCESS_TOKEN
func merchantGroupDeleteURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/group/del?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/group/propertymod?access_token=ACCESS_TOKEN
func merchantGroupPropertyModifyURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/group/propertymod?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/group/productmod?access_token=ACCESS_TOKEN
func merchantGroupProductModifyURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/group/productmod?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/group/getall?access_token=ACCESS_TOKEN
func merchantGroupGetAllURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/group/getall?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/group/getbyid?access_token=ACCESS_TOKEN
func merchantGroupGetByIdURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/group/getbyid?access_token=" +
		accesstoken
}

// =============================================================================

// https://api.weixin.qq.com/merchant/shelf/add?access_token=ACCESS_TOKEN
func merchantShelfAddURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/shelf/add?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/shelf/del?access_token=ACCESS_TOKEN
func merchantShelfDeleteURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/shelf/del?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/shelf/mod?access_token=ACCESS_TOKEN
func merchantShelfModifyURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/shelf/mod?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/shelf/getall?access_token=ACCESS_TOKEN
func merchantShelfGetAllURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/shelf/getall?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/shelf/getbyid?access_token=ACCESS_TOKEN
func merchantShelfGetByIdURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/shelf/getbyid?access_token=" +
		accesstoken
}

// =============================================================================

// https://api.weixin.qq.com/merchant/order/getbyid?access_token=ACCESS_TOKEN
func merchantOrderGetByIdURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/order/getbyid?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/order/getbyfilter?access_token=ACCESS_TOKEN
func merchantOrderGetByFilterURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/order/getbyfilter?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/order/setdelivery?access_token=ACCESS_TOKEN
func merchantOrderSetDeliveryURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/order/setdelivery?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/order/close?access_token=ACCESS_TOKEN
func merchantOrderCloseURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/order/close?access_token=" +
		accesstoken
}
