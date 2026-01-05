package server

import "sync"

type LcServerPool struct {
	Backends  []*Backend
	mu        sync.RWMutex
	// crnPeer *Backend
}

func (lc *LcServerPool) GetServerPool() []*Backend{
	lc.mu.RLock()
	defer lc.mu.RUnlock()
	return lc.Backends
}

func (lc *LcServerPool) GetValidPeer() *Backend{
	lc.mu.RLock()
	defer lc.mu.RUnlock()
	
	var targetBackend *Backend;
	for _,backend := range(lc.Backends){
		if targetBackend == nil && backend.IsAlive(){
			targetBackend = backend
		}else{
			if backend.IsAlive() && targetBackend.GetActiveConnection() > backend.GetActiveConnection(){
				targetBackend = backend
			}
		}
	}

	return targetBackend
}

func (lc *LcServerPool) AddPeer(b *Backend) {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	lc.Backends = append(lc.Backends, b)
}

func (lc *LcServerPool) GetServerPoolSize() int{
	lc.mu.RLock()
	defer lc.mu.RUnlock()
	return len(lc.Backends)
}