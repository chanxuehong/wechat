// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TokenGetter interface {
	GetToken() (token string, err error)
}

var _ TokenGetter = new(DefaultTokenGetter)

type DefaultTokenGetter struct {
	corpId     string
	corpSecret string
	httpClient *http.Client
}

func NewDefaultTokenGetter(corpId, corpSecret string, httpClient *http.Client) *DefaultTokenGetter {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &DefaultTokenGetter{
		corpId:     corpId,
		corpSecret: corpSecret,
		httpClient: httpClient,
	}
}

// 从微信服务器获取新的 access_token
func (getter *DefaultTokenGetter) GetToken() (token string, err error) {
	url_ := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" +
		getter.corpId + "&corpsecret=" + getter.corpSecret

	httpResp, err := getter.httpClient.Get(url_)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	var result struct {
		Error
		Token string `json:"access_token"`
	}

	if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
		return
	}

	if result.ErrCode != errCodeOK {
		err = &result.Error
		return
	}

	token = result.Token
	return
}
