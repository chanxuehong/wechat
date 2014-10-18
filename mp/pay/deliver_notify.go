// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

// 为了更好地跟踪订单的情况, 需要第三方在收到最终支付通知之后, 调用发货通知API告知微信后台该订单的发货状态.
// 发货时间限制: 虚拟, 服务类24小时内, 实物类72小时内.
//
// 请在收到支付通知后, 按时发货, 并使用发货通知接口将相关信息同步到微信后台.
// 若平台在规定时间内没有收到, 将视作发货超时处理.
//
// 发货通知的的数据是放在PostData中的, 格式为 JSON.
type DeliverNotifyData struct {
	AppId            string `json:"appid"`                    // 公众平台账户的 AppId
	OpenId           string `json:"openid"`                   // 购买用户的 OpenId, 这个已经放在最终支付结果通知的 PostData 里了
	TransactionId    string `json:"transid"`                  // 交易单号
	OutTradeNo       string `json:"out_trade_no"`             // 第三方订单号
	DeliverTimeStamp int64  `json:"deliver_timestamp,string"` // 发货时间戳, unixtime;
	DeliverStatus    int    `json:"deliver_status,string"`    // 发货状态, 1表明成功, 0表明失败, 失败时需要在deliver_msg填上失败原因;
	DeliverMessage   string `json:"deliver_msg"`              // 发货状态信息, 失败时可以填上UTF8编码的错误提示信息, 比如"该商品已退款";
	Signature        string `json:"app_signature"`            // 签名
	SignMethod       string `json:"sign_method"`              // 签名方法
}

// 设置签名字段.
//  @paySignKey: 公众号支付请求中用于加密的密钥 Key, 对应于支付场景中的 appKey
//  NOTE: 要求在 data *DeliverNotifyData 其他字段设置完毕后才能调用这个函数, 否则签名就不正确.
func (data *DeliverNotifyData) SetSignature(paySignKey string) (err error) {
	var sumFunc hashSumFunc

	switch {
	case data.SignMethod == SIGN_METHOD_SHA1:
		sumFunc = sha1Sum

	case strings.ToLower(data.SignMethod) == "sha1":
		data.SignMethod = SIGN_METHOD_SHA1
		sumFunc = sha1Sum

	default:
		err = fmt.Errorf(`not implement for "%s" sign method`, data.SignMethod)
		return
	}

	DeliverTimeStampStr := strconv.FormatInt(data.DeliverTimeStamp, 10)
	DeliverStatusStr := strconv.FormatInt(int64(data.DeliverStatus), 10)

	const keysLen = len(`appid=&appkey=&deliver_msg=&deliver_status=&deliver_timestamp=&openid=&out_trade_no=&transid=`)
	n := keysLen + len(data.AppId) + len(paySignKey) + len(data.DeliverMessage) +
		len(DeliverStatusStr) + len(DeliverTimeStampStr) + len(data.OpenId) +
		len(data.OutTradeNo) + len(data.TransactionId)

	string1 := make([]byte, 0, n)

	// 字典序
	// appid
	// appkey
	// deliver_msg
	// deliver_status
	// deliver_timestamp
	// openid
	// out_trade_no
	// transid
	string1 = append(string1, "appid="...)
	string1 = append(string1, data.AppId...)
	string1 = append(string1, "&appkey="...)
	string1 = append(string1, paySignKey...)
	string1 = append(string1, "&deliver_msg="...)
	string1 = append(string1, data.DeliverMessage...)
	string1 = append(string1, "&deliver_status="...)
	string1 = append(string1, DeliverStatusStr...)
	string1 = append(string1, "&deliver_timestamp="...)
	string1 = append(string1, DeliverTimeStampStr...)
	string1 = append(string1, "&openid="...)
	string1 = append(string1, data.OpenId...)
	string1 = append(string1, "&out_trade_no="...)
	string1 = append(string1, data.OutTradeNo...)
	string1 = append(string1, "&transid="...)
	string1 = append(string1, data.TransactionId...)

	data.Signature = hex.EncodeToString(sumFunc(string1))
	return
}
