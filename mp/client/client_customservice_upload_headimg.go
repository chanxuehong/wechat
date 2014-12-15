// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// 上传客服头像
func (c *Client) CustomServiceKFAccountUploadHeadImage(kfAccount, imagePath string) (err error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return
	}
	defer file.Close()

	return c.CustomServiceKFAccountUploadHeadImageFromReader(kfAccount, filepath.Base(imagePath), file)
}

// 上传客服头像
func (c *Client) CustomServiceKFAccountUploadHeadImageFromReader(kfAccount, filename string, reader io.Reader) (err error) {
	filename = escapeQuotes(filename)

	switch v := reader.(type) {
	case *os.File:
		return c.customServiceKFAccountUploadHeadImageFromOSFile(kfAccount, filename, v)
	case *bytes.Buffer:
		return c.customServiceKFAccountUploadHeadImageFromBytesBuffer(kfAccount, filename, v)
	case *bytes.Reader:
		return c.customServiceKFAccountUploadHeadImageFromBytesReader(kfAccount, filename, v)
	case *strings.Reader:
		return c.customServiceKFAccountUploadHeadImageFromStringsReader(kfAccount, filename, v)
	default:
		return c.customServiceKFAccountUploadHeadImageFromIOReader(kfAccount, filename, v)
	}
}

func (c *Client) customServiceKFAccountUploadHeadImageFromOSFile(kfAccount, filename string, file *os.File) (err error) {
	fi, err := file.Stat()
	if err != nil {
		return
	}

	// 非常规文件, FileInfo.Size() 不一定准确
	if !fi.Mode().IsRegular() {
		return c.customServiceKFAccountUploadHeadImageFromIOReader(kfAccount, filename, file)
	}

	originalOffset, err := file.Seek(0, 1)
	if err != nil {
		return
	}
	ContentLength := int64(multipart_constPartLen+len(filename)) +
		fi.Size() - originalOffset

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := customServiceKFAccountUploadHeadImgURL(token, kfAccount)

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

	var result Error

	if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return
	case errCodeInvalidCredential, errCodeTimeout:
		if !hasRetry {
			hasRetry = true

			if token, err = getNewToken(c.tokenService, token); err != nil {
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

func (c *Client) customServiceKFAccountUploadHeadImageFromBytesBuffer(kfAccount, filename string, buffer *bytes.Buffer) (err error) {
	fileBytes := buffer.Bytes()
	ContentLength := int64(multipart_constPartLen + len(filename) + len(fileBytes))

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := customServiceKFAccountUploadHeadImgURL(token, kfAccount)

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

	var result Error

	if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return
	case errCodeInvalidCredential, errCodeTimeout:
		if !hasRetry {
			hasRetry = true

			if token, err = getNewToken(c.tokenService, token); err != nil {
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

func (c *Client) customServiceKFAccountUploadHeadImageFromBytesReader(kfAccount, filename string, reader *bytes.Reader) (err error) {
	originalOffset, err := reader.Seek(0, 1)
	if err != nil {
		return
	}
	ContentLength := int64(multipart_constPartLen + len(filename) + reader.Len())

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := customServiceKFAccountUploadHeadImgURL(token, kfAccount)

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

	var result Error

	if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return
	case errCodeInvalidCredential, errCodeTimeout:
		if !hasRetry {
			hasRetry = true

			if token, err = getNewToken(c.tokenService, token); err != nil {
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

func (c *Client) customServiceKFAccountUploadHeadImageFromStringsReader(kfAccount, filename string, reader *strings.Reader) (err error) {
	originalOffset, err := reader.Seek(0, 1)
	if err != nil {
		return
	}
	ContentLength := int64(multipart_constPartLen + len(filename) + reader.Len())

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := customServiceKFAccountUploadHeadImgURL(token, kfAccount)

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

	var result Error

	if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return
	case errCodeInvalidCredential, errCodeTimeout:
		if !hasRetry {
			hasRetry = true

			if token, err = getNewToken(c.tokenService, token); err != nil {
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

func (c *Client) customServiceKFAccountUploadHeadImageFromIOReader(kfAccount, filename string, reader io.Reader) (err error) {
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

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := customServiceKFAccountUploadHeadImgURL(token, kfAccount)

	httpResp, err := c.httpClient.Post(url_, multipart_ContentType, bytes.NewReader(bodyBytes))
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	var result Error

	if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return
	case errCodeInvalidCredential, errCodeTimeout:
		if !hasRetry {
			hasRetry = true

			if token, err = getNewToken(c.tokenService, token); err != nil {
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
