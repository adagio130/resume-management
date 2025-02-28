// internal/middleware/error_handler.go
package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	customerror "resume/internal/errors"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if c.Writer.Written() {
			return
		}

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			var customErr *customerror.CustomError
			if errors.As(err, &customErr) {
				c.JSON(customErr.StatusCode, map[string]string{"error": customErr.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			}
			c.Abort()
		}
	}
}
