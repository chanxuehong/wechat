## 简介

这个 package 封装了微信公众平台的网页授权获取用户信息功能。

## 示例

```Go
package main

import (
	"github.com/chanxuehong/wechat/oauth2"
	"net/http"
)

// 一个应用只用一个全局变量（实际应用中根据自己的参数填写）
var oauth2Config = oauth2.NewOAuth2Config("appid", "appsecret", "redirectURL", "scope0", "scope1")

// 在需要用户授权的时候引导用户到授权页面
func SomeHandler(w http.ResponseWriter, r *http.Request) {
	var hasAuth bool // 判断是否授权
	// TODO: 判断 session 里是否有 openid 字段，如果有则表明已经授权，没有则没有授权

	if hasAuth {
		var info *oauth2.OAuth2Info
		// TODO: 根据 openid 字段 找到 info(type == *oauth2.OAuth2Info)

		client := oauth2.Client{
			OAuth2Config: oauth2Config,
			OAuth2Info:   info,
		}

		// 可以根据这个 info 做相应的操作，比如下面的获取用户信息
		userinfo, err := client.UserInfo("zh_CN")
		if err != nil {
			// TODO: ...
			return
		}

		// 处理 userinfo
		_ = userinfo // NOTE: 实际开发中可不是这样的

	} else {
		var state = "state" // NOTE: 实际上是一串不可预测的随机数

		// TODO: state 做相应处理，好识别下次跳转回来的 state 参数

		http.Redirect(w, r, oauth2Config.AuthCodeURL(state), http.StatusFound)
	}
}

// 跳转后的页面请求处理.
// redirectURL?code=CODE&state=STATE // 授权
// redirectURL?state=STATE           // 不授权
func RedirectURLHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		// TODO: 处理 error
		return
	}

	state := r.FormValue("state")

	// TODO: 处理这个 state 参数, 判断是否是有效的

	_ = state // NOTE: 实际开发中不要有这个

	// 假定 state 有效

	if code := r.FormValue("code"); code != "" { // 授权

		client := oauth2.Client{
			OAuth2Config: oauth2Config,
		}
		info, err := client.Exchange(code)
		if err != nil {
			// TODO: ...
			return
		}

		// TODO: 这里把 info 根据 info.OpenId 缓存起来，以后可以直接用
		// 做相应的 session 处理。
		_ = info // NOTE: 示例代码

	} else { // 不授权
		// TODO: 不授权的相应代码
	}
}

func main() {
	// TODO: 为 http 添加路由处理，然后在运行 http service
}
```