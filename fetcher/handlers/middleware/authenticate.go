package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gsstoykov/go-ethereum-fetcher/fetcher/handlers"
)

// AuthenticateMiddleware is a middleware function for JWT token authentication
func AuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		// Check if the Authorization header is in the correct format
		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		// Extract the token string from the header
		tokenString := headerParts[1]

		// Parse the token
		token, err := jwt.ParseWithClaims(tokenString, &handlers.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_STRING")), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse token"})
			c.Abort()
			return
		}

		// Check if the token is valid
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Get the claims from the token
		claims, ok := token.Claims.(*handlers.JWTClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to get token claims"})
			c.Abort()
			return
		}

		// Set the user information in the context
		c.Set("username", claims.Username)

		// Continue to the next handler
		c.Next()
	}
}
