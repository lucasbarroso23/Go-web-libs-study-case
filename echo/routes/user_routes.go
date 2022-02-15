package routes

import (
	"echo/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	e.POST("/user", controllers.CreateUser)
	e.GET("/user/:userId", controllers.GetUser)
	e.PUT("/user/:userId", controllers.EditUser)
	e.DELETE("/user/:userId", controllers.DeteleUser)
	e.GET("/users", controllers.GetAllUsers)
}
