package security

import (
	"github.com/gin-gonic/gin"
)

//IUsernamePasswordResolver resolver username and password from gin context
type IUsernamePasswordResolver interface {
	ObtainUsernamePassword(c *gin.Context) (string, string)
}

//FUsernamePasswordResolverFactory FUsernamePasswordResolverFactory
type FUsernamePasswordResolverFactory func() IUsernamePasswordResolver

var usernamePasswordResolverFactories = make(map[string]FUsernamePasswordResolverFactory)

//RegisterUsernamePasswordResolver RegisterUsernamePasswordResolver
func RegisterUsernamePasswordResolver(name string, factory FUsernamePasswordResolverFactory) {
	usernamePasswordResolverFactories[name] = factory
}

//UsernamePasswordResolverInstance UsernamePasswordResolverInstance
func UsernamePasswordResolverInstance(name string) IUsernamePasswordResolver {
	factory := usernamePasswordResolverFactories[name]
	return factory()
}
