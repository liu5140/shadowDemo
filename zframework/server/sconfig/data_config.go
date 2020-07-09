package sconfig

type TDataSourceConfig struct {
	Key          string
	Username     string
	Password     string
	URL          string
	Driver       string
	IdlePoolSize int
	MaxPoolSize  int
	MaxLifeTime  int64
	SqlDebug     int8
	AutoCreate   bool
	Models       []interface{}
}

var DataSourceConfigureFactories = make(map[string][]TDataSourceConfig)

func RegisterDataSourceConfigure(name string, factory []TDataSourceConfig) {
	DataSourceConfigureFactories[name] = factory
}

func DataSourceConfigureInstance(name string) []TDataSourceConfig {
	factory := DataSourceConfigureFactories[name]
	return factory
}
