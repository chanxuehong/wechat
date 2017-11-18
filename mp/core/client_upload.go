package core

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"gopkg.in/chanxuehong/wechat.v2/internal/debug/api"
	"gopkg.in/chanxuehong/wechat.v2/internal/debug/api/retry"
	"gopkg.in/chanxuehong/wechat.v2/util"
)

type MultipartFormField struct {
	IsFile   bool
	Name     string
	FileName string
	Value    io.Reader
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
//  1. 一般不需要调用这个方法, 请直接调用高层次的封装函数;
//  2. 最终的 URL == incompleteURL + access_token;
//  3. response 格式有要求, 要么是 *Error, 要么是下面结构体的指针(注意 Error 必须是第一个 Field):
//      struct {
//          Error
//          ...
//      }
func (clt *Client) PostMultipartForm(incompleteURL string, fields []MultipartFormField, response interface{}) (err error) {
	ErrorStructValue, ErrorErrCodeValue := checkResponse(response)

	buffer := mediaBufferPool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer mediaBufferPool.Put(buffer)

	multipartWriter := multipart.NewWriter(buffer)
	for i := 0; i < len(fields); i++ {
		if field := &fields[i]; field.IsFile {
			partWriter, err3 := multipartWriter.CreateFormFile(field.Name, field.FileName)
			if err3 != nil {
				return err3
			}
			if _, err3 = io.Copy(partWriter, field.Value); err3 != nil {
				return err3
			}
		} else {
			partWriter, err3 := multipartWriter.CreateFormField(field.Name)
			if err3 != nil {
				return err3
			}
			if _, err3 = io.Copy(partWriter, field.Value); err3 != nil {
				return err3
			}
		}
	}
	if err = multipartWriter.Close(); err != nil {
		return
	}
	requestBodyBytes := buffer.Bytes()
	requestBodyType := multipartWriter.FormDataContentType()

	httpClient := clt.HttpClient
	if httpClient == nil {
		httpClient = util.DefaultMediaHttpClient
	}

	token, err := clt.Token()
	if err != nil {
		return
	}

	hasRetried := false
RETRY:
	finalURL := incompleteURL + url.QueryEscape(token)
	if err = httpPostMultipartForm(httpClient, finalURL, requestBodyType, requestBodyBytes, response); err != nil {
		return
	}

	switch errCode := ErrorErrCodeValue.Int(); errCode {
	case ErrCodeOK:
		return
	case ErrCodeInvalidCredential, ErrCodeAccessTokenExpired:
		errMsg := ErrorStructValue.Field(errorErrMsgIndex).String()
		retry.DebugPrintError(errCode, errMsg, token)
		if !hasRetried {
			hasRetried = true
			ErrorStructValue.Set(errorZeroValue)
			if token, err = clt.RefreshToken(token); err != nil {
				return
			}
			retry.DebugPrintNewToken(token)
			goto RETRY
		}
		retry.DebugPrintFallthrough(token)
		fallthrough
	default:
		return
	}
}

func httpPostMultipartForm(clt *http.Client, url, bodyType string, body []byte, response interface{}) error {
	api.DebugPrintPostMultipartRequest(url, body)
	httpResp, err := clt.Post(url, bodyType, bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}
	return api.DecodeJSONHttpResponse(httpResp.Body, response)
}
