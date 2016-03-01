package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
)

type MultipartFormField struct {
	ContentType int // 0:文件field, 1:普通的文本field
	FieldName   string
	FileName    string // ContentType == 0 时有效
	Value       io.Reader
}

// PostMultipartForm 通用上传接口.
//
//  --BOUNDARY
//  Content-Disposition: form-data; name="FIELDNAME"; filename="FILENAME"
//  Content-Type: application/octet-stream
//
//  FILE-CONTENT
//  --BOUNDARY
//  Content-Disposition: form-data; name="FIELDNAME"
//
//  JSON-DESCRIPTION
//  --BOUNDARY--
//
//
//  NOTE:
//  1. 一般不需要调用这个方法, 请直接调用高层次的封装方法;
//  2. 最终的 URL == incompleteURL + access_token;
//  3. response 格式有要求, 要么是 *Error, 要么是下面结构体的指针(注意 Error 必须是第一个 Field):
//      struct {
//          Error
//          ...
//      }
func (clt *Client) PostMultipartForm(incompleteURL string, fields []MultipartFormField, response interface{}) (err error) {
	responseValue := reflect.ValueOf(response)
	if responseValue.Kind() != reflect.Ptr {
		panic("the type of response is incorrect")
	}
	responseStructValue := responseValue.Elem()
	if responseStructValue.Kind() != reflect.Struct {
		panic("the type of response is incorrect")
	}

	var ErrorStructValue reflect.Value // Error
	if t := responseStructValue.Type(); t == errorType {
		ErrorStructValue = responseStructValue
	} else {
		if t.NumField() == 0 {
			panic("the type of response is incorrect")
		}
		v := responseStructValue.Field(0)
		if v.Type() != errorType {
			panic("the type of response is incorrect")
		}
		ErrorStructValue = v
	}
	ErrorErrCodeValue := ErrorStructValue.Field(errorErrCodeIndex)

	bodyBuf := mediaBufferPool.Get().(*bytes.Buffer)
	bodyBuf.Reset()
	defer mediaBufferPool.Put(bodyBuf)

	multipartWriter := multipart.NewWriter(bodyBuf)
	for _, field := range fields {
		switch field.ContentType {
		case 0: // 文件
			partWriter, err := multipartWriter.CreateFormFile(field.FieldName, field.FileName)
			if err != nil {
				return err
			}
			if _, err = io.Copy(partWriter, field.Value); err != nil {
				return err
			}
		case 1: // 文本
			partWriter, err := multipartWriter.CreateFormField(field.FieldName)
			if err != nil {
				return err
			}
			if _, err = io.Copy(partWriter, field.Value); err != nil {
				return err
			}
		}
	}
	if err = multipartWriter.Close(); err != nil {
		return
	}
	bodyBytes := bodyBuf.Bytes()

	token, err := clt.Token()
	if err != nil {
		return
	}

	hasRetried := false
RETRY:
	finalURL := incompleteURL + url.QueryEscape(token)

	httpResp, err := clt.HttpClient.Post(finalURL, multipartWriter.FormDataContentType(), bytes.NewReader(bodyBytes))
	if err != nil {
		return
	}
	defer httpResp.Body.Close() // 在发生 RETRY 的场景下可以正确工作

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}

	if err = json.NewDecoder(httpResp.Body).Decode(response); err != nil {
		return
	}

	switch errCode := ErrorErrCodeValue.Int(); errCode {
	case ErrCodeOK:
		return
	case ErrCodeInvalidCredential, ErrCodeAccessTokenExpired:
		if !hasRetried {
			hasRetried = true
			responseStructValue.Set(reflect.Zero(responseStructValue.Type()))
			if token, err = clt.TokenRefresh(); err != nil {
				return
			}
			goto RETRY
		}
		fallthrough
	default:
		return
	}
}
