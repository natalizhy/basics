package main // server

import (
	"log"
	"net"

	"github.com/natalizhy/basics/grpcadder/pkg/adder"
	"github.com/natalizhy/basics/grpcadder/pkg/api"
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
