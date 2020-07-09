package aloss

import (
	"sync"

	"shadowDemo/shadow-framework/logger"
)

var (
	Log *logger.Logger
	l   sync.Mutex
)

const (
	OSS_MANAGER = "OSSManager"
)

func init() {
	Log = logger.InitLog()
	Log.Infoln("OssManager init")
	RegisterOssManager(OSS_MANAGER, OssInstance)
}
