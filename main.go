package main

import (
	"fmt"
	"net/http"
)

func main() {
	app,err := NewApplication()
	if err != nil{
		fmt.Println("Error occured while initializing the application,err")
	}

	http.HandleFunc("/",app.requestHandler)
	http.ListenAndServe(app.config.Server.Port,nil)
}
