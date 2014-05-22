// 高级群发接口

package message

// 上传图文消息里的 item
type NewsUploadArticle struct {
	ThumbMediaId     string `json:"thumb_media_id"`               // 图文消息缩略图的media_id，可以在基础支持-上传多媒体文件接口中获得
	Author           string `json:"author,omitempty"`             // 图文消息的作者
	Title            string `json:"title"`                        // 图文消息的标题
	ContentSourceUrl string `json:"content_source_url,omitempty"` // 在图文消息页面点击“阅读原文”后的页面
	Content          string `json:"content"`                      // 图文消息页面的内容，支持HTML标签
	Digest           string `json:"digest,,omitempty"`            // 图文消息的描述
}

// 上传图文消息
type NewsUploadMsg struct {
	Articles []*NewsUploadArticle `json:"articles"` // 图文消息，一个图文消息支持1到10条图文
}

type NewsMassByGroupMsg struct {
	Filter struct {
		GroupId string `json:"group_id"`
	} `json:"filter"`
	MPNews struct {
		MediaId string `json:"media_id"`
	} `json:"mpnews"`
	MsgType string `json:"msgtype"`
}

type TextMassByGroupMsg struct {
	Filter struct {
		GroupId string `json:"group_id"`
	} `json:"filter"`
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
	MsgType string `json:"msgtype"`
}

type VoiceMassByGroupMsg struct {
	Filter struct {
		GroupId string `json:"group_id"`
	} `json:"filter"`
	Voice struct {
		MediaId string `json:"media_id"`
	} `json:"voice"`
	MsgType string `json:"msgtype"`
}

type ImageMassByGroupMsg struct {
	Filter struct {
		GroupId string `json:"group_id"`
	} `json:"filter"`
	Image struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
	MsgType string `json:"msgtype"`
}

type VideoMassByGroupMsg struct {
	Filter struct {
		GroupId string `json:"group_id"`
	} `json:"filter"`
	MPVideo struct {
		MediaId string `json:"media_id"`
	} `json:"mpvideo"`
	MsgType string `json:"msgtype"`
}

type NewsMassByOpenIdListMsg struct {
	Touser []string `json:"touser"`
	MPNews struct {
		MediaId string `json:"media_id"`
	} `json:"mpnews"`
	MsgType string `json:"msgtype"`
}

type TextMassByOpenIdListMsg struct {
	Touser []string `json:"touser"`
	Text   struct {
		Content string `json:"content"`
	} `json:"text"`
	MsgType string `json:"msgtype"`
}

type VoiceMassByOpenIdListMsg struct {
	Touser []string `json:"touser"`
	Voice  struct {
		MediaId string `json:"media_id"`
	} `json:"voice"`
	MsgType string `json:"msgtype"`
}

type ImageMassByOpenIdListMsg struct {
	Touser []string `json:"touser"`
	Image  struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
	MsgType string `json:"msgtype"`
}

type VideoMassByOpenIdListMsg struct {
	Touser []string `json:"touser"`
	Video  struct {
		MediaId     string `json:"media_id"`
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"video"`
	MsgType string `json:"msgtype"`
}
