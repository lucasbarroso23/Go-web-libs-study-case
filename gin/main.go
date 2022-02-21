package main

import (
	"context"
	"gin/configs"
	"gin/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// setup gracefully shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	router := gin.Default()

	// run database
	configs.ConnectDB()

	// routes
	routes.UserRoute(router)

	srv := &http.Server{
		Addr:    ":6000",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server err: %s", err)
		}
	}()

	<-shutdown

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown falied: %+v", err)
	}
	log.Print("Shutdown properly")
}
