package oauth2

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/bububa/wechat/internal/debug/api"
	"github.com/bububa/wechat/oauth2"
	"github.com/bububa/wechat/util"
)

type UserInfo struct {
	UserId string `json:"UserId"` // 成员UserID。若需要获得用户详情信息，可调用通讯录接口：读取成员
	OpenId string `json:"OpenId"` // 非企业成员的标识，对当前企业唯一
}

// GetUserInfo 获取用户信息.
//
//	accessToken: 网页授权接口调用凭证
//	code:     通过成员授权获取到的code，最大为512字节。每次成员授权带上的code将不一样，code只能使用一次，5分钟未被使用自动过期。
//	httpClient:  如果不指定则默认为 util.DefaultHttpClient
func GetUserInfo(accessToken, code string, httpClient *http.Client) (info *UserInfo, err error) {
	if httpClient == nil {
		httpClient = util.DefaultHttpClient
	}

	_url := "https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=" + url.QueryEscape(accessToken) +
		"&code=" + url.QueryEscape(code)
	api.DebugPrintGetRequest(_url, false)
	httpResp, err := httpClient.Get(_url)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	var result struct {
		oauth2.Error
		UserInfo
	}
	if err = api.DecodeJSONHttpResponse(httpResp.Body, &result, false); err != nil {
		return
	}
	if result.ErrCode != oauth2.ErrCodeOK {
		err = &result.Error
		return
	}
	info = &result.UserInfo
	return
}
