package main

import (
	"errors"
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

	Pipelines []*pipeline
	Env       []*env
	Services  []*service
}

type env struct {
	ID        int32 `storm:"id,increment"`
	Name      string
	Key       string
	Value     string
	ConfigMap string
	Secret    string
}

type service struct {
	ID    int32 `storm:"id,increment"`
	Name  string
	Type  string
	Ports []*servicePort
}

type servicePort struct {
	ID         int32 `storm:"id,increment"`
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
	Stages        []stage
}

type stage struct {
	ID               int32 `storm:"id,increment"`
	Name             string
	Type             string
	Order            int32
	Strategy         string
	Cluster          string
	Namespace        string
	Memory           string
	Proc             string
	Replicas         int32
	MaxUnavailable   int32
	MaxSurge         int32
	NotifyEmail      bool
	NotifySlack      bool
	SmokeHeaderKey   string
	SmokeHeaderVal   string
	CanaryPercentage int32
	WaitDuration     int64
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
	var dbapp application
	result := &pp.Application{}
	err := db.One("Name", ag.Name, &dbapp)
	if err != nil {
		return result, err
	}
	result.ID = dbapp.ID
	result.Name = dbapp.Name
	result.Owner = dbapp.Owner
	result.Repo = dbapp.Repo

	for _, dbe := range dbapp.Env {
		result.Env = append(result.Env, &pp.Env{
			Name:      dbe.Name,
			ID:        dbe.ID,
			Key:       dbe.Key,
			Value:     dbe.Value,
			ConfigMap: dbe.ConfigMap,
			Secret:    dbe.Secret,
		})
	}

	for _, dbs := range dbapp.Services {
		t, ok := pp.ServiceType_value[dbs.Type]
		if !ok {
			return result, errors.New("Service type " + dbs.Type + " from database isn't valid")
		}
		s := &pp.Service{
			Name: dbs.Name,
			Type: pp.ServiceType(t),
		}

		for _, dbsp := range dbs.Ports {
			p := &pp.Port{
				Name:       dbsp.Name,
				Port:       dbsp.Port,
				TargetPort: dbsp.TargetPort,
				Protocol:   dbsp.Protocol,
				NodePort:   dbsp.NodePort,
			}
			s.Ports = append(s.Ports, p)
		}

		result.Services = append(result.Services, s)
	}
	var dbpipelines []pipeline
	err = db.Find("ApplicationID", dbapp.ID, &dbpipelines)
	if err != nil {
		return result, err
	}
	for _, dbp := range dbpipelines {
		protoPipeline := &pp.Pipeline{}
		protoPipeline.ID = dbp.ID
		protoPipeline.Name = dbp.Name
		for _, dbstage := range dbp.Stages {
			castStage := pp.Stage(dbstage)
			protoPipeline.Stages = append(protoPipeline.Stages, &castStage)
		}

		result.Pipelines = append(result.Pipelines, protoPipeline)
	}
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

	if err != nil {
		return result, err
	}

	for _, e := range ac.Env {
		dbe := &env{
			Name:      e.Name,
			Key:       e.Key,
			Value:     e.Value,
			ConfigMap: e.ConfigMap,
			Secret:    e.Secret,
		}
		dbobject.Env = append(dbobject.Env, dbe)

	}

	for _, s := range ac.Services {
		dbs := &service{
			Name: s.Name,
			Type: s.Type.String(),
		}

		for _, sp := range s.Ports {
			dbsp := &servicePort{
				Name:       sp.Name,
				Protocol:   sp.Protocol,
				Port:       sp.Port,
				TargetPort: sp.TargetPort,
				NodePort:   sp.NodePort,
			}
			dbs.Ports = append(dbs.Ports, dbsp)
		}
		dbobject.Services = append(dbobject.Services, dbs)
	}

	err = db.Save(&dbobject)
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
	var err error
	result := &pp.Pipeline{}
	dbp := &pipeline{}
	dbapplication := &application{}
	dbp.Name = pc.Name
	err = db.One("Name", pc.Application, dbapplication)
	if err != nil {
		return result, errors.New("Could not load application with name `" + pc.Application + "`")
	}
	dbp.ApplicationID = dbapplication.ID
	for _, s := range pc.Stages {
		dbp.Stages = append(dbp.Stages, stage(*s))
	}
	err = db.Save(dbp)
	if err != nil {
		return result, err
	}
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
