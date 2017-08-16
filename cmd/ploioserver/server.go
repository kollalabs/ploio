package main

import (
	"fmt"

	"github.com/asdine/storm"
	"golang.org/x/net/context"

	pp "github.com/ploio/ploio/pkg/api/ploioproto"
)

var db *storm.DB

func init() {
	var err error
	db, err = storm.Open("ploio.db")
	if err != nil {
		fmt.Print(err)
		panic("Error initializing database on startup")
	}
}

type ploioserver struct{}

// DB Object structs

type application struct {
	ID       int32  `storm:"id,increment"`
	Name     string `storm:"unique"`
	Owner    string
	Repo     string
	Stateful bool
}

type pipeline struct {
	ID            int32 `storm:"id,increment"`
	Name          string
	ApplicationID int32
	Order         int32
}

type step struct {
	ID         int32
	Name       string
	PipelineID int32
	Type       string
	Order      int32
	Memory     string
	Proc       string
	Replicas   int32
}

//ListApplications gets a list of all the applications from the ploio server, or supply a filter by name or ID
func (p *ploioserver) ListApplications(c context.Context, ag *pp.ApplicationGet) (*pp.ApplicationList, error) {
	var apps []application
	result := &pp.ApplicationList{}
	err := db.All(&apps)
	if err != nil {
		return result, err
	}
	for _, a := range apps {
		ap := &pp.Application{
			ID:    a.ID,
			Name:  a.Name,
			Owner: a.Owner,
			Repo:  a.Repo,
		}
		result.Applications = append(result.Applications, ap)
	}

	return result, nil
}

//GetApplication gets a single application by either name or ID
func (p *ploioserver) GetApplication(c context.Context, ag *pp.ApplicationGet) (*pp.Application, error) {
	result := &pp.Application{}
	return result, nil
}

//CreateApplication creates a new application
func (p *ploioserver) CreateApplication(c context.Context, ac *pp.ApplicationCreate) (*pp.Application, error) {

	result := &pp.Application{
		Name:     ac.Name,
		Owner:    ac.Owner,
		Repo:     ac.Repo,
		Stateful: ac.Stateful,
	}

	dbobject := application{
		Name:     ac.Name,
		Owner:    ac.Owner,
		Repo:     ac.Repo,
		Stateful: ac.Stateful,
	}

	err := db.Save(&dbobject)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (p *ploioserver) ListClusters(c context.Context, cg *pp.ClusterGet) (*pp.ClusterList, error) {
	result := &pp.ClusterList{}
	return result, nil
}
func (p *ploioserver) ListPipelines(c context.Context, pg *pp.PipelineGet) (*pp.PipelineList, error) {
	result := &pp.PipelineList{}
	return result, nil
}
func (p *ploioserver) CreatePipeline(c context.Context, pc *pp.PipelineCreate) (*pp.Pipeline, error) {
	result := &pp.Pipeline{}
	return result, nil
}
func (p *ploioserver) CreateStep(c context.Context, s *pp.Step) (*pp.Step, error) {
	result := &pp.Step{}
	return result, nil
}
func (p *ploioserver) ListStep(c context.Context, sg *pp.StepGet) (*pp.StepList, error) {
	result := &pp.StepList{}
	return result, nil
}
