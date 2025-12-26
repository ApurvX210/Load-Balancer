package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			fmt.Println("still running")
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("main exiting")
}