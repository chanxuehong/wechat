package wechat

import (
	"net/url"
)

// https://api.weixin.qq.com/merchant/common/upload_img?access_token=ACCESS_TOKEN&filename=test.png
func clientMerchantUploadImageURL(accesstoken, filename string) string {
	return "https://api.weixin.qq.com/merchant/common/upload_img?access_token=" +
		accesstoken +
		"&filename=" +
		url.QueryEscape(filename)
}

// =============================================================================

// https://api.weixin.qq.com/merchant/create?access_token=ACCESS_TOKEN
func clientMerchantProductAddURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/create?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/del?access_token=ACCESS_TOKEN
func clientMerchantProductDeleteURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/del?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/update?access_token=ACCESS_TOKEN
func clientMerchantProductUpdateURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/update?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/get?access_token=ACCESS_TOKEN
func clientMerchantProductGetURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/get?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/getbystatus?access_token=ACCESS_TOKEN
func clientMerchantProductGetByStatusURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/getbystatus?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/modproductstatus?access_token=ACCESS_TOKEN
func clientMerchantProductModifyStatusURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/modproductstatus?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/category/getsub?access_token=ACCESS_TOKEN
func clientMerchantCategoryGetSubURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/category/getsub?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/category/getsku?access_token=ACCESS_TOKEN
func clientMerchantCategoryGetSKUURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/category/getsku?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/category/getproperty?access_token=ACCESS_TOKEN
func clientMerchantCategoryGetPropertyURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/category/getproperty?access_token=" +
		accesstoken
}

// =============================================================================

// https://api.weixin.qq.com/merchant/stock/add?access_token=ACCESS_TOKEN
func clientMerchantStockAddURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/stock/add?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/stock/reduce?access_token=ACCESS_TOKEN
func clientMerchantStockReduceURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/stock/reduce?access_token=" +
		accesstoken
}

// =============================================================================

// https://api.weixin.qq.com/merchant/express/add?access_token=ACCESS_TOKEN
func clientMerchantExpressAddURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/express/add?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/express/del?access_token=ACCESS_TOKEN
func clientMerchantExpressDeleteURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/express/del?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/express/update?access_token=ACCESS_TOKEN
func clientMerchantExpressUpdateURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/express/update?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/express/getbyid?access_token=ACCESS_TOKEN
func clientMerchantExpressGetByIdURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/express/getbyid?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/express/getall?access_token=ACCESS_TOKEN
func clientMerchantExpressGetAllURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/express/getall?access_token=" +
		accesstoken
}

// =============================================================================

// https://api.weixin.qq.com/merchant/group/add?access_token=ACCESS_TOKEN
func clientMerchantGroupAddURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/group/add?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/group/del?access_token=ACCESS_TOKEN
func clientMerchantGroupDeleteURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/group/del?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/group/propertymod?access_token=ACCESS_TOKEN
func clientMerchantGroupPropertyModifyURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/group/propertymod?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/group/productmod?access_token=ACCESS_TOKEN
func clientMerchantGroupProductModifyURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/group/productmod?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/group/getall?access_token=ACCESS_TOKEN
func clientMerchantGroupGetAllURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/group/getall?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/group/getbyid?access_token=ACCESS_TOKEN
func clientMerchantGroupGetByIdURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/group/getbyid?access_token=" +
		accesstoken
}

// =============================================================================

// https://api.weixin.qq.com/merchant/shelf/add?access_token=ACCESS_TOKEN
func clientMerchantShelfAddURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/shelf/add?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/shelf/del?access_token=ACCESS_TOKEN
func clientMerchantShelfDeleteURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/shelf/del?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/shelf/mod?access_token=ACCESS_TOKEN
func clientMerchantShelfModifyURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/shelf/mod?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/shelf/getall?access_token=ACCESS_TOKEN
func clientMerchantShelfGetAllURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/shelf/getall?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/shelf/getbyid?access_token=ACCESS_TOKEN
func clientMerchantShelfGetByIdURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/shelf/getbyid?access_token=" +
		accesstoken
}

// =============================================================================

// https://api.weixin.qq.com/merchant/order/getbyid?access_token=ACCESS_TOKEN
func clientMerchantOrderGetByIdURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/order/getbyid?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/order/getbyfilter?access_token=ACCESS_TOKEN
func clientMerchantOrderGetByFilterURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/order/getbyfilter?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/order/setdelivery?access_token=ACCESS_TOKEN
func clientMerchantOrderSetDeliveryURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/order/setdelivery?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/merchant/order/close?access_token=ACCESS_TOKEN
func clientMerchantOrderCloseURL(accesstoken string) string {
	return "https://api.weixin.qq.com/merchant/order/close?access_token=" +
		accesstoken
}
