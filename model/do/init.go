package do

import "shadowDemo/zframework/logger"

var Log *logger.Logger = nil

func init() {
	Log = logger.InitLog()
}
