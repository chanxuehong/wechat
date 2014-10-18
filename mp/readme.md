# 微信公众平台 订阅号、服务号 golang SDK

Version:   0.8.16

NOTE:      在 v1.0.0 之前 API 都有可能微调

## change log:

#### v0.8.16
server WriteXXX 系列方法独立成函数
#### v0.8.15
server 增加消息体签名及加解密
#### v0.8.14
client 包里更改了qrcode相关的函数和方法
#### v0.8.13
server 里的数据结构命名做了调整
#### v0.8.12
server 消息处理方式重构了
#### v0.8.11
client 的 Media 系列方法名称做了个调整

## 简介
<<<<<<< HEAD
wechat 包主要分为三个部分，client、server 和 oauth2
=======
mp 包主要分为三个部分，client、server 和 oauth2
>>>>>>> origin/dev

client 主要实现的是“主动”请求功能，如发送客服消息，群发消息，创建菜单，创建二维码等等，
详见 https://github.com/chanxuehong/wechat/blob/master/mp/client/readme.md

server 主要实现的是“被动”接收消息和处理功能，如被动接收文本消息及回复，被动接收语音消息及回复等等，
详见 https://github.com/chanxuehong/wechat/blob/master/mp/server/readme.md

oauth2 主要实现的是网页授权获取用户基本信息功能，微信公众号可以引导（自定义菜单或者网页）到一个页面，
请求用户授权，详见 https://github.com/chanxuehong/wechat/blob/master/mp/oauth2/readme.md

