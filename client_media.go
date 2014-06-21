package wechat

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

// 上传多媒体文件.
//  NOTE:
//  1. 媒体文件在后台保存时间为3天，即3天后 media_id 失效。
//  2. 返回的 media_id 是可复用的;
//  3. 图片（image）: 256K，支持JPG格式
//  4. 语音（voice）：256K，播放长度不超过60s，支持AMR\MP3格式
//  5. 视频（video）：1MB，支持MP4格式
//  6. 缩略图（thumb）：64KB，支持JPG格式
func (c *Client) MediaUploadFromFile(mediaType, filePath string) (*media.UploadResponse, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return c.MediaUpload(mediaType, filepath.Base(filePath), file)
}

// 上传多媒体文件.
//  NOTE:
//  1. 媒体文件在后台保存时间为3天，即3天后 media_id 失效。
//  2. 返回的 media_id 是可复用的;
//  3. 图片（image）: 256K，支持JPG格式
//  4. 语音（voice）：256K，播放长度不超过60s，支持AMR\MP3格式
//  5. 视频（video）：1MB，支持MP4格式
//  6. 缩略图（thumb）：64KB，支持JPG格式
func (c *Client) MediaUpload(mediaType, filename string, mediaReader io.Reader) (*media.UploadResponse, error) {
	switch mediaType {
	case media.MEDIA_TYPE_IMAGE,
		media.MEDIA_TYPE_VOICE,
		media.MEDIA_TYPE_VIDEO,
		media.MEDIA_TYPE_THUMB:
	default:
		return nil, errors.New("MediaUpload: 错误的 mediaType")
	}
	if filename == "" {
		return nil, errors.New(`MediaUpload: filename == ""`)
	}
	if mediaReader == nil {
		return nil, errors.New("MediaUpload: mediaReader == nil")
	}

	token, err := c.Token()
	if err != nil {
		return nil, err
	}

	bodyBuf := c.getBufferFromPool() // io.ReadWriter
	defer c.putBufferToPool(bodyBuf) // important!

	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(fileWriter, mediaReader); err != nil {
		return nil, err
	}

	bodyContentType := bodyWriter.FormDataContentType()

	if err = bodyWriter.Close(); err != nil {
		return nil, err
	}

	_url := clientMediaUploadURL(token, mediaType)
	resp, err := c.httpClient.Post(_url, bodyContentType, bodyBuf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("MediaUpload: %s", resp.Status)
	}

	var result struct {
		media.UploadResponse
		Error
	}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		return nil, &result.Error
	}
	return &result.UploadResponse, nil
}

// 下载多媒体文件.
//  NOTE: 视频文件不支持下载.
func (c *Client) MediaDownloadToFile(mediaId, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return c.MediaDownload(mediaId, file)
}

// 下载多媒体文件.
//  NOTE: 视频文件不支持下载.
func (c *Client) MediaDownload(mediaId string, writer io.Writer) error {
	if mediaId == "" {
		return errors.New(`MediaDownload: mediaId == ""`)
	}
	if writer == nil {
		return errors.New("MediaDownload: writer == nil")
	}

	token, err := c.Token()
	if err != nil {
		return err
	}

	_url := clientMediaDownloadURL(token, mediaId)
	resp, err := c.httpClient.Get(_url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("MediaDownload: %s", resp.Status)
	}

	contentType, _, _ := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if contentType != "text/plain" { // 如果下载失败返回的是 Content-Type: text/plain, 下载成功是其他的 Content-Type
		_, err = io.Copy(writer, resp.Body)
		return err
	}

	// 返回的是错误信息

	var result Error
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}
	return &result
}

// 上传图文消息素材
func (c *Client) MediaUploadNews(news *media.News) (*media.UploadResponse, error) {
	if news == nil {
		return nil, errors.New("MediaUploadNews: news == nil")
	}

	token, err := c.Token()
	if err != nil {
		return nil, err
	}

	buf := c.getBufferFromPool()
	// defer c.putBufferToPool(buf) // buf 要快速迭代, 所以不用 defer, 要提前释放

	if err = json.NewEncoder(buf).Encode(news); err != nil {
		c.putBufferToPool(buf) ////
		return nil, err
	}

	_url := clientMediaUploadNewsURL(token)
	resp, err := c.httpClient.Post(_url, postJSONContentType, buf)
	c.putBufferToPool(buf) ////
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("MediaUploadNews: %s", resp.Status)
	}

	var result struct {
		media.UploadResponse
		Error
	}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		return nil, &result.Error
	}
	return &result.UploadResponse, nil
}

// 上传视频消息
func (c *Client) MediaUploadVideo(video *media.Video) (*media.UploadResponse, error) {
	if video == nil {
		return nil, errors.New("MediaUploadVideo: video == nil")
	}

	token, err := c.Token()
	if err != nil {
		return nil, err
	}

	buf := c.getBufferFromPool()
	// defer c.putBufferToPool(buf) // buf 要快速迭代, 所以不用 defer, 要提前释放

	if err = json.NewEncoder(buf).Encode(video); err != nil {
		c.putBufferToPool(buf) ////
		return nil, err
	}

	_url := clientMediaUploadVideoURL(token)
	resp, err := c.httpClient.Post(_url, postJSONContentType, buf)
	c.putBufferToPool(buf) ////
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("MediaUploadVideo: %s", resp.Status)
	}

	var result struct {
		media.UploadResponse
		Error
	}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		return nil, &result.Error
	}
	return &result.UploadResponse, nil
}
