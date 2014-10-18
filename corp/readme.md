# 首先来一个大广告哈哈
我的个人订阅号，刚开始做，先帮我凑齐500人先，不管你有没有需求哈，感谢了！

![产先生二维码](https://github.com/chanxuehong/wechat/corp/blob/master/qrcode_cxs0556.jpg)

# 微信公众平台 企业号 golang SDK


Version:   0.1.0

NOTE:      在 v1.0.0 之前 API 都有可能微调

联系方式： chanxuehong@gmail.com / 15967396@qq.com

QQ群：     297489459

## 简介
wechatcorp 包主要分为2个部分，client 和 server

client 主要实现的是“主动”请求功能，如创建自定义菜单，创建部门等等，
详见 https://github.com/chanxuehong/wechat/corp/blob/master/client/readme.md

server 主要实现的是“被动”接收消息和处理功能，如被动接收文本消息及回复，被动接收语音消息及回复等等，
详见 https://github.com/chanxuehong/wechat/corp/blob/master/server/readme.md

## 安装
通过执行下列语句就可以完成安装

	go get -u github.com/chanxuehong/wechat/corp/...

## 文档

### [在线文档](http://godoc.org/github.com/chanxuehong/wechat/corp)

### 离线文档
通过上面步骤下载下来后，可以在shell(windows 下面是 cmd) 里运行

	godoc -http=:8080
	
然后在浏览器里地址栏输入 

	http://localhost:8080/
	
即可查看文档

## 授权(LICENSE)

wechat is licensed under the Apache Licence, Version 2.0
(http://www.apache.org/licenses/LICENSE-2.0.html).