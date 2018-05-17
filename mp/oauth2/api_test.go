package oauth2

import (
	"fmt"
	"testing"
)

const (
	wxAppId     = "" // 填上自己的参数
	wxAppSecret = "" // 填上自己的参数
	//oauth2RedirectURI = "http://192.168.1.129:8080/page2" // 填上自己的参数
	//oauth2Scope       = "snsapi_userinfo"                 // 填上自己的参数
)

var (
	oauth2Endpoint *Endpoint = NewEndpoint(wxAppId, wxAppSecret)
)

func TestGetSession(t *testing.T) {

	GetSession(oauth2Endpoint, "013Rc7FP0lmgxb2lRIIP0VefFP0Rc7FW")
}

func TestGetUserInfoBySession(t *testing.T) {

	sessionKey := ""

	iv := ""

	encrypt := ""

	info, err := GetSessionInfo(encrypt, sessionKey, iv)

	fmt.Println(info)

	fmt.Println(err)
}
