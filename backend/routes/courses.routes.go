package routes

import (
	"github.com/ahmadammarm/courses-management/backend/handlers/courses"
	"github.com/ahmadammarm/courses-management/backend/middlewares"
	"github.com/gin-gonic/gin"
)

func CoursesRoutesSetup(router *gin.RouterGroup) {
	// Handle both with and without trailing slash
	router.GET("", middlewares.JWTAuthMiddleware(), courses.GetAllCoursesHandler)
	router.POST("", middlewares.JWTAuthMiddleware(), courses.CreateCourseHandler)

	// Get courses by instructor ID (for dashboard)
	router.GET("/instructor/:instructor_id", middlewares.JWTAuthMiddleware(), courses.GetAllCoursesByInstructorHandler)

	// Single course operations - must be after /instructor/:id to avoid conflict
	router.GET("/:id", middlewares.JWTAuthMiddleware(), courses.GetCourseByIDHandler)
	router.PUT("/:id", middlewares.JWTAuthMiddleware(), courses.UpdateCourseHandler)
	router.DELETE("/:id", middlewares.JWTAuthMiddleware(), courses.DeleteCourseHandler)
}
