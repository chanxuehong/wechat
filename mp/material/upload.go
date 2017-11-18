package material

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"gopkg.in/chanxuehong/wechat.v2/mp/core"
)

const (
	MaterialTypeImage = "image"
	MaterialTypeVoice = "voice"
	MaterialTypeVideo = "video"
	MaterialTypeThumb = "thumb"
	MaterialTypeNews  = "news"
)

// image ===============================================================================================================

// UploadImage 上传多媒体图片
func UploadImage(clt *core.Client, _filepath string) (mediaId, url string, err error) {
	file, err := os.Open(_filepath)
	if err != nil {
		return
	}
	defer file.Close()

	return UploadImageFromReader(clt, filepath.Base(_filepath), file)
}

// UploadImageFromReader 上传多媒体图片
//  NOTE: 参数 filename 不是文件路径, 是 multipart/form-data 里面 filename 的值.
func UploadImageFromReader(clt *core.Client, filename string, reader io.Reader) (mediaId, url string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/material/add_material?type=image&access_token="

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
		MediaId string `json:"media_id"`
		URL     string `json:"url"`
	}
	if err = clt.PostMultipartForm(incompleteURL, fields, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	mediaId = result.MediaId
	url = result.URL
	return
}

// thumb ===============================================================================================================

// UploadThumb 上传多媒体缩略图
func UploadThumb(clt *core.Client, _filepath string) (mediaId, url string, err error) {
	file, err := os.Open(_filepath)
	if err != nil {
		return
	}
	defer file.Close()

	return UploadThumbFromReader(clt, filepath.Base(_filepath), file)
}

// UploadThumbFromReader 上传多媒体缩略图
//  NOTE: 参数 filename 不是文件路径, 是 multipart/form-data 里面 filename 的值.
func UploadThumbFromReader(clt *core.Client, filename string, reader io.Reader) (mediaId, url string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/material/add_material?type=thumb&access_token="

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
		MediaId string `json:"media_id"`
		URL     string `json:"url"`
	}
	if err = clt.PostMultipartForm(incompleteURL, fields, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	mediaId = result.MediaId
	url = result.URL
	return
}

// voice ===============================================================================================================

// UploadVoice 上传多媒体语音
func UploadVoice(clt *core.Client, _filepath string) (mediaId string, err error) {
	file, err := os.Open(_filepath)
	if err != nil {
		return
	}
	defer file.Close()

	return UploadVoiceFromReader(clt, filepath.Base(_filepath), file)
}

// UploadVoiceFromReader 上传多媒体语音
//  NOTE: 参数 filename 不是文件路径, 是 multipart/form-data 里面 filename 的值.
func UploadVoiceFromReader(clt *core.Client, filename string, reader io.Reader) (mediaId string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/material/add_material?type=voice&access_token="

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
		MediaId string `json:"media_id"`
	}
	if err = clt.PostMultipartForm(incompleteURL, fields, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	mediaId = result.MediaId
	return
}

// video ===============================================================================================================

// UploadVideo 上传多媒体视频.
func UploadVideo(clt *core.Client, _filepath string, title, introduction string) (mediaId string, err error) {
	file, err := os.Open(_filepath)
	if err != nil {
		return
	}
	defer file.Close()

	return UploadVideoFromReader(clt, filepath.Base(_filepath), file, title, introduction)
}

// UploadVideoFromReader 上传多媒体缩视频.
//  NOTE: 参数 filename 不是文件路径, 是 multipart/form-data 里面 filename 的值.
func UploadVideoFromReader(clt *core.Client, filename string, reader io.Reader, title, introduction string) (mediaId string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/material/add_material?type=video&access_token="

	buffer := bytes.NewBuffer(make([]byte, 0, 256))
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)

	var description = struct {
		Title        string `json:"title"`
		Introduction string `json:"introduction"`
	}{
		Title:        title,
		Introduction: introduction,
	}
	if err = encoder.Encode(&description); err != nil {
		return
	}
	descriptionBytes := buffer.Bytes()

	var fields = []core.MultipartFormField{
		{
			IsFile:   true,
			Name:     "media",
			FileName: filename,
			Value:    reader,
		},
		{
			IsFile: false,
			Name:   "description",
			Value:  bytes.NewReader(descriptionBytes),
		},
	}
	var result struct {
		core.Error
		MediaId string `json:"media_id"`
	}
	if err = clt.PostMultipartForm(incompleteURL, fields, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	mediaId = result.MediaId
	return
}
