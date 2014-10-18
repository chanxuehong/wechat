## 简介

这个 package 封装了微信公众平台的主动请求功能，如发送客服消息、群发消息、创建自定义菜单、
创建二维码等等...

## access token 逻辑图
![产先生二维码](https://github.com/chanxuehong/wechat/blob/dev/mp/client/token_service.png)

## 使用方法

大部分功能都是 Client 对象的方法, 根据对应的功能调用对应的方法.

NewClient 函数的参数 TokenService 可以自己实现, 也可以用默认实现 *DefaultTokenService, 

采用默认实现的时候要注意, 对于一个特定的 appid, 只能有一个 DefaultTokenService 的实例,
一般的做法就是把这个实例作为全局对象!!!

## 示例

这个实例是创建一个临时的二维码, TokenService 采用默认的实现.
```golang
package main

import (
	"fmt"
	"github.com/chanxuehong/wechat/mp/client"
)

// *DefaultTokenService 实现了 TokenService 接口.
// 当然你也可以不用默认的实现, 这个时候就需要你自己实现 TokenService 接口了,
// 根据你自己的实现, tokenService 不一定要求作为全局变量,
// 但是如果用默认的实现 client.NewDefaultTokenService, 一个 appid 只能有一个实例.
var tokenService = client.NewDefaultTokenService(
	"xxxxxxxxxxxxxxxxxx",               // 公众号 appid
	"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", // 公众号 appsecret
	nil,
)

func main() {
	wechatClient := client.NewClient(tokenService, nil)

	qrcode, err := wechatClient.QRCodeTemporaryCreate(100, 1000)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(qrcode)
}
```
