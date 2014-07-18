# 微信公众平台 golang SDK

#### 要求 go1.3+，如果你的环境是 go1.3 以下的，可以参考 github.com/chanxuehong/util/pool 来修改 client.Client 和 server.Handler，或者直接用 sync.pool.patch 里的文件覆盖。

#### 因为目前我的公众号只有基本接口权限，所以大部分功能（特别是微信小店）没有经过测试，所以请大家使用过程中发现问题及时通知我，谢谢！

Version:   0.8.5

NOTE:      在 v1.0.0 之前 API 都有可能微调

联系方式： chanxuehong@gmail.com / 15967396@qq.com

QQ群：     297489459

## 简介
wechat 包主要分为三个部分，client、server 和 sns。

client 主要实现的是“主动”请求功能，如发送客服消息，群发消息，创建菜单，创建二维码等等，
详见 https://github.com/chanxuehong/wechat/blob/master/client/readme.md

server 主要实现的是“被动”接收消息和处理功能，如被动接收文本消息及回复，被动接收语音消息及回复等等，
详见 https://github.com/chanxuehong/wechat/blob/master/server/readme.md

sns    主要实现的是网页授权获取用户基本信息功能，即微信扫描网页上的二维码实现 OAuth2 授权登录和获取用户信息，
详见 https://github.com/chanxuehong/wechat/blob/master/sns/readme.md

## 安装
通过执行下列语句就可以完成安装

	go get -u github.com/chanxuehong/wechat/...

## 文档
通过上面步骤下载下来后，可以在shell(windows 下面是 cmd) 里运行

	godoc -http=:8080
	
然后在浏览器里地址栏输入 

	http://localhost:8080/
	
即可查看文档

## 授权(LICENSE)

wechat is licensed under the Apache Licence, Version 2.0
(http://www.apache.org/licenses/LICENSE-2.0.html).