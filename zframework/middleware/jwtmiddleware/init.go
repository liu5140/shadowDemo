package jwtmiddleware

import (
	"shadowDemo/zframework/logger"
)

var (
	Log *logger.Logger
)

func init() {
	Log = logger.InitLog()
	Log.Info("JwtParser init")
}
