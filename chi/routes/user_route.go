package routes

import (
	"chi/controllers"

	"github.com/go-chi/chi/v5"
)

func UserRoute(router chi.Router) {
	router.Post("/user", controllers.CreateUser())
	router.Get("/user/{userId}", controllers.GetAUser())
	router.Put("/user/{userId}", controllers.EditUser())
	router.Delete("/user/{userId}", controllers.DeleteUser())
	router.Get("/users", controllers.GetAllUser())
}
