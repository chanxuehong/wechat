package client

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/chanxuehong/wechat/media"
	"io/ioutil"
	"testing"
)

func TestMediaUploadDownload(t *testing.T) {
	c := NewClient(appid, appsecret, nil)
	fmt.Println(c.Token())

	resp, err := c.MediaUploadFromFile(media.MEDIA_TYPE_IMAGE, "../test_data/media/upload.jpg")
	if err != nil {
		t.Error(err)
		return
	}

	imgMediaId := resp.MediaId

	err = c.MediaDownloadToFile(imgMediaId, "../test_data/media/download.jpg")
	if err != nil {
		t.Error(err)
		return
	}

	file0Bytes, err := ioutil.ReadFile(`../test_data/media/upload.jpg`)
	if err != nil {
		t.Error(err)
		return
	}
	file1Bytes, err := ioutil.ReadFile(`../test_data/media/download.jpg`)
	if err != nil {
		t.Error(err)
		return
	}

	file0MD5Sum := md5.Sum(file0Bytes)
	file1MD5Sum := md5.Sum(file1Bytes)

	if !bytes.Equal(file0MD5Sum[:], file1MD5Sum[:]) {
		t.Error("上传的文件和下载下来的文件签名不一致")
		return
	}

	// 上传缩略图
	resp, err = c.MediaUploadFromFile(media.MEDIA_TYPE_THUMB, "../test_data/media/thumb.jpg")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(resp)
}
