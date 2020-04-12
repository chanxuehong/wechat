package wxacode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/util"
	//"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

func PostJSON(clt *core.Client, incompleteURL string, request interface{}) (data []byte, err error) {
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

	httpClient := clt.HttpClient
	if httpClient == nil {
		httpClient = util.DefaultHttpClient
	}

	token, err := clt.Token()
	if err != nil {
		return
	}

	hasRetried := false
RETRY:
	finalURL := incompleteURL + url.QueryEscape(token)
	data, err = httpPostJSON(httpClient, finalURL, requestBodyBytes)
	if err != nil {
		return
	}

	var result core.Error
	json.Unmarshal(data, &result)

	switch result.ErrCode {
	case core.ErrCodeOK:
		return
	case core.ErrCodeInvalidCredential, core.ErrCodeAccessTokenExpired:
		if !hasRetried {
			hasRetried = true
			if token, err = clt.RefreshToken(token); err != nil {
				return
			}
			goto RETRY
		}
		fallthrough
	default:
		return
	}
}

func httpPostJSON(clt *http.Client, url string, body []byte) ([]byte, error) {
	httpResp, err := clt.Post(url, "application/json; charset=utf-8", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http.Status: %s", httpResp.Status)
	}
	return ioutil.ReadAll(httpResp.Body)
}

var textBufferPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 4<<10)) // 4KB
	},
}
