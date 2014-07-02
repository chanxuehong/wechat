// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"github.com/chanxuehong/wechat/media"
	"testing"
)

func TestMediaUploadImageFromFileAndDownload(t *testing.T) {
	resp, err := _test_client.MediaUploadFromFile(media.MEDIA_TYPE_IMAGE, "testdata/upload_image.jpg")
	if err != nil {
		t.Error("上传图片失败,", err)
		return
	}

	if resp == nil {
		t.Error(`resp == ""`)
		return
	}

	if resp.MediaType != media.MEDIA_TYPE_IMAGE || resp.MediaId == "" ||
		resp.CreatedAt == 0 {

		t.Error(`返回的 resp 不合法`)
		return
	}

	err = _test_client.MediaDownloadToFile(resp.MediaId, "testdata/download_image.jpg")
	if err != nil {
		t.Error("下载图片失败,", err)
		return
	}

	isEqual, err := fileEqual("testdata/upload_image.jpg", "testdata/download_image.jpg")
	if err != nil {
		t.Error(err)
		return
	}

	if !isEqual {
		t.Error("上传和下载的文件不相等")
		return
	}
}

func TestMediaUploadThumbFromFileAndDownload(t *testing.T) {
	resp, err := _test_client.MediaUploadFromFile(media.MEDIA_TYPE_THUMB, "testdata/upload_thumb.jpg")
	if err != nil {
		t.Error("上传缩略图失败,", err)
		return
	}

	if resp == nil {
		t.Error(`resp == ""`)
		return
	}

	if resp.MediaType != media.MEDIA_TYPE_THUMB || resp.MediaId == "" ||
		resp.CreatedAt == 0 {

		t.Error(`返回的 resp 不合法`)
		return
	}

	err = _test_client.MediaDownloadToFile(resp.MediaId, "testdata/download_thumb.jpg")
	if err != nil {
		t.Error("下载缩略图失败,", err)
		return
	}

	isEqual, err := fileEqual("testdata/upload_thumb.jpg", "testdata/download_thumb.jpg")
	if err != nil {
		t.Error(err)
		return
	}

	if !isEqual {
		t.Error("上传和下载的文件不相等")
		return
	}
}

func TestMediaUploadVoiceFromFileAndDownload(t *testing.T) {
	resp, err := _test_client.MediaUploadFromFile(media.MEDIA_TYPE_VOICE, "testdata/upload_voice.amr")
	if err != nil {
		t.Error("上传语音失败,", err)
		return
	}

	if resp == nil {
		t.Error(`resp == ""`)
		return
	}

	if resp.MediaType != media.MEDIA_TYPE_VOICE || resp.MediaId == "" ||
		resp.CreatedAt == 0 {

		t.Error(`返回的 resp 不合法`)
		return
	}

	err = _test_client.MediaDownloadToFile(resp.MediaId, "testdata/download_voice.amr")
	if err != nil {
		t.Error("下载语音失败,", err)
		return
	}

	isEqual, err := fileEqual("testdata/upload_voice.amr", "testdata/download_voice.amr")
	if err != nil {
		t.Error(err)
		return
	}

	if !isEqual {
		t.Error("上传和下载的文件不相等")
		return
	}
}

// 视频文件不能下载
func TestMediaUploadVideoFromFile(t *testing.T) {
	resp, err := _test_client.MediaUploadFromFile(media.MEDIA_TYPE_VIDEO, "testdata/upload_video.mp4")
	if err != nil {
		t.Error("上传视频媒体失败,", err)
		return
	}

	if resp == nil {
		t.Error(`resp == ""`)
		return
	}

	if resp.MediaType != media.MEDIA_TYPE_VIDEO || resp.MediaId == "" ||
		resp.CreatedAt == 0 {

		t.Error(`返回的 resp 不合法`)
		return
	}
}
