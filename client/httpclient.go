// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"net"
	"net/http"
	"time"
)

/*
net/http:

var DefaultTransport RoundTripper = &Transport{
    Proxy: ProxyFromEnvironment,
    Dial: (&net.Dialer{
        Timeout:   30 * time.Second,
        KeepAlive: 30 * time.Second,
    }).Dial,
    TLSHandshakeTimeout: 10 * time.Second,
}
*/

// 一般的请求
var CommonHttpClient = &http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   5 * time.Second, // 连接超时设置为 5 秒
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second, // TLS 握手超时设置为 5 秒
	},
	Timeout: 15 * time.Second, // 请求超时时间设置为 15 秒
}

// 多媒体上传下载请求
var MediaHttpClient = &http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   5 * time.Second, // 连接超时设置为 5 秒
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second, // TLS 握手超时设置为 5 秒
	},
	// 因为目前微信支持最大的文件是 10MB, 请求超时时间保守设置为 100 秒
	Timeout: 300 * time.Second,
}
