package aloss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type IOssManager interface {
	Client() *oss.Client
}

type FOssManagerFactory func() IOssManager

var OssManagerFactories = make(map[string]FOssManagerFactory)

func RegisterOssManager(name string, factory FOssManagerFactory) {
	OssManagerFactories[name] = factory
}

func OssManagerInstance(name string) IOssManager {
	factory := OssManagerFactories[name]
	return factory()
}
