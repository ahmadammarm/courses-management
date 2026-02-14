package main

import "github.com/gin-gonic/gin"

func main() {
	mainRouter := gin.Default()
	apiPrefix := mainRouter.Group("/api")


	// helath check endpoint
	apiPrefix.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"status": "OK",
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
			"error": "Endpoint not found",
		})
	});

	mainRouter.Run(":8080")
}
