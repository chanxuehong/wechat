package media

import (
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"

	"github.com/chanxuehong/wechat/mp/core"
)

// Download 下载多媒体到文件.
//  请注意, 视频文件不支持下载
func Download(clt *core.Client, mediaId, filepath string) (written int64, err error) {
	file, err := os.Create(filepath)
	if err != nil {
		return
	}
	defer func() {
		file.Close()
		if err != nil {
			os.Remove(filepath)
		}
	}()

	return DownloadToWriter(clt, mediaId, file)
}

// DownloadToWriter 下载多媒体到 io.Writer.
//  请注意, 视频文件不支持下载
func DownloadToWriter(clt *core.Client, mediaId string, writer io.Writer) (written int64, err error) {
	httpClient := clt.HttpClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	var incompleteURL = "https://api.weixin.qq.com/cgi-bin/media/get?media_id=" + url.QueryEscape(mediaId) + "&access_token="
	var result core.Error

	token, err := clt.Token()
	if err != nil {
		return
	}

	hasRetried := false
RETRY:
	finalURL := incompleteURL + url.QueryEscape(token)

	written, err = func() (int64, error) {
		httpResp, err := httpClient.Get(finalURL)
		if err != nil {
			return 0, err
		}
		defer httpResp.Body.Close()

		if httpResp.StatusCode != http.StatusOK {
			return 0, fmt.Errorf("http.Status: %s", httpResp.Status)
		}

		ContentDisposition := httpResp.Header.Get("Content-Disposition")
		ContentType, _, _ := mime.ParseMediaType(httpResp.Header.Get("Content-Type"))
		if ContentDisposition != "" && ContentType != "text/plain" && ContentType != "application/json" {
			// 返回的是媒体流
			return io.Copy(writer, httpResp.Body)
		} else {
			// 返回的是错误信息
			return 0, json.NewDecoder(httpResp.Body).Decode(&result)
		}
	}()
	if err != nil {
		return
	}
	if written > 0 {
		return
	}

	switch result.ErrCode {
	case core.ErrCodeOK:
		return // 基本不会出现
	case core.ErrCodeInvalidCredential, core.ErrCodeAccessTokenExpired:
		if !hasRetried {
			hasRetried = true
			result = core.Error{}
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
