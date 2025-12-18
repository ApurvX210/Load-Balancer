package server

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

type Backend struct{
	url 			*url.URL
	alive			bool
	mu      		sync.RWMutex
	connections		int
	reverseProxy 	*httputil.ReverseProxy
}

func (b *Backend) IsAlive(){
	
}

func (b *Backend) SetAlive(){
	b.alive = true
}

func (b *Backend) GetUrl() *url.URL{
	return b.url
}

func (b *Backend) GetActiveConnection() int{
	b.mu.RLock()
	defer b.mu.RLocker().Unlock()

	return b.connections
}

func (b *Backend) Serve(rw http.ResponseWriter,req *http.Request){
	b.reverseProxy.ServeHTTP(rw,req)
}

