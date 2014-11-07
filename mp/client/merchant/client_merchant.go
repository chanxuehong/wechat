// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package merchant

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	multipart_boundary    = "--------wvm6LNx=y4rEq?BUD(k_:0Pj2V.M'J)t957K-Sh/Q1ZA+ceWFunTRdfGaXgY"
	multipart_ContentType = "multipart/form-data; boundary=" + multipart_boundary

	// ----------wvm6LNx=y4rEq?BUD(k_:0Pj2V.M'J)t957K-Sh/Q1ZA+ceWFunTRdfGaXgY
	// Content-Disposition: form-data; name="upload"; filename="filename"
	// Content-Type: application/octet-stream
	//
	// mediaReader
	// ----------wvm6LNx=y4rEq?BUD(k_:0Pj2V.M'J)t957K-Sh/Q1ZA+ceWFunTRdfGaXgY--
	//
	multipart_formDataFront = "--" + multipart_boundary +
		"\r\nContent-Disposition: form-data; name=\"upload\"; filename=\""
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

// 上传图片
func (c *Client) MerchantUploadImage(filepath_ string) (imageURL string, err error) {
	file, err := os.Open(filepath_)
	if err != nil {
		return
	}
	defer file.Close()

	return c.merchantUploadImageFromReader(filepath.Base(filepath_), file)
}

// 上传图片
//  NOTE: 参数 filename 不是文件路径, 是指定 multipart form 里面文件名称
func (c *Client) MerchantUploadImageFromReader(filename string, imageReader io.Reader) (imageURL string, err error) {
	if filename == "" {
		err = errors.New(`filename == ""`)
		return
	}
	if imageReader == nil {
		err = errors.New("imageReader == nil")
		return
	}

	return c.merchantUploadImageFromReader(filename, imageReader)
}

// 上传图片
func (c *Client) merchantUploadImageFromReader(filename string, reader io.Reader) (imageURL string, err error) {
	switch v := reader.(type) {
	case *os.File:
		return c.merchantUploadImageFromOSFile(filename, v)
	case *bytes.Buffer:
		return c.merchantUploadImageFromBytesBuffer(filename, v)
	case *bytes.Reader:
		return c.merchantUploadImageFromBytesReader(filename, v)
	case *strings.Reader:
		return c.merchantUploadImageFromStringsReader(filename, v)
	default:
		return c.merchantUploadImageFromIOReader(filename, v)
	}
}

func (c *Client) merchantUploadImageFromOSFile(filename string, file *os.File) (imageURL string, err error) {
	fi, err := file.Stat()
	if err != nil {
		return
	}

	// 非常规文件, FileInfo.Size() 不一定准确
	if !fi.Mode().IsRegular() {
		return c.merchantUploadImageFromIOReader(filename, file)
	}

	originalOffset, err := file.Seek(0, 1)
	if err != nil {
		return
	}

	FormDataFileName := escapeQuotes(filename)
	ContentLength := int64(multipart_constPartLen+len(FormDataFileName)) +
		fi.Size() - originalOffset

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	url_ := merchantUploadImageURL(token, filename)

	if hasRetry {
		if _, err = file.Seek(originalOffset, 0); err != nil {
			return
		}
	}
	mr := io.MultiReader(
		strings.NewReader(multipart_formDataFront),
		strings.NewReader(FormDataFileName),
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

	var result struct {
		Error
		ImageURL string `json:"image_url"`
	}
	if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		imageURL = result.ImageURL
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

func (c *Client) merchantUploadImageFromBytesBuffer(filename string, buffer *bytes.Buffer) (imageURL string, err error) {
	fileBytes := buffer.Bytes()

	FormDataFileName := escapeQuotes(filename)
	ContentLength := int64(multipart_constPartLen + len(FormDataFileName) + len(fileBytes))

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	url_ := merchantUploadImageURL(token, filename)

	mr := io.MultiReader(
		strings.NewReader(multipart_formDataFront),
		strings.NewReader(FormDataFileName),
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

	var result struct {
		Error
		ImageURL string `json:"image_url"`
	}
	if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		imageURL = result.ImageURL
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

func (c *Client) merchantUploadImageFromBytesReader(filename string, reader *bytes.Reader) (imageURL string, err error) {
	originalOffset, err := reader.Seek(0, 1)
	if err != nil {
		return
	}

	FormDataFileName := escapeQuotes(filename)
	ContentLength := int64(multipart_constPartLen + len(FormDataFileName) + reader.Len())

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	url_ := merchantUploadImageURL(token, filename)

	if hasRetry {
		if _, err = reader.Seek(originalOffset, 0); err != nil {
			return
		}
	}
	mr := io.MultiReader(
		strings.NewReader(multipart_formDataFront),
		strings.NewReader(FormDataFileName),
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

	var result struct {
		Error
		ImageURL string `json:"image_url"`
	}
	if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		imageURL = result.ImageURL
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

func (c *Client) merchantUploadImageFromStringsReader(filename string, reader *strings.Reader) (imageURL string, err error) {
	originalOffset, err := reader.Seek(0, 1)
	if err != nil {
		return
	}

	FormDataFileName := escapeQuotes(filename)
	ContentLength := int64(multipart_constPartLen + len(FormDataFileName) + reader.Len())

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	url_ := merchantUploadImageURL(token, filename)

	if hasRetry {
		if _, err = reader.Seek(originalOffset, 0); err != nil {
			return
		}
	}
	mr := io.MultiReader(
		strings.NewReader(multipart_formDataFront),
		strings.NewReader(FormDataFileName),
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

	var result struct {
		Error
		ImageURL string `json:"image_url"`
	}
	if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		imageURL = result.ImageURL
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

func (c *Client) merchantUploadImageFromIOReader(filename string, reader io.Reader) (imageURL string, err error) {
	bodyBuf := mediaBufferPool.Get().(*bytes.Buffer) // io.ReadWriter
	bodyBuf.Reset()                                  // important
	defer mediaBufferPool.Put(bodyBuf)               // important

	bodyBuf.WriteString(multipart_formDataFront)
	bodyBuf.WriteString(escapeQuotes(filename))
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
	url_ := merchantUploadImageURL(token, filename)

	httpResp, err := c.httpClient.Post(url_, multipart_ContentType, bytes.NewReader(bodyBytes))
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	var result struct {
		Error
		ImageURL string `json:"image_url"`
	}
	if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		imageURL = result.ImageURL
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
