package db

import (
	"github.com/ploio/ploio/pkg/api/ploioproto"
)

var DB Database

type Database interface {
	Initialize() error
	Close()
	ApplicationSave(*ploioproto.Application) error
}

func RegisterDB(d Database) error {
	err := d.Initialize()
	DB = d
	return err
}