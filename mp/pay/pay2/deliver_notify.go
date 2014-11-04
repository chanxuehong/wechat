// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"strconv"
	"time"
)

// 为了更好地跟踪订单的情况，需要第三方在收到最终支付通知之后，调用发货通知API
// 告知微信后台该订单的发货状态。
// 发货时间限制：虚拟、服务类24小时内，实物类72小时内。
// 请在收到支付通知后，按时发货，并使用发货通知接口将相关信息同步到微信后台。若
// 平台在规定时间内没有收到，将视作发货超时处理。
//
// 发货通知的真正的数据是放在PostData 中的，格式为json
type DeliverNotifyData map[string]string

func (data DeliverNotifyData) SetAppId(str string) {
	data["appid"] = str
}
func (data DeliverNotifyData) SetOpenId(str string) {
	data["openid"] = str
}
func (data DeliverNotifyData) SetTransactionId(str string) {
	data["transid"] = str
}
func (data DeliverNotifyData) SetOutTradeNo(str string) {
	data["out_trade_no"] = str
}
func (data DeliverNotifyData) SetDeliverTimeStamp(t time.Time) {
	data["deliver_timestamp"] = strconv.FormatInt(t.Unix(), 10)
}
func (data DeliverNotifyData) SetDeliverStatus(status int) {
	data["deliver_status"] = strconv.FormatInt(int64(status), 10)
}
func (data DeliverNotifyData) SetDeliverMsg(str string) {
	data["deliver_msg"] = str
}
func (data DeliverNotifyData) SetSignMethod(str string) {
	data["sign_method"] = str
}

// 设置签名字段.
//  appKey: 即 paySignKey, 公众号支付请求中用于加密的密钥 Key
//
//  NOTE: 要求在 DeliverNotifyData 其他字段设置完毕后才能调用这个函数, 否则签名就不正确.
func (data DeliverNotifyData) SetSignature(appKey string) (err error) {
	return OrderQueryRequest(data).SetSignature(appKey)
}
