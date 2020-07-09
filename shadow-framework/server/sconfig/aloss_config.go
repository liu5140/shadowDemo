package sconfig

type TAlOssConfig struct {
	AccesskeyID     string
	AccesskeySecret string
	BucketName      string
	EndPoint        string
}

var AlOssConfigFactories = make(map[string]TAlOssConfig)

func RegisterAlOssConfigure(name string, factory TAlOssConfig) {
	AlOssConfigFactories[name] = factory
}

func AlOssConfigureInstance(name string) TAlOssConfig {
	factory := AlOssConfigFactories[name]
	return factory
}
