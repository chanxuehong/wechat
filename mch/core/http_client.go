package core

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

// NewTLSHttpClient 创建支持双向证书认证的 http.Client
func NewTLSHttpClient(certFile, keyFile string) (httpClient *http.Client, err error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	httpClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			TLSClientConfig:       tlsConfig,
		},
		Timeout: 60 * time.Second,
	}
	return
}
