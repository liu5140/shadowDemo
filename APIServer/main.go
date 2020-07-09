package main

import (
	"context"
	"encoding/hex"
	"log"
	"net/http"
	"os"
	"os/signal"
	"shadowDemo/APIServer/router"
	"shadowDemo/APIServer/vo"
	"shadowDemo/model"
	"shadowDemo/service"
	"shadowDemo/zframework/credis"
	"shadowDemo/zframework/datasource"
	"shadowDemo/zframework/logger"
	"shadowDemo/zframework/security"
	"shadowDemo/zframework/server"

	"shadowDemo/zframework/utils/aes"
	"shadowDemo/zframework/utils/encoder"
	"shadowDemo/zframework/utils/query"
	"strings"
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

	Log.Infoln("model 初始化")
	model.ModelInit()

	Log.Infoln("service 初始化")
	service.ServiceInit()

	BuildVersion = server.ServerConfigInstance().BuildVersion

	ss := vo.APIChannelReqData{UserLevel: "1"}

	plain, err := encoder.JSONMarshal(ss)

	Log.Debug("plain: ", string(plain))

	mm, err := aes.AesEcbEncrypt(plain, []byte("g3im_lF46_h6eRmU1FmyS4AKVBZ8QWwo"))
	Log.Debug("err: ", err)

	mmstr := hex.EncodeToString(mm)
	Log.Debug("EncryptData: ", mmstr)
	Result := strings.ToUpper(string(mmstr))
	Log.Debug("EncryptData: ", Result)

	vo := vo.APIPublicRequest{
		MerchantNo: "liujian",
		Action:     "query_banklist",
		Data:       Result,
		Sign:       "",
	}

	values, _, _ := query.Values(vo)
	param := values.Encode() + "&Key=" + "tbKAdNFwroxe5acEee98T0ZZXenoXC18sdtr4lwIILY="
	Log.Debug("param: ", param)

	sign := strings.ToUpper(encoder.MD5(param))

	Log.Debug("sign: ", sign)

	//tbKAdNFwroxe5acEee98T0ZZXenoXC18sdtr4lwIILY=
	// mm, err := aes.InternalEncryptStr("111111")
	// Log.Infoln("=======aes mm", mm)

	// mmstr, err := aes.InternalDecryptStr(mm)

	// Log.Infoln("=======aes reslut", mmstr)
	//	os.Exit(-1)

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
		Addr:         server.ServerConfigInstance().PlayerServer.Port,
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
