package routes

import (
	"github.com/ahmadammarm/courses-management/backend/handlers/auth"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {
    router.POST("/register", auth.InstructorRegisterHandler)
    router.POST("/login", auth.InstructorLoginHandler)
}