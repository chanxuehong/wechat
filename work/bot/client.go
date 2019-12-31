package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/chanxuehong/wechat/internal/debug/api"
	"github.com/chanxuehong/wechat/internal/debug/api/retry"
	"github.com/chanxuehong/wechat/util"
)

func PostJSON(webhook string, request interface{}, response interface{}) (err error) {
	ErrorStructValue, ErrorErrCodeValue := checkResponse(response)

	buffer := textBufferPool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer textBufferPool.Put(buffer)

	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	if err = encoder.Encode(request); err != nil {
		return
	}
	requestBodyBytes := buffer.Bytes()
	if i := len(requestBodyBytes) - 1; i >= 0 && requestBodyBytes[i] == '\n' {
		requestBodyBytes = requestBodyBytes[:i] // 去掉最后的 '\n', 这样能统一log格式, 不然可能多一个空白行
	}

	hasRetried := false
RETRY:
	if err = httpPostJSON(webhook, requestBodyBytes, response); err != nil {
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

func httpPostJSON(url string, body []byte, response interface{}) error {
	api.DebugPrintPostJSONRequest(url, body)
	httpResp, err := util.DefaultHttpClient.Post(url, "application/json; charset=utf-8", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}
	return api.DecodeJSONHttpResponse(httpResp.Body, response)
}

// checkResponse 检查 response 参数是否满足特定的结构要求, 如果不满足要求则会 panic, 否则返回相应的 reflect.Value.
func checkResponse(response interface{}) (ErrorStructValue, ErrorErrCodeValue reflect.Value) {
	responseValue := reflect.ValueOf(response)
	if responseValue.Kind() != reflect.Ptr {
		panic("the type of response is incorrect")
	}
	responseStructValue := responseValue.Elem()
	if responseStructValue.Kind() != reflect.Struct {
		panic("the type of response is incorrect")
	}

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
	ErrorErrCodeValue = ErrorStructValue.Field(errorErrCodeIndex)
	return
}
