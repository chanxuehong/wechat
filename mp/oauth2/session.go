package oauth2

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gopkg.in/chanxuehong/wechat.v2/internal/debug/api"
	util2 "gopkg.in/chanxuehong/wechat.v2/internal/util"
	"gopkg.in/chanxuehong/wechat.v2/oauth2"
	"gopkg.in/chanxuehong/wechat.v2/util"
	"net/http"
)

type Session struct {
	OpenId     string `json:"openid"`
	UnionId    string `json:"unionid,omitempty"`
	SessionKey string `json:"session_key"` // 用户授权的作用域, 使用逗号(,)分隔
}

type SessionInfo struct {
	OpenId   string `json:"openId"`   // 用户的唯一标识
	Nickname string `json:"nickName"` // 用户昵称
	Gender   int    `json:"gender"`   // 用户的性别, 值为1时是男性, 值为2时是女性, 值为0时是未知
	Language string `json:"language"` // 用户的性别, 值为1时是男性, 值为2时是女性, 值为0时是未知
	City     string `json:"city"`     // 普通用户个人资料填写的城市
	Province string `json:"province"` // 用户个人资料填写的省份
	Country  string `json:"country"`  // 国家, 如中国为CN

	// 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），
	// 用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	AvatarUrl string `json:"avatarUrl"`
	UnionId   string `json:"unionId"` // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
}

func GetSession(Endpoint *Endpoint, code string) (session *Session, err error) {
	session = &Session{}
	if err = getSession(session, Endpoint.SessionCodeUrl(code)); err != nil {
		return
	}
	return
}

func getSession(session *Session, url string) (err error) {

	httpClient := util.DefaultHttpClient

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
		Session
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

func GetSessionInfo(EncryptedData, sessionKey, iv string) (info *SessionInfo, err error) {

	cipherText, err := base64.StdEncoding.DecodeString(EncryptedData)

	aesKey, err := base64.StdEncoding.DecodeString(sessionKey)
	aesIv, err := base64.StdEncoding.DecodeString(iv)

	if err != nil {
		return
	}

	raw, err := util2.AESDecryptData(cipherText, aesKey, aesIv)

	if err != nil {
		return
	}

	if err = json.Unmarshal(raw, &info); err != nil {
		return
	}
	return
}
