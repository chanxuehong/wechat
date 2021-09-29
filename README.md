# wechat SDK for golang
https://github.com/chanxuehong/wechat

## 简介
| 模块  | 描述                     |
|-----:|:-------------------------|
| mp   | 微信公众平台 SDK           |
| mch  | 微信商户平台(微信支付) SDK   |

## 安装
go get -u github.com/chanxuehong/wechat/...

## 一点简单的帮助文档, 也许对你有作用
* [微信公众号 SDK 核心 package](/mp/core/README.md)
* [基本的 api 调用](/mp/README.md)
* [微信网页授权](/mp/oauth2/README.md)

## 联系方式
QQ群: 297489459

## 文档
代码下载下来后，放入 GOPATH 路径的 src 下面，可以在shell(windows 下面是 cmd) 里运行
```sh
godoc -http=:8080
```

然后在浏览器里地址栏输入
```sh
http://localhost:8080/
```
即可查看文档

## 授权(LICENSE)
[wechat is licensed under the Apache Licence, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0.html)
