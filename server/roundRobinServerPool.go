package server

import "sync"

type RrServerPool struct {
	Backends  []*Backend
	mu        sync.RWMutex
	crnIndex  int
}

func (lc *RrServerPool) GetServerPool() []*Backend{
	return lc.Backends
}

func (lc *RrServerPool) Rotate() *Backend{
	lc.mu.Lock()
	defer lc.mu.Unlock()

	lc.crnIndex = (lc.crnIndex+1) % lc.GetServerPoolSize()

	return lc.Backends[lc.crnIndex]
}

func (lc *RrServerPool) GetValidPeer() *Backend{
	crnLen := lc.GetServerPoolSize()

	for idx:=0;idx<crnLen;idx++{
		targetServer := lc.Rotate()
		if targetServer.IsAlive(){
			return targetServer
		}
	}
	return nil
}

func (lc *RrServerPool) AddPeer(b *Backend) {
	lc.Backends = append(lc.Backends, b)
}

func (lc *RrServerPool) GetServerPoolSize() int{
	return len(lc.Backends)
}