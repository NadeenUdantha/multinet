package multinet

import (
	"math/rand"
	"net/netip"
	"sync"

	lru "github.com/hashicorp/golang-lru/v2"
)

type weightedLRUSelector struct {
	selectorBase
	c *lru.Cache[string, int]
	l sync.Mutex
}

func newWeightedLRUSelector(cfg *Config, ds []Dialer) Selector {
	c, err := lru.New[string, int](1000)
	if err != nil {
		panic(err) //wtf?
	}
	return &weightedLRUSelector{
		selectorBase: selectorBase{cfg, ds},
		c:            c,
	}
}

func (s *weightedLRUSelector) Select(r netip.AddrPort) Dialer {
	k := s.addrHashKey(r)
	s.l.Lock()
	defer s.l.Unlock()
	v, z := s.c.Get(k)
	if !z {
		v = rand.Int()
	}
	d := s.ds[v%len(s.ds)]
	v++
	s.c.Add(k, v)
	return d
}
