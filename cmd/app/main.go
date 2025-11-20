package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"qa-service/internal/app"
)

func main() {
	ctx := context.Background()

	app, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed init app: %s", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.Run(); err != nil {
			log.Printf("failed to run app: %s", err)
		}
	}()

	log.Println("app started successfully")

	<-done
	log.Println("shutdown server")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.Stop(shutdownCtx); err != nil {
		log.Fatalf("Failed to stop app: %v", err)
	}

	log.Println("app is stopped")
}
