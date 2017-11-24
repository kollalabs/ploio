package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"golang.org/x/net/context"

	pp "github.com/ploio/ploio/pkg/api/ploioproto"
	"github.com/ploio/ploio/pkg/db"
	"github.com/ploio/ploio/pkg/db/gorm"	

)

const (
	port = ":9000"
)

func main() {
	gdb := &gorm.GormDB{}
	db.RegisterDB(gdb)
	
	fmt.Println("Starting ploio server on " + port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pp.RegisterPloioAPIServer(s, &ploioserver{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


type ploioserver struct {}

func (p ploioserver) GetApplication(c context.Context, ag *pp.ApplicationGet) (*pp.Application, error) {
	aNew := &pp.Application{}
	return aNew, nil
}