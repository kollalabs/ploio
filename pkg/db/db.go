package db

import (
	"fmt"
	"github.com/ploio/ploio/pkg/api/ploioproto"
)

var DB Database

type Database interface {
	Initialize() error
	Close()
	ApplicationSave(*ploioproto.Application) error
}

func RegisterDB(d Database) error {
	fmt.Printf("\nRegisterDB: %#v", d)
	DB = d
	err := d.Initialize()
	return err
}