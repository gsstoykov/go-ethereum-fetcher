package middleware

import (
	"encoding/hex"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestRLPDecodeMiddleware tests the RLPDecodeMiddleware function.
func TestRLPDecodeMiddleware(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create a new Gin router
	r := gin.New()
	r.Use(RLPDecodeMiddleware())

	// Define a route to test middleware functionality
	r.GET("/test/:rlphex", func(c *gin.Context) {
		txHashes, exists := c.Get("transactionHashes")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction hashes not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"transactionHashes": txHashes})
	})

	t.Run("Success case", func(t *testing.T) {
		// Prepare sample data and encode it to RLP
		txHashes := []string{"0xc5f96bf1b54d3314425d2379bd77d7ed4e644f7c6e849a74832028b328d4d798"}
		rlpData, err := rlp.EncodeToBytes(txHashes)
		if err != nil {
			t.Fatalf("Failed to encode to RLP: %v", err)
		}
		rlpHex := hex.EncodeToString(rlpData)

		// Create a request and record response
		req, _ := http.NewRequest("GET", "/test/"+rlpHex, nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// Assert response
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `{"transactionHashes":["0xc5f96bf1b54d3314425d2379bd77d7ed4e644f7c6e849a74832028b328d4d798"]}`, rec.Body.String())
	})

	t.Run("Failure case - invalid RLP data", func(t *testing.T) {
		// Create a request with invalid RLP data
		req, _ := http.NewRequest("GET", "/test/invalidRLP", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		// Assert response
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}
