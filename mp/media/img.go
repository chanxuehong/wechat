package media

import (
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/chanxuehong/wechat/mp/core"
)

// 上传到微信服务器的图片信息.
type ImageInfo struct {
	URL string `json:"url"`
}

// 上传图片到微信服务器, 给其他场景使用, 比如卡卷, POI.
func UploadImagePermanent(clt *core.Client, imgPath string) (info ImageInfo, err error) {
	file, err := os.Open(imgPath)
	if err != nil {
		return
	}
	defer file.Close()

	return uploadImagePermanentFromReader(clt, filepath.Base(imgPath), file)
}

// 上传图片到微信服务器, 给其他场景使用, 比如卡卷, POI.
func UploadImagePermanentFromReader(clt *core.Client, filename string, reader io.Reader) (info ImageInfo, err error) {
	if filename == "" {
		err = errors.New("empty filename")
		return
	}
	if reader == nil {
		err = errors.New("nil reader")
		return
	}

	return uploadImagePermanentFromReader(clt, filename, reader)
}

func uploadImagePermanentFromReader(clt *core.Client, filename string, reader io.Reader) (info ImageInfo, err error) {
	var result struct {
		core.Error
		ImageInfo
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token="
	fields := []core.MultipartFormField{{
		IsFile:   true,
		Name:     "buffer",
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
	info = result.ImageInfo
	return
}
