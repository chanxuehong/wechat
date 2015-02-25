# 群发消息接口

### 根据分组进行群发消息示例
```Go
package main

import (
	"fmt"

	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/message/mass/masstogroup"
)

var TokenServer = mp.NewDefaultTokenServer("appid", "appsecret", nil)

func main() {
	text := masstogroup.NewText(1 /* groupid */, "content")

	clt := masstogroup.NewClient(TokenServer, nil)

	msgId, err := clt.SendText(text)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("msgId:", msgId)
}
```
