// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chanxuehong/wechat/corp/media"
	"io"
	"mime"
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

// 上传多媒体缩略图（目前文档还没有）
func (c *Client) MediaUploadThumb(filepath_ string) (info *media.MediaInfo, err error) {
	return c.mediaUpload(media.MEDIA_TYPE_THUMB, filepath_)
}

// 上传普通文件
func (c *Client) MediaUploadFile(filepath_ string) (info *media.MediaInfo, err error) {
	return c.mediaUpload(media.MEDIA_TYPE_FILE, filepath_)
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

// 上传多媒体缩略图（目前文档还没有）
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

// 上传普通文件
//  NOTE: 参数 filename 不是文件路径, 是指定 multipart form 里面文件名称
func (c *Client) MediaUploadFileFromReader(filename string, mediaReader io.Reader) (info *media.MediaInfo, err error) {
	if filename == "" {
		err = errors.New(`filename == ""`)
		return
	}
	if mediaReader == nil {
		err = errors.New("mediaReader == nil")
		return
	}
	return c.mediaUploadFromReader(media.MEDIA_TYPE_FILE, filename, mediaReader)
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
	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	httpResp, err := c.httpClient.Get(_MediaDownloadURL(token, mediaId))
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

			if token, err = c.TokenRefresh(); err != nil {
				return
			}
			goto RETRY
		}
		fallthrough

	default:
		err = &result
		return
	}
}
