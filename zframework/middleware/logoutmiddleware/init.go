package logoutmiddleware

import (
	"shadowDemo/zframework/logger"
	"shadowDemo/zframework/middleware"

	"github.com/astaxie/beego/session"
)

var (
	Log            *logger.Logger
	globalSessions *session.Manager
)

const (
	LOGOUT         = "logout"
	LOGOUT_HANDLER = "LogoutHandler"
)

func init() {
	Log = logger.InitLog()
	Log.Info("DefaultLogoutUrlRegistry init")
	middleware.RegisterMiddlewareHandler(LOGOUT_HANDLER, newDefaultLogoutHandler)
}
