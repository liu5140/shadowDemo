package service

import (
	"shadowDemo/model"
	"shadowDemo/model/dao"
	"shadowDemo/model/do"
	modelc "shadowDemo/shadow-framework/model"
)

type APIAccessResLogService struct{}

var aPIAccessResLogService *APIAccessResLogService

func NewAPIAccessResLogService() *APIAccessResLogService {
	if aPIAccessResLogService == nil {
		l.Lock()
		if aPIAccessResLogService == nil {
			aPIAccessResLogService = &APIAccessResLogService{}
		}
		l.Unlock()
	}
	return aPIAccessResLogService
}

//创建
func (service *APIAccessResLogService) CreateAPIAccessResLog(aPIAccessResLog *do.APIAccessResLog) (err error) {
	return model.GetModel().APIAccessResLogDao.Create(aPIAccessResLog)
}

//通过id获取详情
func (service *APIAccessResLogService) GetAPIAccessResLogByID(id int64) (aPIAccessResLog *do.APIAccessResLog, err error) {
	aPIAccessResLog.ID = id
	err = model.GetModel().APIAccessResLogDao.Get(aPIAccessResLog)
	if err != nil {
		Log.Error(err)
		return aPIAccessResLog, err
	}
	return aPIAccessResLog, err
}

//通过id删除
func (service *APIAccessResLogService) DeleteAPIAccessResLogByID(id int64) (err error) {
	if model.GetModel().APIAccessResLogDao.Delete(&do.APIAccessResLog{ID: id}); err != nil {
		Log.Error(err)
		return err
	}
	return err
}

//通过id更新
func (service *APIAccessResLogService) UpdateAPIAccessResLog(id int64, attrs map[string]interface{}) (err error) {
	if err = model.GetModel().APIAccessResLogDao.Updates(id, attrs); err != nil {
		Log.Error(err)
		return err
	}
	return err
}

//查询
func (service *APIAccessResLogService) SearchAPIAccessResLogPaging(condition *dao.APIAccessResLogSearchCondition, pageNum int, pageSize int) (request []do.APIAccessResLog, count int, err error) {
	rowbound := modelc.NewRowBound(pageNum, pageSize)
	return service.searchAPIAccessResLog(condition, &rowbound)
}

func (service *APIAccessResLogService) SearchAPIAccessResLogWithOutPaging(condition *dao.APIAccessResLogSearchCondition) (request []do.APIAccessResLog, count int, err error) {
	return service.searchAPIAccessResLog(condition, nil)
}

func (service *APIAccessResLogService) searchAPIAccessResLog(condition *dao.APIAccessResLogSearchCondition, rowbound *modelc.RowBound) (request []do.APIAccessResLog, count int, err error) {
	result, count, err := model.GetModel().APIAccessResLogDao.SearchAPIAccessResLogs(condition, rowbound)
	if err != nil {
		Log.Error(err)
		return nil, 0, err
	}
	return result, count, err
}
