package parser


import (
	"log"
	"gopkg.in/yaml.v2"
	"github.com/ploio/ploio/pkg/api/ploioproto"
)


func Marshal(config []byte) (*ploioproto.Application, error) {
	a := &ploioproto.Application{}
	err := yaml.Unmarshal(config, a)
	if err != nil {
		log.Fatalf("error: %v", err)
		return a, err
	}
	return a, nil
}


func Unmarshal(a *ploioproto.Application) ([]byte, error) {
	yaml, err := yaml.Marshal(a)
	return yaml, err
}