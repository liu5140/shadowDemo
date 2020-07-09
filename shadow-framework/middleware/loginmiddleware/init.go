package loginmiddleware

import (
	"shadowDemo/shadow-framework/logger"
	"shadowDemo/shadow-framework/middleware"
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
