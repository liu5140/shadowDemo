package concurrentlimit

import (
	"shadowDemo/shadow-framework/logger"
)

var (
	Log *logger.Logger
)

func init() {
	Log = logger.InitLog()
	Log.Info("concurrentlimit init")
}
