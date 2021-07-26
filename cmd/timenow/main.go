package main

import (
	"log"
	"net"

	"github.com/littlefut/grpc-test/pkg/api"
	"github.com/littlefut/grpc-test/timenow"
	"google.golang.org/grpc"
)

func main() {
	const port = ":8000"

	server := grpc.NewServer()
	service := timenow.NewService()

	api.RegisterTimenowServiceServer(server, service)
	log.Printf("Server listening at %s...\n", port)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("server crashed: %v", err)
	}
	if err = server.Serve(listener); err != nil {
		log.Fatalf("server crashed: %v", err)
	}
}
