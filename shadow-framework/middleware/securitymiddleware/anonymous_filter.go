package securitymiddleware

import (
	"shadowDemo/shadow-framework/middleware/sessionmiddleware"
	"shadowDemo/shadow-framework/security"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// AnonymousFilter is a middleware for anonymous user request
func AnonymousFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		//检查request上下文中是否已经存在一个已经认证的security token，如果不存在则创建一个未认证token放入request上下文中，
		authentication, exist := c.Get(security.SHADOW_SECURITY_TOKEN)
		if !exist {
			anonymousToken := security.NewAnonymousAuthenticationToken()
			anonymousToken.SetDetails(&security.TWebAuthenticationDetails{
				RemoteAddress: c.ClientIP(),
				SessionID:     sessionmiddleware.GetCurrentSession(c).SessionID(),
				RequestURI:    c.Request.RequestURI,
			})
			Log.Debug("set anonymous token")
			Log.Debug(anonymousToken)
			c.Set(security.SHADOW_SECURITY_TOKEN, anonymousToken)
		}

		if auth, ok := authentication.(security.IAuthentication); ok {
			Log.WithFields(logrus.Fields{
				"Authenticated": auth.IsAuthenticated(),
				"Details":       auth.GetDetails(),
				"Principal":     auth.GetPrincipal(),
			}).Debug("AnonymousFilter")
		}
		c.Next()
	}
}
