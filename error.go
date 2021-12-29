package ginerror

import (
	"github.com/gin-gonic/gin"
)

// RespondError is a gin middleware which write error string to response in text/plain
func RespondError(c *gin.Context) {
	// execute pending handlers
	c.Next()

	// do nothing if no error
	if len(c.Errors) == 0 {
		return
	}
	// write error string to response
	c.String(c.Writer.Status(), c.Errors.ByType(gin.ErrorTypePrivate).String())
}
