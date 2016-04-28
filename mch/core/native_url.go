package core

import (
	"net/url"
)

// 扫码原生支付模式1的地址
func NativeURL1(appId, mchId, productId, timestamp, nonceStr, apiKey string) string {
	m := make(map[string]string, 5)
	m["appid"] = appId
	m["mch_id"] = mchId
	m["product_id"] = productId
	m["time_stamp"] = timestamp
	m["nonce_str"] = nonceStr

	return "weixin://wxpay/bizpayurl?sign=" + Sign(m, apiKey, nil) +
		"&appid=" + url.QueryEscape(appId) +
		"&mch_id=" + url.QueryEscape(mchId) +
		"&product_id=" + url.QueryEscape(productId) +
		"&time_stamp=" + url.QueryEscape(timestamp) +
		"&nonce_str=" + url.QueryEscape(nonceStr)
}
