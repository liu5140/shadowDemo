package idgenerator

import (
	"shadowDemo/shadow-framework/logger"
)

var Log *logger.Logger

const ID_GENERATOR = "IDGenerator"

func init() {
	Log = logger.InitLog()

	Log.Infoln("IDGenerator init")
	RegisterService(ID_GENERATOR, newSQLIdGenerator)
}
