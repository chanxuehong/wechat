// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/chanxuehong/wechat/mp/qrcode"
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
	url_ := qrcodeCreateURL(token)

	if err = c.postJSON(url_, &request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		result.TemporaryQRCode.SceneId = sceneId
		_qrcode = &result.TemporaryQRCode
		return

	case errCodeInvalidCredential, errCodeTimeout:
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
	url_ := qrcodeCreateURL(token)

	if err = c.postJSON(url_, &request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		result.PermanentQRCode.SceneId = sceneId
		_qrcode = &result.PermanentQRCode
		return

	case errCodeInvalidCredential, errCodeTimeout:
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
//  如果 httpClient == nil 则默认用 http.DefaultClient
func QRCodeDownloadToWriter(ticket string, writer io.Writer, httpClient *http.Client) (err error) {
	if writer == nil {
		return errors.New("writer == nil")
	}
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	url_ := qrcodeURL(ticket)
	resp, err := httpClient.Get(url_)
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
func (c *Client) QRCodeDownloadToWriter(ticket string, writer io.Writer) (err error) {
	if writer == nil {
		return errors.New("writer == nil")
	}

	url_ := qrcodeURL(ticket)
	resp, err := c.httpClient.Get(url_)
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

// 通过 ticket 换取二维码到文件 filepath_
//  如果 httpClient == nil 则默认用 http.DefaultClient
func QRCodeDownload(ticket, filepath_ string, httpClient *http.Client) (err error) {
	file, err := os.Create(filepath_)
	if err != nil {
		return
	}
	defer file.Close()

	return QRCodeDownloadToWriter(ticket, file, httpClient)
}

// 通过 ticket 换取二维码到文件 filepath_
func (c *Client) QRCodeDownload(ticket, filepath_ string) (err error) {
	file, err := os.Create(filepath_)
	if err != nil {
		return
	}
	defer file.Close()

	return c.QRCodeDownloadToWriter(ticket, file)
}
