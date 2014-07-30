// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package media

import (
	"errors"
	"fmt"
)

// 上传媒体成功时的回复报文
type UploadResponse struct {
	// 媒体文件类型，分别有图片（image）、语音（voice）、视频（video）、
	// 缩略图（thumb，主要用于视频与音乐格式的缩略图）和 图文消息（news）
	MediaType string `json:"type"`

	// 媒体文件上传后，获取时的唯一标识.
	//  NOTE:
	//  1. 每个多媒体文件（media_id）会在上传、用户发送到微信服务器3天后自动删除，以节省服务器资源。
	//  2. media_id是可复用的.
	MediaId string `json:"media_id"`
	// 媒体文件上传时间戳
	CreatedAt int64 `json:"created_at"`
}

// 图文消息里的 Article
type NewsArticle struct {
	ThumbMediaId     string `json:"thumb_media_id"`                  // 图文消息缩略图的media_id，可以在基础支持-上传多媒体文件接口中获得
	Author           string `json:"author,omitempty"`                // 图文消息的作者
	Title            string `json:"title"`                           // 图文消息的标题
	ContentSourceURL string `json:"content_source_url,omitempty"`    // 在图文消息页面点击“阅读原文”后的页面
	Content          string `json:"content"`                         // 图文消息页面的内容，支持HTML标签
	Digest           string `json:"digest,omitempty"`                // 图文消息的描述
	ShowCoverPic     int    `json:"show_cover_pic,string,omitempty"` // 是否显示封面，1为显示，0为不显示
}

// 图文消息
type News struct {
	Articles []NewsArticle `json:"articles,omitempty"` // 图文消息，一个图文消息支持1到10条图文
}

// 检查 News 是否有效，有效返回 nil，否则返回错误信息
func (n *News) CheckValid() (err error) {
	articleNum := len(n.Articles)
	if articleNum == 0 {
		err = errors.New("图文消息是空的")
		return
	}
	if articleNum > NewsArticleCountLimit {
		err = fmt.Errorf("图文消息的文章个数不能超过 %d, 现在为 %d", NewsArticleCountLimit, articleNum)
		return
	}
	return
}
