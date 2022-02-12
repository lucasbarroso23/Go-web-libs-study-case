package main

import (
	"gin/configs"
	"gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// run database
	configs.ConnectDB()

	// routes
	routes.UserRoute(router)

	router.Run("localhost:6000")
}
