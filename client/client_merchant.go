package client

import (
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
func (c *Client) MerchantUploadImageFromFile(filePath string) (imageURL string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	return c.MerchantUploadImage(filepath.Base(filePath), file)
}

// 上传图片
func (c *Client) MerchantUploadImage(filename string, imageReader io.Reader) (imageURL string, err error) {
	if filename == "" {
		err = errors.New(`filename == ""`)
		return
	}
	if imageReader == nil {
		err = errors.New("imageReader == nil")
		return
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	bodyBuf := c.getBufferFromPool() // io.ReadWriter
	defer c.putBufferToPool(bodyBuf) // important!

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

	_url := merchantUploadImageURL(token, filename)
	resp, err := c.httpClient.Post(_url, bodyContentType, bodyBuf)
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
	if result.ErrCode != 0 {
		err = &result.Error
		return
	}

	imageURL = result.ImageURL
	return
}
