package server

import (
	"context"
	"net/url"
	"time"
	"fmt"
)

type ServerPool interface {
	GetServerPool() []*Backend
	GetValidPeer() *Backend
	AddPeer(*Backend)
	GetServerPoolSize() int
}

func checkServer(ctx context.Context, c chan bool, url *url.URL){

}

func HeathCheck(ctx context.Context, s ServerPool) {
	activeChan := make(chan bool, 1)

	for _, backend := range s.GetServerPool() {
		requestCtx,stop := context.WithTimeout(ctx, 10*time.Second)
		defer stop()
		go checkServer(requestCtx,activeChan,backend.url)

		select{
		case active := <- activeChan:
			backend.SetAlive(active)
		case <- ctx.Done():
			fmt.Println("Gracefully shutting down health check")
			return
		}
	}
}
