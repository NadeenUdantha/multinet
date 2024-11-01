package multinet

import (
	"fmt"
	"net/netip"
)

type selectorBase struct {
	cfg *Config
	ds  []Dialer
}

func (s *selectorBase) addrHashKey(r netip.AddrPort) string {
	if s.cfg.HashPort {
		return r.String()
	}
	return r.Addr().String()
}

type Selector interface {
	Select(r netip.AddrPort) Dialer
}

type SelectorBuilder func(cfg *Config, ds []Dialer) Selector

var sbm = map[string]SelectorBuilder{
	"random":       newRandomSelector,
	"hash":         newAddrHashSelector,
	"round-robin":  newRoundRobinSelector,
	"weighted-lru": newWeightedLRUSelector,
	//todo: least utilized path selector
}

func GetSelector(cfg *Config) (Selector, error) {
	sb, z := sbm[cfg.Algorithm]
	if !z {
		return nil, fmt.Errorf("Unknown selector: %v", cfg.Algorithm)
	}
	ds, err := BuildDialers(cfg)
	if err != nil {
		return nil, err
	}
	return sb(cfg, ds), nil
}
