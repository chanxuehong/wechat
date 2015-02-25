### 发送客服消息示例
```Go
package main

import (
	"fmt"

	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/message/custom"
)

var TokenServer = mp.NewDefaultTokenServer("appid", "appsecret", nil)

func main() {
	text := custom.NewText("touser", "content", "")

	clt := custom.NewClient(TokenServer, nil)
	if err := clt.SendText(text); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ok")
}
```
