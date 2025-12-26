package server

type ServerPool interface{
	GetServerPool()  	[]*Backend
	GetValidPeer()	 	*Backend
	AddPeer(*Backend)
	GetServerPoolSize()	int
}
