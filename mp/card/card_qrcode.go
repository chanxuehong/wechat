// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     gaowenbin(gaowenbinmarr@gmail.com)
package card

// 卡券二维码
type CardQRCode struct {
	CardId        string `json:"card_id"` // 卡券ID
	Code          string `json:"code"`    // 指定卡券code码
	OpenID        string `json:"openid"`
	ExpireSeconds string `json:"expire_seconds"`
	IsUniqueCode  bool   `json:"is_unique_code "`
	OuterId       int    `json:"outer_id"`
}
