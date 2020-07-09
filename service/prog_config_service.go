package service

import (
	"errors"
	"shadowDemo/model"
	"shadowDemo/model/dao"
	"shadowDemo/model/do"
	"shadowDemo/zframework/credis"
	modelc "shadowDemo/zframework/model"

	"github.com/gomodule/redigo/redis"
)

type ProgConfigService struct{}

var progConfigService *ProgConfigService

const configKey = "prog:config:"

func NewProgConfigService() *ProgConfigService {
	if progConfigService == nil {
		l.Lock()
		if progConfigService == nil {
			progConfigService = &ProgConfigService{}
		}
		l.Unlock()
		progConfigService.init()
	}
	return progConfigService
}

//加载时候查询所有可以用进行初始化
func (service *ProgConfigService) init() {
	result, err := model.GetModel().ProgConfigDao.Find(&do.ProgConfig{Disabled: true})
	if err != nil {
		Log.Errorln(err)
	}
	for _, v := range result {
		if err = service.Set(v.ParamName, v.ParamValue); err != nil {
			Log.Errorln(err)
		}
	}
}

//创建的时候如果是可用状态则，需要插入到redis中
func (service *ProgConfigService) CreateProgConfig(progConfig *do.ProgConfig) (err error) {
	if err = model.GetModel().ProgConfigDao.Create(progConfig); err != nil {
		Log.Errorln(err)
		return
	}
	//可用则插入到redis
	if progConfig.Disabled == true {
		if err = service.Set(progConfig.ParamName, progConfig.ParamValue); err != nil {
			Log.Errorln(err)
			return
		}
	}
	return err
}

//通过id获取详情
func (service *ProgConfigService) GetProgConfigByID(id int64) (progConfig *do.ProgConfig, err error) {
	progConfig.ID = id
	err = model.GetModel().ProgConfigDao.Get(progConfig)
	if err != nil {
		Log.Error(err)
		return progConfig, err
	}
	return progConfig, err
}

//通过id删除
func (service *ProgConfigService) DeleteProgConfigByID(id int64) (err error) {
	progConfig, err := service.GetProgConfigByID(id)
	if err != nil {
		Log.Error(err)
		return err
	}
	if model.GetModel().ProgConfigDao.Delete(progConfig); err != nil {
		Log.Error(err)
		return err
	}

	if err = service.Remove(progConfig.ParamName); err != nil {
		Log.Error(err)
		return err
	}
	return err
}

//通过id更新
func (service *ProgConfigService) UpdateProgConfig(id int64, attrs map[string]interface{}) (err error) {
	if err = model.GetModel().ProgConfigDao.Updates(id, attrs); err != nil {
		Log.Error(err)
		return err
	}
	progConfig, err := service.GetProgConfigByID(id)
	if err != nil {
		Log.Error(err)
		return err
	}
	//可用则插入到redis
	if progConfig.Disabled == true {
		if err = service.Set(progConfig.ParamName, progConfig.ParamValue); err != nil {
			Log.Errorln(err)
			return
		}
	}
	return err
}

//查询
func (service *ProgConfigService) SearchProgConfigPaging(condition *dao.ProgConfigSearchCondition, pageNum int, pageSize int) (request []do.ProgConfig, count int, err error) {
	rowbound := modelc.NewRowBound(pageNum, pageSize)
	return service.searchProgConfig(condition, &rowbound)
}

func (service *ProgConfigService) SearchProgConfigWithOutPaging(condition *dao.ProgConfigSearchCondition) (request []do.ProgConfig, count int, err error) {
	return service.searchProgConfig(condition, nil)
}

func (service *ProgConfigService) searchProgConfig(condition *dao.ProgConfigSearchCondition, rowbound *modelc.RowBound) (request []do.ProgConfig, count int, err error) {
	result, count, err := model.GetModel().ProgConfigDao.SearchProgConfigs(condition, rowbound)
	if err != nil {
		Log.Error(err)
		return nil, 0, err
	}
	return result, count, err
}

//Get 根据配置key查询一个配置
func (service *ProgConfigService) Get(key string) (string, error) {
	conn := credis.RedisManagerInstance(credis.REDIS_MANAGER).Client().Get()
	defer conn.Close()

	raw, err := redis.String(conn.Do("HGET", configKey, key))
	if err != nil {
		err := errors.New("key not found")
		Log.Error(err)
		return raw, err
	}
	return raw, nil
}

//GetAll 获取所有配置
func (service *ProgConfigService) GetAll() (map[string]string, error) {
	conn := credis.RedisManagerInstance(credis.REDIS_MANAGER).Client().Get()
	defer conn.Close()

	configs, err := redis.StringMap(conn.Do("HGETALL", configKey))
	if err != nil {
		err := errors.New("key not found")
		Log.Error(err)
		return nil, err
	}

	return configs, nil
}

//Set 添加配置，如果配置的key已经存在，则更新配置, dataType: 1：json, 2:string
func (service *ProgConfigService) Set(key string, data string) (err error) {
	conn := credis.RedisManagerInstance(credis.REDIS_MANAGER).Client().Get()
	defer conn.Close()
	_, err = conn.Do("HSET", configKey, key, data)
	if err != nil {
		Log.Error(err)
		return err
	}
	return nil
}

func (service *ProgConfigService) Remove(key string) (err error) {
	conn := credis.RedisManagerInstance(credis.REDIS_MANAGER).Client().Get()
	defer conn.Close()
	_, err = conn.Do("HDEL", configKey, key)
	if err != nil {
		Log.Error(err)
		return err
	}
	return nil
}
