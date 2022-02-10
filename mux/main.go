package main

import (
	"log"
	"mux/configs"
	"mux/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(router)

	log.Fatal(http.ListenAndServe(":6000", router))
}
