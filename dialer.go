package multinet

import (
	"fmt"
	"io"
	"net/netip"
)

type Dialer interface {
	Dial(remote netip.AddrPort) (c io.ReadWriteCloser, local netip.AddrPort, err error)
	String() string
}

type DialerBuilder func(x string) (Dialer, error)

var dbm = map[string]DialerBuilder{
	"direct": newDirectDialer,
	"proxy":  newProxyDialer,
}

func BuildDialers(cfg *Config) ([]Dialer, error) {
	var ds []Dialer
	for _, d := range cfg.Paths {
		df, ok := dbm[d.Type]
		if !ok {
			return nil, fmt.Errorf("Unknown dialer type: %s", d.Type)
		}
		d, err := df(d.Addr)
		if err != nil {
			return nil, fmt.Errorf("Failed to create dialer: %v", err)
		}
		ds = append(ds, d)
	}
	return ds, nil
}
