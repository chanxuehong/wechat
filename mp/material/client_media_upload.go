// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package material

import (
	"encoding/json"
	"errors"
	"io"
	"net/url"
	"os"
	"path/filepath"

	"github.com/chanxuehong/wechat/mp"
)

const (
	MediaTypeImage = "image"
	MediaTypeVoice = "voice"
	MediaTypeVideo = "video"
	MediaTypeThumb = "thumb"
)

// 上传多媒体图片
func (clt *Client) UploadImage(filepath string) (mediaId string, err error) {
	return clt.uploadMedia(MediaTypeImage, filepath)
}

// 上传多媒体缩略图
func (clt *Client) UploadThumb(filepath string) (mediaId string, err error) {
	return clt.uploadMedia(MediaTypeThumb, filepath)
}

// 上传多媒体语音
func (clt *Client) UploadVoice(filepath string) (mediaId string, err error) {
	return clt.uploadMedia(MediaTypeVoice, filepath)
}

// 上传多媒体
func (clt *Client) uploadMedia(mediaType, _filepath string) (mediaId string, err error) {
	file, err := os.Open(_filepath)
	if err != nil {
		return
	}
	defer file.Close()

	return clt.uploadMediaFromReader(mediaType, filepath.Base(_filepath), file)
}

// 上传多媒体图片
//  NOTE: 参数 filename 不是文件路径, 是指定 multipart/form-data 里面文件名称
func (clt *Client) UploadImageFromReader(filename string, reader io.Reader) (mediaId string, err error) {
	if filename == "" {
		err = errors.New("empty filename")
		return
	}
	if reader == nil {
		err = errors.New("nil reader")
		return
	}
	return clt.uploadMediaFromReader(MediaTypeImage, filename, reader)
}

// 上传多媒体缩略图
//  NOTE: 参数 filename 不是文件路径, 是指定 multipart/form-data 里面文件名称
func (clt *Client) UploadThumbFromReader(filename string, reader io.Reader) (mediaId string, err error) {
	if filename == "" {
		err = errors.New("empty filename")
		return
	}
	if reader == nil {
		err = errors.New("nil reader")
		return
	}
	return clt.uploadMediaFromReader(MediaTypeThumb, filename, reader)
}

// 上传多媒体语音
//  NOTE: 参数 filename 不是文件路径, 是指定 multipart/form-data 里面文件名称
func (clt *Client) UploadVoiceFromReader(filename string, reader io.Reader) (mediaId string, err error) {
	if filename == "" {
		err = errors.New("empty filename")
		return
	}
	if reader == nil {
		err = errors.New("nil reader")
		return
	}
	return clt.uploadMediaFromReader(MediaTypeVoice, filename, reader)
}

func (clt *Client) uploadMediaFromReader(mediaType, filename string, reader io.Reader) (mediaId string, err error) {
	var result struct {
		mp.Error
		MediaId string `json:"media_id"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/material/add_material?type=" +
		url.QueryEscape(mediaType) + "&access_token="
	if err = clt.UploadFromReader(incompleteURL, "media", filename, reader, "", nil, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	mediaId = result.MediaId
	return
}

// =============================================================================

// 上传多媒体视频
func (clt *Client) UploadVideo(_filepath string, title, introduction string) (mediaId string, err error) {
	file, err := os.Open(_filepath)
	if err != nil {
		return
	}
	defer file.Close()

	return clt.uploadVideoFromReader(filepath.Base(_filepath), file, title, introduction)
}

// 上传多媒体缩视频
//  NOTE: 参数 filename 不是文件路径, 是指定 multipart/form-data 里面文件名称
func (clt *Client) UploadVideoFromReader(filename string, reader io.Reader, title, introduction string) (mediaId string, err error) {
	if filename == "" {
		err = errors.New("empty filename")
		return
	}
	if reader == nil {
		err = errors.New("nil reader")
		return
	}
	return clt.uploadVideoFromReader(filename, reader, title, introduction)
}

func (clt *Client) uploadVideoFromReader(filename string, reader io.Reader,
	title, introduction string) (mediaId string, err error) {

	var desc = struct {
		Title        string `json:"title"`
		Introduction string `json:"introduction"`
	}{
		Title:        title,
		Introduction: introduction,
	}

	descBytes, err := json.Marshal(&desc)
	if err != nil {
		return
	}

	var result struct {
		mp.Error
		MediaId string `json:"media_id"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/material/add_material?type=video&access_token="
	if err = clt.UploadFromReader(incompleteURL, "media", filename, reader, "description", descBytes, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	mediaId = result.MediaId
	return
}
