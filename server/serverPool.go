package server

import (
	"context"
	"net/http"
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

var healthCheckClient = &http.Client{
	Timeout: 5 * time.Second,
}

func checkServer(ctx context.Context, c chan bool, backendUrl *url.URL, endpoint string) {
	healthURL := backendUrl.String() + endpoint
	req, err := http.NewRequestWithContext(ctx, "GET", healthURL, nil)
	if err != nil {
		c <- false
		return
	}

	resp, err := healthCheckClient.Do(req)
	if err != nil {
		c <- false
		return
	}
	defer resp.Body.Close()

	c <- resp.StatusCode == http.StatusOK
}

func HealthCheck(ctx context.Context, s ServerPool, interval time.Duration, endpoint string) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// Run initial health check
	checkAllServers(ctx, s, endpoint)

	for {
		select {
		case <-ticker.C:
			checkAllServers(ctx, s, endpoint)
		case <-ctx.Done():
			fmt.Println("Gracefully shutting down health check")
			return
		}
	}
}

func checkAllServers(ctx context.Context, s ServerPool, endpoint string) {
	backends := s.GetServerPool()
	for _, backend := range backends {
		requestCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		activeChan := make(chan bool, 1)
		
		go checkServer(requestCtx, activeChan, backend.GetUrl(), endpoint)

		select {
		case active := <-activeChan:
			backend.SetAlive(active)
		case <-requestCtx.Done():
			backend.SetAlive(false)
		}
		cancel()
	}
}
