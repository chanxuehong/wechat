package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/bububa/wechat/internal/debug/api"
	"github.com/bububa/wechat/util"
)

type SimpleAccessTokenServer struct {
	httpClient *http.Client
	token      string
	corpId     string
	corpSecret string
	debug      bool
}

func NewSimpleAccessTokenServer(token string, httpClient *http.Client) (srv *SimpleAccessTokenServer) {
	if httpClient == nil {
		httpClient = util.DefaultHttpClient
	}

	srv = &SimpleAccessTokenServer{
		token:      token,
		httpClient: httpClient,
	}
	return
}

func (srv *SimpleAccessTokenServer) IID01332E16DF5011E5A9D5A4DB30FED8E1() {}

func (srv *SimpleAccessTokenServer) SetDebug(debug bool) {
	srv.debug = debug
}

func (srv *SimpleAccessTokenServer) Debug() bool {
	return srv.debug
}

func (srv *SimpleAccessTokenServer) SetSecret(corpId string, corpSecret string) {
	srv.corpId = url.QueryEscape(corpId)
	srv.corpSecret = url.QueryEscape(corpSecret)
}

func (srv *SimpleAccessTokenServer) Token() (token string, err error) {
	return srv.token, nil
}

func (srv *SimpleAccessTokenServer) RefreshToken(currentToken string) (token string, err error) {
	if currentToken == "" && srv.corpId != "" && srv.corpSecret != "" {
		currentToken, _, err = srv.updateToken()
	}
	srv.token = currentToken
	return srv.token, nil
}

func (srv *SimpleAccessTokenServer) RefreshTokenWithExpires(currentToken string) (token string, expiresIn int64, err error) {
	if currentToken == "" && srv.corpId != "" && srv.corpSecret != "" {
		currentToken, expiresIn, err = srv.updateToken()
	}
	srv.token = currentToken
	return srv.token, expiresIn, err
}

func (srv *SimpleAccessTokenServer) updateToken() (token string, expiresIn int64, err error) {
	url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + srv.corpId +
		"&corpsecret=" + srv.corpSecret
	api.DebugPrintGetRequest(url, srv.debug)
	httpResp, err := srv.httpClient.Get(url)
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
		Token     string `json:"access_token"`
		ExpiresIn int64  `json:"expires_in"`
	}

	if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
		return
	}
	if result.ErrCode != ErrCodeOK {
		err = &result.Error
		return
	}
	switch {
	case result.ExpiresIn > 31556952: // 60*60*24*365.2425
		err = errors.New("expires_in too large: " + strconv.FormatInt(result.ExpiresIn, 10))
		return
	case result.ExpiresIn > 60*60:
		result.ExpiresIn -= 60 * 10
	case result.ExpiresIn > 60*30:
		result.ExpiresIn -= 60 * 5
	case result.ExpiresIn > 60*5:
		result.ExpiresIn -= 60
	case result.ExpiresIn > 60:
		result.ExpiresIn -= 10
	default:
		err = errors.New("expires_in too small: " + strconv.FormatInt(result.ExpiresIn, 10))
		return
	}
	return result.Token, result.ExpiresIn, nil
}
