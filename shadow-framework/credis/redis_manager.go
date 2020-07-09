package credis

import (
	"github.com/gomodule/redigo/redis"
)

type IRedisManager interface {
	Client() *redis.Pool
}

type FRedisManagerFactory func() IRedisManager

var RedisManagerFactories = make(map[string]FRedisManagerFactory)

func RegisterRedisManager(name string, factory FRedisManagerFactory) {
	RedisManagerFactories[name] = factory
}

func RedisManagerInstance(name string) IRedisManager {
	factory := RedisManagerFactories[name]
	return factory()
}

