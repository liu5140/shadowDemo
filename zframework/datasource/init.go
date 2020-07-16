package datasource

import (
	"shadowDemo/zframework/logger"
	"sync"
)

var (
	Log *logger.Logger
	l   sync.Mutex
)

const (
	DATASOURCE_MANAGER = "DataSourceManager"
)

func init() {
	Log = logger.InitLog()
	Log.Infoln("==============datasource=======================")
}
