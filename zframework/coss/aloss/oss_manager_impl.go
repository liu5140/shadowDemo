package aloss

import (
	"shadowDemo/zframework/server"
	"shadowDemo/zframework/server/sconfig"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var dbManger IOssManager

type TOssManager struct {
	config sconfig.TAlOssConfig
	client *oss.Client
}

func OssInstance() IOssManager {
	return newOssManager()
}

func newOssManager() IOssManager {
	if dbManger == nil {
		l.Lock()
		defer l.Unlock()
		if dbManger == nil {
			dbManger = &TOssManager{}
		}
	}
	return dbManger
}

func (manager *TOssManager) Client() *oss.Client {
	if manager.client == nil {
		l.Lock()
		defer l.Unlock()
		if manager.client == nil {
			if manager.config.AccesskeyID == "" {
				manager.config = sconfig.AlOssConfigureInstance(server.ALOSS_CONFIGURE)
			}
			client, err := oss.New(manager.config.EndPoint, manager.config.AccesskeyID, manager.config.AccesskeySecret, oss.Timeout(30, 120))
			if err != nil {

			}
			manager.client = client
		}
	}
	return manager.client
}
