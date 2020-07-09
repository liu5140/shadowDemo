package adapter

import (
	"shadowDemo/shadow-framework/logger"
)

var Log *logger.Logger

func init() {
	Log = logger.InitLog()
}
