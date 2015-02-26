# 微信企业号 golang SDK

## 整体架构图

![架构图](https://github.com/chanxuehong/wechat/blob/master/corp/corp.png)

## 示例

### 一个 URL 监听一个企业号应用的消息
```Go
package main

import (
	"fmt"
	"github.com/chanxuehong/wechat/corp"
	"github.com/chanxuehong/wechat/corp/menu"
	"github.com/chanxuehong/wechat/util"
	"net/http"
)

func MenuClickEventHandler(w http.ResponseWriter, r *corp.Request) {
	event := menu.GetClickEvent(r.MixedMsg)
	fmt.Println(event.EventKey)
	return
}

func main() {
	aesKey, err := util.AESKeyDecode("encodedAESKey")
	if err != nil {
		panic(err)
	}

	messageServeMux := corp.NewMessageServeMux()
	messageServeMux.EventHandleFunc(menu.EventTypeClick, MenuClickEventHandler)

	agentServer := corp.NewDefaultAgentServer("corpId", 0 /* agentId */, "token", aesKey, messageServeMux)

	agentServerFrontend := corp.NewAgentServerFrontend(agentServer, nil)

	http.Handle("/agent", agentServerFrontend)
	http.ListenAndServe(":80", nil)
}
```

### 一个 URL 监听多个企业号应用的消息
```Go
package main

import (
	"fmt"
	"github.com/chanxuehong/wechat/corp"
	"github.com/chanxuehong/wechat/corp/menu"
	"github.com/chanxuehong/wechat/util"
	"net/http"
)

func MenuClickEventHandler(w http.ResponseWriter, r *corp.Request) {
	event := menu.GetClickEvent(r.MixedMsg)
	fmt.Println(event.EventKey)
	return
}

func main() {
	aesKey1, err := util.AESKeyDecode("encodedAESKey1")
	if err != nil {
		panic(err)
	}

	messageServeMux1 := corp.NewMessageServeMux()
	messageServeMux1.EventHandleFunc(menu.EventTypeClick, MenuClickEventHandler)

	agentServer1 := corp.NewDefaultAgentServer("corpId", 1 /* agentId */, "token1", aesKey1, messageServeMux1)

	aesKey2, err := util.AESKeyDecode("encodedAESKey2")
	if err != nil {
		panic(err)
	}

	messageServeMux2 := corp.NewMessageServeMux()
	messageServeMux2.EventHandleFunc(menu.EventTypeClick, MenuClickEventHandler)

	agentServer2 := corp.NewDefaultAgentServer("corpId", 2 /* agentId */, "token2", aesKey2, messageServeMux2)

	var multiAgentServerFrontend corp.MultiAgentServerFrontend
	multiAgentServerFrontend.SetAgentServer("agent1", agentServer1) // 在回调 url 里要设置相应的参数
	multiAgentServerFrontend.SetAgentServer("agent2", agentServer2) // 在回调 url 里要设置相应的参数

	http.Handle("/wechat", &multiAgentServerFrontend)
	http.ListenAndServe(":80", nil)
}
```