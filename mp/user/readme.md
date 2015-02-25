# 用户管理接口

### 打印所有关注用户的示例
```Go
package main

import (
	"fmt"

	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/user"
)

var TokenServer = mp.NewDefaultTokenServer("appid", "appsecret", nil)

func main() {
	clt := user.NewClient(TokenServer, nil)

	iter, err := clt.UserIterator("")
	if err != nil {
		fmt.Println(err)
		return
	}

	for iter.HasNext() {
		openIds, err := iter.NextPage()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, openId := range openIds {
			userinfo, err := clt.UserInfo(openId, user.Language_zh_CN)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("%+v\r\n", userinfo)
		}
	}
}
```