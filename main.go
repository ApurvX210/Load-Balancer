package main

import (
	"LoadBalancer/loadBalancer"
	"LoadBalancer/server"
	"fmt"
	"net/http"
)

func requestHandler(rq http.ResponseWriter, req *http.Request) {

}



func main() {
	sp := &server.RrServerPool{}
	lb := loadbalancer.NewLoadBalancer(sp)
	fmt.Println(lb)
}
