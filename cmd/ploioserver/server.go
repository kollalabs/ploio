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
	ID    int32  `storm:"id,increment"`
	Name  string `storm:"unique"`
	Owner string
	Repo  string
}

type env struct {
	ID            int32 `storm:"id,increment"`
	ApplicationID int32 `storm:"index"`
	Name          string
	Key           string
	Value         string
	ConfigMap     string
	Secret        string
}

type service struct {
	ID            int32 `storm:"id,increment"`
	ApplicationID int32 `storm:"index"`
	Type          string
}

type servicePort struct {
	ID         int32 `storm:"id,increment"`
	ServiceID  int32 `storm:"index"`
	Name       string
	Protocol   string
	Port       string
	TargetPort string
	NodePort   string
}

type cluster struct {
	ID   int32  `storm:"id,increment"`
	Name string `storm:"unique"`
	URL  string
}

type pipeline struct {
	ID            int32 `storm:"id,increment"`
	Name          string
	ApplicationID int32
	Order         int32
}

type step struct {
	ID         int32 `storm:"id,increment"`
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
	var err error

	result := &pp.Application{
		Name:  ac.Name,
		Owner: ac.Owner,
		Repo:  ac.Repo,
	}

	dbobject := application{
		Name:  ac.Name,
		Owner: ac.Owner,
		Repo:  ac.Repo,
	}

	tx, err := db.Begin(true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	err = tx.Save(&dbobject)
	if err != nil {
		return result, err
	}

	for _, e := range ac.Env {
		dbe := env{
			ApplicationID: dbobject.ID,
			Name:          e.Name,
			Key:           e.Key,
			Value:         e.Value,
			ConfigMap:     e.ConfigMap,
			Secret:        e.Secret,
		}
		err = tx.Save(&dbe)
		if err != nil {
			return nil, err
		}
	}

	for _, s := range ac.Services {
		dbs := service{
			ApplicationID: dbobject.ID,
			Type:          s.Type.String(),
		}
		err = tx.Save(&dbs)
		if err != nil {
			return nil, err
		}
		for _, sp := range s.Ports {
			dbsp := servicePort{
				ServiceID:  dbs.ID,
				Name:       sp.Name,
				Protocol:   sp.Protocol,
				Port:       sp.Port,
				TargetPort: sp.TargetPort,
				NodePort:   sp.NodePort,
			}
			err = tx.Save(&dbsp)
			if err != nil {
				return nil, err
			}
		}
	}
	err = tx.Commit()
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
func (p *ploioserver) CreateStage(c context.Context, s *pp.Stage) (*pp.Stage, error) {
	result := &pp.Stage{}
	return result, nil
}
func (p *ploioserver) ListStage(c context.Context, sg *pp.StageGet) (*pp.StageList, error) {
	result := &pp.StageList{}
	return result, nil
}
