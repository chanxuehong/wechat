package wechat

import (
	"github.com/chanxuehong/wechat/media"
	"io"
)

// 上传多媒体文件
//  NOTE:
//  1. media_id是可复用的，调用该接口需http协议;
//  2. 媒体文件在后台保存时间为3天，即3天后media_id失效。
//  3. 图片（image）: 256K，支持JPG格式
//  4. 语音（voice）：256K，播放长度不超过60s，支持AMR\MP3格式
//  5. 视频（video）：1MB，支持MP4格式
//  6. 缩略图（thumb）：64KB，支持JPG格式
func (c *Client) UploadMedia(mediaType string, media io.Reader) (*media.UploadResponse, error) {
	return nil, nil
}

// 下载多媒体文件
//  NOTE: 视频文件不支持下载，调用该接口需http协议。
func (c *Client) DownloadMedia(mediaId string) (io.ReadCloser, error) {
	return nil, nil
}

// 上传图文消息素材
func (c *Client) UploadNewsMsg(news *media.UploadNewsMsg) (*media.UploadResponse, error) {
	return nil, nil
}

// 上传视频消息
func (c *Client) UploadVideoMsg(video *media.UploadVideoMsg) (*media.UploadResponse, error) {
	return nil, nil
}
