package multinet

import (
	"fmt"
	"io"
	"net"
	"net/netip"
)

type directDialer struct {
	x string
	l *net.TCPAddr
}

func newDirectDialer(x string) (Dialer, error) {
	a, err := netip.ParseAddr(x)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse local address: %v", err)
	}
	ap := netip.AddrPortFrom(a, 0)
	return &directDialer{
		x: a.String(),
		l: net.TCPAddrFromAddrPort(ap),
	}, nil
}

func (d *directDialer) Dial(r netip.AddrPort) (io.ReadWriteCloser, netip.AddrPort, error) {
	c, err := net.DialTCP("tcp", d.l, net.TCPAddrFromAddrPort(r))
	if err != nil {
		return nil, netip.AddrPort{}, fmt.Errorf("Failed to dial: %v", err)
	}
	l := c.LocalAddr().(*net.TCPAddr).AddrPort()
	return c, l, nil
}

func (d *directDialer) String() string {
	return "d:" + d.x
}
