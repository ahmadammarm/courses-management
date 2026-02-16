package routes

import (
	"github.com/ahmadammarm/courses-management/backend/handlers/courses"
	"github.com/gin-gonic/gin"
)

func CoursesRoutesSetup(router *gin.RouterGroup) {
    router.GET("/", courses.GetAllCoursesHandler)
}