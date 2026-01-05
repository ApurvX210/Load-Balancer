package main

import (
	loadbalancer "LoadBalancer/loadBalancer"
	"LoadBalancer/server"
	"context"
	"fmt"
	"net/http"
	"time"
)

type Application struct {
	config *Config
	lb     loadbalancer.LoadBalancer
}

func (app *Application) requestHandler(rq http.ResponseWriter, req *http.Request) {
	app.lb.Serve(rq, req)
}

func NewApplication() (*Application, error) {
	config, err := parseYaml("config.yml")
	if err != nil {
		fmt.Printf("Error occurred while parsing config file: %v\n", err)
		return nil, err
	}
	backendList := config.Backends
	backends := []*server.Backend{}
	for _, backendInfo := range backendList {
		backend, err := server.NewBackend(backendInfo.Url, backendInfo.InitialAlive)
		if err != nil {
			fmt.Printf("Error occurred while registering the server %s: %v\n", backendInfo.Url, err)
			continue
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
		return nil, fmt.Errorf("unsupported load balancer algorithm: %s", config.LoadBalancer.Algorithm)
	}

	lb := loadbalancer.NewLoadBalancer(serverPool)

	// Parse health check interval
	healthInterval, err := time.ParseDuration(config.HealthCheck.Interval)
	if err != nil {
		return nil, fmt.Errorf("invalid health check interval: %w", err)
	}

	// Start health check in background
	ctx := context.Background()
	go server.HealthCheck(ctx, serverPool, healthInterval, config.HealthCheck.Endpoint)

	return &Application{
		config: config,
		lb:     lb,
	}, nil
}
