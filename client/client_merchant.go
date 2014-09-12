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
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
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
	bodyBuf := mediaBufferPool.Get().(*bytes.Buffer) // io.ReadWriter
	bodyBuf.Reset()                                  // important
	defer mediaBufferPool.Put(bodyBuf)

	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		return
	}
	if _, err = io.Copy(fileWriter, imageReader); err != nil {
		return
	}

	bodyContentType := bodyWriter.FormDataContentType()

	if err = bodyWriter.Close(); err != nil {
		return
	}

	postContent := bodyBuf.Bytes() // 这么绕一下是为了 RETRY 的时候能正常工作

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := merchantUploadImageURL(token, filename)
	resp, err := c.httpClient.Post(_url, bodyContentType, bytes.NewReader(postContent))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", resp.Status)
		return
	}

	var result struct {
		Error
		ImageURL string `json:"image_url"`
	}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
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
