package oauth2

import (
	"fmt"
	"testing"
)

const (
	wxAppId     = "wxa33cba2b69f869f3"               // 填上自己的参数
	wxAppSecret = "ca202c5592c3bebd9f9dfa1ee9e39847" // 填上自己的参数
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

	sessionKey := "HogNecGqZeDxFIDGjBwWKw=="

	iv := "aoZqkfGDWwqj6rFgWdafyw=="

	encrypt := "4B9B1aknFM6yQjAh9mFxH3iN4PZXUGfpBLB98CRAVZzNfUI4J1WUur70+NSH/5MXcCidrt44hi6dkByTtRPxrTIV1BOOTvaa2G5NWunTXZJ37/Oq0ezydbDal5v+X3bvVVeFR6MkhYI+hT9xVl2XnE/Bzon2gIq9F9Fy8Yny0VPqsQ95xUyXnN3/IuhiquR1pAgKjDK3kgCoqhUVNa0dRRQQgTNIpy1djbLfyErPXGTXe1qhAj7RvDdJtRloEfg63JgaB/QTR2BLEGT7/GfHnwROfngxa3esGDeBr9Mtav67R4PjYESVFtH2Yf2npaDWAvAMvr+8hNVd2tjy/3HgImctZWo7bh1OHa/ktH4wKYVbTgxZPg/lgTWC+zsl1Z9g07vWQyGpaA11HODDGn1Kdvh6esiY7T3JQ15nKs1rdyXigNQFb9+kicPiInKVfOiuMcGEazli5/UQHF5rnuWhkg/wy+PRbZybR18lLh5d+EW1CnK8gNcQblDJPvx6QFcFhPxr28WkEd7ys0BZQGCIkQ=="

	info, err := GetSessionInfo(encrypt, sessionKey, iv)

	fmt.Println(info)

	fmt.Println(err)
}
