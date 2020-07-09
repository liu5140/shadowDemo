package middleware

import (
	"shadowDemo/shadow-framework/logger"
)

var (
	Log *logger.Logger
)

const (
	LOGIN  = "login"
	LOGOUT = "logout"
)

func init() {
	Log = logger.InitLog()
	Log.Info("DefaultLoginUrlRegistry init")
	RegisterUrlRegistry(LOGIN, newDefaultLoginUrlRegistry)
	Log.Info("DefaultLogoutUrlRegistry init")
	RegisterUrlRegistry(LOGOUT, newDefaultLogoutUrlRegistry)
}
