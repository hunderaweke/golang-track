package infrastructure

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

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
		claims, valid := ValidateToken(authParts[1])
		if !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization is needed"})
			c.Abort()
			return
		}
		userClaims := claims.(*customClaims)
		if !userClaims.IsAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "admin access needed"})
			c.Abort()
			return
		}
		c.Next()
	}
}
