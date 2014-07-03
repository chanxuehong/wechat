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
	"time"
)

// 创建临时二维码
func (c *Client) QRCodeTemporaryCreate(sceneId uint32, expireSeconds int) (*qrcode.TemporaryQRCode, error) {
	token, err := c.Token()
	if err != nil {
		return nil, err
	}
	_url := qrcodeCreateURL(token)

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
		Ticket        string `json:"ticket"`
		ExpireSeconds int64  `json:"expire_seconds"`
		Error
	}
	if err = c.postJSON(_url, &request, &result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}

	var ret qrcode.TemporaryQRCode
	ret.SceneId = sceneId
	ret.Ticket = result.Ticket
	ret.Expiry = time.Now().Unix() + result.ExpireSeconds

	return &ret, nil
}

// 创建永久二维码
func (c *Client) QRCodePermanentCreate(sceneId uint32) (*qrcode.PermanentQRCode, error) {
	token, err := c.Token()
	if err != nil {
		return nil, err
	}
	_url := qrcodeCreateURL(token)

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
	if err = c.postJSON(_url, &request, &result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}
	result.PermanentQRCode.SceneId = sceneId
	return &result.PermanentQRCode, nil
}

// 根据 qrcode ticket 得到 qrcode 图片的 url
func QRCodeURL(ticket string) string {
	return qrcodeURL(ticket)
}

// 通过 ticket 换取二维码到 writer
func QRCodeDownload(ticket string, writer io.Writer) error {
	if writer == nil {
		return errors.New("writer == nil")
	}

	_url := qrcodeURL(ticket)
	resp, err := http.Get(_url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// ticket正确情况下，http 返回码是200，是一张图片，可以直接展示或者下载。
	if resp.StatusCode == http.StatusOK {
		_, err = io.Copy(writer, resp.Body)
		return err
	}

	// 错误情况下（如ticket非法）返回HTTP错误码404。
	return fmt.Errorf("qrcode with ticket %s not found", ticket)
}

// 通过 ticket 换取二维码到 writer
func (c *Client) QRCodeDownload(ticket string, writer io.Writer) error {
	if writer == nil {
		return errors.New("writer == nil")
	}

	_url := qrcodeURL(ticket)
	resp, err := c.httpClient.Get(_url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// ticket正确情况下，http 返回码是200，是一张图片，可以直接展示或者下载。
	if resp.StatusCode == http.StatusOK {
		_, err = io.Copy(writer, resp.Body)
		return err
	}

	// 错误情况下（如ticket非法）返回HTTP错误码404。
	return fmt.Errorf("qrcode with ticket %s not found", ticket)
}

// 通过 ticket 换取二维码到文件 filePath
func QRCodeDownloadToFile(ticket, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return QRCodeDownload(ticket, file)
}

// 通过 ticket 换取二维码到文件 filePath
func (c *Client) QRCodeDownloadToFile(ticket, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return c.QRCodeDownload(ticket, file)
}
