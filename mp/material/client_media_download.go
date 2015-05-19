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
	"net/http"
	"net/url"
	"os"

	"github.com/chanxuehong/wechat/mp"
)

// 下载多媒体到文件.
func (clt Client) DownloadMaterial(mediaId, filepath string) (err error) {
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

	return clt.downloadMaterialToWriter(mediaId, file)
}

// 下载多媒体到 io.Writer.
func (clt Client) DownloadMaterialToWriter(mediaId string, writer io.Writer) error {
	if writer == nil {
		return errors.New("nil writer")
	}
	return clt.downloadMaterialToWriter(mediaId, writer)
}

var (
	errRespBeginCode = []byte(`{"errcode":`)
	errRespBeginMsg  = []byte(`{"errmsg":"`)
)

// 下载多媒体到 io.Writer.
func (clt Client) downloadMaterialToWriter(mediaId string, writer io.Writer) (err error) {
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
	finalURL := "https://api.weixin.qq.com/cgi-bin/material/get_material?access_token=" + url.QueryEscape(string(token))

	httpResp, err := clt.HttpClient.Post(finalURL, "application/json; charset=utf-8", bytes.NewReader(requestBody))
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}

	// fuck, 騰訊這次又蛋疼了, Content-Type 不能區分返回的是媒體類型還是錯誤
	var respBegin [11]byte // {"errcode": or {"errmsg":"

	n, err := io.ReadFull(httpResp.Body, respBegin[:])
	switch {
	case err == nil:
		break
	case err == io.ErrUnexpectedEOF:
		_, err = writer.Write(respBegin[:n])
		return
	case err == io.EOF:
		err = nil
		return
	default:
		return
	}

	httpRespBody := io.MultiReader(bytes.NewReader(respBegin[:]), httpResp.Body)

	if !bytes.Equal(respBegin[:], errRespBeginCode) && !bytes.Equal(respBegin[:], errRespBeginMsg) { // 返回的是媒體內容
		_, err = io.Copy(writer, httpRespBody)
		return
	}

	// 返回的是错误信息
	var result mp.Error
	if err = json.NewDecoder(httpRespBody).Decode(&result); err != nil {
		return
	}

	switch result.ErrCode {
	case mp.ErrCodeOK:
		return // 基本不会出现
	case mp.ErrCodeInvalidCredential, mp.ErrCodeAccessTokenExpired: // 失效(过期)重试一次
		mp.LogInfoln("[WECHAT_RETRY] err_code:", result.ErrCode, ", err_msg:", result.ErrMsg)
		mp.LogInfoln("[WECHAT_RETRY] current token:", token)

		if !hasRetried {
			hasRetried = true

			if token, err = clt.TokenRefresh(); err != nil {
				return
			}
			mp.LogInfoln("[WECHAT_RETRY] new token:", token)

			result = mp.Error{}
			goto RETRY
		}
		mp.LogInfoln("[WECHAT_RETRY] fallthrough, current token:", token)
		fallthrough
	default:
		err = &result
		return
	}
}
