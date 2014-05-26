package wechat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chanxuehong/wechat/media"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// 上传多媒体文件
//  NOTE:
//  1. media_id是可复用的，调用该接口需http协议;
//  2. 媒体文件在后台保存时间为3天，即3天后media_id失效。
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

// 上传多媒体文件
//  NOTE:
//  1. media_id是可复用的，调用该接口需http协议;
//  2. 媒体文件在后台保存时间为3天，即3天后media_id失效。
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
		return nil, errors.New("错误的 mediaType")
	}

	if mediaReader == nil {
		return nil, errors.New("mediaReader == nil")
	}

	token, err := c.Token()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf(mediaUploadUrlFormat, token, mediaType)

	bodyBuf := c.getBuffer()
	defer c.putBuffer(bodyBuf)

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

	resp, err := http.Post(url, bodyContentType, bodyBuf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		media.UploadResponse
		Error
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		return nil, &result.Error
	}
	return &result.UploadResponse, nil
}

// 下载多媒体文件
//  NOTE: 视频文件不支持下载，调用该接口需http协议。
func (c *Client) MediaDownloadToFile(mediaId, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return c.MediaDownload(mediaId, file)
}

// 下载多媒体文件
//  NOTE: 视频文件不支持下载，调用该接口需http协议。
func (c *Client) MediaDownload(mediaId string, writer io.Writer) error {
	token, err := c.Token()
	if err != nil {
		return err
	}

	url := fmt.Sprintf(mediaDownloadUrlFormat, token, mediaId)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.Header.Get("Content-Type") != "text/plain" {
		_, err = io.Copy(writer, resp.Body)
		return err
	}

	// 返回的是错误信息
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result Error
	if err = json.Unmarshal(body, &result); err != nil {
		return err
	}
	return &result
}

// 上传图文消息素材
func (c *Client) MediaUploadNews(news *media.News) (*media.UploadResponse, error) {
	if news == nil {
		return nil, errors.New("news == nil")
	}

	token, err := c.Token()
	if err != nil {
		return nil, err
	}
	jsonData, err := json.Marshal(news)
	if err != nil {
		return nil, err
	}

	url := mediaUploadNewsUrlPrefix + token
	resp, err := http.Post(url, "application/json; charset=utf-8", bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		media.UploadResponse
		Error
	}
	if err = json.Unmarshal(body, &result); err != nil {
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
		return nil, errors.New("video == nil")
	}

	token, err := c.Token()
	if err != nil {
		return nil, err
	}
	jsonData, err := json.Marshal(video)
	if err != nil {
		return nil, err
	}

	url := mediaUploadVideoUrlPrefix + token
	resp, err := http.Post(url, "application/json; charset=utf-8", bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		media.UploadResponse
		Error
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		return nil, &result.Error
	}
	return &result.UploadResponse, nil
}
