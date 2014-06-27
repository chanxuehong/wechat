// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong@gmail.com

package qrcode

const (
	// 临时二维码 expire seconds 限制
	TemporaryQRCodeExpireSecondsLimit = 1800
	// 永久二维码 scene id 限制
	PermanentQRCodeSceneIdLimit = 100000
)
