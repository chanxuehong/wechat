package component

import (
	"fmt"
	"github.com/chanxuehong/wechat/internal/debug/api"
	mpoauth2 "github.com/chanxuehong/wechat/mp/oauth2"
	"github.com/chanxuehong/wechat/oauth2"
	"github.com/chanxuehong/wechat/util"
	"net/http"
)

// GetSession 获取小程序会话
func GetSession(Endpoint *Endpoint, code string) (session *mpoauth2.Session, err error) {
	session = &mpoauth2.Session{}
	if err = getSession(session, Endpoint.SessionCodeUrl(code), nil); err != nil {
		return
	}
	return
}

// GetSessionWithClient 获取小程序会话
func GetSessionWithClient(Endpoint *Endpoint, code string, httpClient *http.Client) (session *mpoauth2.Session, err error) {
	session = &mpoauth2.Session{}
	if err = getSession(session, Endpoint.SessionCodeUrl(code), httpClient); err != nil {
		return
	}
	return
}

func getSession(session *mpoauth2.Session, url string, httpClient *http.Client) (err error) {

	if httpClient == nil {
		httpClient = util.DefaultHttpClient
	}

	api.DebugPrintGetRequest(url)

	httpResp, err := httpClient.Get(url)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}

	var result struct {
		oauth2.Error
		mpoauth2.Session
	}

	if err = api.DecodeJSONHttpResponse(httpResp.Body, &result); err != nil {
		return
	}

	if result.ErrCode != oauth2.ErrCodeOK {
		return &result.Error
	}

	*session = result.Session

	return
}
