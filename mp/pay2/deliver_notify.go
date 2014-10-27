// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"hash"
	"strconv"
)

// 为了更好地跟踪订单的情况，需要第三方在收到最终支付通知之后，调用发货通知API
// 告知微信后台该订单的发货状态。
// 发货时间限制：虚拟、服务类24小时内，实物类72小时内。
// 请在收到支付通知后，按时发货，并使用发货通知接口将相关信息同步到微信后台。若
// 平台在规定时间内没有收到，将视作发货超时处理。
//
// 发货通知的真正的数据是放在PostData 中的，格式为json
type DeliverNotifyData struct {
	AppId            string `json:"appid"`                    // 公众平台账户的AppId
	OpenId           string `json:"openid"`                   // 购买用户的OpenId，这个已经放在最终支付结果通知的PostData 里了
	TransactionId    string `json:"transid"`                  // 交易单号
	OutTradeNo       string `json:"out_trade_no"`             // 第三方订单号
	DeliverTimeStamp int64  `json:"deliver_timestamp,string"` // 发货时间戳, unixtime;
	DeliverStatus    int    `json:"deliver_status,string"`    // 发货状态, 1表明成功, 0表明失败, 失败时需要在deliver_msg填上失败原因;
	DeliverMessage   string `json:"deliver_msg"`              // 发货状态信息, 失败时可以填上UTF8编码的错误提示信息, 比如"该商品已退款";
	Signature        string `json:"app_signature"`            // 签名
	SignMethod       string `json:"sign_method"`              // 签名方法
}

// 设置签名字段.
//  appKey: 即 paySignKey, 公众号支付请求中用于加密的密钥 Key
//
//  NOTE: 要求在 data *DeliverNotifyData 其他字段设置完毕后才能调用这个函数, 否则签名就不正确.
func (data *DeliverNotifyData) SetSignature(appKey string) (err error) {
	var Hash hash.Hash

	switch data.SignMethod {
	case "sha1", "SHA1":
		Hash = sha1.New()

	default:
		err = fmt.Errorf(`unknown sign method: %q`, data.SignMethod)
		return
	}

	// 字典序
	// appid
	// appkey
	// deliver_msg
	// deliver_status
	// deliver_timestamp
	// openid
	// out_trade_no
	// transid
	Hash.Write([]byte("appid="))
	Hash.Write([]byte(data.AppId))
	Hash.Write([]byte("&appkey="))
	Hash.Write([]byte(appKey))
	Hash.Write([]byte("&deliver_msg="))
	Hash.Write([]byte(data.DeliverMessage))
	Hash.Write([]byte("&deliver_status="))
	Hash.Write([]byte(strconv.FormatInt(int64(data.DeliverStatus), 10)))
	Hash.Write([]byte("&deliver_timestamp="))
	Hash.Write([]byte(strconv.FormatInt(data.DeliverTimeStamp, 10)))
	Hash.Write([]byte("&openid="))
	Hash.Write([]byte(data.OpenId))
	Hash.Write([]byte("&out_trade_no="))
	Hash.Write([]byte(data.OutTradeNo))
	Hash.Write([]byte("&transid="))
	Hash.Write([]byte(data.TransactionId))

	data.Signature = hex.EncodeToString(Hash.Sum(nil))
	return
}
