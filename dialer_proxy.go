package multinet

import (
	"fmt"
	"io"
	"net"
	"net/netip"
	"net/url"

	"golang.org/x/net/proxy"
)

type proxyDialer struct {
	x string
	d proxy.Dialer
}

func newProxyDialer(x string) (Dialer, error) {
	u, err := url.Parse(x)
	if err != nil {
		return nil, fmt.Errorf("failed to parse proxy URL: %v", err)
	}
	d, err := proxy.FromURL(u, proxy.Direct)
	if err != nil {
		return nil, fmt.Errorf("failed to create proxy dialer: %v", err)
	}
	return &proxyDialer{u.Host, d}, nil
}

type socksConn interface {
	BoundAddr() net.Addr
}

func (d *proxyDialer) Dial(r netip.AddrPort) (io.ReadWriteCloser, netip.AddrPort, error) {
	c, err := d.d.Dial("tcp", r.String())
	if err != nil {
		return nil, netip.AddrPort{}, fmt.Errorf("failed to dial: %v", err)
	}
	l := c.(socksConn).BoundAddr().(*net.TCPAddr).AddrPort()
	return c, l, nil
}

func (d *proxyDialer) String() string {
	return "p:" + d.x
}
