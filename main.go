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
	config,err := parseYaml("config.yml")
	if err != nil{
		fmt.Println("Error occured while parsing config file",err)
		return
	}
	backendList := config.Backends
	backends := []*server.Backend{}
	for _,backendInfo := range backendList{
		backend,err := server.NewBackend(backendInfo.Url,backendInfo.InitialAlive)
		if err != nil{
			fmt.Println("Error occurred while Registering the server",backendInfo.Url,err)
		}
		backends = append(backends, backend)
	}

	sp := &server.RrServerPool{}
	lb := loadbalancer.NewLoadBalancer(sp)
	fmt.Println(lb)
}
