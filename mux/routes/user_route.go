package routes

import (
	"mux/controllers"

	"github.com/gorilla/mux"
)

func UserRoute(router *mux.Router) {
	router.HandleFunc("/user", controllers.CreateUser()).Methods("POST")
	router.HandleFunc("/user/{userId}", controllers.GetAUser()).Methods("GET")
}
