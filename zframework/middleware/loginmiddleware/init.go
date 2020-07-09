package loginmiddleware

import (
	"shadowDemo/zframework/logger"
	"shadowDemo/zframework/middleware"
)

var (
	Log *logger.Logger
)

const (
	LOGIN_HANDLER = "LoginHandler"
)

func init() {
	Log = logger.InitLog()
	Log.Info("LoginHandler init")
	middleware.RegisterMiddlewareHandler(LOGIN_HANDLER, newDefaultLoginHandler)
}
