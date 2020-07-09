package securitymiddleware

import (
	"shadowDemo/shadow-framework/middleware/sessionmiddleware"
	"shadowDemo/shadow-framework/security"

	"github.com/gin-gonic/gin"
)

//SecurityTokenResolver 从session中解析出authentication token，用于认证当前的请求上下文
func SecurityTokenResolver() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从session中取出token, 放入request上下文中
		sess := sessionmiddleware.GetCurrentSession(c)
		authentication := sess.Get(security.SHADOW_SECURITY_TOKEN)

		Log.Debug("authentication", authentication)

		if authentication != nil {
			c.Set(security.SHADOW_SECURITY_TOKEN, authentication)
		}

		c.Next()

		//请求返回前将token再写回session中
		Log.Debug("rewrite security token to session")
		if auth, exist := c.Get(security.SHADOW_SECURITY_TOKEN); exist {
			Log.Debugf("authentication2 : %+v", auth)
			Log.Debugf("authentication2 detail: %+v", auth.(security.IAuthentication).GetDetails())
			sess.Set(security.SHADOW_SECURITY_TOKEN, auth)
		}

	}
}
