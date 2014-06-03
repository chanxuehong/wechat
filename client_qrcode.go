package wechat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chanxuehong/wechat/qrcode"
	"io"
	"math"
	"net/http"
	"os"
)

// 创建临时二维码
func (c *Client) QRCodeCreate(sceneId int, expireSeconds int) (*qrcode.QRCode, error) {
	if sceneId == 0 {
		return nil, errors.New("QRCodeCreate: sceneId 应该是个32位非0整型")
	}
	if sceneId < math.MinInt32 || sceneId > math.MaxUint32 { // 包括了 int32, uint32
		return nil, errors.New("QRCodeCreate: sceneId 应该是个32位非0整型")
	}
	if expireSeconds <= 0 || expireSeconds > qrcode.QRCodeExpireSecondsLimit {
		return nil, fmt.Errorf("QRCodeCreate: expireSeconds 应该在 (0,%d] 之间", qrcode.QRCodeExpireSecondsLimit)
	}

	token, err := c.Token()
	if err != nil {
		return nil, err
	}

	var request struct {
		ExpireSeconds int    `json:"expire_seconds"`
		ActionName    string `json:"action_name"`
		ActionInfo    struct {
			Scene struct {
				SceneId int `json:"scene_id"`
			} `json:"scene"`
		} `json:"action_info"`
	}

	request.ExpireSeconds = expireSeconds
	request.ActionName = "QR_SCENE"
	request.ActionInfo.Scene.SceneId = sceneId

	jsonData, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}

	_url := clientQRCodeCreateURL(token)
	resp, err := c.httpClient.Post(_url, postJSONContentType, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("QRCodeCreate: %s", resp.Status)
	}

	var result struct {
		qrcode.QRCode
		Error
	}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		return nil, &result.Error
	}
	result.QRCode.SceneId = sceneId
	return &result.QRCode, nil
}

// 创建永久二维码
func (c *Client) QRCodeLimitCreate(sceneId int) (*qrcode.QRCode, error) {
	if sceneId <= 0 || sceneId > qrcode.QRCodeLimitSceneIdLimit {
		return nil, fmt.Errorf("QRCodeLimitCreate: sceneId 应该在 (0,%d] 之间", qrcode.QRCodeLimitSceneIdLimit)
	}

	token, err := c.Token()
	if err != nil {
		return nil, err
	}

	var request struct {
		ActionName string `json:"action_name"`
		ActionInfo struct {
			Scene struct {
				SceneId int `json:"scene_id"`
			} `json:"scene"`
		} `json:"action_info"`
	}

	request.ActionName = "QR_LIMIT_SCENE"
	request.ActionInfo.Scene.SceneId = sceneId

	jsonData, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}

	_url := clientQRCodeCreateURL(token)
	resp, err := c.httpClient.Post(_url, postJSONContentType, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("QRCodeLimitCreate: %s", resp.Status)
	}

	var result struct {
		qrcode.QRCode
		Error
	}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		return nil, &result.Error
	}
	result.QRCode.SceneId = sceneId
	result.QRCode.ExpireSeconds = 0 // 强制为 0
	return &result.QRCode, nil
}

// 根据 qrcode ticket 得到 qrcode 图片的 url
func QRCodeUrl(ticket string) string {
	return clientQRCodeURL(ticket)
}

// 通过 ticket 换取二维码到 writer
func QRCodeDownload(ticket string, writer io.Writer) error {
	if len(ticket) == 0 {
		return errors.New(`QRCodeDownload: ticket == ""`)
	}
	if writer == nil {
		return errors.New("QRCodeDownload: writer == nil")
	}

	_url := clientQRCodeURL(ticket)
	resp, err := http.Get(_url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		_, err = io.Copy(writer, resp.Body)
		return err
	}

	return fmt.Errorf("QRCodeDownload: qrcode with ticket %s not found", ticket)
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
