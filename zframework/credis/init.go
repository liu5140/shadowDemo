package credis

import (
	"shadowDemo/zframework/logger"
	"sync"
)

var (
	Log *logger.Logger
	l   sync.Mutex
)

const (
	REDIS_MANAGER = "RedisManager"
)

func init() {
	Log = logger.InitLog()
}
