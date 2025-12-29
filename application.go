package main

import (
	loadbalancer "LoadBalancer/loadBalancer"
	"LoadBalancer/server"
	"fmt"
	"net/http"
)

type Application struct {
	config *Config
	lb     loadbalancer.LoadBalancer
}

func (app *Application) requestHandler(rq http.ResponseWriter, req *http.Request) {
	app.lb.Serve(rq, req)
}

func NewApplication() (*Application,error){
	config, err := parseYaml("config.yml")
	if err != nil {
		fmt.Println("Error occured while parsing config file", err)
		return nil,err
	}
	backendList := config.Backends
	backends := []*server.Backend{}
	for _, backendInfo := range backendList {
		backend, err := server.NewBackend(backendInfo.Url, backendInfo.InitialAlive)
		if err != nil {
			fmt.Println("Error occurred while Registering the server", backendInfo.Url, err)
		}
		backends = append(backends, backend)
	}

	var serverPool server.ServerPool
	switch config.LoadBalancer.Algorithm {
	case "round_robin":
		serverPool = &server.RrServerPool{
			Backends: backends,
		}
	case "least_connection":
		serverPool = &server.LcServerPool{
			Backends: backends,
		}
	default:
		return nil,fmt.Errorf("Unsupported load balancer algorithm:", config.LoadBalancer.Algorithm)
	}

	lb := loadbalancer.NewLoadBalancer(serverPool)
	return &Application{
		config: config,
		lb: lb,
	},nil
}