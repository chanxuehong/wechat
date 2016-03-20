package account

import (
	"io"
	"os"
	"path/filepath"

	"github.com/chanxuehong/wechat/mp/core"
)

// UploadHeadImage 上传客服头像.
func UploadHeadImage(clt *core.Client, kfAccount, imageFilePath string) (err error) {
	file, err := os.Open(imageFilePath)
	if err != nil {
		return
	}
	defer file.Close()

	return UploadHeadImageFromReader(clt, kfAccount, filepath.Base(imageFilePath), file)
}

// UploadHeadImageFromReader 上传客服头像.
//  NOTE: 参数 filename 不是文件路径, 是 multipart/form-data 里面 filename 的值.
func UploadHeadImageFromReader(clt *core.Client, kfAccount, filename string, reader io.Reader) (err error) {
	// TODO
	//	incompleteURL := "https://api.weixin.qq.com/customservice/kfaccount/uploadheadimg?kf_account=" +
	//		url.QueryEscape(kfAccount) + "&access_token="
	incompleteURL := "https://api.weixin.qq.com/customservice/kfaccount/uploadheadimg?kf_account=" +
		kfAccount + "&access_token="

	var fields = []core.MultipartFormField{{
		IsFile:   true,
		Name:     "media",
		FileName: filename,
		Value:    reader,
	}}
	var result core.Error
	if err = clt.PostMultipartForm(incompleteURL, fields, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}
