package sconfig

type TPlayerConfig struct {
	Host        string
	Port        string
	ContextPath string
	CurrentSite string
}

var PlayerConfigureFactories = make(map[string]TPlayerConfig)

func RegisterPlayerConfigure(name string, factory TPlayerConfig) {
	PlayerConfigureFactories[name] = factory
}

func PlayerConfigureInstance(name string) TPlayerConfig {
	factory := PlayerConfigureFactories[name]
	return factory
}
