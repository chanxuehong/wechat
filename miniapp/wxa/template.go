package wxa

type TemplateDraft struct {
	Id          uint64 `json:"draft_id"`
	CreateTime  int64  `json:"create_time"`
	UserVersion string `json:"user_version"`
	UserDesc    string `json:"user_desc"`
	AppId       string `json:"source_miniprogram_appid"`
	Developer   string `json:"developer"`
}

type Template struct {
	Id          uint64 `json:"template_id"`
	CreateTime  int64  `json:"create_time"`
	UserVersion string `json:"user_version"`
	UserDesc    string `json:"user_desc"`
}
