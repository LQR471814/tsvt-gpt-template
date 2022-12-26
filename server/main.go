package main

import (
	"log"
	"net"
	"tsvt-gp-template/server/api"

	"github.com/LQR471814/grpcboot"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":7000")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	server := &APIServer{}
	api.RegisterServiceServer(grpcServer, server)

	config, err := grpcboot.InitializeConfig(grpcboot.Config{
		GRPCServer: grpcServer,
		Listener:   listener,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = grpcboot.Serve(config)
	if err != nil {
		log.Fatal(err)
	}
}
