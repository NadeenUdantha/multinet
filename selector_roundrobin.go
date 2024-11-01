package multinet

import (
	"net/netip"
	"sync/atomic"
)

type roundRobinSelector struct {
	selectorBase
	c atomic.Uint32
}

func newRoundRobinSelector(cfg *Config, ds []Dialer) Selector {
	return &roundRobinSelector{selectorBase: selectorBase{cfg, ds}}
}

func (s *roundRobinSelector) Select(r netip.AddrPort) Dialer {
	return s.ds[s.c.Add(1)%uint32(len(s.ds))]
}
