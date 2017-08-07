package apiclient

import (
	"log"

	pp "github.com/weave-lab/ploio/pkg/api/ploioproto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9000"
)

//Client is the grpc client that the other packages use
var Client pp.PloioAPIClient

func init() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	Client = pp.NewPloioAPIClient(conn)
}
