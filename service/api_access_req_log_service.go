package service

import (
	"shadowDemo/model"
	"shadowDemo/model/dao"
	"shadowDemo/model/do"
	modelc "shadowDemo/shadow-framework/model"
)

type APIAccessReqLogService struct{}

var aPIAccessReqLogService *APIAccessReqLogService

func NewAPIAccessReqLogService() *APIAccessReqLogService {
	if aPIAccessReqLogService == nil {
		l.Lock()
		if aPIAccessReqLogService == nil {
			aPIAccessReqLogService = &APIAccessReqLogService{}
		}
		l.Unlock()
	}
	return aPIAccessReqLogService
}

//创建
func (service *APIAccessReqLogService) CreateAPIAccessReqLog(aPIAccessReqLog *do.APIAccessReqLog) (err error) {
	return model.GetModel().APIAccessReqLogDao.Create(aPIAccessReqLog)
}

//通过id获取详情
func (service *APIAccessReqLogService) GetAPIAccessReqLogByID(id int64) (aPIAccessReqLog *do.APIAccessReqLog, err error) {
	aPIAccessReqLog.ID = id
	err = model.GetModel().APIAccessReqLogDao.Get(aPIAccessReqLog)
	if err != nil {
		Log.Error(err)
		return aPIAccessReqLog, err
	}
	return aPIAccessReqLog, err
}

//通过id删除
func (service *APIAccessReqLogService) DeleteAPIAccessReqLogByID(id int64) (err error) {
	if model.GetModel().APIAccessReqLogDao.Delete(&do.APIAccessReqLog{ID: id}); err != nil {
		Log.Error(err)
		return err
	}
	return err
}

//通过id更新
func (service *APIAccessReqLogService) UpdateAPIAccessReqLog(id int64, attrs map[string]interface{}) (err error) {
	if err = model.GetModel().APIAccessReqLogDao.Updates(id, attrs); err != nil {
		Log.Error(err)
		return err
	}
	return err
}

//查询
func (service *APIAccessReqLogService) SearchAPIAccessReqLogPaging(condition *dao.APIAccessReqLogSearchCondition, pageNum int, pageSize int) (request []do.APIAccessReqLog, count int, err error) {
	rowbound := modelc.NewRowBound(pageNum, pageSize)
	return service.searchAPIAccessReqLog(condition, &rowbound)
}

func (service *APIAccessReqLogService) SearchAPIAccessReqLogWithOutPaging(condition *dao.APIAccessReqLogSearchCondition) (request []do.APIAccessReqLog, count int, err error) {
	return service.searchAPIAccessReqLog(condition, nil)
}

func (service *APIAccessReqLogService) searchAPIAccessReqLog(condition *dao.APIAccessReqLogSearchCondition, rowbound *modelc.RowBound) (request []do.APIAccessReqLog, count int, err error) {
	result, count, err := model.GetModel().APIAccessReqLogDao.SearchAPIAccessReqLogs(condition, rowbound)
	if err != nil {
		Log.Error(err)
		return nil, 0, err
	}
	return result, count, err
}
