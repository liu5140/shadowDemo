package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"shadowDemo/ACCServer/router"
	"shadowDemo/model"
	"shadowDemo/shadow-framework/credis"
	"shadowDemo/shadow-framework/datasource"
	"shadowDemo/shadow-framework/logger"
	"shadowDemo/shadow-framework/security"
	"shadowDemo/shadow-framework/server"
	"syscall"
	"time"

	rediswatcher "github.com/billcobbler/casbin-redis-watcher"
	"golang.org/x/sync/errgroup"
)

var (
	g            errgroup.Group
	Log          *logger.Logger = logger.InitLog()
	BuildVersion string
	serverConfig *server.TServerConfigure
)

//go:generate swagger generate spec
func main() {
	Log.Infoln("主数据库初始化")
	datasource.RegisterDatasourceManager(datasource.DATASOURCE_MANAGER, datasource.DataSourceInstance)
	Log.Infoln("注册表")
	datasource.DatasourceManagerInstance(datasource.DATASOURCE_MANAGER).RegisterModels(model.GetInitialModels()...)
	//Log.Infoln("从数据库初始化")
	//datasource.RegisterShardingDatasourceManager(datasource.DATASOURCE_MANAGER, datasource.ShardindDataSourceInstance)
	Log.Infoln("redis 初始化")
	credis.RegisterRedisManager(credis.REDIS_MANAGER, credis.RedisInstance)
	//dao初始化
	model.ModelInit()

	BuildVersion = server.ServerConfigInstance().BuildVersion

	Log.Info("#######################################################")
	Log.Infof("###### %s ######", BuildVersion)
	Log.Info("#######################################################")
	//set casbin watcher
	w, err := rediswatcher.NewWatcher(server.ServerConfigInstance().RedisSource.Host + server.ServerConfigInstance().RedisSource.Port)
	if err != nil {
		Log.Error(err)
		return
	}
	e := security.GetCasbinEnforcer()
	e.SetWatcher(w)

	HttpPprof()

	var server = &http.Server{
		Addr:         server.ServerConfigInstance().AdminServer.Port,
		Handler:      router.MainRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	g.Go(func() error {
		return server.ListenAndServe()
	})

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown: ", err)
	}

	log.Println("server exiting...")

}

func HttpPprof() {
	go func() {
		log.Println(http.ListenAndServe(":8524", nil))
	}()
}
