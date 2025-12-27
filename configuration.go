package main

type Config struct{
	server 			map[string] string
	load_balancer	map[string] string
	backends		map[string] map[string]string
	health_check	map[string] string
}