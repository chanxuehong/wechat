下面是一个简单的实现

```golang
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/chanxuehong/session"
	"github.com/chanxuehong/util/random"
	"github.com/chanxuehong/wechat/mp/oauth2"
)

var (
	sessionStorage = session.New(20*60, 5*60)

	oauth2Config = oauth2.NewOAuth2Config(
		"appid",     // 填上自己的参数
		"appsecret", // 填上自己的参数
		"http://192.168.1.253/page2",
		"snsapi_base",
	)
)

func Page1Handler(w http.ResponseWriter, r *http.Request) {
	sid := string(random.NewSessionId())
	state := string(random.NewToken())

	if err := sessionStorage.Add(sid, state); err != nil {
		io.WriteString(w, err.Error())
		return
	}

	cookie := http.Cookie{
		Name:     "sid",
		Value:    sid,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, oauth2Config.AuthCodeURL(state), http.StatusFound)
}

func Page2Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL == nil {
		io.WriteString(w, "r.URL == nil")
		return
	}

	urlValues, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	cookie, err := r.Cookie("sid")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	sid := cookie.Value
	session, err := sessionStorage.Get(sid)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	savedState := session.(string)

	code := urlValues.Get("code")
	urlState := urlValues.Get("state")

	if savedState != urlState {
		io.WriteString(w, fmt.Sprintf("state 不匹配, session 中的为 %q, url 传递过来的是 %q", savedState, urlState))
		return
	}

	if code == "" {
		io.WriteString(w, "客户禁止授权")
		return
	}

	oauth2Client := oauth2.Client{
		OAuth2Config: oauth2Config,
	}

	token, err := oauth2Client.Exchange(code)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	io.WriteString(w, "客户的 openid 为: "+token.OpenId)
}

func init() {
	http.HandleFunc("/page1", Page1Handler)
	http.HandleFunc("/page2", Page2Handler)
}

func main() {
	http.ListenAndServe(":80", nil)
}
```
