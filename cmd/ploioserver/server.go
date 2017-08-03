package main

import (
	"golang.org/x/net/context"

	"github.com/weave-lab/ploio/pkg/api/ploioproto"
)

type ploioserver struct{}

//ListApplications gets a list of all the applications from the ploio server, or supply a filter by name or ID
func (p *ploioserver) ListApplications(context.Context, *ploioproto.ApplicationGet) (*ploioproto.ApplicationList, error) {
	result := &ploioproto.ApplicationList{}
	return result, nil
}

//GetApplication gets a single application by either name or ID
func (p *ploioserver) GetApplication(context.Context, *ploioproto.ApplicationGet) (*ploioproto.Application, error) {
	result := &ploioproto.Application{}
	return result, nil
}

//CreateApplication creates a new application
func (p *ploioserver) CreateApplication(context.Context, *ploioproto.ApplicationCreate) (*ploioproto.Application, error) {
	result := &ploioproto.Application{}
	return result, nil
}
