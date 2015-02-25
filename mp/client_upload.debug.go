// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

// +build wechatdebug

package mp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
)

const (
	multipartBoundary    = "--------wvm6LNx=y4rEq?BUD(k_:0Pj2V.M'J)t957K-Sh/Q1ZA+ceWFunTRdfGaXgY"
	multipartContentType = "multipart/form-data; boundary=" + multipartBoundary

	// ----------wvm6LNx=y4rEq?BUD(k_:0Pj2V.M'J)t957K-Sh/Q1ZA+ceWFunTRdfGaXgY
	// Content-Disposition: form-data; name="file"; filename="filename"
	// Content-Type: application/octet-stream
	//
	// filecontent
	// ----------wvm6LNx=y4rEq?BUD(k_:0Pj2V.M'J)t957K-Sh/Q1ZA+ceWFunTRdfGaXgY--
	//
	multipartFormDataFront = "--" + multipartBoundary +
		"\r\nContent-Disposition: form-data; name=\"file\"; filename=\""
	// filename
	multipartFormDataMiddle = "\"\r\nContent-Type: application/octet-stream\r\n\r\n"
	// filecontent
	multipartFormDataEnd = "\r\n--" + multipartBoundary + "--\r\n"

	multipartConstPartLen = len(multipartFormDataFront) +
		len(multipartFormDataMiddle) + len(multipartFormDataEnd)
)

// copy from mime/multipart/writer.go
var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

// copy from mime/multipart/writer.go
func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

// 通用上传接口.
//
//  NOTE:
//  1. 一般不用调用这个方法, 请直接调用高层次的封装方法;
//  2. 最终的 URL == incompleteURL + access_token;
//  3. 参数 filename 不是文件路径, 是指定 multipart/form-data 里面文件名称
//  4. response 要求是 struct 的指针, 并且该 struct 拥有属性:
//     ErrCode int `json:"errcode"` (可以是直接属性, 也可以是匿名属性里的属性)
func (clt *WechatClient) UploadFromReader(incompleteURL, filename string,
	reader io.Reader, response interface{}) (err error) {

	filename = escapeQuotes(filename)
	switch v := reader.(type) {
	case *os.File:
		return clt.uploadFromOSFile(incompleteURL, filename, v, response)
	case *bytes.Buffer:
		return clt.uploadFromBytesBuffer(incompleteURL, filename, v, response)
	case *bytes.Reader:
		return clt.uploadFromBytesReader(incompleteURL, filename, v, response)
	case *strings.Reader:
		return clt.uploadFromStringsReader(incompleteURL, filename, v, response)
	default:
		return clt.uploadFromIOReader(incompleteURL, filename, v, response)
	}
}

