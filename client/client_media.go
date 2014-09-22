// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chanxuehong/wechat/media"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// 上传多媒体图片
func (c *Client) MediaUploadImage(filepath_ string) (info *media.MediaInfo, err error) {
	return c.mediaUpload(media.MEDIA_TYPE_IMAGE, filepath_)
}

// 上传多媒体语音
func (c *Client) MediaUploadVoice(filepath_ string) (info *media.MediaInfo, err error) {
	return c.mediaUpload(media.MEDIA_TYPE_VOICE, filepath_)
}

// 上传多媒体视频
func (c *Client) MediaUploadVideo(filepath_ string) (info *media.MediaInfo, err error) {
	return c.mediaUpload(media.MEDIA_TYPE_VIDEO, filepath_)
}

// 上传多媒体缩略图
func (c *Client) MediaUploadThumb(filepath_ string) (info *media.MediaInfo, err error) {
	return c.mediaUpload(media.MEDIA_TYPE_THUMB, filepath_)
}

// 上传多媒体
func (c *Client) mediaUpload(mediaType, filepath_ string) (info *media.MediaInfo, err error) {
	file, err := os.Open(filepath_)
	if err != nil {
		return
	}
	defer file.Close()

	return c.mediaUploadFromReader(mediaType, filepath.Base(filepath_), file)
}

// 上传多媒体图片
//  NOTE: 参数 filename 不是文件路径, 是指定 multipart form 里面文件名称
func (c *Client) MediaUploadImageFromReader(filename string, mediaReader io.Reader) (info *media.MediaInfo, err error) {
	if filename == "" {
		err = errors.New(`filename == ""`)
		return
	}
	if mediaReader == nil {
		err = errors.New("mediaReader == nil")
		return
	}
	return c.mediaUploadFromReader(media.MEDIA_TYPE_IMAGE, filename, mediaReader)
}

// 上传多媒体语音
//  NOTE: 参数 filename 不是文件路径, 是指定 multipart form 里面文件名称
func (c *Client) MediaUploadVoiceFromReader(filename string, mediaReader io.Reader) (info *media.MediaInfo, err error) {
	if filename == "" {
		err = errors.New(`filename == ""`)
		return
	}
	if mediaReader == nil {
		err = errors.New("mediaReader == nil")
		return
	}
	return c.mediaUploadFromReader(media.MEDIA_TYPE_VOICE, filename, mediaReader)
}

// 上传多媒体视频
//  NOTE: 参数 filename 不是文件路径, 是指定 multipart form 里面文件名称
func (c *Client) MediaUploadVideoFromReader(filename string, mediaReader io.Reader) (info *media.MediaInfo, err error) {
	if filename == "" {
		err = errors.New(`filename == ""`)
		return
	}
	if mediaReader == nil {
		err = errors.New("mediaReader == nil")
		return
	}
	return c.mediaUploadFromReader(media.MEDIA_TYPE_VIDEO, filename, mediaReader)
}

// 上传多媒体缩略图
//  NOTE: 参数 filename 不是文件路径, 是指定 multipart form 里面文件名称
func (c *Client) MediaUploadThumbFromReader(filename string, mediaReader io.Reader) (info *media.MediaInfo, err error) {
	if filename == "" {
		err = errors.New(`filename == ""`)
		return
	}
	if mediaReader == nil {
		err = errors.New("mediaReader == nil")
		return
	}
	return c.mediaUploadFromReader(media.MEDIA_TYPE_THUMB, filename, mediaReader)
}

