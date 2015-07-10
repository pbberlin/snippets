package main

import "fmt"

type Server struct {
	addr     string
	timeout  int
	listener *string
	valx     string
}

func NewServer(addr string, options ...func(*Server)) (*Server, error) {
	l := &addr
	srv := Server{listener: l}

	for _, option := range options {
		option(&srv)
	}
	return &srv, nil
}

func ValX(newValX string) func(*Server) {
	return func(s *Server) {
		s.valx = newValX
	}
}

func main() {
	srv, _ := NewServer("localhost")

	timeout := func(srv *Server) {
		srv.timeout = 60
	}
	tls := func(srv *Server) {
		config := "xxx"
		*srv.listener += config
	}
	srv2, _ := NewServer(
		"localhost", timeout, tls, ValX("ss"))

	fmt.Printf("%+v\n", srv)
	fmt.Printf("%+v\n", srv2)
}
