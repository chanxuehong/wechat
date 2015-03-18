// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package material

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"

	"github.com/chanxuehong/wechat/mp"
)

// 下载多媒体到文件.
//  请注意，视频文件不支持下载
func (clt *Client) DownloadMedia(mediaId, filepath string) (err error) {
	file, err := os.Create(filepath)
	if err != nil {
		return
	}
	defer file.Close()

	return clt.downloadMediaToWriter(mediaId, file)
}

// 下载多媒体到 io.Writer.
//  请注意，视频文件不支持下载
func (clt *Client) DownloadMediaToWriter(mediaId string, writer io.Writer) error {
	if writer == nil {
		return errors.New("nil writer")
	}
	return clt.downloadMediaToWriter(mediaId, writer)
}

// 下载多媒体到 io.Writer.
func (clt *Client) downloadMediaToWriter(mediaId string, writer io.Writer) (err error) {
	request := struct {
		MediaId string `json:"media_id"`
	}{
		MediaId: mediaId,
	}

	requestBody, err := json.Marshal(&request)
	if err != nil {
		return
	}

	token, err := clt.Token()
	if err != nil {
		return
	}

	hasRetried := false
RETRY:
	finalURL := "https://api.weixin.qq.com/cgi-bin/material/get_material?access_token=" + url.QueryEscape(token)

	httpResp, err := clt.HttpClient.Post(finalURL, "application/json; charset=utf-8", bytes.NewReader(requestBody))
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}

	ContentType, _, _ := mime.ParseMediaType(httpResp.Header.Get("Content-Type"))
	if ContentType != "text/plain" && ContentType != "application/json" { // 返回的是媒体流
		_, err = io.Copy(writer, httpResp.Body)
		return
	}

	// 返回的是错误信息
	var result mp.Error
	if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
		return
	}

	switch result.ErrCode {
	case mp.ErrCodeOK:
		return // 基本不会出现
	case mp.ErrCodeInvalidCredential, mp.ErrCodeTimeout: // 失效(过期)重试一次
		if !hasRetried {
			hasRetried = true

			if token, err = clt.TokenRefresh(); err != nil {
				return
			}
			goto RETRY
		}
		fallthrough
	default:
		err = &result
		return
	}
}
