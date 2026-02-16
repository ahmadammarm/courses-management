package main

import (
	"github.com/ahmadammarm/courses-management/backend/config"
	"github.com/ahmadammarm/courses-management/backend/db"
	"github.com/ahmadammarm/courses-management/backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()

    db.DatabaseConnect()

	appPort := config.GetEnv("APP_PORT")

	mainRouter := gin.Default()
	apiPrefix := mainRouter.Group("/api")

	// helath check endpoint
	apiPrefix.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"status": "OK bang",
		})
	})

	// root endpoint
	apiPrefix.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Welcome to the Courses Management API",
		})
	})

	// not found handler
	mainRouter.NoRoute(func(context *gin.Context) {
		context.JSON(404, gin.H{
			"error": "Endpoint not found woii",
		})
	})

    // auth routes
    authGroup := apiPrefix.Group("/auth")
    routes.AuthRoutes(authGroup)


	mainRouter.Run(":" + appPort)
}
