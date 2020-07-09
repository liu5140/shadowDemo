// Package router shadowDemo restful API.
//     Version: 1.0
//     Host: localhost:8030
//     BasePath: /acc
// SecurityDefinitions:
//     api_key:
//          type: apiKey
//          name: Authorization
//          in: header
// swagger:meta
package router

import (
	"net/http"
	"shadowDemo/APIServer/apimiddleware"
	"shadowDemo/shadow-framework/middleware/concurrentlimit"
	"shadowDemo/shadow-framework/server"

	"github.com/gin-gonic/gin"
)

const (
	hmacSampleSecret = "tnb9Y0du$2a$10$KmatydruRTKlaUwErUOtNOXiPHVPunb9Y0dup9newm"
	hmacSecureSecret = "Tnb9Y0du$2a$10$KmatydruRTKlaUwErUOtNOXiPHVPunb9Y0dup9newm"
)

func MainRouter() http.Handler {
	//loginPath := server.ServerConfigInstance().AdminServer.ContextPath + "/login"
	engine := gin.New()
	engine.Use(concurrentlimit.ConcurrentLimit(20, 0))
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	engine.Use(apimiddleware.RequestResolver())
	engine.Use(apimiddleware.RequestSignValidator())
	engine.Use(apimiddleware.MerchantIPFilter(apimiddleware.VIEW_DEPOSIT_URL))
	//engine.Use(apimiddleware.QueryOrderHanlder(apimiddleware.VIEW_DEPOSIT_URL, apimiddleware.QUERY_DEPOSIT))
	engine.Use(apimiddleware.ErrorHanlder())
	//	engine.Use(apimiddleware.CreateOrderHanlder(apimiddleware.DEPOSIT, apimiddleware.WITHDRAW, 2))
	engine.Use(apimiddleware.AccessLogHandler())

	router := engine.Group(server.ServerConfigInstance().PlayerServer.ContextPath)
	actionRouter(router)
	return engine
}
