package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

type Server interface {
	IsAlive()
	SetAlive()
	GetUrl() *url.URL
	GetActiveConnection() int
	Serve(http.ResponseWriter,*http.Request)
}

type backend struct{
	url 			*url.URL
	alive			bool
	mu      		sync.RWMutex
	connections		int
	reverseProxy 	*httputil.ReverseProxy
}

func (b *backend) IsAlive(){
	
}

func (b *backend) SetAlive(){
	b.alive = true
}

func (b *backend) GetUrl() *url.URL{
	return b.url
}

func (b *backend) GetActiveConnection() int{
	b.mu.RLock()
	defer b.mu.RLocker().Unlock()

	return b.connections
}

func (b *backend) Serve(rw http.ResponseWriter,req *http.Request){
	b.reverseProxy.ServeHTTP(rw,req)
}

