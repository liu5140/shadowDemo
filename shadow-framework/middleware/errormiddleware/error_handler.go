package errormiddleware

import (
	"github.com/gin-gonic/gin"
)

//ErrorHandler is a gin middleware used to handle errors
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Errors) > 0 {
			c.Abort()
		} else {
			c.Next()
		}

	}
}
