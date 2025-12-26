package loadbalancer

import (
	"LoadBalancer/server"
	"net/http"
)

type LoadBalancer interface {
	Serve(http.ResponseWriter, *http.Request)
}

type loadBalancer struct {
	ServerPool server.ServerPool
}
