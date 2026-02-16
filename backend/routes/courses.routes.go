package routes

import (
	"github.com/ahmadammarm/courses-management/backend/handlers/courses"
	"github.com/ahmadammarm/courses-management/backend/middlewares"
	"github.com/gin-gonic/gin"
)

func CoursesRoutesSetup(router *gin.RouterGroup) {
	router.GET("/", middlewares.JWTAuthMiddleware(), courses.GetAllCoursesHandler)
	router.GET("/:instructor_id", middlewares.JWTAuthMiddleware(), courses.GetAllCoursesByInstructorHandler)
	router.POST("/", middlewares.JWTAuthMiddleware(), courses.CreateCourseHandler)
}