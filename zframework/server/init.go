package server

import (
	"shadowDemo/zframework/logger"
	"shadowDemo/zframework/server/sconfig"
)

var Log *logger.Logger

const (
	DATASOURCE_CONFIGURE = "DataSourceConfigure"
	REDIS_CONFIGURE      = "RedisConfigure"
	ALOSS_CONFIGURE      = "ALossConfigure"
)

var (
	DefaultConfigurePath string = "config/server.json"
	serverConfig         *TServerConfigure
)

func init() {
	Log = logger.InitLog()
	Log.Infoln("======初始化server.json========")
	serverConfig := newServerConfigure()
	Log.Infoln("======加载数据库配置=========")
	sconfig.RegisterDataSourceConfigure(DATASOURCE_CONFIGURE, serverConfig.DataSource)
	Log.Infoln("======加载redis配置========")
	sconfig.RegisterRedisConfigure(REDIS_CONFIGURE, serverConfig.RedisSource)
	Log.Infoln("======加载阿里oss配置=======")
	sconfig.RegisterAlOssConfigure(ALOSS_CONFIGURE, serverConfig.AlossServer)
}
