package multinet

import (
	"fmt"
	"net"

	"github.com/cloudflare/cloudflared/socks"
)

type Server struct {
	l   net.Listener
	cfg *Config
	h   socks.ConnectionHandler
	d   Dialer
}

func NewServer(cfg *Config) (*Server, error) {
	return NewServerWithListener(cfg, nil)
}

func NewServerWithListener(cfg *Config, l net.Listener) (*Server, error) {
	s, err := GetSelector(cfg)
	if err != nil {
		return nil, err
	}
	d := NewSocksDialer(newSelectorDialer(s))
	rh := socks.NewRequestHandler(d, nil)
	return &Server{
		cfg: cfg,
		l:   l,
		h:   socks.NewConnectionHandler(rh),
	}, nil
}

func (s *Server) Listen() error {
	if s.l != nil {
		return fmt.Errorf("Already listening")
	}

	l, err := net.Listen("tcp", s.cfg.Listen)
	if err != nil {
		return fmt.Errorf("Failed to listen: %v\n", err)
	}
	s.l = l
	fmt.Println("Listening on", s.cfg.Listen)

	return nil
}

func (s *Server) Serve() error {
	if s.l == nil {
		return fmt.Errorf("Not listening")
	}
	fmt.Println("Serving")
	for {
		c, err := s.l.Accept()
		if err != nil {
			return fmt.Errorf("Failed to accept: %v", err)
		}
		go handle(s.h, c)
	}
}

func (s *Server) Close() error {
	return s.l.Close()
}

func handle(h socks.ConnectionHandler, c net.Conn) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	defer c.Close()
	err := h.Serve(c)
	if err != nil {
		fmt.Println(err)
	}
}
