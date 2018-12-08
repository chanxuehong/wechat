package material

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"unicode"

	"github.com/chanxuehong/wechat/internal/debug/api"
	"github.com/chanxuehong/wechat/internal/debug/api/retry"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/util"
)

// Download 下载多媒体到文件.
//  对于视频素材, 先通过 GetVideo 得到 Video 信息, 然后通过 Video.DownloadURL 来下载
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
//  对于视频素材, 先通过 GetVideo 得到 Video 信息, 然后通过 Video.DownloadURL 来下载
func DownloadToWriter(clt *core.Client, mediaId string, writer io.Writer) (written int64, err error) {
	httpClient := clt.HttpClient
	if httpClient == nil {
		httpClient = util.DefaultMediaHttpClient
	}

	buffer := bytes.NewBuffer(make([]byte, 0, 256))
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)

	var request = struct {
		MediaId string `json:"media_id"`
	}{
		MediaId: mediaId,
	}
	if err = encoder.Encode(&request); err != nil {
		return
	}
	requestBodyBytes := buffer.Bytes()

	var errorResult core.Error

	// 先读取 64bytes 内容来判断返回的是不是错误信息
	// {"errcode":40007,"errmsg":"invalid media_id"}
	var buf = make([]byte, 64)

	token, err := clt.Token()
	if err != nil {
		return
	}

	hasRetried := false
RETRY:
	finalURL := "https://api.weixin.qq.com/cgi-bin/material/get_material?access_token=" + url.QueryEscape(token)
	written, err = httpDownloadToWriter(httpClient, finalURL, requestBodyBytes, buf, writer, &errorResult)
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
		retry.DebugPrintError(errorResult.ErrCode, errorResult.ErrMsg, token)
		if !hasRetried {
			hasRetried = true
			errorResult = core.Error{}
			if token, err = clt.RefreshToken(token); err != nil {
				return
			}
			retry.DebugPrintNewToken(token)
			goto RETRY
		}
		retry.DebugPrintFallthrough(token)
		fallthrough
	default:
		err = &errorResult
		return
	}
}

var (
	// {"errcode":40007,"errmsg":"invalid media_id"}
	errRespBeginWithCode = []byte(`{"errcode":`)
	errRespBeginWithMsg  = []byte(`{"errmsg":"`)
)

func httpDownloadToWriter(clt *http.Client, url string, body []byte, buf []byte, writer io.Writer, errorResult *core.Error) (written int64, err error) {
	api.DebugPrintPostJSONRequest(url, body)
	httpResp, err := clt.Post(url, "application/json; charset=utf-8", bytes.NewReader(body))
	if err != nil {
		return 0, err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("http.Status: %s", httpResp.Status)
	}

	buf2 := buf // 保存预先读取的少量头部信息
	switch n, err := io.ReadFull(httpResp.Body, buf2); err {
	case nil:
		break
	case io.ErrUnexpectedEOF:
		buf2 = buf2[:n]
		break
	case io.EOF: // 基本不会出现
		return 0, nil
	default:
		return 0, err
	}

	var httpRespBody io.Reader
	if len(buf2) < len(buf) {
		httpRespBody = bytes.NewReader(buf2)
	} else {
		httpRespBody = io.MultiReader(bytes.NewReader(buf2), httpResp.Body)
	}

	buf3 := trimLeft(buf2)
	if bytes.HasPrefix(buf3, errRespBeginWithCode) || bytes.HasPrefix(buf3, errRespBeginWithMsg) {
		// 返回的是错误信息
		return 0, api.DecodeJSONHttpResponse(httpRespBody, errorResult)
	} else {
		// 返回的是媒体流
		return io.Copy(writer, httpRespBody)
	}
}

func trimLeft(s []byte) []byte {
	for i := 0; i < len(s); i++ {
		if isSpace(s[i]) {
			continue
		}
		return s[i:]
	}
	return s
}

func isSpace(b byte) bool {
	if b > unicode.MaxASCII {
		return false
	}
	return unicode.IsSpace(rune(b))
}
