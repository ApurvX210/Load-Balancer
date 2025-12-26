package server

import "sync"

type RrServerPool struct {
	backends  []*Backend
	mu        sync.RWMutex
	crnIndex  int
}

func (lc *RrServerPool) GetServerPool() []*Backend{
	return lc.backends
}

func (lc *RrServerPool) GetValidPeer() *Backend{
	
}

func (lc *RrServerPool) AddPeer(b *Backend) {
	lc.backends = append(lc.backends, b)
}

func (lc *RrServerPool) GetServerPoolSize() int{
	return len(lc.backends)
}