package application

import (
	"log"

	"gopkg.in/yaml.v2"
)

type Version struct {
	Config *Application
	Tag    string
	Type   string
}

// Application is a struct representing a parsed ploiofile
type Application struct {
	Schema    string `yaml:"schema"`
	Name      string `yaml:"name"`
	Owner     string `yaml:"owner"`
	Repo      string `yaml:"repo"`
	Namespace string `yaml:"namespace"`
	Ports     []struct {
		Name     string `yaml:"name"`
		Protocol string `yaml:"protocol"`
		Number   int    `yaml:"number"`
	} `yaml:"ports"`
	Resources struct {
		RAM         string `yaml:"ram"`
		CPU         string `yaml:"cpu"`
		MinReplicas int    `yaml:"minReplicas"`
		MaxReplicas int    `yaml:"maxReplicas"`
	} `yaml:"resources"`
	Env []struct {
		Name          string `yaml:"name"`
		Value         int    `yaml:"value,omitempty"`
		ConfigMapName string `yaml:"configMapName,omitempty"`
		Key           string `yaml:"key,omitempty"`
	} `yaml:"env"`
	Expose []struct {
		Name        string   `yaml:"name"`
		Type        string   `yaml:"type"`
		Domain      string   `yaml:"domain,omitempty"`
		Paths       []string `yaml:"paths,omitempty"`
		Destination string   `yaml:"destination"`
		Port        int      `yaml:"port,omitempty"`
	} `yaml:"expose"`
}

func Unmarshal(config []byte) (*Application, error) {
	c := &Application{}
	err := yaml.Unmarshal(config, c)
	if err != nil {
		log.Fatalf("error: %v", err)
		return c, err
	}
	return c, nil
}
