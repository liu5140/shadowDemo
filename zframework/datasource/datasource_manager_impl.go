package datasource

import (
	"regexp"
	"strings"
	"time"

	"shadowDemo/zframework/logger"
	"shadowDemo/zframework/server"
	"shadowDemo/zframework/server/sconfig"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

type TSQLLogger struct{}

func (slog TSQLLogger) Print(values ...interface{}) {
	vals := gorm.LogFormatter(values...)
	logger.Log.SqlDebug(vals...)
}

var dbManger IDatasourceManager
var shardingManager IShardingDatasourceManager

type TGormDataSourceManager struct {
	configs    []sconfig.TDataSourceConfig
	masterDB   *gorm.DB
	slaveDB    *gorm.DB
	shardingDB map[string]*gorm.DB
	Models     []interface{}
}

func ShardindDataSourceInstance() IShardingDatasourceManager {
	return newGormShardingDatasourceManager()
}

func DataSourceInstance() IDatasourceManager {
	return newGormDataSourceManager()
}
func newGormDataSourceManager() IDatasourceManager {
	if dbManger == nil {
		l.Lock()
		defer l.Unlock()
		if dbManger == nil {
			dbManger = &TGormDataSourceManager{}
		}
	}
	return dbManger
}

func newGormShardingDatasourceManager() IShardingDatasourceManager {
	if shardingManager == nil {
		l.Lock()
		defer l.Unlock()
		if shardingManager == nil {
			shardingManager = &TGormDataSourceManager{}
		}
	}
	return shardingManager
}

func (manager *TGormDataSourceManager) Configs(configs []sconfig.TDataSourceConfig) {
	manager.configs = configs
}

// auto create/migrate table if you want
func (manager *TGormDataSourceManager) RegisterModels(models ...interface{}) {
	manager.Models = append(manager.Models, models...)
}

func (manager *TGormDataSourceManager) Datasource() *gorm.DB {
	return manager.Master()
}

func (manager *TGormDataSourceManager) SDatasource(key string) *gorm.DB {
	if manager.shardingDB == nil {
		l.Lock()
		defer l.Unlock()
		if manager.shardingDB == nil {
			m := make(map[string]*gorm.DB)
			if manager.configs == nil {
				panic("configs is nil")
			}
			for _, config := range manager.configs {
				db := manager.openConn(config)
				m[config.Key] = db
			}
			manager.shardingDB = m
		}
	}

	return manager.shardingDB[key]
}

// lazy init
func (manager *TGormDataSourceManager) Master() *gorm.DB {
	if manager.masterDB == nil {
		l.Lock()
		defer l.Unlock()
		if manager.masterDB == nil {
			if manager.configs == nil {
				manager.configs = sconfig.DataSourceConfigureInstance(server.DATASOURCE_CONFIGURE)
			}
			config := manager.configs[0]
			db := manager.openConn(config)
			manager.masterDB = db
		}
	}

	return manager.masterDB
}

func (manager *TGormDataSourceManager) Slave() *gorm.DB {
	if len(manager.configs) == 1 {
		return manager.Master()
	}
	if manager.slaveDB == nil {
		l.Lock()
		defer l.Unlock()
		if manager.slaveDB == nil {
			if manager.configs == nil {
				manager.configs = sconfig.DataSourceConfigureInstance(server.DATASOURCE_CONFIGURE)
			}
			config := manager.configs[1]
			db := manager.openConn(config)
			manager.slaveDB = db
		}
	}
	return manager.slaveDB
}

func (manager *TGormDataSourceManager) openConn(config sconfig.TDataSourceConfig) *gorm.DB {
	db, err := gorm.Open(config.Driver, config.URL)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"username": config.Username,
			"password": config.Password,
			"url":      config.URL,
			"driver":   config.Driver,
		}).Error("DataSourceManager init error")
		if error := manager.tryToCreateDatabase(); error != nil {
			panic(error)
		}
	}

	db.DB().SetMaxIdleConns(config.IdlePoolSize)
	db.DB().SetMaxOpenConns(config.MaxPoolSize)
	db.DB().SetConnMaxLifetime(time.Duration(config.MaxLifeTime) * time.Second)

	// 设置字符编码
	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8")
	db.SingularTable(true)
	if config.SqlDebug == 1 {
		db.LogMode(true)
		db.SetLogger(TSQLLogger{})
	}

	if config.AutoCreate {
		// 如果有注册models, 则进行建表同步
		if len(manager.Models) > 0 {
			for _, m := range manager.Models {
				if !db.HasTable(m) {
					err := db.CreateTable(m).Error
					if err != nil {
						Log.Error(err)
					}
				}
			}
			db.AutoMigrate(manager.Models...)
		}
		//从config注册的model进行建表， 为了兼容老版本，将同时使用两种方式
		if len(config.Models) > 0 {
			for _, m := range manager.Models {
				if !db.HasTable(m) {
					err := db.CreateTable(m).Error
					if err != nil {
						Log.Error(err)
					}
				}
			}
			db.AutoMigrate(config.Models...)
		}
	}

	Log.WithField("db", db).Debug("create a new db connetion")
	return db
}

func (manager *TGormDataSourceManager) tryToCreateDatabase() error {
	Log.Info("Try to create a new db accrodding to url")
	for _, config := range manager.configs {
		driver := config.Driver
		url := config.URL

		if config.Driver == "mysql" {
			var dbname string
			if strings.Contains(url, "?") {
				reg := regexp.MustCompile("/(.*)\\?")
				result := reg.FindStringSubmatch(url)
				if len(result) == 2 {
					dbname = string(result[1])
				}
			} else {
				dbname = strings.Split(url, "/")[1]
			}

			rootURL := strings.Replace(url, dbname, "", 1)
			db, error := gorm.Open(driver, rootURL)
			if error != nil {
				Log.WithFields(logrus.Fields{
					"dbname": dbname,
					"driver": driver,
					"url":    rootURL,
				}).Error("create database failed")
				return error
			}

			if error = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname).Error; error != nil {
				Log.WithFields(logrus.Fields{
					"dbname": dbname,
					"driver": driver,
					"url":    rootURL,
				}).Error("create database failed")
				return error
			}
		}
	}
	return nil
}
