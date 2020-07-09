package sconfig

type TRedisConfig struct {
	Password    string
	Host        string
	Port        string
	DB          string
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
}

var RedisConfigureFactories = make(map[string]TRedisConfig)

func RegisterRedisConfigure(name string, factory TRedisConfig) {
	RedisConfigureFactories[name] = factory
}

func RedisConfigureInstance(name string) TRedisConfig {
	factory := RedisConfigureFactories[name]
	return factory
}
