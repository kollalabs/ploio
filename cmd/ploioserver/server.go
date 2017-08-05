package main

import (
	"github.com/asdine/storm"
	"golang.org/x/net/context"

	"github.com/weave-lab/ploio/pkg/api/ploioproto"
)

var db *storm.DB

func init() {
	db, err := storm.Open("ploio.db")
}

type ploioserver struct{}

//ListApplications gets a list of all the applications from the ploio server, or supply a filter by name or ID
func (p *ploioserver) ListApplications(c context.Context, ag *ploioproto.ApplicationGet) (*ploioproto.ApplicationList, error) {
	result := &ploioproto.ApplicationList{}
	return result, nil
}

//GetApplication gets a single application by either name or ID
func (p *ploioserver) GetApplication(c context.Context, ag *ploioproto.ApplicationGet) (*ploioproto.Application, error) {
	result := &ploioproto.Application{}
	return result, nil
}

//CreateApplication creates a new application
func (p *ploioserver) CreateApplication(c context.Context, ac *ploioproto.ApplicationCreate) (*ploioproto.Application, error) {
	result := &ploioproto.Application{
		Name:     ac.Name,
		Owner:    ac.Owner,
		Repo:     ac.Repo,
		Stateful: ac.Stateful,
	}
	return result, nil
}
