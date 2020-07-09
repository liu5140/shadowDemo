package middleware

import (
	"shadowDemo/zframework/security"

	"github.com/gin-gonic/gin"
)

func ProfileResolver(callback func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Errors) > 0 {
			return
		}
		token, _ := c.Get(security.SHADOW_SECURITY_TOKEN)
		if token == nil {
			Log.Panic("authentication is nil")
		}
		if secureToken, ok := token.(*security.TUsernamePasswordAuthenticationToken); ok && secureToken.IsAuthenticated() {
			callback(c)
		}
	}
}
