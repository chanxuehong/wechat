// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay3

// 支付完成后，微信会把相关支付和用户信息发送到该 notify URL,
// 商户处理后同步返回给微信的参数
type OrderNotifyResponse struct {
	XMLName struct{} `xml:"xml" json:"-"`

	RetCode string `xml:"return_code"           json:"return_code"`          // 必须, SUCCESS/FAIL; 此字段是通信标识，非交易标识，交易是否成功需要查看result_code 来判断
	RetMsg  string `xml:"return_msg,omitempty"  json:"return_msg,omitempty"` // 可选, 返回信息，如非空，为错误原因: 签名失败, 参数格式校验错误
}
