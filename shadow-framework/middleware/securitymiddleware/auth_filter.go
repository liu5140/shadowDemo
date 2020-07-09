package securitymiddleware

import (
	"errors"

	"shadowDemo/shadow-framework/security"

	"github.com/gin-gonic/gin"
)

// Authorizer is a gin midlleware for authorizer request operation
func Authorizer() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, _ := c.Get(security.SHADOW_SECURITY_TOKEN)
		request := c.Request
		if token == nil {
			Log.Panic("authentication is nil")
		}

		authRequest := new(security.TRequestAuthenticationToken)

		if usernamePasswordAuthenticationToken, ok := token.(*security.TUsernamePasswordAuthenticationToken); ok && usernamePasswordAuthenticationToken.IsAuthenticated() {
			authRequest.SetDetails(security.CasbinAuthenticationRequestResolveInstance(security.CASBIN_AUTHENTICATION_REQUEST_RESOLVER).ObtainCasbinRequest(request.Host, request.RequestURI, request.Method, usernamePasswordAuthenticationToken.GetPrincipal()))
		} else if anonymousAuthenticationToken, ok := token.(*security.TAnonymousAuthenticationToken); ok {
			authRequest.SetDetails(security.CasbinAuthenticationRequestResolveInstance(security.CASBIN_AUTHENTICATION_REQUEST_RESOLVER).ObtainCasbinRequest(request.Host, request.RequestURI, request.Method, anonymousAuthenticationToken.GetPrincipal()))
		} else {
			err := security.NotPromissionError{
				Err: errors.New("Not promission"),
			}
			c.Error(err)
		}

		newAuthentication := security.AuthenticationManagerInstance(security.PROVIDER_MANAGER).Authenticate(authRequest)
		if !newAuthentication.IsAuthenticated() {
			err := security.NotPromissionError{
				Err: errors.New("Not promission"),
			}
			c.Error(err)
		}
		c.Next()
	}
}
