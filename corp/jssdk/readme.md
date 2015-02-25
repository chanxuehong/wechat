### 获取 jsapi_ticket 示例
```Go
package main

import (
	"fmt"

	"github.com/chanxuehong/wechat/corp"
	"github.com/chanxuehong/wechat/corp/jssdk"
)

var TokenServer = corp.NewDefaultTokenServer("corpId", "corpSecret", nil)
var TicketServer = jssdk.NewDefaultTicketServer(TokenServer, nil)

func main() {
	fmt.Println(TicketServer.Ticket())
}
```