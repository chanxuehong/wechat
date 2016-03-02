package material

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/url"
	"os"
	"path/filepath"

	"github.com/chanxuehong/wechat/mp/core"
)

const (
	MaterialTypeImage = "image"
	MaterialTypeVoice = "voice"
	MaterialTypeVideo = "video"
	MaterialTypeThumb = "thumb"
	MaterialTypeNews  = "news"
)

// 上传多媒体图片
func UploadImage(clt *core.Client, filepath string) (mediaId, _url string, err error) {
	return uploadMaterial(clt, MaterialTypeImage, filepath)
}

// 上传多媒体缩略图
func UploadThumb(clt *core.Client, filepath string) (mediaId, _url string, err error) {
	return uploadMaterial(clt, MaterialTypeThumb, filepath)
}

// 上传多媒体
func uploadMaterial(clt *core.Client, materialType, _filepath string) (mediaId, _url string, err error) {
	file, err := os.Open(_filepath)
	if err != nil {
		return
	}
	defer file.Close()

	return uploadMaterialFromReader(clt, materialType, filepath.Base(_filepath), file)
}

// 上传多媒体图片
//  NOTE: 参数 filename 不是文件路径, 是指定 multipart/form-data 里面文件名称
func UploadImageFromReader(clt *core.Client, filename string, reader io.Reader) (mediaId, _url string, err error) {
	if filename == "" {
		err = errors.New("empty filename")
		return
	}
	if reader == nil {
		err = errors.New("nil reader")
		return
	}
	return uploadMaterialFromReader(clt, MaterialTypeImage, filename, reader)
}

// 上传多媒体缩略图
//  NOTE: 参数 filename 不是文件路径, 是指定 multipart/form-data 里面文件名称
func UploadThumbFromReader(clt *core.Client, filename string, reader io.Reader) (mediaId, _url string, err error) {
	if filename == "" {
		err = errors.New("empty filename")
		return
	}
	if reader == nil {
		err = errors.New("nil reader")
		return
	}
	return uploadMaterialFromReader(clt, MaterialTypeThumb, filename, reader)
}

func uploadMaterialFromReader(clt *core.Client, materialType, filename string, reader io.Reader) (mediaId, _url string, err error) {
	var result struct {
		core.Error
		MediaId string `json:"media_id"`
		URL     string `json:"url"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/material/add_material?type=" +
		url.QueryEscape(materialType) + "&access_token="
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
	mediaId = result.MediaId
	_url = result.URL
	return
}

// voice =======================================================================

// 上传多媒体语音
func UploadVoice(clt *core.Client, _filepath string) (mediaId string, err error) {
	file, err := os.Open(_filepath)
	if err != nil {
		return
	}
	defer file.Close()

	return uploadVoiceFromReader(clt, filepath.Base(_filepath), file)
}

// 上传多媒体语音
//  NOTE: 参数 filename 不是文件路径, 是指定 multipart/form-data 里面文件名称
func UploadVoiceFromReader(clt *core.Client, filename string, reader io.Reader) (mediaId string, err error) {
	if filename == "" {
		err = errors.New("empty filename")
		return
	}
	if reader == nil {
		err = errors.New("nil reader")
		return
	}
	return uploadVoiceFromReader(clt, filename, reader)
}

func uploadVoiceFromReader(clt *core.Client, filename string, reader io.Reader) (mediaId string, err error) {
	var result struct {
		core.Error
		MediaId string `json:"media_id"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/material/add_material?type=voice&access_token="
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
	mediaId = result.MediaId
	return
}

// video =======================================================================

// 上传多媒体视频
func UploadVideo(clt *core.Client, _filepath string, title, introduction string) (mediaId string, err error) {
	file, err := os.Open(_filepath)
	if err != nil {
		return
	}
	defer file.Close()

	return uploadVideoFromReader(clt, filepath.Base(_filepath), file, title, introduction)
}

// 上传多媒体缩视频
//  NOTE: 参数 filename 不是文件路径, 是指定 multipart/form-data 里面文件名称
func UploadVideoFromReader(clt *core.Client, filename string, reader io.Reader, title, introduction string) (mediaId string, err error) {
	if filename == "" {
		err = errors.New("empty filename")
		return
	}
	if reader == nil {
		err = errors.New("nil reader")
		return
	}
	return uploadVideoFromReader(clt, filename, reader, title, introduction)
}

func uploadVideoFromReader(clt *core.Client, filename string, reader io.Reader,
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
		core.Error
		MediaId string `json:"media_id"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/material/add_material?type=video&access_token="
	fields := []core.MultipartFormField{
		{
			IsFile:   true,
			Name:     "media",
			FileName: filename,
			Value:    reader,
		},
		{
			IsFile: false,
			Name:   "description",
			Value:  bytes.NewReader(descBytes),
		},
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
