package sconfig

type TAdminConfig struct {
	Host        string
	Port        string
	ContextPath string
	CurrentSite string
}

var AdminConfigureFactories = make(map[string]TAdminConfig)

func RegisterAdminConfigure(name string, factory TAdminConfig) {
	AdminConfigureFactories[name] = factory
}

func AdminConfigureInstance(name string) TAdminConfig {
	factory := AdminConfigureFactories[name]
	return factory
}
