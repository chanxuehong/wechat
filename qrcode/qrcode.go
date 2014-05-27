package qrcode

type QRCode struct {
	SceneId int    `json:"scene_id"`
	Ticket  string `json:"ticket"`

	// ExpireSeconds == 0 则表示永久二维码
	ExpireSeconds int `json:"expire_seconds,omitempty"`
}
