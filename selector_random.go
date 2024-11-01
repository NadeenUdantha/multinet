package multinet

import (
	"math/rand"
	"net/netip"
)

type randomSelector selectorBase

func newRandomSelector(cfg *Config, ds []Dialer) Selector {
	return randomSelector(selectorBase{cfg, ds})
}

func (s randomSelector) Select(r netip.AddrPort) Dialer {
	return s.ds[rand.Intn(len(s.ds))]
}
