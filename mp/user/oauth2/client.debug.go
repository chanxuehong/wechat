// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

// +build wechatdebug

package oauth2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/chanxuehong/wechat/mp"
)

type Client struct {
	Config Config
	Token  *Token // 程序会自动更新最新的 Token 到这个字段, 如有必要该字段可以保存起来

	HttpClient *http.Client // 如果 HttpClient == nil 则默认用 http.DefaultClient
}

func (clt *Client) httpClient() *http.Client {
	if clt.HttpClient != nil {
		return clt.HttpClient
	}
	return http.DefaultClient
}

func (clt *Client) getJSON(url string, response interface{}) (err error) {
	mp.LogInfoln("[WECHAT_DEBUG] request url:", url)

	httpResp, err := clt.httpClient().Get(url)
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
	mp.LogInfoln("[WECHAT_DEBUG] response json:", string(respBody))

	return json.Unmarshal(respBody, response)
}
