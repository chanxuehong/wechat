# 微信公众平台 golang SDK
Version:   0.8.16

NOTE:      在 v1.0.0 之前 API 都有可能微调

联系方式： chanxuehong@gmail.com / 15967396@qq.com

QQ群：     297489459

企业号移步到 https://github.com/chanxuehong/wechat/tree/dev/corp

# 首先来一个大广告哈哈
我的个人订阅号，刚开始做，先帮我凑齐500人先，不管你有没有需求哈，感谢了！

![产先生二维码](https://github.com/chanxuehong/wechat/blob/master/qrcode_cxs0556.jpg)

# change log:

### v0.8.16
server WriteXXX 系列方法独立成函数
### v0.8.15
server 增加消息体签名及加解密
### v0.8.14
client 包里更改了qrcode相关的函数和方法
### v0.8.13
server 里的数据结构命名做了调整
### v0.8.12
server 消息处理方式重构了
### v0.8.11
client 的 Media 系列方法名称做了个调整

# 简介
wechat 包主要分为三个部分，client、server 和 oauth2

client 主要实现的是“主动”请求功能，如发送客服消息，群发消息，创建菜单，创建二维码等等，
详见 https://github.com/chanxuehong/wechat/blob/dev/mp/client/readme.md

server 主要实现的是“被动”接收消息和处理功能，如被动接收文本消息及回复，被动接收语音消息及回复等等，
详见 https://github.com/chanxuehong/wechat/blob/dev/mp/server/readme.md

oauth2 主要实现的是网页授权获取用户基本信息功能，微信公众号可以引导（自定义菜单或者网页）到一个页面，
请求用户授权，详见 https://github.com/chanxuehong/wechat/blob/dev/mp/oauth2/readme.md


# 安装
通过执行下列语句就可以完成安装

	go get -u github.com/chanxuehong/wechat/mp/...

# 文档

### [在线文档](http://godoc.org/github.com/chanxuehong/wechat/mp)

### 离线文档
通过上面步骤下载下来后，可以在shell(windows 下面是 cmd) 里运行

	godoc -http=:8080
	
然后在浏览器里地址栏输入 

	http://localhost:8080/
	
即可查看文档

# 授权(LICENSE)

wechat is licensed under the Apache Licence, Version 2.0
(http://www.apache.org/licenses/LICENSE-2.0.html).
