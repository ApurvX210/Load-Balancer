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

func (lc *RrServerPool) Rotate() *Backend{
	lc.mu.Lock()
	defer lc.mu.Unlock()

	lc.crnIndex = (lc.crnIndex+1) % lc.GetServerPoolSize()

	return lc.backends[lc.crnIndex]
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
	lc.backends = append(lc.backends, b)
}

func (lc *RrServerPool) GetServerPoolSize() int{
	return len(lc.backends)
}