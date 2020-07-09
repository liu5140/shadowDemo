package security

import (
	io "io/ioutil"
	"os"
	"path"
	"runtime"

	"shadowDemo/zframework/datasource"
	adapter "shadowDemo/zframework/security/adapter"

	"github.com/casbin/casbin"
)

var enforcer *casbin.Enforcer

func GetCasbinEnforcer(config ...string) *casbin.Enforcer {
	if enforcer == nil {
		var cnf string
		if len(config) > 0 {
			cnf = config[0]
		} else {
			cnf = casbinModelFilePath()
		}
		Log.Info("Casbin Enforcer init")
		adapter := adapter.NewAdapter(datasource.DatasourceManagerInstance(datasource.DATASOURCE_MANAGER).Datasource())
		enforcer = casbin.NewEnforcer(cnf, adapter)
		enforcer.LoadPolicy()
	}
	return enforcer
}

func casbinModelFilePath() string {
	conf := "./config/rbac/rbac_model.conf"
	if _, err := os.Stat(conf); os.IsNotExist(err) {
		_, filename, _, _ := runtime.Caller(1)
		conf = path.Join(path.Dir(filename), "../", conf)
	}
	data, err := io.ReadFile(conf)
	if err != nil {
		Log.WithField("rbac_model.conf", conf).Error("server configure init failed, file doesn't exist")
		Log.Panic(err)
	}
	Log.Info("rbac_model.conf", string(data))
	return conf
}
