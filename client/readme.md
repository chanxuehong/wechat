## 简介

这个 package 封装了微信公众平台的主动请求功能，如发送客服消息、群发消息、创建自定义菜单、
创建二维码等等。

## 使用方法

一般先全局创建一个 Client 实例，然后根据需要的功能调用这个实例相应的方法；
详细可参见文档。

## 示例

这个实例是创建一个临时的二维码
```Go
package main

import (
	"fmt"
	"github.com/chanxuehong/wechat/client"
)

// 一个应用一个实例
var wechatClient *client.Client

func init() {
	// TODO: 获取必要数据的代码

	// 初始化 wechatClient
	wechatClient = client.NewClient("你的公众号-appid", "你的公众号-appsecret", nil)
}

func main() {
	qrcode, err := wechatClient.QRCodeTemporaryCreate(100, 1000)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(qrcode)
}
```

## 测试信息

https://github.com/chanxuehong/wechat/blob/master/client/testinfo.txt