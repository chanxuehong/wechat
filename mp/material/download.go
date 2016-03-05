package material

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/chanxuehong/wechat/mp/core"
)

// Download 下载多媒体到文件.
//  对于视频素材, 先通过 GetVideo 得到 VideoInfo, 然后通过 VideoInfo.DownURL 来下载
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

var (
	errRespBeginCode = []byte(`{"errcode":`)
	errRespBeginMsg  = []byte(`{"errmsg":"`)
)

// DownloadToWriter 下载多媒体到 io.Writer.
//  对于视频素材, 先通过 GetVideo 得到 VideoInfo, 然后通过 VideoInfo.DownURL 来下载
func DownloadToWriter(clt *core.Client, mediaId string, writer io.Writer) (written int64, err error) {
	var request = struct {
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
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	// fuck, 騰訊這次又蛋疼了, Content-Type 不能區分返回的是媒體類型還是錯誤
	var respBegin [11]byte // {"errcode": or {"errmsg":"

	n, err := io.ReadFull(httpResp.Body, respBegin[:])
	switch {
	case err == nil:
		break
	case err == io.ErrUnexpectedEOF: // 很小的媒体, 基本不会出现
		n, err = writer.Write(respBegin[:n])
		written = int64(n)
		return
	case err == io.EOF: // 返回空的body, 基本不会出现
		err = nil
		return
	default:
		return
	}

	httpRespBody := io.MultiReader(bytes.NewReader(respBegin[:]), httpResp.Body)

	if !bytes.Equal(respBegin[:], errRespBeginCode) && !bytes.Equal(respBegin[:], errRespBeginMsg) { // 返回的是媒體內容
		return io.Copy(writer, httpRespBody)
	}

	// 返回的是错误信息
	var result core.Error
	if err = json.NewDecoder(httpRespBody).Decode(&result); err != nil {
		return
	}

	switch result.ErrCode {
	case core.ErrCodeOK:
		return // 基本不会出现
	case core.ErrCodeInvalidCredential, core.ErrCodeAccessTokenExpired: // 失效(过期)重试一次
		if !hasRetried {
			hasRetried = true

			if token, err = clt.TokenRefresh(); err != nil {
				return
			}

			result = core.Error{}
			goto RETRY
		}
		fallthrough
	default:
		err = &result
		return
	}
}
