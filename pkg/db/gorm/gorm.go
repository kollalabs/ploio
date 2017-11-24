package gorm

import(
	"github.com/ploio/ploio/pkg/api/ploioproto"
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)


type GormDB struct {
	c *gorm.DB
}


func (g GormDB) Initialize() error{
	var err error
	g.c, err = gorm.Open("sqlite3", "ploio.db")
	if err != nil {
		return err
	}
	g.c.AutoMigrate(&ploioproto.Application{}, &ploioproto.Service{}, &ploioproto.Port{}, &ploioproto.Env{})
	return nil
}

func (g *GormDB) Close() {
	g.c.Close()
}