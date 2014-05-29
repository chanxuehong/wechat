package qrcode

type QRCode struct {
	// 场景值ID，临时二维码时为32位非0整型，
	// 永久二维码时最大值为100000（目前参数只支持1--100000）
	SceneId int    `json:"scene_id"`
	Ticket  string `json:"ticket"` // 获取的二维码ticket，凭借此ticket可以在有效时间内换取二维码。

	// 二维码的有效时间，以秒为单位。最大不超过1800。
	// ExpireSeconds == 0 则表示永久二维码
	ExpireSeconds int `json:"expire_seconds,omitempty"`
}
