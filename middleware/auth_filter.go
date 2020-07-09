package middleware

import (
	"errors"
	"fmt"

	"shadowDemo/shadow-framework/security"

	"github.com/gin-gonic/gin"
)

// Authorizer is a gin midlleware for authorizer request operation
func Authorizer(site string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Errors) > 0 {
			return
		}

		token, _ := c.Get(security.SHADOW_SECURITY_TOKEN)
		if token == nil {
			Log.Panic("authentication is nil")
			return
		}

		authRequest := new(security.TRequestAuthenticationToken)
		if usernamePasswordAuthenticationToken, ok := token.(*security.TUsernamePasswordAuthenticationToken); ok && usernamePasswordAuthenticationToken.IsAuthenticated() {
			//获取profile
			profile := c.MustGet("profile").(Profile)
			authRequest.SetDetails(security.CasbinAuthenticationRequestResolveInstance(security.CASBIN_AUTHENTICATION_REQUEST_RESOLVER).ObtainCasbinRequest(c.Request.Host, c.Request.URL.Path, c.Request.Method, fmt.Sprintf("%d", profile.ID)))
		} else if anonymousAuthenticationToken, ok := token.(*security.TAnonymousAuthenticationToken); ok {
			authRequest.SetDetails(security.CasbinAuthenticationRequestResolveInstance(security.CASBIN_AUTHENTICATION_REQUEST_RESOLVER).ObtainCasbinRequest(c.Request.Host, c.Request.URL.Path, c.Request.Method, anonymousAuthenticationToken.GetPrincipal()))
		} else {
			err := security.NotPromissionError{
				Err: errors.New("Not promission"),
			}
			c.Error(err)
			return
		}

		newAuthentication := security.AuthenticationManagerInstance(security.PROVIDER_MANAGER).Authenticate(authRequest)
		if !newAuthentication.IsAuthenticated() {
			err := security.NotPromissionError{
				Err: errors.New("Not promission"),
			}
			c.Error(err)
			return
		}
		c.Next()
	}
}
