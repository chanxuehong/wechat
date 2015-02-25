# 群发消息接口

### 根据分组进行群发消息示例
```Go
package main

import (
	"fmt"

	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/message/mass/mass2group"
)

var TokenServer = mp.NewDefaultTokenServer("appid", "appsecret", nil)

func main() {
	text := mass2group.NewText(1 /* groupid */, "content")

	clt := mass2group.NewClient(TokenServer, nil)

	msgId, err := clt.SendText(text)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("msgId:", msgId)
}
```
