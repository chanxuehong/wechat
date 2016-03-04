## 微信网页授权

```Go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/chanxuehong/rand"
	"github.com/chanxuehong/session"
	"github.com/chanxuehong/sid"
	mpoauth2 "github.com/chanxuehong/wechat/mp/oauth2"
	"github.com/chanxuehong/wechat/oauth2"
)

const (
	AppId       = "appid"        // 填上自己的参数
	AppSecret   = "app_secret"   // 填上自己的参数
	RedirectURI = "redirect_uri" // 填上自己的参数
	Scope       = "scope"        // 填上自己的参数
)

var (
	sessionStorage               = session.New(20*60, 60*60)
	oauth2Config   oauth2.Config = mpoauth2.NewConfig(AppId, AppSecret)
)

func init() {
	http.HandleFunc("/page1", Page1Handler)
	http.HandleFunc("/page2", Page2Handler)
}

// 建立必要的 session, 然后跳转到授权页面
func Page1Handler(w http.ResponseWriter, r *http.Request) {
	sid := sid.New()
	state := string(rand.NewHex())

	if err := sessionStorage.Add(sid, state); err != nil {
		io.WriteString(w, err.Error())
		log.Println(err)
		return
	}

	cookie := http.Cookie{
		Name:     "sid",
		Value:    sid,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	AuthCodeURL := mpoauth2.AuthCodeURL(AppId, RedirectURI, Scope, state)
	log.Println("AuthCodeURL:", AuthCodeURL)

	http.Redirect(w, r, AuthCodeURL, http.StatusFound)
}

// 授权后回调页面
func Page2Handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RequestURI)

	cookie, err := r.Cookie("sid")
	if err != nil {
		io.WriteString(w, err.Error())
		log.Println(err)
		return
	}

	session, err := sessionStorage.Get(cookie.Value)
	if err != nil {
		io.WriteString(w, err.Error())
		log.Println(err)
		return
	}

	savedState := session.(string) // 一般是要序列化的, 这里保存在内存所以可以这么做

	queryValues, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		io.WriteString(w, err.Error())
		log.Println(err)
		return
	}

	code := queryValues.Get("code")
	if code == "" {
		log.Println("用户禁止授权")
		return
	}

	queryState := queryValues.Get("state")
	if queryState == "" {
		log.Println("state 参数为空")
		return
	}
	if savedState != queryState {
		str := fmt.Sprintf("state 不匹配, session 中的为 %q, url 传递过来的是 %q", savedState, queryState)
		io.WriteString(w, str)
		log.Println(str)
		return
	}

	oauth2Client := oauth2.Client{
		Config: oauth2Config,
	}
	token, err := oauth2Client.Exchange(code)
	if err != nil {
		io.WriteString(w, err.Error())
		log.Println(err)
		return
	}
	log.Printf("token: %+v\r\n", token)

	userinfo, err := mpoauth2.GetUserInfo(token.AccessToken, token.OpenId, "", nil)
	if err != nil {
		io.WriteString(w, err.Error())
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(userinfo)
	log.Printf("userinfo: %+v\r\n", userinfo)
	return
}

func main() {
	fmt.Println(http.ListenAndServe(":8080", nil))
}
```

#### 上面的程序基本上会打印下面的内容
```
2016/03/04 23:38:25 AuthCodeURL: https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx7aa843bc1a414ca4&redirect_uri=http%3A%2F%2F192.168.1.129%3A8080%2Fpage2&response_type=code&scope=snsapi_userinfo&state=948f34d6cbc66ba2d9a619726bca76df#wechat_redirect
2016/03/04 23:38:40 /page2?code=02194667b305a5891d57571ab0a7d5bT&state=948f34d6cbc66ba2d9a619726bca76df
2016/03/04 23:38:40 token: &{AccessToken:OezXcEiiBSKSxW0eoylIeBXr5MlyOHqjC6Db82eo0Txphdi2X9hT8lI2PqjpL0rCovT8Upm2zk7eqhS9HH6lva1VkLDVtC5wXq2rSKmCBrggAaiBtcpiQYkHSIx6Sz602U90uJB2pHqeTDgAG7TLFA CreatedAt:1457105920 ExpiresIn:6000 RefreshToken:OezXcEiiBSKSxW0eoylIeBXr5MlyOHqjC6Db82eo0Txphdi2X9hT8lI2PqjpL0rCVp03vRNwbZ0oT2qX-NPEj50whGAz0-W4OzVtwS1sLr189SjMCrCpMXtKPr9K1hT5HfGgiCV77TUiUENH5u8R1w OpenId:os-IKuHd9pJ6xsn4mS7GyL4HxqI4 UnionId: Scope:snsapi_userinfo}
2016/03/04 23:38:40 userinfo: &{OpenId:os-IKuHd9pJ6xsn4mS7GyL4HxqI4 Nickname:产学红 Sex:1 City:安庆 Province:安徽 Country:中国 HeadImageURL:http://wx.qlogo.cn/mmopen/O1HUibMqqHXduhNiagwbE0m4zgJU2YbFkyZPG6VoH8IP2wEdFuWcnjUtrXHNl1OmCsoffYBBnkC0cy1yfsOibcenaAn2SeRNKYw/0 Privilege:[] UnionId:}
```
