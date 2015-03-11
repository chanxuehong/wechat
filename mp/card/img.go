// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     gaowenbin(gaowenbinmarr@gmail.com), chanxuehong(chanxuehong@gmail.com)

package card

import (
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/chanxuehong/wechat/mp"
)

// 上传到微信服务器的图片信息，用于卡卷的 logo_url.
type ImageInfo struct {
	URL string `json:"url"`
}

// 上传图片, 用于卡卷的 logo_url.
//  1.上传的图片限制文件大小限制1MB，像素为300*300，支持JPG 格式。
//  2.调用接口获取的logo_url 进支持在微信相关业务下使用，否则会做相应处理。
func (clt *Client) UploadImage(imgPath string) (info ImageInfo, err error) {
	file, err := os.Open(imgPath)
	if err != nil {
		return
	}
	defer file.Close()

	return clt.uploadImageFromReader(filepath.Base(imgPath), file)
}

// 上传图片, 用于卡卷的 logo_url.
//  1.上传的图片限制文件大小限制1MB，像素为300*300，支持JPG 格式。
//  2.调用接口获取的logo_url 进支持在微信相关业务下使用，否则会做相应处理。
//  3.注意参数 filename 不是文件路径, 是指定 multipart/form-data 里面文件名称
func (clt *Client) UploadImageFromReader(filename string, reader io.Reader) (info ImageInfo, err error) {
	if filename == "" {
		err = errors.New("empty filename")
		return
	}
	if reader == nil {
		err = errors.New("nil reader")
		return
	}

	return clt.uploadImageFromReader(filename, reader)
}

func (clt *Client) uploadImageFromReader(filename string, reader io.Reader) (info ImageInfo, err error) {
	var result struct {
		mp.Error
		ImageInfo
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token="
	if err = clt.UploadFromReader(incompleteURL, filename, reader, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	info = result.ImageInfo
	return
}
