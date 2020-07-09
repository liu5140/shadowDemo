package errormiddleware

import (
	"shadowDemo/zframework/logger"
)

var (
	Log *logger.Logger
)

func init() {
	Log = logger.InitLog()
	Log.Info("ErrorHandler init")
}
