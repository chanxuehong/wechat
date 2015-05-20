// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com) Harry Rong(harrykobe@gmail.com)
package shakearound

import (
	"github.com/chanxuehong/wechat/mp"
	"io"
	"os"
	"path/filepath"
	"errors"
)

type PicInfo struct {
	PicUrl string `json:"pic_url"`
}

func (clt Client) AddMaterial(_filepath string) (info *PicInfo, err error) {
	file, err := os.Open(_filepath)
	if err != nil {
		return
	}
	defer file.Close()

	return clt.AddMaterialFromReader(filepath.Base(_filepath), file)
}

func (clt Client) AddMaterialFromReader(filename string, reader io.Reader) (info *PicInfo, err error) {
	if filename == "" {
		err = errors.New("empty filename")
		return
	}
	if reader == nil {
		err = errors.New("nil reader")
		return
	}

	var result struct {
		mp.Error
		Data struct{
				 PicInfo
			 }	`json:"data"`

	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/material/add?access_token="
	fields := []mp.MultipartFormField{{
		ContentType: 0,
		FieldName:   "media",
		FileName:    filename,
		Value:       reader,
	}}
	if err = clt.PostMultipartForm(incompleteURL, fields, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	info = &result.Data.PicInfo
	return
}