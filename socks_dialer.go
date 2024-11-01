package multinet

import (
	"io"
	"net"

	"github.com/cloudflare/cloudflared/socks"
)

type SocksDialer struct {
	d Dialer
}

func NewSocksDialer(d Dialer) socks.Dialer {
	return &SocksDialer{d}
}

func (d *SocksDialer) Dial(x string) (io.ReadWriteCloser, *socks.AddrSpec, error) {
	a, err := net.ResolveTCPAddr("tcp", x)
	if err != nil {
		return nil, nil, err
	}
	ap := a.AddrPort()
	c, l, err := d.d.Dial(ap)
	if err != nil {
		return nil, nil, err
	}
	return c, &socks.AddrSpec{
		IP:   l.Addr().AsSlice(),
		Port: int(l.Port()),
	}, nil
}
