package media

// 上传媒体成功时的回复报文
type UploadResponse struct {
	// 媒体文件类型，分别有图片（image）、语音（voice）、视频（video）和
	// 缩略图（thumb，主要用于视频与音乐格式的缩略图），次数为news，即图文消息
	MediaType string `json:"type"`
	MediaId   string `json:"media_id"`   // 媒体文件上传后，获取时的唯一标识
	CreatedAt string `json:"created_at"` // 媒体文件上传时间戳
}

// 上传图文消息里的 item
type NewsArticle struct {
	ThumbMediaId     string `json:"thumb_media_id"`               // 图文消息缩略图的media_id，可以在基础支持-上传多媒体文件接口中获得
	Author           string `json:"author,omitempty"`             // 图文消息的作者
	Title            string `json:"title"`                        // 图文消息的标题
	ContentSourceUrl string `json:"content_source_url,omitempty"` // 在图文消息页面点击“阅读原文”后的页面
	Content          string `json:"content"`                      // 图文消息页面的内容，支持HTML标签
	Digest           string `json:"digest,,omitempty"`            // 图文消息的描述
	ShowCoverPic     string `json:"show_cover_pic"`               // 是否显示封面，1为显示，0为不显示
}

// 上传图文消息
type News struct {
	// Article 的个数不能超过 NewsArticleCountLimit
	Articles []*NewsArticle `json:"articles"` // 图文消息，一个图文消息支持1到10条图文
}

// 如果总的按钮数超过限制, 则截除多余的.
func (news *News) AppendArticle(article ...*NewsArticle) {
	if len(article) <= 0 {
		return
	}

	switch n := NewsArticleCountLimit - len(news.Articles); {
	case n > 0:
		if len(article) > n {
			article = article[:n]
		}
		news.Articles = append(news.Articles, article...)
	case n == 0:
	default: // n < 0
		news.Articles = news.Articles[:NewsArticleCountLimit]
	}
}

// 上传视频消息
type Video struct {
	MediaId     string `json:"media_id"` // 此处media_id需通过基础支持中的上传下载多媒体文件来得到
	Title       string `json:"title"`
	Description string `json:"description"`
}
