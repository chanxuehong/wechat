// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay3

// 下载历史交易清单 失败时返回参数
type DownloadBillFailResponse struct {
	XMLName struct{} `xml:"xml" json:"-"`

	RetCode string `xml:"return_code"           json:"return_code"`          // 必须, FAIL
	RetMsg  string `xml:"return_msg,omitempty"  json:"return_msg,omitempty"` // 返回信息，如非空，为错误原因：签名失败、参数格式校验错误、该日期订单未生成
}
