package server

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	PORT = 3000
)

func HandleServer(router *chi.Mux) {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", PORT),
		Handler:      router,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	// create a channel to receive signal
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

	log.Println("Server Logics starting")
	// start the server in a separate go routine.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Error:%s\n", err)
		}
	}()
	log.Println("Server started on port ", PORT)
	// wait for a signal to shut down the server
	sig := <-stopChan
	log.Printf("signal recieved: %v\n", sig)

	// create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	log.Println("Server shutting down gracefully")
	// shutdown the server gracefully
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v\n", err)
	}
}
