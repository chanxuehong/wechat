package core

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/chanxuehong/wechat/internal"
	wechatjson "github.com/chanxuehong/wechat/json"
)

type Client struct {
	AccessTokenServer
	HttpClient *http.Client
}

// NewClient 创建一个新的 Client.
//  如果 clt == nil 则默认用 http.DefaultClient
func NewClient(srv AccessTokenServer, clt *http.Client) *Client {
	if srv == nil {
		panic("nil AccessTokenServer")
	}
	if clt == nil {
		clt = http.DefaultClient
	}
	return &Client{
		AccessTokenServer: srv,
		HttpClient:        clt,
	}
}

// GetJSON HTTP GET 微信资源, 然后将微信服务器返回的 JSON 用 encoding/json 解析到 response.
//
//  NOTE:
//  1. 一般不需要调用这个方法, 请直接调用高层次的封装函数;
//  2. 最终的 URL == incompleteURL + access_token;
//  3. response 格式有要求, 要么是 *Error, 要么是下面结构体的指针(注意 Error 必须是第一个 Field):
//      struct {
//          Error
//          ...
//      }
func (clt *Client) GetJSON(incompleteURL string, response interface{}) (err error) {
	ErrorStructValue, ErrorErrCodeValue := checkResponse(response)

	httpClient := clt.HttpClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	token, err := clt.Token()
	if err != nil {
		return
	}

	hasRetried := false
RETRY:
	finalURL := incompleteURL + url.QueryEscape(token)

	err = func() error {
		internal.DebugPrintGetRequest(finalURL)
		httpResp, err := httpClient.Get(finalURL)
		if err != nil {
			return err
		}
		defer httpResp.Body.Close()

		if httpResp.StatusCode != http.StatusOK {
			return fmt.Errorf("http.Status: %s", httpResp.Status)
		}
		return internal.JsonHttpResponseUnmarshal(httpResp.Body, response)
	}()
	if err != nil {
		return
	}

	switch errCode := ErrorErrCodeValue.Int(); errCode {
	case ErrCodeOK:
		return
	case ErrCodeInvalidCredential, ErrCodeAccessTokenExpired:
		errMsg := ErrorStructValue.Field(errorErrMsgIndex).String()
		internal.DebugPrintRetryError(errCode, errMsg, token)
		if !hasRetried {
			hasRetried = true
			ErrorStructValue.Set(errorZeroValue)
			if token, err = clt.TokenRefresh(); err != nil {
				return
			}
			internal.DebugPrintRetryNewToken(token)
			goto RETRY
		}
		internal.DebugPrintRetryFallthrough(token)
		fallthrough
	default:
		return
	}
}

// PostJSON 用 encoding/json 把 request marshal 为 JSON, HTTP POST 到微信服务器,
// 然后将微信服务器返回的 JSON 用 encoding/json 解析到 response.
//
//  NOTE:
//  1. 一般不需要调用这个方法, 请直接调用高层次的封装函数;
//  2. 最终的 URL == incompleteURL + access_token;
//  3. response 格式有要求, 要么是 *Error, 要么是下面结构体的指针(注意 Error 必须是第一个 Field):
//      struct {
//          Error
//          ...
//      }
func (clt *Client) PostJSON(incompleteURL string, request interface{}, response interface{}) (err error) {
	ErrorStructValue, ErrorErrCodeValue := checkResponse(response)

	bodyBuf := textBufferPool.Get().(*bytes.Buffer)
	bodyBuf.Reset()
	defer textBufferPool.Put(bodyBuf)

	if err = wechatjson.NewEncoder(bodyBuf).Encode(request); err != nil {
		return
	}
	requestBodyBytes := bodyBuf.Bytes()
	const requestBodyType = "application/json; charset=utf-8"

	httpClient := clt.HttpClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	token, err := clt.Token()
	if err != nil {
		return
	}

	hasRetried := false
RETRY:
	finalURL := incompleteURL + url.QueryEscape(token)

	err = func() error {
		internal.DebugPrintPostJSONRequest(finalURL, requestBodyBytes)
		httpResp, err := httpClient.Post(finalURL, requestBodyType, bytes.NewReader(requestBodyBytes))
		if err != nil {
			return err
		}
		defer httpResp.Body.Close()

		if httpResp.StatusCode != http.StatusOK {
			return fmt.Errorf("http.Status: %s", httpResp.Status)
		}
		return internal.JsonHttpResponseUnmarshal(httpResp.Body, response)
	}()
	if err != nil {
		return
	}

	switch errCode := ErrorErrCodeValue.Int(); errCode {
	case ErrCodeOK:
		return
	case ErrCodeInvalidCredential, ErrCodeAccessTokenExpired:
		errMsg := ErrorStructValue.Field(errorErrMsgIndex).String()
		internal.DebugPrintRetryError(errCode, errMsg, token)
		if !hasRetried {
			hasRetried = true
			ErrorStructValue.Set(errorZeroValue)
			if token, err = clt.TokenRefresh(); err != nil {
				return
			}
			internal.DebugPrintRetryNewToken(token)
			goto RETRY
		}
		internal.DebugPrintRetryFallthrough(token)
		fallthrough
	default:
		return
	}
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
