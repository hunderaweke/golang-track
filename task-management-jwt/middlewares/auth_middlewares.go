package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type customClaims struct {
	Name   string
	Email  string
	UsedID string
	jwt.StandardClaims
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authString := c.GetHeader("Authorization")
		if authString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
			c.Abort()
			return
		}
		authParts := strings.Split(authString, " ")
		if len(authParts) != 2 || authParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
			c.Abort()
			return
		}
		claims := &customClaims{}
		jwtKey := []byte(os.Getenv("JWT_SECRET"))
		token, err := jwt.ParseWithClaims(authParts[1], claims, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
