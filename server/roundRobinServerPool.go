package server

import "sync"

type RrServerPool struct {
	backends  []*Backend
	mu        sync.RWMutex
	crnPeer *Backend
}

func (lc *RrServerPool) GetServerPool() []*Backend{
	return lc.backends
}

func (lc *RrServerPool) GetValidPeer() *Backend{
	
}

func (lc *LcServerPool) AddPeer(b *Backend) {
	lc.backends = append(lc.backends, b)
}

func (lc *LcServerPool) GetServerPoolSize() int{
	return len(lc.backends)
}