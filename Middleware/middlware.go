package Middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"time"
)

type CustomClaims struct {
	jwt.StandardClaims
	Role string `json:"role"`
}

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the request header
		tokenString := c.GetHeader("Authorization")

		// Parse and validate the token with custom claims
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			// Verify the token signing method
			if token.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("invalid token signing method")
			}
			return []byte(os.Getenv("JWT_PRIVATE_KEY")), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// Extract the claims from the token
		claims, ok := token.Claims.(*CustomClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		 // Check if the token is expired
		if time.Unix(claims.ExpiresAt, 0).Before(time.Now()) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
			return
		}

		// Get the role from the claims
		role := claims.Role

		// Check if the role has sufficient permissions
		if role != requiredRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient role permissions"})
			return
		}

		// Pass the role to the next handler
		c.Set("role", role)
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return RoleMiddleware("Admin")
}
