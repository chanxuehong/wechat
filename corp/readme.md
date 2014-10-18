# 微信公众平台 企业号 golang SDK

Version:   0.1.0

NOTE:      在 v1.0.0 之前 API 都有可能微调

## 简介

corp 包主要分为2个部分，client 和 server

client 主要实现的是“主动”请求功能，如创建自定义菜单，创建部门等等，
详见 https://github.com/chanxuehong/wechat/blob/master/corp/client/readme.md

server 主要实现的是“被动”接收消息和处理功能，如被动接收文本消息及回复，被动接收语音消息及回复等等，
详见 https://github.com/chanxuehong/wechat/blob/master/corp/server/readme.md

