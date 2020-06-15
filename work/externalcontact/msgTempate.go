package externalcontact

type TextMessage struct {
	Content string `json:"content"`
}

type ImageMessage struct {
	MediaId string `json:"media_id,omitempty"`
	PicUrl  string `json:"pic_url,omitempty"`
}

type LinkMessage struct {
	Title  string `json:"title"`
	PicUrl string `json:"picurl,omitempty"`
	Desc   string `json:"desc,omitempty"`
	Url    string `json:"url"`
}

type MiniProgramMessage struct {
	Title      string `json:"title"`
	PicMediaId string `json:"pic_media_id"`
	AppId      string `json:"appid"`
	Page       string `json:"page"`
}
