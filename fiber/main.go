package main

import (
	"fiber/configs"
	"fiber/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// setup gracefully shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	app := fiber.New()

	// run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(app)

	go func() {
		if err := app.Listen(":6000"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server Err: %s", err)
		}
	}()

	<-shutdown

	if err := app.Shutdown(); err != nil {
		log.Fatalf("Server Shutdown failed:%+v", err)
	}
	log.Print("Shutdown properly")

}
