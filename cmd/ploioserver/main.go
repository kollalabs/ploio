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
	gdb := gorm.GormDB{}
	db.RegisterDB(&gdb)
	
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

func (p *ploioserver) ListApplications(c context.Context, alf *pp.ApplicationListFilter) (*pp.ApplicationList, error) {
	al := &pp.ApplicationList{}
	return al, nil
}

func (p *ploioserver) GetApplication(c context.Context, ag *pp.ApplicationID) (*pp.Application, error) {
	aNew := &pp.Application{}
	return aNew, nil
}

func (p *ploioserver) UpsertApplication(c context.Context, a *pp.Application) (*pp.Application, error) {
	err := db.DB.ApplicationSave(a)
	if err != nil {
		return a, err
	}
	return a, nil
}

func (p *ploioserver) DeleteApplication(c context.Context, aid *pp.ApplicationID) (*pp.ApplicationID, error) {
	aidNew := &pp.ApplicationID{}
	return aidNew, nil
}