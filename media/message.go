package media

// 上传媒体成功时的回复报文
type UploadResponse struct {
	MediaType string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

// 上传图文消息里的 item
type UploadNewsMsgArticle struct {
	ThumbMediaId     string `json:"thumb_media_id"`               // 图文消息缩略图的media_id，可以在基础支持-上传多媒体文件接口中获得
	Author           string `json:"author,omitempty"`             // 图文消息的作者
	Title            string `json:"title"`                        // 图文消息的标题
	ContentSourceUrl string `json:"content_source_url,omitempty"` // 在图文消息页面点击“阅读原文”后的页面
	Content          string `json:"content"`                      // 图文消息页面的内容，支持HTML标签
	Digest           string `json:"digest,,omitempty"`            // 图文消息的描述
}

// 上传图文消息
type UploadNewsMsg struct {
	Articles []*UploadNewsMsgArticle `json:"articles"` // 图文消息，一个图文消息支持1到10条图文
}

// 上传视频消息
type UploadVideoMsg struct {
	MediaId     string `json:"media_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
