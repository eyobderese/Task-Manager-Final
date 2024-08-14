package infrastructure

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func stringInSlice(str interface{}, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func AuthMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// JWT validation logic
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		claims, err := NewJwtService().TotokenParser(authParts[1])

		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid JWT claims"})
			c.Abort()
			return
		}

		if len(roles) != 0 && !stringInSlice(claims["role"], roles) {

			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		fmt.Println("clamis", claims)

		c.Next()
	}
}
