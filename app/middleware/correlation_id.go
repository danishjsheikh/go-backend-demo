package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AddCorrelationID is a Gin middleware that ensures a valid correlation ID is present
func AddCorrelationID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		correlationID := ctx.GetHeader("x-correlationid")

		if correlationID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "CorrelationId is mandatory"})
			ctx.Abort()
			return
		}

		if _, err := uuid.Parse(correlationID); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "CorrelationId is not a valid GUID"})
			ctx.Abort()
			return
		}

		// Store correlation ID in the context
		ctx.Set("correlationId", correlationID)
		ctx.Next()
	}
}
