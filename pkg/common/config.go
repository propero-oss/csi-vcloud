package common

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)


type Config struct {
	VCloud struct{
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		ORG      string `yaml:"org"`
		API      string `yaml:"api"`
		VDC      string `yaml:"vdc"`
		Insecure bool `yaml:"insecure"`
	}
}



func ParseConfig() (*Config, error){
	f, err := os.Open("config.yml")
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	defer f.Close()
	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return &cfg, nil
}


