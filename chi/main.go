package main

import (
	"chi/configs"
	"chi/routes"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func main() {
	// setup gracefully shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	router := chi.NewRouter()
	router.Use(render.SetContentType(render.ContentTypeJSON))

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(router)

	srv := &http.Server{
		Addr:    ":6000",
		Handler: router,
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
