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
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"

	wechatjson "github.com/chanxuehong/wechat/json"
)

// 企业号"主动"请求功能的基本封装.
type CorpClient struct {
	TokenServer
	HttpClient *http.Client
}

// 创建一个新的 CorpClient.
//  如果 HttpClient == nil 则默认用 http.DefaultClient
func NewCorpClient(TokenServer TokenServer, HttpClient *http.Client) *CorpClient {
	if TokenServer == nil {
		panic("TokenServer == nil")
	}
	if HttpClient == nil {
		HttpClient = http.DefaultClient
	}

	return &CorpClient{
		TokenServer: TokenServer,
		HttpClient:  HttpClient,
	}
}

// 用 encoding/json 把 request marshal 为 JSON, 放入 http 请求的 body 中,
// POST 到微信服务器, 然后将微信服务器返回的 JSON 用 encoding/json 解析到 response.
//
//  NOTE:
//  1. 一般不用调用这个方法, 请直接调用高层次的封装方法;
//  2. 最终的 URL == incompleteURL + access_token;
//  3. response 要求是 struct 的指针, 并且该 struct 拥有属性:
//     ErrCode int `json:"errcode"` (可以是直接属性, 也可以是匿名属性里的属性)
func (clt *CorpClient) PostJSON(incompleteURL string, request interface{}, response interface{}) (err error) {
	buf := textBufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer textBufferPool.Put(buf)

	if err = wechatjson.NewEncoder(buf).Encode(request); err != nil {
		return
	}
	requestBytes := buf.Bytes()

	token, err := clt.Token()
	if err != nil {
		return
	}

	hasRetried := false
RETRY:
	finalURL := incompleteURL + url.QueryEscape(token)

	log.Println("request url:", finalURL)
	log.Println("request json:", string(requestBytes))

	httpResp, err := clt.HttpClient.Post(finalURL, "application/json; charset=utf-8", bytes.NewReader(requestBytes))
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}

	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return
	}
	log.Println("response json:", string(respBody))

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
		log.Println("RETRY, err_code:", ErrCode, ", err_msg:", ErrMsg)
		log.Println("RETRY, current token:", token)

		if !hasRetried {
			hasRetried = true

			if token, err = clt.TokenRefresh(); err != nil {
				return
			}
			log.Println("RETRY, new token:", token)

			responseStructValue.Set(reflect.New(responseStructValue.Type()).Elem())
			goto RETRY
		}
		log.Println("RETRY fallthrough, current token:", token)
		fallthrough
	default:
		return
	}
}

// GET 微信资源, 然后将微信服务器返回的 JSON 用 encoding/json 解析到 response.
//
//  NOTE:
//  1. 一般不用调用这个方法, 请直接调用高层次的封装方法;
//  2. 最终的 URL == incompleteURL + access_token;
//  3. response 要求是 struct 的指针, 并且该 struct 拥有属性:
//     ErrCode int `json:"errcode"` (可以是直接属性, 也可以是匿名属性里的属性)
func (clt *CorpClient) GetJSON(incompleteURL string, response interface{}) (err error) {
	token, err := clt.Token()
	if err != nil {
		return
	}

	hasRetried := false
RETRY:
	finalURL := incompleteURL + url.QueryEscape(token)

	httpResp, err := clt.HttpClient.Get(finalURL)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}

	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return
	}
	log.Println("request url:", finalURL)
	log.Println("response json:", string(respBody))

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
		log.Println("RETRY, err_code:", ErrCode, ", err_msg:", ErrMsg)
		log.Println("RETRY, current token:", token)

		if !hasRetried {
			hasRetried = true

			if token, err = clt.TokenRefresh(); err != nil {
				return
			}
			log.Println("RETRY, new token:", token)

			responseStructValue.Set(reflect.New(responseStructValue.Type()).Elem())
			goto RETRY
		}
		log.Println("RETRY fallthrough, current token:", token)
		fallthrough
	default:
		return
	}
}
