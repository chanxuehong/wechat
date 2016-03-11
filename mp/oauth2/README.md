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
	wxAppId           = "appid"        // 填上自己的参数
	wxAppSecret       = "app_secret"   // 填上自己的参数
	oauth2RedirectURI = "redirect_uri" // 填上自己的参数
	oauth2Scope       = "scope"        // 填上自己的参数
)

var (
	sessionStorage                 = session.New(20*60, 60*60)
	oauth2Endpoint oauth2.Endpoint = mpoauth2.NewEndpoint(wxAppId, wxAppSecret)
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

	AuthCodeURL := mpoauth2.AuthCodeURL(wxAppId, oauth2RedirectURI, oauth2Scope, state)
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
		Endpoint: oauth2Endpoint,
	}
	token, err := oauth2Client.ExchangeToken(code)
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
2016/03/11 13:49:56 AuthCodeURL: https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx7aa843bc1a414ca4&redirect_uri=http%3A%2F%2F192.168.1.129%3A8080%2Fpage2&response_type=code&scope=snsapi_userinfo&state=6172a4d8fc646efebb0589a8425e24ad#wechat_redirect
2016/03/11 13:50:05 AuthCodeURL: https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx7aa843bc1a414ca4&redirect_uri=http%3A%2F%2F192.168.1.129%3A8080%2Fpage2&response_type=code&scope=snsapi_userinfo&state=25a2520bb45194e219baa75c518fdd4c#wechat_redirect
2016/03/11 13:50:08 /page2?code=011e84a8d7d5cbf8f92cb0a7d3ac74eH&state=25a2520bb45194e219baa75c518fdd4c
2016/03/11 13:50:08 [WECHAT_DEBUG] [API] GET https://api.weixin.qq.com/sns/oauth2/access_token?appid=wx7aa843bc1a414ca4&secret=03a8ffc8fec1fa6a3d810d7f2c60c133&code=011e84a8d7d5cbf8f92cb0a7d3ac74eH&grant_type=authorization_code
2016/03/11 13:50:08 [WECHAT_DEBUG] [API] http response body:
{"access_token":"OezXcEiiBSKSxW0eoylIeBXr5MlyOHqjC6Db82eo0Txphdi2X9hT8lI2PqjpL0rC_ZstDSrFZQfXzElMr7hXecvs5dy1kf2JdbvTMimOkt16K2yr2F7VtvCP1zcpDUXglDF6M6Wg6SRlZgo2mfv7pw","expires_in":7200,"refresh_token":"OezXcEiiBSKSxW0eoylIeBXr5MlyOHqjC6Db82eo0Txphdi2X9hT8lI2PqjpL0rCgJ9YKeHtmxFMbXmnxm_Vo3PrjEWso7DOtvBxtvQYq6K7l9Mrk59JgqGGYpwCSUYbkn76JJCGvL0JjH5pywINug","openid":"os-IKuHd9pJ6xsn4mS7GyL4HxqI4","scope":"snsapi_userinfo"}
2016/03/11 13:50:08 token: &{AccessToken:OezXcEiiBSKSxW0eoylIeBXr5MlyOHqjC6Db82eo0Txphdi2X9hT8lI2PqjpL0rC_ZstDSrFZQfXzElMr7hXecvs5dy1kf2JdbvTMimOkt16K2yr2F7VtvCP1zcpDUXglDF6M6Wg6SRlZgo2mfv7pw CreatedAt:1457675408 ExpiresIn:6000 RefreshToken:OezXcEiiBSKSxW0eoylIeBXr5MlyOHqjC6Db82eo0Txphdi2X9hT8lI2PqjpL0rCgJ9YKeHtmxFMbXmnxm_Vo3PrjEWso7DOtvBxtvQYq6K7l9Mrk59JgqGGYpwCSUYbkn76JJCGvL0JjH5pywINug OpenId:os-IKuHd9pJ6xsn4mS7GyL4HxqI4 UnionId: Scope:snsapi_userinfo}
2016/03/11 13:50:08 [WECHAT_DEBUG] [API] GET https://api.weixin.qq.com/sns/userinfo?access_token=OezXcEiiBSKSxW0eoylIeBXr5MlyOHqjC6Db82eo0Txphdi2X9hT8lI2PqjpL0rC_ZstDSrFZQfXzElMr7hXecvs5dy1kf2JdbvTMimOkt16K2yr2F7VtvCP1zcpDUXglDF6M6Wg6SRlZgo2mfv7pw&openid=os-IKuHd9pJ6xsn4mS7GyL4HxqI4&lang=zh_CN
2016/03/11 13:50:08 [WECHAT_DEBUG] [API] http response body:
{"openid":"os-IKuHd9pJ6xsn4mS7GyL4HxqI4","nickname":"产学红","sex":1,"language":"zh_CN","city":"安庆","province":"安徽","country":"中国","headimgurl":"http:\/\/wx.qlogo.cn\/mmopen\/O1HUibMqqHXduhNiagwbE0m4zgJU2YbFkyZPG6VoH8IP2wEdFuWcnjUtrXHNl1OmCsoffYBBnkC0cy1yfsOibcenaAn2SeRNKYw\/0","privilege":[]}
2016/03/11 13:50:08 userinfo: &{OpenId:os-IKuHd9pJ6xsn4mS7GyL4HxqI4 Nickname:产学红 Sex:1 City:安庆 Province:安徽 Country:中国 HeadImageURL:http://wx.qlogo.cn/mmopen/O1HUibMqqHXduhNiagwbE0m4zgJU2YbFkyZPG6VoH8IP2wEdFuWcnjUtrXHNl1OmCsoffYBBnkC0cy1yfsOibcenaAn2SeRNKYw/0 Privilege:[] UnionId:}
```
