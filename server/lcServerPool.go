package server

import "sync"

type LcServerPool struct {
	Backends  []*Backend
	mu        sync.RWMutex
	// crnPeer *Backend
}

func (lc *LcServerPool) GetServerPool() []*Backend{
	return lc.Backends
}

func (lc *LcServerPool) GetValidPeer() *Backend{
	var targetBackend *Backend;
	for _,backend := range(lc.Backends){
		if targetBackend == nil && backend.IsAlive(){
			targetBackend = backend
		}else{
			if backend.IsAlive() && targetBackend.connections > backend.connections{
				targetBackend = backend
			}
		}
	}

	return targetBackend
}

func (lc *LcServerPool) AddPeer(b *Backend) {
	lc.Backends = append(lc.Backends, b)
}

func (lc *LcServerPool) GetServerPoolSize() int{
	return len(lc.Backends)
}