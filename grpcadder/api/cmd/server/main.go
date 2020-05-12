package main

import (
	"github.com/natalizhy/basics/grpcadder/pkg/adder"
	"github.com/natalizhy/basics/grpcadder/pkg/api"
	"log"
	"net"
	qrpc "google.golang.org/grpc"
)

func main() {
	s := qrpc.NewServer()
	srv := &adder.GRPCServer{}
	api.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != err {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
