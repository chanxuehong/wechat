<<<<<<< HEAD
// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com), magicshui(shuiyuzhe@gmail.com), Harry Rong(harrykobe@gmail.com)

=======
>>>>>>> github/v2
package material

import (
	"errors"
	"io"
	"net/url"
	"os"
	"path/filepath"

<<<<<<< HEAD
	"github.com/chanxuehong/wechat/mp"
=======
	"github.com/chanxuehong/wechat/mp/core"
>>>>>>> github/v2
)

type ImageInfo struct {
	PicURL string `json:"pic_url"`
}

<<<<<<< HEAD
func Add(clt *mp.Client, imagePath, _type string) (info ImageInfo, err error) {
=======
func Add(clt *core.Client, imagePath, _type string) (info ImageInfo, err error) {
>>>>>>> github/v2
	file, err := os.Open(imagePath)
	if err != nil {
		return
	}
	defer file.Close()

	return addFromReader(clt, filepath.Base(imagePath), file, _type)
}

<<<<<<< HEAD
func AddFromReader(clt *mp.Client, filename string, reader io.Reader, _type string) (info ImageInfo, err error) {
=======
func AddFromReader(clt *core.Client, filename string, reader io.Reader, _type string) (info ImageInfo, err error) {
>>>>>>> github/v2
	if filename == "" {
		err = errors.New("empty filename")
		return
	}
	if reader == nil {
		err = errors.New("nil reader")
		return
	}

	return addFromReader(clt, filename, reader, _type)
}

<<<<<<< HEAD
func addFromReader(clt *mp.Client, filename string, reader io.Reader, _type string) (info ImageInfo, err error) {
	var result struct {
		mp.Error
=======
func addFromReader(clt *core.Client, filename string, reader io.Reader, _type string) (info ImageInfo, err error) {
	var result struct {
		core.Error
>>>>>>> github/v2
		ImageInfo `json:"data"`
	}

	var incompleteURL string
	if _type != "" {
		incompleteURL = "https://api.weixin.qq.com/shakearound/material/add?type=" + url.QueryEscape(_type) +
			"&access_token="
	} else {
		incompleteURL = "https://api.weixin.qq.com/shakearound/material/add?access_token="
	}
<<<<<<< HEAD
	fields := []mp.MultipartFormField{{
		ContentType: 0,
		FieldName:   "media",
		FileName:    filename,
		Value:       reader,
=======
	fields := []core.MultipartFormField{{
		IsFile:   true,
		Name:     "media",
		FileName: filename,
		Value:    reader,
>>>>>>> github/v2
	}}
	if err = clt.PostMultipartForm(incompleteURL, fields, &result); err != nil {
		return
	}

<<<<<<< HEAD
	if result.ErrCode != mp.ErrCodeOK {
=======
	if result.ErrCode != core.ErrCodeOK {
>>>>>>> github/v2
		err = &result.Error
		return
	}
	info = result.ImageInfo
	return
}
