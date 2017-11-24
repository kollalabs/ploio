package db


type Database interface {
	Initialize() error
	Close()
}

func RegisterDB(d Database) error {
	err := d.Initialize()
	return err
}