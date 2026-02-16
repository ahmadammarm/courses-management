package middlewares

import (
	"net/http"
	"strings"

	"github.com/ahmadammarm/courses-management/backend/config"
	"github.com/ahmadammarm/courses-management/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")

		if tokenString == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token is required",
			})
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")


		jwtKey := []byte(config.GetEnv("JWT_SECRET"))
		claims := &utils.CustomClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
			})
			return
		}
		
        
		context.Set("instructor_id", claims.InstructorID)
		context.Set("username", claims.Username)

		context.Next()
	}
}