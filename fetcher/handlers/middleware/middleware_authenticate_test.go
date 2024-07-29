package middleware

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestAuthenticateMiddleware tests the AuthenticateMiddleware function.
func TestAuthenticateMiddleware(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create a new Gin router
	r := gin.New()
	r.Use(AuthenticateMiddleware())

	// Define a route to test middleware functionality
	r.GET("/protected", func(c *gin.Context) {
		username, _ := c.Get("username")
		c.JSON(http.StatusOK, gin.H{"username": username})
	})

	// Set the JWT secret key for testing
	os.Setenv("JWT_STRING", "secret")

	// Generate a valid token for testing
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": "testuser",
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		t.Fatalf("Failed to sign token: %v", err)
	}

	t.Run("Success case - valid token", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/protected", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// Assert response
		assert.Equal(t, http.StatusOK, rec.Code)
		expectedResponse := `{"username":"testuser"}`
		assert.JSONEq(t, expectedResponse, rec.Body.String())
	})

	t.Run("Failure case - missing token", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/protected", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// Assert response
		assert.Equal(t, http.StatusOK, rec.Code)
		expectedResponse := `{"username":null}`
		assert.JSONEq(t, expectedResponse, rec.Body.String())
	})

	t.Run("Failure case - invalid token", func(t *testing.T) {
		// Provide an invalid token
		invalidToken := "invalidTokenString"

		req, _ := http.NewRequest("GET", "/protected", nil)
		req.Header.Set("Authorization", "Bearer "+invalidToken)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// Assert response
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		expectedResponse := `{"error":"Failed to parse token"}`
		assert.JSONEq(t, expectedResponse, rec.Body.String())
	})
}
