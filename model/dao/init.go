package dao

import "shadowDemo/shadow-framework/logger"

var Log *logger.Logger = nil

func init() {
	Log = logger.InitLog()
}
