package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ploio/ploio/pkg/api/ploioproto"
	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

func main() {
	fmt.Println("Starting ploio server on " + port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	ploioproto.RegisterPloioAPIServer(s, &ploioserver{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
