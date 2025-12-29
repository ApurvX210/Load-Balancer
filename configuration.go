package main

import (
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Server struct {
		Port         string `yaml:"port"`
		ReadTimeout  string `yaml:"readTimeout"`
		WriteTimeout string `yaml:"rriteTimeout"`
	} `yaml:"server"`
	LoadBalancer struct {
		Algorithm string `yaml:"algorithm"`
	} `yaml:"load_balancer"`
	Backends []struct {
		Url          string `yaml:"url"`
		InitialAlive bool   `yaml:"initial_alive"`
	} `yaml:"backends"`
	HealthCheck struct {
		Interval string `yaml:"interval"`
		Endpoint string `yaml:"endpoint"`
	} `yaml:"health_check"`
}

func parseYaml(path string) (*Config,error) {
	yamlFile, err := os.ReadFile(path)
	if err != nil{
		return nil,err
	}
	var config Config;
	err = yaml.Unmarshal(yamlFile,&config)
	if err != nil{
		return nil,err
	}

	return &config,nil
}