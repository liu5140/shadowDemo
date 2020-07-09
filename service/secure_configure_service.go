package service

import (
	"errors"

	"shadowDemo/zframework/credis"
	"shadowDemo/zframework/utils/encrypt"

	"github.com/gomodule/redigo/redis"
)

type SecureConfigureService struct {
	encryptKey []byte
}

var secureConfigureService *SecureConfigureService

func NewSecureConfigureService() *SecureConfigureService {
	if secureConfigureService == nil {
		l.Lock()
		if secureConfigureService == nil {
			secureConfigureService = &SecureConfigureService{
				encryptKey: []byte("TvgRZ5GwVgrXyaWLLNWpvjtPMQxuXe28"),
			}
		}
		l.Unlock()
	}
	return secureConfigureService
}

//Get 根据配置key查询一个配置
func (service *SecureConfigureService) Get(key string) (string, error) {
	conn := credis.RedisManagerInstance(credis.REDIS_MANAGER).Client().Get()
	defer conn.Close()

	raw, err := redis.String(conn.Do("HGET", configKey, key))
	if err != nil {
		err := errors.New("key not found")
		Log.Error(err)
		return raw, err
	}
	return encrypt.Decrypt(service.encryptKey, raw), nil
}

//GetAll 获取所有配置
func (service *SecureConfigureService) GetAll() (map[string]string, error) {
	conn := credis.RedisManagerInstance(credis.REDIS_MANAGER).Client().Get()
	defer conn.Close()

	configs, err := redis.StringMap(conn.Do("HGETALL", configKey))
	if err != nil {
		err := errors.New("key not found")
		Log.Error(err)
		return nil, err
	}

	for k, v := range configs {
		configs[k] = encrypt.Decrypt(service.encryptKey, v)
	}
	return configs, nil
}

//Set 添加配置，如果配置的key已经存在，则更新配置, dataType: 1：json, 2:string
func (service *SecureConfigureService) Set(key string, data string) (err error) {
	conn := credis.RedisManagerInstance(credis.REDIS_MANAGER).Client().Get()
	defer conn.Close()

	//加密存储
	secureValue := encrypt.Encrypt(service.encryptKey, data)
	if err != nil {
		Log.Error(err)
		return err
	}

	_, err = conn.Do("HSET", configKey, key, secureValue)
	if err != nil {
		Log.Error(err)
		return err
	}
	return nil
}
