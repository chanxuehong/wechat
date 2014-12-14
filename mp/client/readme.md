## 简介

这个 package 封装了微信公众平台的主动请求功能，如发送客服消息、群发消息、创建自定义菜单、
创建二维码等等...

## 使用方法

大部分功能都是 Client 对象的方法, 根据对应的功能调用对应的方法.

## 示例

这个实例是创建一个临时的二维码, wechat.TokenService 采用默认的实现.
```golang
package main

import (
	"fmt"

	"github.com/chanxuehong/wechat/mp/client"
)

// *client.DefaultTokenService 实现了 wechat.TokenService 接口.
// 当然你也可以不用默认的实现, 这个时候就需要你自己实现 wechat.TokenService 接口了,
// 根据你自己的实现, TokenService 不一定要求作为全局变量,
// 但是如果用默认的实现 client.NewDefaultTokenService, 一个 appid 只能有一个实例.
var TokenService = client.NewDefaultTokenService(
	"xxxxxxxxxxxxxxxxxx",               // 公众号 appid
	"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", // 公众号 appsecret
	nil,
)

func main() {
	wechatClient := client.NewClient(TokenService, nil)

	qrcode, err := wechatClient.QRCodeTemporaryCreate(100, 1000)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(qrcode)
}
```
