package req

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/url"

	"golang.org/x/net/proxy"
)

const (
	protocolHTTP  = "http"
	protocolHTTPS = "https"
	protocolSOCKS = "socks5"
)

var ErrInvalidProxyURL = errors.New("invalid proxyURL format (It should be http[s]/socks5://[username:password@]host:port)")

// socks5Dialer is a dialer for socks5 proxy
// It uses a socks5 proxy to dial the network connection
// TODO: handle error when dialing
func socks5DialerContext(proxyURL string) func(ctx context.Context, network, addr string) (net.Conn, error) {
	var proxyAuth *proxy.Auth
	socksURL, err := url.Parse(proxyURL)
	if err == nil {
		proxyAuth = &proxy.Auth{}
		proxyAuth.User = socksURL.User.Username()
		proxyAuth.Password, _ = socksURL.User.Password()
	}
	proxyAuth = &proxy.Auth{}
	proxyAuth.User = socksURL.User.Username()
	proxyAuth.Password, _ = socksURL.User.Password()
	dialer, proxyErr := proxy.SOCKS5("tcp", fmt.Sprintf("%s:%s", socksURL.Hostname(), socksURL.Port()), proxyAuth, proxy.Direct)
	dc := dialer.(interface {
		DialContext(ctx context.Context, network, addr string) (net.Conn, error)
	})
	if proxyErr == nil {
		return dc.DialContext
	}
	return nil
}

func validateProxyURL(proxyURL string) (*url.URL, error) {
	if u, err := url.Parse(proxyURL); err == nil && isSupportedProtocol(u.Scheme) {
		return u, nil
	}
	return nil, ErrInvalidProxyURL
}

// isSupportedProtocol checks given protocols are supported
func isSupportedProtocol(value string) bool {
	return value == protocolHTTP || value == protocolHTTPS || value == protocolSOCKS
}
