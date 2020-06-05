package media

import (
	"io"
	"os"
	"path/filepath"

	"github.com/chanxuehong/wechat/work/core"
)

const (
	MediaTypeImage = "image"
	MediaTypeVoice = "voice"
	MediaTypeVideo = "video"
	MediaTypeFile  = "file"
)

type MediaInfo struct {
	MediaType string `json:"type"`       // 媒体文件类型，分别有图片（image）、语音（voice）、视频（video），普通文件（file）
	MediaId   string `json:"media_id"`   // 媒体文件上传后，获取时的唯一标识
	CreatedAt int64  `json:"created_at"` // 媒体文件上传时间戳
}

// UploadImage 上传多媒体图片
func UploadImage(clt *core.Client, filepath string) (info *MediaInfo, err error) {
	return upload(clt, MediaTypeImage, filepath)
}

// UploadImageFromReader 上传多媒体图片
//  NOTE: 参数 filename 不是文件路径, 是 multipart/form-data 里面 filename 的值.
func UploadImageFromReader(clt *core.Client, filename string, reader io.Reader) (info *MediaInfo, err error) {
	return uploadFromReader(clt, MediaTypeImage, filename, reader)
}

// UploadVoice 上传多媒体语音
func UploadVoice(clt *core.Client, filepath string) (info *MediaInfo, err error) {
	return upload(clt, MediaTypeVoice, filepath)
}

// UploadVoiceFromReader 上传多媒体语音
//  NOTE: 参数 filename 不是文件路径, 是 multipart/form-data 里面 filename 的值.
func UploadVoiceFromReader(clt *core.Client, filename string, reader io.Reader) (info *MediaInfo, err error) {
	return uploadFromReader(clt, MediaTypeVoice, filename, reader)
}

// UploadVideo 上传多媒体视频
func UploadVideo(clt *core.Client, filepath string) (info *MediaInfo, err error) {
	return upload(clt, MediaTypeVideo, filepath)
}

// UploadVideoFromReader 上传多媒体视频
//  NOTE: 参数 filename 不是文件路径, 是 multipart/form-data 里面 filename 的值.
func UploadVideoFromReader(clt *core.Client, filename string, reader io.Reader) (info *MediaInfo, err error) {
	return uploadFromReader(clt, MediaTypeVideo, filename, reader)
}

// UploadFile 上传普通文件
func UploadFile(clt *core.Client, filepath string) (info *MediaInfo, err error) {
	return upload(clt, MediaTypeFile, filepath)
}

// UploadFileFromReader 上传普通文件
//  NOTE: 参数 filename 不是文件路径, 是 multipart/form-data 里面 filename 的值.
func UploadFileFromReader(clt *core.Client, filename string, reader io.Reader) (info *MediaInfo, err error) {
	return uploadFromReader(clt, MediaTypeFile, filename, reader)
}

// =====================================================================================================================

func upload(clt *core.Client, mediaType, _filepath string) (info *MediaInfo, err error) {
	file, err := os.Open(_filepath)
	if err != nil {
		return
	}
	defer file.Close()

	return uploadFromReader(clt, mediaType, filepath.Base(_filepath), file)
}

func uploadFromReader(clt *core.Client, mediaType, filename string, reader io.Reader) (info *MediaInfo, err error) {
	var incompleteURL = "https://qyapi.weixin.qq.com/cgi-bin/media/upload?type=" + mediaType + "&access_token="

	var fields = []core.MultipartFormField{
		{
			IsFile:   true,
			Name:     "media",
			FileName: filename,
			Value:    reader,
		},
	}
	var result struct {
		core.Error
		MediaInfo
	}
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
