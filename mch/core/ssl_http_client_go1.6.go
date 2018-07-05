// +build go1.6,!go1.7

package core

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

// NewTLSHttpClient 创建支持双向证书认证的 http.Client.
func NewTLSHttpClient(certFile, keyFile string) (httpClient *http.Client, err error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	return newTLSHttpClient(tlsConfig)
}

// NewTLSHttpClient2 创建支持双向证书认证的 http.Client.
func NewTLSHttpClient2(certPEMBlock, keyPEMBlock []byte) (httpClient *http.Client, err error) {
	cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	if err != nil {
		return nil, err
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	return newTLSHttpClient(tlsConfig)
}

func newTLSHttpClient(tlsConfig *tls.Config) (*http.Client, error) {
	dialTLS := func(network, addr string) (net.Conn, error) {
		return tls.DialWithDialer(&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}, network, addr, tlsConfig)
	}
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			Dial: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			DialTLS:               dialTLS,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}, nil
}
