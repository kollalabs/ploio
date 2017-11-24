package parser


import (
	"log"
	"gopkg.in/yaml.v2"
	"github.com/ploio/ploio/pkg/api/ploioproto"
)


func Parse(config []byte) (*ploioproto.Application, error) {
	a := &ploioproto.Application{}
	err := yaml.Unmarshal(config, a)
	if err != nil {
		log.Fatalf("error: %v", err)
		return a, err
	}
	return a, nil
}