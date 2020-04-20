package pkg

import (
	"fmt"
	"gopkg.in/yaml.v2"
	v1 "k8s.io/api/apps/v1"
	"os"
)


type Config struct {
	VCloud struct{
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		ORG      string `yaml:"org"`
		API      string `yaml:"api"`
		VDC      string `yaml:"vdc"`
	}
}


func parseConfig() (error){
	f, err := os.Open("config.yml")
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	defer f.Close()
	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
}


