// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/chanxuehong/wechat/mp/media"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	multipart_boundary    = "--------wvm6LNx=y4rEq?BUD(k_:0Pj2V.M'J)t957K-Sh/Q1ZA+ceWFunTRdfGaXgY"
	multipart_ContentType = "multipart/form-data; boundary=" + multipart_boundary

	// ----------wvm6LNx=y4rEq?BUD(k_:0Pj2V.M'J)t957K-Sh/Q1ZA+ceWFunTRdfGaXgY
	// Content-Disposition: form-data; name="filename"; filename="filename"
	// Content-Type: application/octet-stream
	//
	// mediaReader
	// ----------wvm6LNx=y4rEq?BUD(k_:0Pj2V.M'J)t957K-Sh/Q1ZA+ceWFunTRdfGaXgY--
	//
	multipart_formDataFront = "--" + multipart_boundary +
		"\r\nContent-Disposition: form-data; name=\"filename\"; filename=\""
	multipart_formDataMiddle = "\"\r\nContent-Type: application/octet-stream\r\n\r\n"
	multipart_formDataEnd    = "\r\n--" + multipart_boundary + "--\r\n"

	multipart_constPartLen = len(multipart_formDataFront) +
		len(multipart_formDataMiddle) + len(multipart_formDataEnd)
)

// copy from mime/multipart/writer.go
var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

// copy from mime/multipart/writer.go
func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

// 上传多媒体
func (c *Client) mediaUploadFromReader(mediaType, filename string, reader io.Reader) (info *media.MediaInfo, err error) {
	filename = escapeQuotes(filename)

	switch v := reader.(type) {
	case *os.File:
		return c.mediaUploadFromOSFile(mediaType, filename, v)
	case *bytes.Buffer:
		return c.mediaUploadFromBytesBuffer(mediaType, filename, v)
	case *bytes.Reader:
		return c.mediaUploadFromBytesReader(mediaType, filename, v)
	case *strings.Reader:
		return c.mediaUploadFromStringsReader(mediaType, filename, v)
	default:
		return c.mediaUploadFromIOReader(mediaType, filename, v)
	}
}

