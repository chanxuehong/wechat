// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

// +build wechatdebug

package corp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
)

// 通用上传接口.
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
//  3. part1 是一个文件, part2 是普通的字符串(如果不需要 part2 则把 part2FieldName 留空);
//  4. response 要求是 struct 的指针, 并且该 struct 拥有属性:
//     ErrCode int `json:"errcode"` (可以是直接属性, 也可以是匿名属性里的属性)
func (clt *CorpClient) UploadFromReader(incompleteURL,
	part1FieldName, part1FileName string, part1ValueReader io.Reader,
	part2FieldName string, part2Value []byte,
	response interface{}) (err error) {

	// 构造 multipart/form-data, 存入一个字节数组里

	bodyBuf := mediaBufferPool.Get().(*bytes.Buffer)
	bodyBuf.Reset()
	defer mediaBufferPool.Put(bodyBuf)

	multipartWriter := multipart.NewWriter(bodyBuf)

	part1Writer, err := multipartWriter.CreateFormFile(part1FieldName, part1FileName)
	if err != nil {
		return
	}
	if _, err = io.Copy(part1Writer, part1ValueReader); err != nil {
		return
	}

	if part2FieldName != "" && len(part2Value) > 0 {
		part2Writer, err := multipartWriter.CreateFormField(part2FieldName)
		if err != nil {
			return err
		}
		if _, err = part2Writer.Write(part2Value); err != nil {
			return err
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
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return
	}
	LogInfoln("[WECHAT_DEBUG] request url:", finalURL)
	LogInfoln("[WECHAT_DEBUG] response json:", string(respBody))

	if err = json.Unmarshal(respBody, response); err != nil {
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
	responseStructValue := reflect.ValueOf(response).Elem()
	ErrCode := responseStructValue.FieldByName("ErrCode").Int()

	switch ErrCode {
	case ErrCodeOK:
		return
	case ErrCodeTimeout, ErrCodeInvalidCredential:
		ErrMsg := responseStructValue.FieldByName("ErrMsg").String()
		LogInfoln("[WECHAT_RETRY] err_code:", ErrCode, ", err_msg:", ErrMsg)
		LogInfoln("[WECHAT_RETRY] current token:", token)

		if !hasRetried {
			hasRetried = true

			if token, err = clt.TokenRefresh(); err != nil {
				return
			}
			LogInfoln("[WECHAT_RETRY] new token:", token)

			responseStructValue.Set(reflect.New(responseStructValue.Type()).Elem())
			goto RETRY
		}
		LogInfoln("[WECHAT_RETRY] fallthrough, current token:", token)
		fallthrough
	default:
		return
	}
}
