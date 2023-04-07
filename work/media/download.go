package media

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"

	"github.com/bububa/wechat/internal/debug/api"
	"github.com/bububa/wechat/internal/debug/api/retry"
	"github.com/bububa/wechat/util"
	"github.com/bububa/wechat/work/core"
)

// Get 获取临时素材.
func Get(clt *core.Client, mediaId, filepath string) (written int64, err error) {
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

	return GetToWriter(clt, mediaId, file)
}

// GetToWriter 获取临时素材 io.Writer.
func GetToWriter(clt *core.Client, mediaId string, writer io.Writer) (written int64, err error) {
	httpClient := clt.HttpClient
	if httpClient == nil {
		httpClient = util.DefaultMediaHttpClient
	}

	var incompleteURL = "https://qyapi.weixin.qq.com/cgi-bin/media/get?media_id=" + url.QueryEscape(mediaId) + "&access_token="
	var errorResult core.Error

	token, err := clt.Token()
	if err != nil {
		return
	}

	hasRetried := false
RETRY:
	finalURL := incompleteURL + url.QueryEscape(token)
	written, err = httpDownloadToWriter(httpClient, finalURL, writer, &errorResult, clt.Debug())
	if err != nil {
		return
	}
	if written > 0 {
		return
	}

	switch errorResult.ErrCode {
	case core.ErrCodeOK:
		return // 基本不会出现
	case core.ErrCodeInvalidCredential, core.ErrCodeAccessTokenExpired:
		retry.DebugPrintError(errorResult.ErrCode, errorResult.ErrMsg, token, clt.Debug())
		if !hasRetried {
			hasRetried = true
			errorResult = core.Error{}
			if token, err = clt.RefreshToken(token); err != nil {
				return
			}
			retry.DebugPrintNewToken(token, clt.Debug())
			goto RETRY
		}
		retry.DebugPrintFallthrough(token, clt.Debug())
		fallthrough
	default:
		err = &errorResult
		return
	}
}

func httpDownloadToWriter(clt *http.Client, url string, writer io.Writer, errorResult *core.Error, debug bool) (written int64, err error) {
	api.DebugPrintGetRequest(url, debug)
	httpResp, err := clt.Get(url)
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
		return 0, api.DecodeJSONHttpResponse(httpResp.Body, errorResult, debug)
	}
}