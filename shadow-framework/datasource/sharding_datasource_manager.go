package datasource

import (
	"github.com/jinzhu/gorm"
	"shadowDemo/shadow-framework/server/sconfig"
)

type IShardingDatasourceManager interface {
	SDatasource(key string) *gorm.DB
	Configs(configs []sconfig.TDataSourceConfig)
	RegisterModels(models ...interface{})
}
type FShardingDatasourceManagerFactory func() IShardingDatasourceManager

var shardingDatasourceManagerFactories = make(map[string]FShardingDatasourceManagerFactory)

func RegisterShardingDatasourceManager(name string, factory FShardingDatasourceManagerFactory) {
	shardingDatasourceManagerFactories[name] = factory
}

func ShardingDatasourceManagerInstance(name string) IShardingDatasourceManager {
	factory := shardingDatasourceManagerFactories[name]
	return factory()
}

func ShardingDatasourceInstance() IShardingDatasourceManager {
	factory := shardingDatasourceManagerFactories[DATASOURCE_MANAGER]
	return factory()
}
