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
	NOTIFY_URL_DATA_CHARSET_GBK          = "GBK"
	NOTIFY_URL_DATA_CHARSET_UTF8         = "UTF-8"
	NOTIFY_URL_DATA_SIGN_METHOD_MD5      = "MD5"
	NOTIFY_URL_DATA_SIGN_METHOD_RSA      = "RSA"
	NOTIFY_URL_DATA_TRADE_MODE_IMMEDIATE = 1 // TradeMode 即时到账
	NOTIFY_URL_DATA_TRADE_SUCCESS        = 0 // TradeState 成功
)

const (
	// 微信后台通过 notify_url 通知商户，商户做业务处理后，需要以字符串的形式反馈处理结果
	// success:       处理成功，微信系统收到此结果后不再进行后续通知
	// fail 或其它字符: 处理不成功，微信收到此结果或者没有收到任何结果，系统通过补单机制再次通知
	NOTIFY_RESPONSE_SUCCESS = "success"
)
