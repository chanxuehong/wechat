package media

import (
	"errors"
	"io"
	"net/url"
	"os"
	"path/filepath"

	"github.com/chanxuehong/wechat/mp/core"
)

// 上传多媒体图片
func UploadImage(clt *core.Client, filepath string) (info *MediaInfo, err error) {
	return uploadMedia(clt, MediaTypeImage, filepath)
}

// 上传多媒体语音
func UploadVoice(clt *core.Client, filepath string) (info *MediaInfo, err error) {
	return uploadMedia(clt, MediaTypeVoice, filepath)
}

// 上传多媒体视频
func UploadVideo(clt *core.Client, filepath string) (info *MediaInfo, err error) {
	return uploadMedia(clt, MediaTypeVideo, filepath)
}

// 上传多媒体
func uploadMedia(clt *core.Client, mediaType, _filepath string) (info *MediaInfo, err error) {
	file, err := os.Open(_filepath)
	if err != nil {
		return
	}
	defer file.Close()

	return uploadMediaFromReader(clt, mediaType, filepath.Base(_filepath), file)
}

// 上传多媒体图片
//  NOTE: 参数 filename 不是文件路径, 是指定 multipart/form-data 里面文件名称
func UploadImageFromReader(clt *core.Client, filename string, reader io.Reader) (info *MediaInfo, err error) {
	if filename == "" {
		err = errors.New("empty filename")
		return
	}
	if reader == nil {
		err = errors.New("nil reader")
		return
	}
	return uploadMediaFromReader(clt, MediaTypeImage, filename, reader)
}

// 上传多媒体语音
//  NOTE: 参数 filename 不是文件路径, 是指定 multipart/form-data 里面文件名称
func UploadVoiceFromReader(clt *core.Client, filename string, reader io.Reader) (info *MediaInfo, err error) {
	if filename == "" {
		err = errors.New("empty filename")
		return
	}
	if reader == nil {
		err = errors.New("nil reader")
		return
	}
	return uploadMediaFromReader(clt, MediaTypeVoice, filename, reader)
}

// 上传多媒体视频
//  NOTE: 参数 filename 不是文件路径, 是指定 multipart/form-data 里面文件名称
func UploadVideoFromReader(clt *core.Client, filename string, reader io.Reader) (info *MediaInfo, err error) {
	if filename == "" {
		err = errors.New("empty filename")
		return
	}
	if reader == nil {
		err = errors.New("nil reader")
		return
	}
	return uploadMediaFromReader(clt, MediaTypeVideo, filename, reader)
}

func uploadMediaFromReader(clt *core.Client, mediaType, filename string, reader io.Reader) (info *MediaInfo, err error) {
	var result struct {
		core.Error
		MediaInfo
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/media/upload?type=" +
		url.QueryEscape(mediaType) + "&access_token="
	fields := []core.MultipartFormField{{
		IsFile:   true,
		Name:     "media",
		FileName: filename,
		Value:    reader,
	}}
	if err = clt.PostMultipartForm(incompleteURL, fields, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	info = &result.MediaInfo
	return
}

// =============================================================================

// 上传多媒体缩略图
func UploadThumb(clt *core.Client, _filepath string) (info *MediaInfo, err error) {
	file, err := os.Open(_filepath)
	if err != nil {
		return
	}
	defer file.Close()

	return uploadThumbFromReader(clt, filepath.Base(_filepath), file)
}

// 上传多媒体缩略图
//  NOTE: 参数 filename 不是文件路径, 是指定 multipart/form-data 里面文件名称
func UploadThumbFromReader(clt *core.Client, filename string, reader io.Reader) (info *MediaInfo, err error) {
	if filename == "" {
		err = errors.New("empty filename")
		return
	}
	if reader == nil {
		err = errors.New("nil reader")
		return
	}
	return uploadThumbFromReader(clt, filename, reader)
}

func uploadThumbFromReader(clt *core.Client, filename string, reader io.Reader) (info *MediaInfo, err error) {
	var result struct {
		core.Error
		MediaType string `json:"type"`
		MediaId   string `json:"thumb_media_id"`
		CreatedAt int64  `json:"created_at"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/media/upload?type=thumb&access_token="
	fields := []core.MultipartFormField{{
		IsFile:   true,
		Name:     "media",
		FileName: filename,
		Value:    reader,
	}}
	if err = clt.PostMultipartForm(incompleteURL, fields, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	info = &MediaInfo{
		MediaType: result.MediaType,
		MediaId:   result.MediaId,
		CreatedAt: result.CreatedAt,
	}
	return
}
