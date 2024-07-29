package middleware

import (
	"encoding/hex"
	"net/http"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/gin-gonic/gin"
)

// RLPDecodeMiddleware decodes RLP data from the URL parameter and adds transaction hashes to the context.
func RLPDecodeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		rlpHex := c.Param("rlphex")
		rlpData, err := hex.DecodeString(rlpHex)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		var txHashes []string
		err = rlp.DecodeBytes(rlpData, &txHashes)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.Set("transactionHashes", txHashes)
		c.Next()
	}
}
