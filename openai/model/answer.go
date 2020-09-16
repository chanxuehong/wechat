package model

type Answer struct {
	News        *News        `json:"news,omitempty"`
	Image       *Image       `json:"image,omitempty"`
	MiniProgram *MiniProgram `json:"miniprogram,omitempty"`
	Multimsg    []Answer     `json:"multimsg,omitempty"`
}

type News struct {
	Articles []Article `json:"articles,omitempty"`
}

type ArticleType = string

const (
	H5 ArticleType = "h5"
	MP ArticleType = "mp"
)

type Article struct {
	Title string      `json:"title,omitempty"`
	Desc  string      `json:"description,omitempty"`
	Url   string      `json:"url,omitempty"`
	Pic   string      `json:"picurl,omitempty"`
	Type  ArticleType `json:"type,omitempty"`
}

type Image struct {
	MediaId string `json:"media_id,omitempty"`
	Url     string `json:"url,omitempty"`
}

type MiniProgram struct {
	Title        string `json:"title,omitempty"`
	AppID        string `json:"appid,omitempty"`
	PagePath     string `json:"pagepath,omitempty"`
	ThumbMediaID string `json:"thumb_media_id,omitempty"`
	ThumbUrl     string `json:"thumb_url,omitempty"`
}
