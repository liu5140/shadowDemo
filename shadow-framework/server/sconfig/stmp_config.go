package sconfig

type TStmpConfig struct {
	SendEmailAccount string
	Host             string
	Port             int
	Username         string
	Password         string
}

var StmpConfigureFactories = make(map[string]TStmpConfig)

func RegisterStmpConfigure(name string, factory TStmpConfig) {
	StmpConfigureFactories[name] = factory
}

func StmpConfigureInstance(name string) TStmpConfig {
	factory := StmpConfigureFactories[name]
	return factory
}
