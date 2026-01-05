package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	app, err := NewApplication()
	if err != nil {
		fmt.Printf("Error occurred while initializing the application: %v\n", err)
		return
	}

	// Parse timeouts from config
	readTimeout, err := time.ParseDuration(app.config.Server.ReadTimeout)
	if err != nil {
		fmt.Printf("Warning: invalid read timeout, using default: %v\n", err)
		readTimeout = 30 * time.Second
	}

	writeTimeout, err := time.ParseDuration(app.config.Server.WriteTimeout)
	if err != nil {
		fmt.Printf("Warning: invalid write timeout, using default: %v\n", err)
		writeTimeout = 30 * time.Second
	}

	http.HandleFunc("/", app.requestHandler)

	server := &http.Server{
		Addr:         app.config.Server.Port,
		Handler:      nil,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		fmt.Printf("Load balancer starting on %s\n", app.config.Server.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting server: %v\n", err)
			quit <- os.Interrupt
		}
	}()

	<-quit
	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown: %v\n", err)
	} else {
		fmt.Println("Server exited gracefully")
	}
}
