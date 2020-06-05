package media

import (
	"io"
	"os"
	"path/filepath"

	"github.com/chanxuehong/wechat/work/core"
)

// UploadImg 上传多媒体图片
func UploadImg(clt *core.Client, filepath string) (link string, err error) {
	return uploadImg(clt, filepath)
}

// UploadImgFromReader 上传多媒体图片
//  NOTE: 参数 filename 不是文件路径, 是 multipart/form-data 里面 filename 的值.
func UploadImgFromReader(clt *core.Client, filename string, reader io.Reader) (link string, err error) {
	return uploadImgFromReader(clt, filename, reader)
}

// =====================================================================================================================

func uploadImg(clt *core.Client, _filepath string) (link string, err error) {
	file, err := os.Open(_filepath)
	if err != nil {
		return
	}
	defer file.Close()

	return uploadImgFromReader(clt, filepath.Base(_filepath), file)
}

func uploadImgFromReader(clt *core.Client, filename string, reader io.Reader) (link string, err error) {
	const incompleteURL = "https://qyapi.weixin.qq.com/cgi-bin/media/uploadimg?access_token="

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
		Url string `json:"url"`
	}
	if err = clt.PostMultipartForm(incompleteURL, fields, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	link = result.Url
	return
}
