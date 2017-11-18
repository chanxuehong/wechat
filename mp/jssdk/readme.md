### 获取 jsapi_ticket 示例
```Go
package main

import (
	"fmt"

	"gopkg.in/chanxuehong/wechat.v1/mp"
	"gopkg.in/chanxuehong/wechat.v1/mp/jssdk"
)

var AccessTokenServer = mp.NewDefaultAccessTokenServer("appid", "appsecret", nil)
var mpClient = mp.NewClient(AccessTokenServer, nil)
var TicketServer = jssdk.NewDefaultTicketServer(mpClient)

func main() {
	fmt.Println(TicketServer.Ticket())
}
```