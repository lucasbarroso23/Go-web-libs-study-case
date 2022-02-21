package main

import (
	"context"
	"echo/configs"
	"echo/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	// setup gracefully shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	e := echo.New()

	// run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(e)

	srv := &http.Server{
		Addr:    ":6000",
		Handler: e,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server Err: %s", err)
		}
	}()

	<-shutdown

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed:%+v", err)
	}
	log.Print("Shutdown properly")
}
