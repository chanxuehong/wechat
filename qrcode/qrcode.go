package qrcode

// 永久二维码
type QRLimitCode struct {
	SceneId int    `json:"scene_id"`
	Ticket  string `json:"ticket"`
}

// 临时二维码
type QRCode struct {
	QRLimitCode
	ExpireSeconds int `json:"expire_seconds"`
}
