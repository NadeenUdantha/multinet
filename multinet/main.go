package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/NadeenUdantha/multinet"
)

func main() {
	if err := main2(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main2() error {
	fp := flag.String("cfg", "", "config file")
	flag.Parse()

	if *fp == "" {
		flag.Usage()
		return nil
	}

	cfg, err := multinet.LoadConfig(*fp)
	if err != nil {
		return err
	}
	s, err := multinet.NewServer(cfg)
	if err != nil {
		return err
	}
	defer s.Close()

	if err := s.Listen(); err != nil {
		return err
	}

	ec := make(chan error)
	go func() { ec <- s.Serve() }()

	//todo: this whole signal bullshit does not work :)
	c := make(chan os.Signal, 10)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	t := 0
	for {
		select {
		case <-c:
		f:
			for {
				select {
				case <-c:
				default:
					break f
				}
			}
			if t == 0 {
				s.Close()
				t++
			} else {
				os.Exit(1)
			}
		case err := <-ec:
			return err
		}
	}
}
