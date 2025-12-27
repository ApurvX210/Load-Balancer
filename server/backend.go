package server

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

type Server interface {
	IsAlive()
	SetAlive(alive bool)
	GetUrl() *url.URL
	GetActiveConnection() int
	IncConnectionCount()
	DecConnectionCount()
	Serve(http.ResponseWriter,*http.Request)
}

type Backend struct{
	url 			*url.URL
	Alive			bool
	mu      		sync.RWMutex
	connections		int
	reverseProxy 	*httputil.ReverseProxy
}

func (b *Backend) IsAlive() bool{
	return b.Alive
}

func (b *Backend) SetAlive(alive bool){
	b.Alive = alive
}

func (b *Backend) GetUrl() *url.URL{
	return b.url
}

func (b *Backend) GetActiveConnection() int{
	b.mu.RLock()
	defer b.mu.RLocker().Unlock()

	return b.connections
}

func (b *Backend) IncConnectionCount(){
	b.mu.Lock()
	defer b.mu.Unlock()
	b.connections += 1
}

func (b *Backend) DecConnectionCount(){
	b.mu.Lock()
	defer b.mu.Unlock()
	b.connections -= 1
}

func (b *Backend) Serve(rw http.ResponseWriter,req *http.Request){
	b.reverseProxy.ServeHTTP(rw,req)
}

func NewBackend(rawUrl string) (*Backend,error){
	url,err := url.Parse(rawUrl)

	if err != nil{
		return nil,err
	}

	reverseProxy := httputil.NewSingleHostReverseProxy(url)

	return &Backend{
		url: url,
		reverseProxy: reverseProxy,
	},nil
}