func (clt *WechatClient) uploadFromOSFile(incompleteURL, filename string,
	file *os.File, response interface{}) (err error) {

	fi, err := file.Stat()
	if err != nil {
		return
	}

	if !fi.Mode().IsRegular() {
		return clt.uploadFromIOReader(incompleteURL, filename, file, response)
	}

	originalOffset, err := file.Seek(0, 1)
	if err != nil {
		return
	}
	ContentLength := int64(multipartConstPartLen+len(filename)) + fi.Size() - originalOffset

	token, err := clt.Token()
	if err != nil {
		return
	}

	hasRetried := false
RETRY:
	finalURL := incompleteURL + url.QueryEscape(token)

	if hasRetried {
		if _, err = file.Seek(originalOffset, 0); err != nil {
			return
		}
	}
	mr := io.MultiReader(
		strings.NewReader(multipartFormDataFront),
		strings.NewReader(filename),
		strings.NewReader(multipartFormDataMiddle),
		file,
		strings.NewReader(multipartFormDataEnd),
	)

	httpReq, err := http.NewRequest("POST", finalURL, mr)
	if err != nil {
		return
	}
	httpReq.Header.Set("Content-Type", multipartContentType)
	httpReq.ContentLength = ContentLength

	httpResp, err := clt.HttpClient.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return
	}
	fmt.Println("mp.WechatClient.UploadFromReader.incompleteURL:", incompleteURL)
	fmt.Println("mp.WechatClient.UploadFromReader.response:", string(body))

	if err = json.Unmarshal(body, response); err != nil {
		return
	}

	// 请注意:
	// 下面获取 ErrCode 的代码不具备通用性!!!
	//
	// 因为本 SDK 的 response 都是
	//  struct {
	//    Error
	//    XXX
	//  }
	// 的结构, 所以用下面简单的方法得到 ErrCode.
	//
	// 如果你是直接调用这个函数, 那么要根据你的 response 数据结构修改下面的代码.
	ErrCode := reflect.ValueOf(response).Elem().FieldByName("ErrCode").Int()

	switch ErrCode {
	case ErrCodeOK:
		return
	case ErrCodeInvalidCredential, ErrCodeTimeout:
		if !hasRetried {
			hasRetried = true

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

func (clt *WechatClient) uploadFromBytesBuffer(incompleteURL, filename string,
	buffer *bytes.Buffer, response interface{}) (err error) {

	fileBytes := buffer.Bytes()
	ContentLength := int64(multipartConstPartLen + len(filename) + len(fileBytes))

	token, err := clt.Token()
	if err != nil {
		return
	}

	hasRetried := false
RETRY:
	finalURL := incompleteURL + url.QueryEscape(token)

	mr := io.MultiReader(
		strings.NewReader(multipartFormDataFront),
		strings.NewReader(filename),
		strings.NewReader(multipartFormDataMiddle),
		bytes.NewReader(fileBytes),
		strings.NewReader(multipartFormDataEnd),
	)

	httpReq, err := http.NewRequest("POST", finalURL, mr)
	if err != nil {
		return
	}
	httpReq.Header.Set("Content-Type", multipartContentType)
	httpReq.ContentLength = ContentLength

	httpResp, err := clt.HttpClient.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return
	}
	fmt.Println("mp.WechatClient.UploadFromReader.incompleteURL:", incompleteURL)
	fmt.Println("mp.WechatClient.UploadFromReader.response:", string(body))

	if err = json.Unmarshal(body, response); err != nil {
		return
	}

	// 请注意:
	// 下面获取 ErrCode 的代码不具备通用性!!!
	//
	// 因为本 SDK 的 response 都是
	//  struct {
	//    Error
	//    XXX
	//  }
	// 的结构, 所以用下面简单的方法得到 ErrCode.
	//
	// 如果你是直接调用这个函数, 那么要根据你的 response 数据结构修改下面的代码.
	ErrCode := reflect.ValueOf(response).Elem().FieldByName("ErrCode").Int()

	switch ErrCode {
	case ErrCodeOK:
		return
	case ErrCodeInvalidCredential, ErrCodeTimeout:
		if !hasRetried {
			hasRetried = true

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

func (clt *WechatClient) uploadFromBytesReader(incompleteURL, filename string,
	reader *bytes.Reader, response interface{}) (err error) {

	originalOffset, err := reader.Seek(0, 1)
	if err != nil {
		return
	}
	ContentLength := int64(multipartConstPartLen + len(filename) + reader.Len())

	token, err := clt.Token()
	if err != nil {
		return
	}

	hasRetried := false
RETRY:
	finalURL := incompleteURL + url.QueryEscape(token)

	if hasRetried {
		if _, err = reader.Seek(originalOffset, 0); err != nil {
			return
		}
	}
	mr := io.MultiReader(
		strings.NewReader(multipartFormDataFront),
		strings.NewReader(filename),
		strings.NewReader(multipartFormDataMiddle),
		reader,
		strings.NewReader(multipartFormDataEnd),
	)

	httpReq, err := http.NewRequest("POST", finalURL, mr)
	if err != nil {
		return
	}
	httpReq.Header.Set("Content-Type", multipartContentType)
	httpReq.ContentLength = ContentLength

	httpResp, err := clt.HttpClient.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return
	}
	fmt.Println("mp.WechatClient.UploadFromReader.incompleteURL:", incompleteURL)
	fmt.Println("mp.WechatClient.UploadFromReader.response:", string(body))

	if err = json.Unmarshal(body, response); err != nil {
		return
	}

	// 请注意:
	// 下面获取 ErrCode 的代码不具备通用性!!!
	//
	// 因为本 SDK 的 response 都是
	//  struct {
	//    Error
	//    XXX
	//  }
	// 的结构, 所以用下面简单的方法得到 ErrCode.
	//
	// 如果你是直接调用这个函数, 那么要根据你的 response 数据结构修改下面的代码.
	ErrCode := reflect.ValueOf(response).Elem().FieldByName("ErrCode").Int()

	switch ErrCode {
	case ErrCodeOK:
		return
	case ErrCodeInvalidCredential, ErrCodeTimeout:
		if !hasRetried {
			hasRetried = true

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

func (clt *WechatClient) uploadFromStringsReader(incompleteURL, filename string,
	reader *strings.Reader, response interface{}) (err error) {

	originalOffset, err := reader.Seek(0, 1)
	if err != nil {
		return
	}
	ContentLength := int64(multipartConstPartLen + len(filename) + reader.Len())

	token, err := clt.Token()
	if err != nil {
		return
	}

	hasRetried := false
RETRY:
	finalURL := incompleteURL + url.QueryEscape(token)

	if hasRetried {
		if _, err = reader.Seek(originalOffset, 0); err != nil {
			return
		}
	}
	mr := io.MultiReader(
		strings.NewReader(multipartFormDataFront),
		strings.NewReader(filename),
		strings.NewReader(multipartFormDataMiddle),
		reader,
		strings.NewReader(multipartFormDataEnd),
	)

	httpReq, err := http.NewRequest("POST", finalURL, mr)
	if err != nil {
		return
	}
	httpReq.Header.Set("Content-Type", multipartContentType)
	httpReq.ContentLength = ContentLength

	httpResp, err := clt.HttpClient.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return
	}
	fmt.Println("mp.WechatClient.UploadFromReader.incompleteURL:", incompleteURL)
	fmt.Println("mp.WechatClient.UploadFromReader.response:", string(body))

	if err = json.Unmarshal(body, response); err != nil {
		return
	}

	// 请注意:
	// 下面获取 ErrCode 的代码不具备通用性!!!
	//
	// 因为本 SDK 的 response 都是
	//  struct {
	//    Error
	//    XXX
	//  }
	// 的结构, 所以用下面简单的方法得到 ErrCode.
	//
	// 如果你是直接调用这个函数, 那么要根据你的 response 数据结构修改下面的代码.
	ErrCode := reflect.ValueOf(response).Elem().FieldByName("ErrCode").Int()

	switch ErrCode {
	case ErrCodeOK:
		return
	case ErrCodeInvalidCredential, ErrCodeTimeout:
		if !hasRetried {
			hasRetried = true

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

func (clt *WechatClient) uploadFromIOReader(incompleteURL, filename string,
	reader io.Reader, response interface{}) (err error) {

	bodyBuf := mediaBufferPool.Get().(*bytes.Buffer)
	bodyBuf.Reset()
	defer mediaBufferPool.Put(bodyBuf)

	bodyBuf.WriteString(multipartFormDataFront)
	bodyBuf.WriteString(filename)
	bodyBuf.WriteString(multipartFormDataMiddle)
	if _, err = io.Copy(bodyBuf, reader); err != nil {
		return
	}
	bodyBuf.WriteString(multipartFormDataEnd)

	bodyBytes := bodyBuf.Bytes()

	token, err := clt.Token()
	if err != nil {
		return
	}

	hasRetried := false
RETRY:
	finalURL := incompleteURL + url.QueryEscape(token)

	httpResp, err := clt.HttpClient.Post(finalURL, multipartContentType, bytes.NewReader(bodyBytes))
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return
	}
	fmt.Println("mp.WechatClient.UploadFromReader.incompleteURL:", incompleteURL)
	fmt.Println("mp.WechatClient.UploadFromReader.response:", string(body))

	if err = json.Unmarshal(body, response); err != nil {
		return
	}

	// 请注意:
	// 下面获取 ErrCode 的代码不具备通用性!!!
	//
	// 因为本 SDK 的 response 都是
	//  struct {
	//    Error
	//    XXX
	//  }
	// 的结构, 所以用下面简单的方法得到 ErrCode.
	//
	// 如果你是直接调用这个函数, 那么要根据你的 response 数据结构修改下面的代码.
	ErrCode := reflect.ValueOf(response).Elem().FieldByName("ErrCode").Int()

	switch ErrCode {
	case ErrCodeOK:
		return
	case ErrCodeInvalidCredential, ErrCodeTimeout:
		if !hasRetried {
			hasRetried = true

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
