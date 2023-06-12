package main

import (
	"context"
	"github.com/Imomali1/platform/internal/platform"
	"github.com/Imomali1/platform/internal/server"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	srv := server.NewServer(platform.NewRouter())
	go func() {
		if err := srv.Run(); err != nil {
			log.Fatal("failed to run http server")
		}
	}()
	log.Println("Server started...")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("failed to stop server")
	}

}
