// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package qrcode

import (
	"net/url"
)

// 永久二维码
type PermanentQRCode struct {
	SceneId uint32 `json:"scene_id"` // 场景值 id, 目前参数只支持1--100000
	Ticket  string `json:"ticket"`   // 二维码ticket, 凭借此ticket可以在有效时间内换取二维码.
}

// 二维码的 URL, 可以 GET 此 URL 下载二维码
func (qrcode *PermanentQRCode) URL() string {
	return "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=" +
		url.QueryEscape(qrcode.Ticket)
}

// 临时二维码
type TemporaryQRCode struct {
	SceneId   uint32 `json:"scene_id"`       // 场景值 id, 32位非0整型
	Ticket    string `json:"ticket"`         // 二维码ticket, 凭借此ticket可以在有效时间内换取二维码.
	ExpiresIn int    `json:"expire_seconds"` // 有效期, 单位为"秒"
}

// 二维码的 URL, 可以 GET 此 URL 下载二维码
func (qrcode *TemporaryQRCode) URL() string {
	return "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=" +
		url.QueryEscape(qrcode.Ticket)
}
