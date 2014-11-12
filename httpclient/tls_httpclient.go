// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package httpclient

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

func NewTLSHTTPClient(certFile, keyFile string) (httpClient *http.Client, err error) {
	var tlsConfig tls.Config
	tlsConfig.Certificates = make([]tls.Certificate, 1)
	tlsConfig.Certificates[0], err = tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return
	}

	httpClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			Dial: (&net.Dialer{
				Timeout:   5 * time.Second, // 连接超时设置为 5 秒
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSClientConfig:     &tlsConfig,
			TLSHandshakeTimeout: 5 * time.Second, // TLS 握手超时设置为 5 秒
		},
		Timeout: 15 * time.Second, // 请求超时时间设置为 15 秒
	}
	return
}
