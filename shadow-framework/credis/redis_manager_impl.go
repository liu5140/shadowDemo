package credis

import (
	"shadowDemo/shadow-framework/server/sconfig"
	"time"

	"github.com/gomodule/redigo/redis"
)

var dbManger IRedisManager

type TRedisManager struct {
	config sconfig.TRedisConfig
	client *redis.Pool
}

func RedisInstance() IRedisManager {
	return newRedisManager()
}

func newRedisManager() IRedisManager {
	if dbManger == nil {
		l.Lock()
		defer l.Unlock()
		if dbManger == nil {
			dbManger = &TRedisManager{}
		}
	}
	return dbManger
}

func (manager *TRedisManager) Client() *redis.Pool {
	if manager.client == nil {
		Log.Info("Redis init")
		manager.client = &redis.Pool{
			MaxIdle:     manager.config.MaxIdle,
			MaxActive:   manager.config.MaxActive,
			IdleTimeout: time.Duration(manager.config.IdleTimeout) * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", manager.config.Host+manager.config.Port)
				if err != nil {
					return nil, err
				}
				c.Do("AUTH", manager.config.Password)
				c.Do("SELECT", manager.config.DB)
				return c, nil
			},
		}
		Log.Debugln("Redis init success", manager.config.Host+manager.config.Port)

	}
	return manager.client
}
