package server

type ServerPool interface{
	GetServerPool()  	[]*backend
	GetValidPeer()	 	*backend
	AddPeer(*backend)
	GetServerPoolSize()	int
}
