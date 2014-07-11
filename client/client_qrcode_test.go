// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"testing"
	"time"
)

func TestQRCodePermanentCreateAndDownload(t *testing.T) {
	qrcode, err := _test_client.QRCodePermanentCreate(100)
	if err != nil {
		t.Error("创建永久二维码失败,", err)
		return
	}
	if qrcode == nil {
		t.Error(`qrcode == nil`)
		return
	}
	if qrcode.SceneId != 100 || qrcode.Ticket == "" {
		t.Error(`返回的 qrcode 不合法`)
		return
	}

	err = QRCodeDownloadToFile(qrcode.Ticket, "testdata/qrcode_permanent0.jpg")
	if err != nil {
		t.Error("下载二维码失败0,", err)
		return
	}

	err = _test_client.QRCodeDownloadToFile(qrcode.Ticket, "testdata/qrcode_permanent1.jpg")
	if err != nil {
		t.Error("下载二维码失败1,", err)
		return
	}

	isEqual, err := fileEqual("testdata/qrcode_permanent0.jpg", "testdata/qrcode_permanent1.jpg")
	if err != nil {
		t.Error(err)
		return
	}
	if !isEqual {
		t.Error("两次下载的二维码不一样")
		return
	}

	t.Log("永久二维码 ticket:", qrcode.Ticket)
}

func TestQRCodeTemporaryCreateAndDownload(t *testing.T) {
	qrcode, err := _test_client.QRCodeTemporaryCreate(1000000, 1000)
	if err != nil {
		t.Error("创建临时二维码失败,", err)
		return
	}
	if qrcode == nil {
		t.Error(`qrcode == nil`)
		return
	}
	if qrcode.SceneId != 1000000 || qrcode.Ticket == "" || qrcode.ExpiresAt != 100+time.Now().Unix() {
		t.Error(`返回的 qrcode 不合法`)
		return
	}

	err = QRCodeDownloadToFile(qrcode.Ticket, "testdata/qrcode_temporary0.jpg")
	if err != nil {
		t.Error("下载二维码失败0,", err)
		return
	}

	err = _test_client.QRCodeDownloadToFile(qrcode.Ticket, "testdata/qrcode_temporary1.jpg")
	if err != nil {
		t.Error("下载二维码失败1,", err)
		return
	}

	isEqual, err := fileEqual("testdata/qrcode_temporary0.jpg", "testdata/qrcode_temporary1.jpg")
	if err != nil {
		t.Error(err)
		return
	}
	if !isEqual {
		t.Error("两次下载的二维码不一样")
		return
	}

	t.Log("临时二维码 ticket:", qrcode.Ticket)
}
