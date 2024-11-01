package multinet

import (
	"fmt"
	"io"
	"net/netip"
)

type selectorDialer struct {
	s Selector
}

func newSelectorDialer(x Selector) Dialer {
	return &selectorDialer{x}
}

func (d *selectorDialer) Dial(r netip.AddrPort) (io.ReadWriteCloser, netip.AddrPort, error) {
	x := d.s.Select(r)
	c, a, err := x.Dial(r)
	//todo: do something with the error
	rs := netip.AddrPortFrom(r.Addr().Unmap(), r.Port()).String()
	as := "error"
	if a.IsValid() {
		as = a.String()
	}
	fmt.Printf("dial %s => %s => %s\n", rs, x.String(), as)
	return c, a, err
}

func (d *selectorDialer) String() string {
	return "selector"
}
