package media

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"

	"github.com/chanxuehong/wechat/mp/core"
)

const (
	NewsArticleCountLimit = 10 // 图文素材里文章的个数限制
)

const (
	MediaTypeImage = "image"
	MediaTypeVoice = "voice"
	MediaTypeVideo = "video"
	MediaTypeThumb = "thumb"
	MediaTypeNews  = "news"
)

type MediaInfo struct {
	MediaType string `json:"type"`       // 图片(image), 语音(voice), 视频(video), 缩略图(thumb)和 图文消息(news)
	MediaId   string `json:"media_id"`   // 媒体的唯一标识
	CreatedAt int64  `json:"created_at"` // 媒体创建的时间戳
}

type Article struct {
	ThumbMediaId     string `json:"thumb_media_id"`               // 必须; 图文消息缩略图的 media_id, 可以在上传多媒体文件接口中获得
	Title            string `json:"title"`                        // 必须; 图文消息的标题
	Author           string `json:"author,omitempty"`             // 可选; 图文消息的作者
	Digest           string `json:"digest,omitempty"`             // 可选; 图文消息的摘要
	Content          string `json:"content"`                      // 必须; 图文消息页面的内容, 支持HTML标签
	ContentSourceURL string `json:"content_source_url,omitempty"` // 可选; 在图文消息页面点击"阅读原文"后的页面
	ShowCoverPic     int    `json:"show_cover_pic"`               // 必须; 是否显示封面, 1为显示, 0为不显示
}

func (article *Article) SetShowCoverPic(b bool) {
	if b {
		article.ShowCoverPic = 1
	} else {
		article.ShowCoverPic = 0
	}
}

// 下载多媒体到文件.
//  请注意, 视频文件不支持下载
func DownloadMedia(clt *core.Client, mediaId, filepath string) (written int64, err error) {
	file, err := os.Create(filepath)
	if err != nil {
		return
	}
	defer func() {
		file.Close()
		if err != nil {
			os.Remove(filepath)
		}
	}()

	return downloadMediaToWriter(clt, mediaId, file)
}

// 下载多媒体到 io.Writer.
//  请注意, 视频文件不支持下载
func DownloadMediaToWriter(clt *core.Client, mediaId string, writer io.Writer) (written int64, err error) {
	if writer == nil {
		err = errors.New("nil writer")
		return
	}
	return downloadMediaToWriter(clt, mediaId, writer)
}

// 下载多媒体到 io.Writer.
func downloadMediaToWriter(clt *core.Client, mediaId string, writer io.Writer) (written int64, err error) {
	token, err := clt.Token()
	if err != nil {
		return
	}

	hasRetried := false
RETRY:
	finalURL := "https://api.weixin.qq.com/cgi-bin/media/get?media_id=" + url.QueryEscape(mediaId) +
		"&access_token=" + url.QueryEscape(token)

	httpResp, err := clt.HttpClient.Get(finalURL)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	ContentType, _, _ := mime.ParseMediaType(httpResp.Header.Get("Content-Type"))
	if ContentType != "text/plain" && ContentType != "application/json" { // 返回的是媒体流
		return io.Copy(writer, httpResp.Body)
	}

	// 返回的是错误信息
	var result core.Error
	if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
		return
	}

	switch result.ErrCode {
	case core.ErrCodeOK:
		return // 基本不会出现
	case core.ErrCodeInvalidCredential, core.ErrCodeAccessTokenExpired: // 失效(过期)重试一次
		if !hasRetried {
			hasRetried = true

			if token, err = clt.TokenRefresh(); err != nil {
				return
			}

			result = core.Error{}
			goto RETRY
		}
		fallthrough
	default:
		err = &result
		return
	}
}

// 创建图文消息素材.
func CreateNews(clt *core.Client, articles []Article) (info *MediaInfo, err error) {
	if len(articles) <= 0 {
		err = errors.New("图文素材是空的")
		return
	}
	if len(articles) > NewsArticleCountLimit {
		err = fmt.Errorf("图文素材的文章个数不能超过 %d, 现在为 %d", NewsArticleCountLimit, len(articles))
		return
	}

	var request = struct {
		Articles []Article `json:"articles,omitempty"`
	}{
		Articles: articles,
	}

	var result struct {
		core.Error
		MediaInfo
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/media/uploadnews?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	info = &result.MediaInfo
	return
}

// 创建视频素材.
//  mediaId:     通过上传视频文件得到
//  title:       标题, 可以为空
//  description: 描述, 可以为空
func CreateVideo(clt *core.Client, mediaId, title, description string) (info *MediaInfo, err error) {
	if mediaId == "" {
		err = errors.New("empty mediaId")
		return
	}
	var request = struct {
		MediaId     string `json:"media_id"`
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
	}{
		MediaId:     mediaId,
		Title:       title,
		Description: description,
	}

	var result struct {
		core.Error
		MediaInfo
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/media/uploadvideo?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	info = &result.MediaInfo
	return
}
