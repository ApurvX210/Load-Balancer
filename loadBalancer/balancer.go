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
	targetServer.IncConnectionCount()
	targetServer.Serve(rw, req)
	targetServer.DecConnectionCount()
}

func NewLoadBalancer(sp server.ServerPool) LoadBalancer {
	return &loadBalancer{
		ServerPool: sp,
	}
}
