package main

import (
	"fmt"
	"log"
	"net"
)

const port = ":4242"

type listener struct {
	l net.Listener
}

func NewTCPListener() listener {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	return listener{l: l}
}

func (l listener) Accept() {
	_, _ = l.l.Accept()
}

type client struct {
	n net.Dialer
}

func NewClient() client {
	return client{
		n: net.Dialer{},
	}
}

func (c client) EstablishConnection() {
	_, err := c.n.Dial("tcp", port)
	if err != nil {
		log.Fatalf("failed to establish connection with %s\n", port)
	}

	fmt.Printf("established connection with %s!\n", port)
}
