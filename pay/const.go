// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

const (
	IS_SUBSCRIBE_TRUE  = 1
	IS_SUBSCRIBE_FALSE = 0

	BANK_TYPE_WX = "WX"
	FEE_TYPE_RMB = 1
)
const (
	BILL_CHARSET_GBK  = "GBK"
	BILL_CHARSET_UTF8 = "UTF-8"
)

const (
	NOTIFICATION_TRADE_MODE_IMMEDIATE = 1 // TradeMode 即时到账
	NOTIFICATION_TRADE_SUCCESS        = 0 // TradeState 成功
)
