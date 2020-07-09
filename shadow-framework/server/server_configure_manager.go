package server

import (
	"encoding/json"
	"io/ioutil"
	"shadowDemo/shadow-framework/server/sconfig"

	"github.com/sirupsen/logrus"
)

type TServerConfigure struct {
	ServerName   string
	Platform     string
	Node         int
	BuildVersion string
	DataSource   []sconfig.TDataSourceConfig
	AlossServer  sconfig.TAlOssConfig
	RedisSource  sconfig.TRedisConfig
	AdminServer  sconfig.TAdminConfig
	PlayerServer sconfig.TPlayerConfig
	StmpServer   sconfig.TStmpConfig
}

func ServerConfigInstance() *TServerConfigure {
	return newServerConfigure()
}

func newServerConfigure() *TServerConfigure {
	if serverConfig == nil {
		config := &TServerConfigure{}
		LoadWithFile(config, DefaultConfigurePath)
		Log.WithFields(logrus.Fields{
			"ServerName": config.ServerName,
			"Platform":   config.Platform,
			"Node":       config.Node,
		}).Debug("TServerConfigure")
		serverConfig = config
	}
	return serverConfig
}

func LoadWithFile(configure interface{}, path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		Log.WithField("file", path).Error("server configure init failed, file doesn't exist")
		Log.Panic(err)
	}
	Log.Info("Server configure", string(data))
	datajson := []byte(data)
	err = json.Unmarshal(datajson, configure)
	if err != nil {
		Log.Panic(err)
	}
	if config, ok := configure.(TServerConfigure); ok {
		serverConfig = &config
	}
}
