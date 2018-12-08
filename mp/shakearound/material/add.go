package material

import (
	"errors"
	"io"
	"net/url"
	"os"
	"path/filepath"

	"github.com/chanxuehong/wechat/mp/core"
)

type ImageInfo struct {
	PicURL string `json:"pic_url"`
}

func Add(clt *core.Client, imagePath, _type string) (info ImageInfo, err error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return
	}
	defer file.Close()

	return addFromReader(clt, filepath.Base(imagePath), file, _type)
}

func AddFromReader(clt *core.Client, filename string, reader io.Reader, _type string) (info ImageInfo, err error) {
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

func addFromReader(clt *core.Client, filename string, reader io.Reader, _type string) (info ImageInfo, err error) {
	var result struct {
		core.Error
		ImageInfo `json:"data"`
	}

	var incompleteURL string
	if _type != "" {
		incompleteURL = "https://api.weixin.qq.com/shakearound/material/add?type=" + url.QueryEscape(_type) +
			"&access_token="
	} else {
		incompleteURL = "https://api.weixin.qq.com/shakearound/material/add?access_token="
	}
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
	info = result.ImageInfo
	return
}
