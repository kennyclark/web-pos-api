package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kennyclark/web-pos-api/internal/db"
	"github.com/kennyclark/web-pos-api/internal/server"
	"github.com/kennyclark/web-pos-api/internal/utils"
)

func main() {
	log.Println("Starting server...")

	dbURL := utils.GetEnvOrDefault("DATABASE_URL", "")
	db.Connect(dbURL)

	app, err := server.New()
	if err != nil {
		log.Fatalf("Failed to create server: %v\n", err)
	}

	// listen from a different goroutine
	go func() {
		port := utils.GetEnvOrDefault("PORT", "3000")
		if err := app.Listen("0.0.0.0:" + port); err != nil {
			log.Panic(err)
		}
	}()

	// create a channel to block main goroutine
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	// cleanup
	log.Println("Running cleanup tasks...")
	// add cleanup tasks here
	db.Disconnect()

	// shutdown the server
	log.Println("Shutting down server...")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Failed to shutdown server: %v\n", err)
	}
	log.Println("Server stopped gracefully.")
}
