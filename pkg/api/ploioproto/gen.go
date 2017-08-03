package ploioproto

//go:generate protoc  -I=./ -I=$GOPATH/src --go_out=plugins=grpc:. api.proto

//protoc -I=./ -I=$GOPATH/src/github.com/weave-lab/ploio/vendor -I=$GOPATH/src --go_out=plugins=grpc:. api.proto
