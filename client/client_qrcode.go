// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"errors"
	"fmt"
	"github.com/chanxuehong/wechat/qrcode"
	"io"
	"net/http"
	"os"
)

// 创建临时二维码
func (c *Client) QRCodeTemporaryCreate(sceneId uint32, expireSeconds int) (_qrcode *qrcode.TemporaryQRCode, err error) {
	var request struct {
		ExpireSeconds int    `json:"expire_seconds"`
		ActionName    string `json:"action_name"`
		ActionInfo    struct {
			Scene struct {
				SceneId uint32 `json:"scene_id"`
			} `json:"scene"`
		} `json:"action_info"`
	}
	request.ExpireSeconds = expireSeconds
	request.ActionName = "QR_SCENE"
	request.ActionInfo.Scene.SceneId = sceneId

	var result struct {
		qrcode.TemporaryQRCode
		Error
	}

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := qrcodeCreateURL(token)
	if err = c.postJSON(_url, &request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		result.TemporaryQRCode.SceneId = sceneId
		_qrcode = &result.TemporaryQRCode
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = &result.Error
		return
	}
}

// 创建永久二维码
func (c *Client) QRCodePermanentCreate(sceneId uint32) (_qrcode *qrcode.PermanentQRCode, err error) {
	var request struct {
		ActionName string `json:"action_name"`
		ActionInfo struct {
			Scene struct {
				SceneId uint32 `json:"scene_id"`
			} `json:"scene"`
		} `json:"action_info"`
	}
	request.ActionName = "QR_LIMIT_SCENE"
	request.ActionInfo.Scene.SceneId = sceneId

	var result struct {
		qrcode.PermanentQRCode
		Error
	}

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := qrcodeCreateURL(token)
	if err = c.postJSON(_url, &request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		result.PermanentQRCode.SceneId = sceneId
		_qrcode = &result.PermanentQRCode
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = &result.Error
		return
	}
}

// 根据 qrcode ticket 得到 qrcode 图片的 url
func QRCodeURL(ticket string) string {
	return qrcodeURL(ticket)
}

// 通过 ticket 换取二维码到 writer
func QRCodeDownload(ticket string, writer io.Writer) (err error) {
	if writer == nil {
		return errors.New("writer == nil")
	}

	_url := qrcodeURL(ticket)
	resp, err := http.Get(_url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// ticket正确情况下，http 返回码是200，是一张图片，可以直接展示或者下载。
	if resp.StatusCode == http.StatusOK {
		_, err = io.Copy(writer, resp.Body)
		return
	}

	// 错误情况下（如ticket非法）返回HTTP错误码404。
	return fmt.Errorf("qrcode with ticket %s not found", ticket)
}

// 通过 ticket 换取二维码到 writer
func (c *Client) QRCodeDownload(ticket string, writer io.Writer) (err error) {
	if writer == nil {
		return errors.New("writer == nil")
	}

	_url := qrcodeURL(ticket)
	resp, err := c.httpClient.Get(_url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// ticket正确情况下，http 返回码是200，是一张图片，可以直接展示或者下载。
	if resp.StatusCode == http.StatusOK {
		_, err = io.Copy(writer, resp.Body)
		return
	}

	// 错误情况下（如ticket非法）返回HTTP错误码404。
	return fmt.Errorf("qrcode with ticket %s not found", ticket)
}

// 通过 ticket 换取二维码到文件 _filepath
func QRCodeDownloadToFile(ticket, _filepath string) (err error) {
	file, err := os.Create(_filepath)
	if err != nil {
		return
	}
	defer file.Close()

	return QRCodeDownload(ticket, file)
}

// 通过 ticket 换取二维码到文件 _filepath
func (c *Client) QRCodeDownloadToFile(ticket, _filepath string) (err error) {
	file, err := os.Create(_filepath)
	if err != nil {
		return
	}
	defer file.Close()

	return c.QRCodeDownload(ticket, file)
}
