// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package response

// 安全模式, 回复微信请求的 http body
type ResponseHttpBody struct {
	XMLName    struct{} `xml:"xml" json:"-"`
	EncryptMsg string   `xml:"Encrypt"`
	Signature  string   `xml:"MsgSignature"`
	TimeStamp  int64    `xml:"TimeStamp"`
	Nonce      string   `xml:"Nonce"`
}
