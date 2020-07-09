package router

import (
	"errors"
	"shadowDemo/middleware"
	"shadowDemo/service"

	"shadowDemo/shadow-framework/security"

	"github.com/gin-gonic/gin"
)

func webSocketRouter(r *gin.RouterGroup) {
	r.GET("/ws", getWebSocket)
	r.GET("/ws/response", responseWebSocket)
}

func responseWebSocket(c *gin.Context) {
	profile := c.MustGet(PROFILE).(middleware.Profile)
	message, _ := c.GetQuery("message")
	Log.Infoln("==============", message)
	err := service.WriteMessage(profile.ID, []byte(message))
	if err != nil {
		newServerError(c, err)
		return
	}
	newSuccess(c)
}

func getWebSocket(c *gin.Context) {
	token, ok := c.GetQuery("token")
	Log.Infoln("=================", token)
	if !ok {
		err := security.JwtExpiredErr{
			Err: errors.New("jwt 失效3"),
		}
		newClientError(c, err)
		return
	}

	session := service.NewSessionService()
	id, err := session.ParseToken(token, hmacSampleSecret)
	if err != nil {
		newServerError(c, err)
		return
	}
	wsconn, err := service.InitWebSocketService(c)
	if err != nil {
		newServerError(c, err)
		return
	}

	err = service.InitWebSocketClientService(id, wsconn)
	if err != nil {
		newServerError(c, err)
		return
	}
	go wsconn.Read()
	go wsconn.WriteText()
}
