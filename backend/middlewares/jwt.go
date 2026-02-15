package middlewares

import (
	"net/http"
	"strings"

	"github.com/ahmadammarm/courses-management/backend/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.GetEnv("JWT_SECRET"))

func JWTAuthMiddleware() gin.HandlerFunc {

	return func(context *gin.Context) {

		tokenString := context.GetHeader("Authorization")

		if tokenString == "" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token is required",
			})
			context.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims := &jwt.RegisteredClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			context.Abort()
			return
		}

		context.Set("username", claims.Subject)

		context.Next()
	}
}
