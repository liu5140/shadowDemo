package middleware

import "github.com/gin-gonic/gin"

//IMiddlewareHandler IMiddlewareHandler
type IMiddlewareHandler interface {
	Handle(c *gin.Context)
}

//FMiddlewareHandlerFactory  FMiddlewareHandlerFactory
type FMiddlewareHandlerFactory func() IMiddlewareHandler

//MiddlewareHandlerFactories MiddlewareHandlerFactories
var MiddlewareHandlerFactories = make(map[string]FMiddlewareHandlerFactory)

//RegisterMiddlewareHandler RegisterMiddlewareHandler
func RegisterMiddlewareHandler(name string, factory FMiddlewareHandlerFactory) {
	MiddlewareHandlerFactories[name] = factory
}

//MiddlewareHandlerInstance MiddlewareHandlerInstance
func MiddlewareHandlerInstance(name string) IMiddlewareHandler {
	factory := MiddlewareHandlerFactories[name]
	return factory()
}
