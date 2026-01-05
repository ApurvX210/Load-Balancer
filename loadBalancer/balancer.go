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

func (lb *loadBalancer) Serve(rw http.ResponseWriter, req *http.Request) {
	targetServer := lb.ServerPool.GetValidPeer()
	if targetServer == nil {
		http.Error(rw, "Service unavailable - no healthy backends", http.StatusServiceUnavailable)
		return
	}
	
	targetServer.IncConnectionCount()
	defer targetServer.DecConnectionCount()
	targetServer.Serve(rw, req)
}

func NewLoadBalancer(sp server.ServerPool) LoadBalancer {
	return &loadBalancer{
		ServerPool: sp,
	}
}
