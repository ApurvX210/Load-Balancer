package server

import "sync"

type LcServerPool struct {
	backends  []*Backend
	mu        sync.RWMutex
	crnPeer *Backend
}

func (lc *LcServerPool) GetServerPool() []*Backend{
	return lc.backends
}

func (lc *LcServerPool) GetValidPeer() *Backend{
	targetBackend := nil
	for backend := range(lc.backends){
		if targetBackend == nil{
			targetBackend = backend
		}else{
			if targetBackend.connections > backend.connections{
				targetBackend = backend
			}
		}
	}

	return targetBackend
}

func (lc *LcServerPool) AddPeer(b *Backend) {
	lc.backends = append(lc.backends, b)
}

func (lc *LcServerPool) GetServerPoolSize() int{
	return len(lc.backends)
}