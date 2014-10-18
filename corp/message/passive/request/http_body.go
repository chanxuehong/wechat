// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package request

// 微信服务器请求 http body
type RequestHttpBody struct {
	XMLName      struct{} `xml:"xml" json:"-"`
	CorpId       string   `xml:"ToUserName"`
	AgentId      int64    `xml:"AgentID"`
	EncryptedMsg string   `xml:"Encrypt"`
}
