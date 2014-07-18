// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

// 封装微信服务器推送过来的常规消息(事件)处理 Handler.
// 对于公众号开发模式, 都会要求提供一个 URL 来处理微信服务器推送过来的消息和事件,
// 这个 package 就是封装推送到这个 URL 上的消息(事件)处理 Handler.
package server
