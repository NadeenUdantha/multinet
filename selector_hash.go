package multinet

import (
	"crypto/sha256"
	"net/netip"
)

type addrHashSelector struct{ selectorBase }

func newAddrHashSelector(cfg *Config, ds []Dialer) Selector {
	return addrHashSelector{selectorBase{cfg, ds}}
}

func (s addrHashSelector) Select(r netip.AddrPort) Dialer {
	h := sha256.Sum256([]byte(s.addrHashKey(r)))
	return s.ds[int(h[0])%len(s.ds)]
}
