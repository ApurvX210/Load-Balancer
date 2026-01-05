package server

import "sync"

type RrServerPool struct {
	Backends  []*Backend
	mu        sync.RWMutex
	crnIndex  int
}

func (rr *RrServerPool) GetServerPool() []*Backend{
	rr.mu.RLock()
	defer rr.mu.RUnlock()
	return rr.Backends
}

func (rr *RrServerPool) Rotate() *Backend{
	rr.mu.Lock()
	defer rr.mu.Unlock()

	poolSize := len(rr.Backends)
	if poolSize == 0 {
		return nil
	}

	rr.crnIndex = (rr.crnIndex+1) % poolSize

	return rr.Backends[rr.crnIndex]
}

func (rr *RrServerPool) GetValidPeer() *Backend{
	rr.mu.RLock()
	crnLen := len(rr.Backends)
	rr.mu.RUnlock()

	if crnLen == 0 {
		return nil
	}

	for idx:=0;idx<crnLen;idx++{
		targetServer := rr.Rotate()
		if targetServer != nil && targetServer.IsAlive(){
			return targetServer
		}
	}
	return nil
}

func (rr *RrServerPool) AddPeer(b *Backend) {
	rr.mu.Lock()
	defer rr.mu.Unlock()
	rr.Backends = append(rr.Backends, b)
}

func (rr *RrServerPool) GetServerPoolSize() int{
	rr.mu.RLock()
	defer rr.mu.RUnlock()
	return len(rr.Backends)
}