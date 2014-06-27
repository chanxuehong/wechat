// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package qrcode

// 永久二维码
type PermanentQRCode struct {
	SceneId uint32 `json:"scene_id"` // 场景值ID, 目前参数只支持1--100000
	Ticket  string `json:"ticket"`   // 获取的二维码ticket，凭借此ticket可以在有效时间内换取二维码。
}

// 临时二维码
type TemporaryQRCode struct {
	SceneId uint32 `json:"scene_id"` // 场景值ID, 32位非0整型
	Ticket  string `json:"ticket"`   // 获取的二维码ticket，凭借此ticket可以在有效时间内换取二维码。
	Expiry  int64  `json:"expiry"`   // 过期时间, unixtime
}
