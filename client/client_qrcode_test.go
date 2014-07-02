// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"bytes"
	"crypto/md5"
	"io/ioutil"
	"testing"
	"time"
)

func fileMd5(filepath string) ([]byte, error) {
	// 因为文件不大, 一次性读入内存
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	sum := md5.Sum(b)
	return sum[:], nil
}

func TestQRCodePermanentCreateAndDownload(t *testing.T) {
	qrcode, err := _test_client.QRCodePermanentCreate(100)
	if err != nil {
		t.Error("创建永久二维码失败,", err)
		return
	}

	if qrcode == nil {
		t.Error(`qrcode == ""`)
		return
	}

	if qrcode.SceneId != 100 || qrcode.Ticket == "" {
		t.Error(`返回的 qrcode 不合法`)
		return
	}

	err = QRCodeDownloadToFile(qrcode.Ticket, "testdata/permanent0.jpg")
	if err != nil {
		t.Error("下载二维码失败,", err)
		return
	}

	err = _test_client.QRCodeDownloadToFile(qrcode.Ticket, "testdata/permanent1.jpg")
	if err != nil {
		t.Error("下载二维码失败,", err)
		return
	}

	sum0, err := fileMd5("testdata/permanent0.jpg")
	if err != nil {
		t.Error("计算 testdata/permanent0.jpg 的 md5 值出错,", err)
		return
	}
	sum1, err := fileMd5("testdata/permanent1.jpg")
	if err != nil {
		t.Error("计算 testdata/permanent1.jpg 的 md5 值出错,", err)
		return
	}

	if !bytes.Equal(sum0, sum1) {
		t.Error("两次下载的二维码不一样")
		return
	}
}

func TestQRCodeTemporaryCreateAndDownload(t *testing.T) {
	qrcode, err := _test_client.QRCodeTemporaryCreate(1000000, 100)
	if err != nil {
		t.Error("创建临时二维码失败,", err)
		return
	}

	if qrcode == nil {
		t.Error(`qrcode == ""`)
		return
	}

	if qrcode.SceneId != 1000000 || qrcode.Ticket == "" || qrcode.Expiry < time.Now().Unix() {
		t.Error(`返回的 qrcode 不合法`)
		return
	}

	err = QRCodeDownloadToFile(qrcode.Ticket, "testdata/temporary0.jpg")
	if err != nil {
		t.Error("下载二维码失败,", err)
		return
	}

	err = _test_client.QRCodeDownloadToFile(qrcode.Ticket, "testdata/temporary1.jpg")
	if err != nil {
		t.Error("下载二维码失败,", err)
		return
	}

	sum0, err := fileMd5("testdata/temporary0.jpg")
	if err != nil {
		t.Error("计算 testdata/temporary0.jpg 的 md5 值出错,", err)
		return
	}
	sum1, err := fileMd5("testdata/temporary1.jpg")
	if err != nil {
		t.Error("计算 testdata/temporary1.jpg 的 md5 值出错,", err)
		return
	}

	if !bytes.Equal(sum0, sum1) {
		t.Error("两次下载的二维码不一样")
		return
	}
}
