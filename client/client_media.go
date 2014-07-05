// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
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
func (c *Client) MediaUploadImageFromFile(_filepath string) (resp *media.UploadResponse, err error) {
	return c.mediaUploadFromFile(media.MEDIA_TYPE_IMAGE, _filepath)
}

// 上传多媒体缩略图
func (c *Client) MediaUploadThumbFromFile(_filepath string) (resp *media.UploadResponse, err error) {
	return c.mediaUploadFromFile(media.MEDIA_TYPE_THUMB, _filepath)
}

// 上传多媒体语音
func (c *Client) MediaUploadVoiceFromFile(_filepath string) (resp *media.UploadResponse, err error) {
	return c.mediaUploadFromFile(media.MEDIA_TYPE_VOICE, _filepath)
}

// 上传多媒体视频
func (c *Client) MediaUploadVideoFromFile(_filepath string) (resp *media.UploadResponse, err error) {
	return c.mediaUploadFromFile(media.MEDIA_TYPE_VIDEO, _filepath)
}

// 上传多媒体
func (c *Client) mediaUploadFromFile(mediaType, _filepath string) (resp *media.UploadResponse, err error) {
	file, err := os.Open(_filepath)
	if err != nil {
		return
	}
	defer file.Close()

	return c.mediaUpload(mediaType, filepath.Base(_filepath), file)
}

// 上传多媒体图片
func (c *Client) MediaUploadImage(filename string, mediaReader io.Reader) (resp *media.UploadResponse, err error) {
	if filename == "" {
		err = errors.New(`filename == ""`)
		return
	}
	if mediaReader == nil {
		err = errors.New("mediaReader == nil")
		return
	}
	return c.mediaUpload(media.MEDIA_TYPE_IMAGE, filename, mediaReader)
}

// 上传多媒体缩略图
func (c *Client) MediaUploadThumb(filename string, mediaReader io.Reader) (resp *media.UploadResponse, err error) {
	if filename == "" {
		err = errors.New(`filename == ""`)
		return
	}
	if mediaReader == nil {
		err = errors.New("mediaReader == nil")
		return
	}
	return c.mediaUpload(media.MEDIA_TYPE_THUMB, filename, mediaReader)
}

// 上传多媒体语音
func (c *Client) MediaUploadVoice(filename string, mediaReader io.Reader) (resp *media.UploadResponse, err error) {
	if filename == "" {
		err = errors.New(`filename == ""`)
		return
	}
	if mediaReader == nil {
		err = errors.New("mediaReader == nil")
		return
	}
	return c.mediaUpload(media.MEDIA_TYPE_VOICE, filename, mediaReader)
}

// 上传多媒体视频
func (c *Client) MediaUploadVideo(filename string, mediaReader io.Reader) (resp *media.UploadResponse, err error) {
	if filename == "" {
		err = errors.New(`filename == ""`)
		return
	}
	if mediaReader == nil {
		err = errors.New("mediaReader == nil")
		return
	}
	return c.mediaUpload(media.MEDIA_TYPE_VIDEO, filename, mediaReader)
}

// 上传多媒体
func (c *Client) mediaUpload(mediaType, filename string, mediaReader io.Reader) (resp *media.UploadResponse, err error) {
	token, err := c.Token()
	if err != nil {
		return
	}

	bodyBuf := c.getBufferFromPool() // io.ReadWriter
	defer c.putBufferToPool(bodyBuf) // important!

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

	_url := mediaUploadURL(token, mediaType)
	httpResp, err := c.httpClient.Post(_url, bodyContentType, bodyBuf)
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
		if result.ErrCode != 0 {
			err = &result.Error
			return
		}

		resp = &media.UploadResponse{
			MediaType: result.MediaType,
			MediaId:   result.MediaId,
			CreatedAt: result.CreatedAt,
		}
		return

	default:
		var result struct {
			media.UploadResponse
			Error
		}
		if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
			return
		}
		if result.ErrCode != 0 {
			err = &result.Error
			return
		}

		resp = &result.UploadResponse
		return
	}
}

// 下载多媒体文件.
//  NOTE: 视频文件不支持下载.
func (c *Client) MediaDownloadToFile(mediaId, _filepath string) (err error) {
	file, err := os.Create(_filepath)
	if err != nil {
		return
	}
	defer file.Close()

	return c.mediaDownload(mediaId, file)
}

// 下载多媒体文件.
//  NOTE: 视频文件不支持下载.
func (c *Client) MediaDownload(mediaId string, writer io.Writer) error {
	if writer == nil {
		return errors.New("writer == nil")
	}
	return c.mediaDownload(mediaId, writer)
}

// 下载多媒体文件.
func (c *Client) mediaDownload(mediaId string, writer io.Writer) (err error) {
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := mediaDownloadURL(token, mediaId)

	resp, err := c.httpClient.Get(_url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", resp.Status)
	}

	contentType, _, _ := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if contentType != "text/plain" && contentType != "application/json" {
		_, err = io.Copy(writer, resp.Body)
		return
	}

	// 返回的是错误信息
	var result Error
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return
	}
	return &result
}

// 根据上传的缩略图媒体创建图文消息素材
func (c *Client) MediaCreateNews(news media.News) (resp *media.UploadResponse, err error) {
	if err = news.CheckValid(); err != nil {
		return
	}

	token, err := c.Token()
	if err != nil {
		return
	}
	_url := mediaCreateNewsURL(token)

	var result struct {
		media.UploadResponse
		Error
	}
	if err = c.postJSON(_url, news, &result); err != nil {
		return
	}

	if result.ErrCode != 0 {
		err = &result.Error
		return
	}

	resp = &result.UploadResponse
	return
}

// 根据上传的视频文件 media_id 创建视频媒体, 群发视频消息应该用这个函数得到的 media_id.
//  NOTE: title, description 可以为空
func (c *Client) MediaCreateVideo(mediaId, title, description string) (resp *media.UploadResponse, err error) {
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := mediaCreateVideoURL(token)

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
		media.UploadResponse
		Error
	}
	if err = c.postJSON(_url, &request, &result); err != nil {
		return
	}

	if result.ErrCode != 0 {
		err = &result.Error
		return
	}

	resp = &result.UploadResponse
	return
}
