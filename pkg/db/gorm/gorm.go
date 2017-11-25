package gorm

import(
	"fmt"

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
	if err = g.c.AutoMigrate(&ploioproto.Application{}, &ploioproto.Service{}, &ploioproto.Port{}, &ploioproto.Env{}).Error; err != nil {
		return err
	}
	fmt.Printf("\nInitialize: %#v", g.c)
	return nil
}

func (g *GormDB) Close() {
	g.c.Close()
}

func (g *GormDB) ApplicationSave(a *ploioproto.Application) error {
	fmt.Printf("\nSave: %#v", g.c)
	aQuery := &ploioproto.Application{}

	if g.c.Where("id = ?", a.ID).First(aQuery).RecordNotFound() { // If no app with that ID exists, then create it
		if err := g.c.Create(a).Error; err != nil {
			return err
		}
		return nil
	}
	if err := g.c.Save(a).Error; err != nil {
		return err
	}
	return nil
}