package main

type ServerPool interface{
	GetServerPool()  	[]*backend
	GetValidPeer()	 	*backend
	AddPeer()
	GetServerPoolSize()	int
}
