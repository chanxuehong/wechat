// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

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

// 上传图片
func (c *Client) MerchantUploadImageFromFile(_filepath string) (imageURL string, err error) {
	file, err := os.Open(_filepath)
	if err != nil {
		return
	}
	defer file.Close()

	return c.merchantUploadImage(filepath.Base(_filepath), file)
}

// 上传图片
//  NOTE: 参数 filename 不是文件路径, 是指定 multipart form 里面文件名称
func (c *Client) MerchantUploadImage(filename string, imageReader io.Reader) (imageURL string, err error) {
	if filename == "" {
		err = errors.New(`filename == ""`)
		return
	}
	if imageReader == nil {
		err = errors.New("imageReader == nil")
		return
	}
	return c.merchantUploadImage(filename, imageReader)
}

// 上传图片
func (c *Client) merchantUploadImage(filename string, imageReader io.Reader) (imageURL string, err error) {
	const boundary = "--------wvm6LNx=y4rEq?BUD(k_:0Pj2V.M'J)t957K-Sh/Q1ZA+ceWFunTRdfGaXgY"
	const ContentType = "multipart/form-data; boundary=" + boundary

	// ----------wvm6LNx=y4rEq?BUD(k_:0Pj2V.M'J)t957K-Sh/Q1ZA+ceWFunTRdfGaXgY
	// Content-Disposition: form-data; name="filename"; filename="filename"
	// Content-Type: application/octet-stream
	//
	// imageReader
	// ----------wvm6LNx=y4rEq?BUD(k_:0Pj2V.M'J)t957K-Sh/Q1ZA+ceWFunTRdfGaXgY--
	//
	const formDataFront = "--" + boundary + "\r\nContent-Disposition: form-data; name=\"filename\"; filename=\""
	filename = escapeQuotes(filename)
	const formDataMiddle = "\"\r\nContent-Type: application/octet-stream\r\n\r\n"
	// imageReader
	const formDataEnd = "\r\n--" + boundary + "--\r\n"

	token, err := c.Token()
	if err != nil {
		return
	}
	url_ := merchantUploadImageURL(token, filename)

	var httpReq *http.Request

	switch v := imageReader.(type) {
	case *os.File:
		var fi os.FileInfo
		if fi, err = v.Stat(); err != nil {
			return
		}

		mr := io.MultiReader(
			strings.NewReader(formDataFront),
			strings.NewReader(filename),
			strings.NewReader(formDataMiddle),
			imageReader,
			strings.NewReader(formDataEnd),
		)

		httpReq, err = http.NewRequest("POST", url_, mr)
		if err != nil {
			return
		}
		httpReq.Header.Set("Content-Type", ContentType)
		httpReq.ContentLength = int64(len(formDataFront)+len(filename)+
			len(formDataMiddle)+len(formDataEnd)) + fi.Size()

	case *bytes.Buffer:
		mr := io.MultiReader(
			strings.NewReader(formDataFront),
			strings.NewReader(filename),
			strings.NewReader(formDataMiddle),
			imageReader,
			strings.NewReader(formDataEnd),
		)

		httpReq, err = http.NewRequest("POST", url_, mr)
		if err != nil {
			return
		}
		httpReq.Header.Set("Content-Type", ContentType)
		httpReq.ContentLength = int64(len(formDataFront) + len(filename) +
			len(formDataMiddle) + len(formDataEnd) + v.Len())

	case *bytes.Reader:
		mr := io.MultiReader(
			strings.NewReader(formDataFront),
			strings.NewReader(filename),
			strings.NewReader(formDataMiddle),
			imageReader,
			strings.NewReader(formDataEnd),
		)

		httpReq, err = http.NewRequest("POST", url_, mr)
		if err != nil {
			return
		}
		httpReq.Header.Set("Content-Type", ContentType)
		httpReq.ContentLength = int64(len(formDataFront) + len(filename) +
			len(formDataMiddle) + len(formDataEnd) + v.Len())

	case *strings.Reader:
		mr := io.MultiReader(
			strings.NewReader(formDataFront),
			strings.NewReader(filename),
			strings.NewReader(formDataMiddle),
			imageReader,
			strings.NewReader(formDataEnd),
		)

		httpReq, err = http.NewRequest("POST", url_, mr)
		if err != nil {
			return
		}
		httpReq.Header.Set("Content-Type", ContentType)
		httpReq.ContentLength = int64(len(formDataFront) + len(filename) +
			len(formDataMiddle) + len(formDataEnd) + v.Len())

	default:
		bodyBuf := mediaBufferPool.Get().(*bytes.Buffer) // io.ReadWriter
		bodyBuf.Reset()                                  // important
		defer mediaBufferPool.Put(bodyBuf)               // important

		bodyBuf.WriteString(formDataFront)
		bodyBuf.WriteString(filename)
		bodyBuf.WriteString(formDataMiddle)
		if _, err = io.Copy(bodyBuf, imageReader); err != nil {
			return
		}
		bodyBuf.WriteString(formDataEnd)

		httpReq, err = http.NewRequest("POST", url_, bodyBuf)
		if err != nil {
			return
		}
		httpReq.Header.Set("Content-Type", ContentType)
	}

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

	default:
		err = &result.Error
		return
	}
}
