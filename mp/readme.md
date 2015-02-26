# 微信公众平台（订阅号、服务号） golang SDK

## 整体架构图

![架构图](https://github.com/chanxuehong/wechat/blob/master/mp/mp.png)

## 示例

### 一个 URL 监听一个公众号的消息
```Go
package main

import (
	"fmt"
	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/menu"
	"github.com/chanxuehong/wechat/util"
	"net/http"
)

func MenuClickEventHandler(w http.ResponseWriter, r *mp.Request) {
	event := menu.GetClickEvent(r.MixedMsg)
	fmt.Println(event.EventKey)
	return
}

func main() {
	aesKey, err := util.AESKeyDecode("encodedAESKey")
	if err != nil {
		panic(err)
	}

	messageServeMux := mp.NewMessageServeMux()
	messageServeMux.EventHandleFunc(menu.EventTypeClick, MenuClickEventHandler)

	wechatServer := mp.NewDefaultWechatServer("id", "token", "appid", aesKey, messageServeMux)

	wechatServerFrontend := mp.NewWechatServerFrontend(wechatServer, nil)

	http.Handle("/wechat", wechatServerFrontend)
	http.ListenAndServe(":80", nil)
}
```

### 一个 URL 监听多个公众号的消息
```Go
package main

import (
	"fmt"
	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/menu"
	"github.com/chanxuehong/wechat/util"
	"net/http"
)

func MenuClickEventHandler(w http.ResponseWriter, r *mp.Request) {
	event := menu.GetClickEvent(r.MixedMsg)
	fmt.Println(event.EventKey)
	return
}

func main() {
	aesKey1, err := util.AESKeyDecode("encodedAESKey1")
	if err != nil {
		panic(err)
	}

	messageServeMux1 := mp.NewMessageServeMux()
	messageServeMux1.EventHandleFunc(menu.EventTypeClick, MenuClickEventHandler)

	wechatServer1 := mp.NewDefaultWechatServer("id1", "token1", "appid1", aesKey1, messageServeMux1)

	aesKey2, err := util.AESKeyDecode("encodedAESKey2")
	if err != nil {
		panic(err)
	}

	messageServeMux2 := mp.NewMessageServeMux()
	messageServeMux2.EventHandleFunc(menu.EventTypeClick, MenuClickEventHandler)

	wechatServer2 := mp.NewDefaultWechatServer("id2", "token2", "appid2", aesKey2, messageServeMux2)

	var multiWechatServerFrontend mp.MultiWechatServerFrontend
	multiWechatServerFrontend.SetWechatServer("wechat1", wechatServer1) // 在回调 url 里要设置相应的参数
	multiWechatServerFrontend.SetWechatServer("wechat2", wechatServer2) // 在回调 url 里要设置相应的参数

	http.Handle("/wechat", &multiWechatServerFrontend)
	http.ListenAndServe(":80", nil)
}
```