func (c *Client) mediaUploadFromOSFile(mediaType, filename string, file *os.File) (info *media.MediaInfo, err error) {
	fi, err := file.Stat()
	if err != nil {
		return
	}

	// 非常规文件, FileInfo.Size() 不一定准确
	if !fi.Mode().IsRegular() {
		return c.mediaUploadFromIOReader(mediaType, filename, file)
	}

	originalOffset, err := file.Seek(0, 1)
	if err != nil {
		return
	}
	ContentLength := int64(multipart_constPartLen+len(filename)) +
		fi.Size() - originalOffset

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	url_ := mediaUploadURL(token, mediaType)

	if hasRetry {
		if _, err = file.Seek(originalOffset, 0); err != nil {
			return
		}
	}
	mr := io.MultiReader(
		strings.NewReader(multipart_formDataFront),
		strings.NewReader(filename),
		strings.NewReader(multipart_formDataMiddle),
		file,
		strings.NewReader(multipart_formDataEnd),
	)

	httpReq, err := http.NewRequest("POST", url_, mr)
	if err != nil {
		return
	}
	httpReq.Header.Set("Content-Type", multipart_ContentType)
	httpReq.ContentLength = ContentLength

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	switch mediaType {
	case media.MEDIA_TYPE_THUMB: // 返回的是 thumb_media_id 而不是 media_id
		var result struct {
			Error
			MediaType string `json:"type"`
			MediaId   string `json:"thumb_media_id"`
			CreatedAt int64  `json:"created_at"`
		}
		if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
			return
		}

		switch result.ErrCode {
		case errCodeOK:
			info = &media.MediaInfo{
				MediaType: result.MediaType,
				MediaId:   result.MediaId,
				CreatedAt: result.CreatedAt,
			}
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

	default:
		var result struct {
			Error
			media.MediaInfo
		}
		if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
			return
		}

		switch result.ErrCode {
		case errCodeOK:
			info = &result.MediaInfo
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
}

func (c *Client) mediaUploadFromBytesBuffer(mediaType, filename string, buffer *bytes.Buffer) (info *media.MediaInfo, err error) {
	fileBytes := buffer.Bytes()
	ContentLength := int64(multipart_constPartLen + len(filename) + len(fileBytes))

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	url_ := mediaUploadURL(token, mediaType)

	mr := io.MultiReader(
		strings.NewReader(multipart_formDataFront),
		strings.NewReader(filename),
		strings.NewReader(multipart_formDataMiddle),
		bytes.NewReader(fileBytes),
		strings.NewReader(multipart_formDataEnd),
	)

	httpReq, err := http.NewRequest("POST", url_, mr)
	if err != nil {
		return
	}
	httpReq.Header.Set("Content-Type", multipart_ContentType)
	httpReq.ContentLength = ContentLength

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	switch mediaType {
	case media.MEDIA_TYPE_THUMB: // 返回的是 thumb_media_id 而不是 media_id
		var result struct {
			Error
			MediaType string `json:"type"`
			MediaId   string `json:"thumb_media_id"`
			CreatedAt int64  `json:"created_at"`
		}
		if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
			return
		}

		switch result.ErrCode {
		case errCodeOK:
			info = &media.MediaInfo{
				MediaType: result.MediaType,
				MediaId:   result.MediaId,
				CreatedAt: result.CreatedAt,
			}
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

	default:
		var result struct {
			Error
			media.MediaInfo
		}
		if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
			return
		}

		switch result.ErrCode {
		case errCodeOK:
			info = &result.MediaInfo
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
}

func (c *Client) mediaUploadFromBytesReader(mediaType, filename string, reader *bytes.Reader) (info *media.MediaInfo, err error) {
	originalOffset, err := reader.Seek(0, 1)
	if err != nil {
		return
	}
	ContentLength := int64(multipart_constPartLen + len(filename) + reader.Len())

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	url_ := mediaUploadURL(token, mediaType)

	if hasRetry {
		if _, err = reader.Seek(originalOffset, 0); err != nil {
			return
		}
	}
	mr := io.MultiReader(
		strings.NewReader(multipart_formDataFront),
		strings.NewReader(filename),
		strings.NewReader(multipart_formDataMiddle),
		reader,
		strings.NewReader(multipart_formDataEnd),
	)

	httpReq, err := http.NewRequest("POST", url_, mr)
	if err != nil {
		return
	}
	httpReq.Header.Set("Content-Type", multipart_ContentType)
	httpReq.ContentLength = ContentLength

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	switch mediaType {
	case media.MEDIA_TYPE_THUMB: // 返回的是 thumb_media_id 而不是 media_id
		var result struct {
			Error
			MediaType string `json:"type"`
			MediaId   string `json:"thumb_media_id"`
			CreatedAt int64  `json:"created_at"`
		}
		if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
			return
		}

		switch result.ErrCode {
		case errCodeOK:
			info = &media.MediaInfo{
				MediaType: result.MediaType,
				MediaId:   result.MediaId,
				CreatedAt: result.CreatedAt,
			}
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

	default:
		var result struct {
			Error
			media.MediaInfo
		}
		if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
			return
		}

		switch result.ErrCode {
		case errCodeOK:
			info = &result.MediaInfo
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
}

func (c *Client) mediaUploadFromStringsReader(mediaType, filename string, reader *strings.Reader) (info *media.MediaInfo, err error) {
	originalOffset, err := reader.Seek(0, 1)
	if err != nil {
		return
	}
	ContentLength := int64(multipart_constPartLen + len(filename) + reader.Len())

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	url_ := mediaUploadURL(token, mediaType)

	if hasRetry {
		if _, err = reader.Seek(originalOffset, 0); err != nil {
			return
		}
	}
	mr := io.MultiReader(
		strings.NewReader(multipart_formDataFront),
		strings.NewReader(filename),
		strings.NewReader(multipart_formDataMiddle),
		reader,
		strings.NewReader(multipart_formDataEnd),
	)

	httpReq, err := http.NewRequest("POST", url_, mr)
	if err != nil {
		return
	}
	httpReq.Header.Set("Content-Type", multipart_ContentType)
	httpReq.ContentLength = ContentLength

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	switch mediaType {
	case media.MEDIA_TYPE_THUMB: // 返回的是 thumb_media_id 而不是 media_id
		var result struct {
			Error
			MediaType string `json:"type"`
			MediaId   string `json:"thumb_media_id"`
			CreatedAt int64  `json:"created_at"`
		}
		if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
			return
		}

		switch result.ErrCode {
		case errCodeOK:
			info = &media.MediaInfo{
				MediaType: result.MediaType,
				MediaId:   result.MediaId,
				CreatedAt: result.CreatedAt,
			}
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

	default:
		var result struct {
			Error
			media.MediaInfo
		}
		if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
			return
		}

		switch result.ErrCode {
		case errCodeOK:
			info = &result.MediaInfo
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
}

func (c *Client) mediaUploadFromIOReader(mediaType, filename string, reader io.Reader) (info *media.MediaInfo, err error) {
	bodyBuf := mediaBufferPool.Get().(*bytes.Buffer) // io.ReadWriter
	bodyBuf.Reset()                                  // important
	defer mediaBufferPool.Put(bodyBuf)               // important

	bodyBuf.WriteString(multipart_formDataFront)
	bodyBuf.WriteString(filename)
	bodyBuf.WriteString(multipart_formDataMiddle)
	if _, err = io.Copy(bodyBuf, reader); err != nil {
		return
	}
	bodyBuf.WriteString(multipart_formDataEnd)

	bodyBytes := bodyBuf.Bytes()

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	url_ := mediaUploadURL(token, mediaType)

	httpResp, err := c.httpClient.Post(url_, multipart_ContentType, bytes.NewReader(bodyBytes))
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	switch mediaType {
	case media.MEDIA_TYPE_THUMB: // 返回的是 thumb_media_id 而不是 media_id
		var result struct {
			Error
			MediaType string `json:"type"`
			MediaId   string `json:"thumb_media_id"`
			CreatedAt int64  `json:"created_at"`
		}
		if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
			return
		}

		switch result.ErrCode {
		case errCodeOK:
			info = &media.MediaInfo{
				MediaType: result.MediaType,
				MediaId:   result.MediaId,
				CreatedAt: result.CreatedAt,
			}
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

	default:
		var result struct {
			Error
			media.MediaInfo
		}
		if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
			return
		}

		switch result.ErrCode {
		case errCodeOK:
			info = &result.MediaInfo
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
}
