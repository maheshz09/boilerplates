package main

import (
	"context"
	"fmt"
	"log"
	"mongoapi/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("MongoDB API")
	fmt.Println("Server is getting started...")
	r := router.Router()
	srv := &http.Server{
		Addr:    ":4000",
		Handler: r,
	}
	// starting server in goroutine so we can listen for shudown signals
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Println("server started on", srv.Addr)

	// create channel to listen for interrupt or terminate singnals from the os
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// block until we receve a signal
	<-quit
	log.Println("Shutdioown signal received, shutting down gracefully......")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// attempt 1
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown, %v", err)
	}

	log.Println("server exiting")

}