// 上传多媒体
func (c *Client) mediaUploadFromReader(mediaType, filename string, mediaReader io.Reader) (info *media.MediaInfo, err error) {
	bodyBuf := mediaBufferPool.Get().(*bytes.Buffer) // io.ReadWriter
	bodyBuf.Reset()                                  // important
	defer mediaBufferPool.Put(bodyBuf)

	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		return
	}
	if _, err = io.Copy(fileWriter, mediaReader); err != nil {
		return
	}

	bodyContentType := bodyWriter.FormDataContentType()

	if err = bodyWriter.Close(); err != nil {
		return
	}

	postContent := bodyBuf.Bytes() // 这么绕一下是为了 RETRY 的时候不会出错

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := mediaUploadURL(token, mediaType)

	httpResp, err := c.httpClient.Post(_url, bodyContentType, bytes.NewReader(postContent))
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	switch mediaType {
	case media.MEDIA_TYPE_THUMB: // 返回的是 thumb_media_id 而不是 media_id
		var result struct {
			MediaType string `json:"type"`
			MediaId   string `json:"thumb_media_id"`
			CreatedAt int64  `json:"created_at"`
			Error
		}
		if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
			return
		}

		switch result.ErrCode {
		case errCodeOK:
			info = &media.MediaInfo{
				MediaType: result.MediaType,
				MediaId:   result.MediaId,
				CreatedAt: result.CreatedAt,
			}
			return

		case errCodeTimeout:
			if !hasRetry {
				hasRetry = true
				timeoutRetryWait()
				goto RETRY
			}
			fallthrough

		default:
			err = &result.Error
			return
		}

	default:
		var result struct {
			media.MediaInfo
			Error
		}
		if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
			return
		}

		switch result.ErrCode {
		case errCodeOK:
			info = &result.MediaInfo
			return

		case errCodeTimeout:
			if !hasRetry {
				hasRetry = true
				timeoutRetryWait()
				goto RETRY
			}
			fallthrough

		default:
			err = &result.Error
			return
		}
	}
}

// 下载多媒体文件
func (c *Client) MediaDownload(mediaId, filepath_ string) (err error) {
	file, err := os.Create(filepath_)
	if err != nil {
		return
	}
	defer file.Close()

	return c.mediaDownloadToWriter(mediaId, file)
}

// 下载多媒体文件
func (c *Client) MediaDownloadToWriter(mediaId string, writer io.Writer) error {
	if writer == nil {
		return errors.New("writer == nil")
	}
	return c.mediaDownloadToWriter(mediaId, writer)
}

// 下载多媒体文件.
func (c *Client) mediaDownloadToWriter(mediaId string, writer io.Writer) (err error) {
	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := mediaDownloadURL(token, mediaId)

	httpResp, err := c.httpClient.Get(_url)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}

	contentType, _, _ := mime.ParseMediaType(httpResp.Header.Get("Content-Type"))
	if contentType != "text/plain" && contentType != "application/json" {
		_, err = io.Copy(writer, httpResp.Body)
		return
	}

	// 返回的是错误信息
	var result Error
	if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = &result
		return
	}
}

// 根据上传的缩略图媒体创建图文消息素材
//  articles 的长度不能大于 media.NewsArticleCountLimit
func (c *Client) MediaCreateNews(articles []media.NewsArticle) (info *media.MediaInfo, err error) {
	if len(articles) == 0 {
		err = errors.New("图文消息是空的")
		return
	}
	if len(articles) > media.NewsArticleCountLimit {
		err = fmt.Errorf("图文消息的文章个数不能超过 %d, 现在为 %d", media.NewsArticleCountLimit, len(articles))
		return
	}

	var request = struct {
		Articles []media.NewsArticle `json:"articles,omitempty"`
	}{
		Articles: articles,
	}

	var result struct {
		media.MediaInfo
		Error
	}

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := mediaCreateNewsURL(token)
	if err = c.postJSON(_url, request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		info = &result.MediaInfo
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = &result.Error
		return
	}
}

// 根据上传的视频文件 media_id 创建视频媒体, 群发视频消息应该用这个函数得到的 media_id.
//  NOTE: title, description 可以为空
func (c *Client) MediaCreateVideo(mediaId, title, description string) (info *media.MediaInfo, err error) {
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
		media.MediaInfo
		Error
	}

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := mediaCreateVideoURL(token)
	if err = c.postJSON(_url, &request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		info = &result.MediaInfo
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = &result.Error
		return
	}
}
