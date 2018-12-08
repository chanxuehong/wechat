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
	wxAppId           = "APPID"                           // 填上自己的参数
	wxAppSecret       = "APPSECRET"                       // 填上自己的参数
	oauth2RedirectURI = "http://192.168.1.129:8080/page2" // 填上自己的参数
	oauth2Scope       = "snsapi_userinfo"                 // 填上自己的参数
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
2016/03/11 14:44:10 AuthCodeURL: https://open.weixin.qq.com/connect/oauth2/authorize?appid=APPID&redirect_uri=http%3A%2F%2F192.168.1.129%3A8080%2Fpage2&response_type=code&scope=snsapi_userinfo&state=12fa275bac6998ba8f89a5baf13f93a0#wechat_redirect
2016/03/11 14:44:12 /page2?code=001e44d3e2972606638027b31e61a8dH&state=12fa275bac6998ba8f89a5baf13f93a0
2016/03/11 14:44:12 [WECHAT_DEBUG] [API] GET https://api.weixin.qq.com/sns/oauth2/access_token?appid=APPID&secret=APPSECRET&code=001e44d3e2972606638027b31e61a8dH&grant_type=authorization_code
2016/03/11 14:44:13 [WECHAT_DEBUG] [API] http response body:
{"access_token":"OezXcEiiBSKSxW0eoylIeBXr5MlyOHqjC6Db82eo0Txphdi2X9hT8lI2PqjpL0rCFVQjLdSs4v2GQXE8BcLRz1hCQHL6wWXl9013zYMOAvE1UGCV4q-xAIlRVuDa85Mqnqw9himuhlgFUP2Kn0qcFg","expires_in":7200,"refresh_token":"OezXcEiiBSKSxW0eoylIeBXr5MlyOHqjC6Db82eo0Txphdi2X9hT8lI2PqjpL0rCZP99UMCpbrq8v2TWR7uxK5fb0ekmVFl9L1kUOsh1mjQy6rhQG5sGqFBKrkzPr9KQbjTxrvFtscFPmCOMuKi9EQ","openid":"os-IKuHd9pJ6xsn4mS7GyL4HxqI4","scope":"snsapi_userinfo"}
2016/03/11 14:44:13 token: &{AccessToken:OezXcEiiBSKSxW0eoylIeBXr5MlyOHqjC6Db82eo0Txphdi2X9hT8lI2PqjpL0rCFVQjLdSs4v2GQXE8BcLRz1hCQHL6wWXl9013zYMOAvE1UGCV4q-xAIlRVuDa85Mqnqw9himuhlgFUP2Kn0qcFg CreatedAt:1457678653 ExpiresIn:6000 RefreshToken:OezXcEiiBSKSxW0eoylIeBXr5MlyOHqjC6Db82eo0Txphdi2X9hT8lI2PqjpL0rCZP99UMCpbrq8v2TWR7uxK5fb0ekmVFl9L1kUOsh1mjQy6rhQG5sGqFBKrkzPr9KQbjTxrvFtscFPmCOMuKi9EQ OpenId:os-IKuHd9pJ6xsn4mS7GyL4HxqI4 UnionId: Scope:snsapi_userinfo}
2016/03/11 14:44:13 [WECHAT_DEBUG] [API] GET https://api.weixin.qq.com/sns/userinfo?access_token=OezXcEiiBSKSxW0eoylIeBXr5MlyOHqjC6Db82eo0Txphdi2X9hT8lI2PqjpL0rCFVQjLdSs4v2GQXE8BcLRz1hCQHL6wWXl9013zYMOAvE1UGCV4q-xAIlRVuDa85Mqnqw9himuhlgFUP2Kn0qcFg&openid=os-IKuHd9pJ6xsn4mS7GyL4HxqI4&lang=zh_CN
2016/03/11 14:44:13 [WECHAT_DEBUG] [API] http response body:
{"openid":"os-IKuHd9pJ6xsn4mS7GyL4HxqI4","nickname":"产学红","sex":1,"language":"zh_CN","city":"安庆","province":"安徽","country":"中国","headimgurl":"http:\/\/wx.qlogo.cn\/mmopen\/O1HUibMqqHXduhNiagwbE0m4zgJU2YbFkyZPG6VoH8IP2wEdFuWcnjUtrXHNl1OmCsoffYBBnkC0cy1yfsOibcenaAn2SeRNKYw\/0","privilege":[]}
2016/03/11 14:44:13 userinfo: &{OpenId:os-IKuHd9pJ6xsn4mS7GyL4HxqI4 Nickname:产学红 Sex:1 City:安庆 Province:安徽 Country:中国 HeadImageURL:http://wx.qlogo.cn/mmopen/O1HUibMqqHXduhNiagwbE0m4zgJU2YbFkyZPG6VoH8IP2wEdFuWcnjUtrXHNl1OmCsoffYBBnkC0cy1yfsOibcenaAn2SeRNKYw/0 Privilege:[] UnionId:}
```
