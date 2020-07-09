package service

import (
	"shadowDemo/model"
	"shadowDemo/model/do"
	"shadowDemo/shadow-framework/idgenerator"
)

type APIAccessLogService struct{}

var apiAccessLogService *APIAccessLogService

func NewAPIAccessLogService() *APIAccessLogService {
	if apiAccessLogService == nil {
		l.Lock()
		if apiAccessLogService == nil {
			apiAccessLogService = &APIAccessLogService{}
		}
		l.Unlock()
	}
	return apiAccessLogService
}

func (manager APIAccessLogService) LogRequest(reqLog *do.APIAccessReqLog) error {
	idGen := idgenerator.Instance()
	reqLog.RequestNo = idGen.GenerateStringID("order")
	m := model.GetModel()
	return m.APIAccessReqLogDao.Create(reqLog)
}

func (manager APIAccessLogService) LogResponse(resLog *do.APIAccessResLog) error {
	m := model.GetModel()
	return m.APIAccessResLogDao.Create(resLog)
}

// func (manager *APIAccessLogService) SearchLogList(username string, startTime time.Time, endTime time.Time, pageNum int, pageSize int) (result []*model.TAPIAccessReqResLog, count int, err error) {
// 	if startTime.IsZero() {
// 		startTime = time.Now()
// 		endTime = time.Now()
// 	}

// 	if pageNum <= 0 {
// 		pageNum = 1
// 	}
// 	if pageSize <= 0 {
// 		pageSize = 20
// 	}
// 	rowbound := &model.RowBound{
// 		Limit:  pageSize,
// 		Offset: (pageNum - 1) * pageSize,
// 	}

// 	db := datasource.DatasourceServiceInstance(datasource.DATASOURCE_MANAGER).Datasource()
// 	dao := dao.NewAPIAccessLogDao(db)

// 	result, count, err = dao.FindReqLogs(&model.TAPIAccessReqLog{
// 		MerchantName: username,
// 	}, startTime, endTime, rowbound)

// 	return
// }

// func (manager *APIAccessLogService) SearchLogByOrder(orderNo string) []*model.TAPIAccessReqResLog {
// 	db := datasource.DatasourceServiceInstance(datasource.DATASOURCE_MANAGER).Datasource()
// 	dao := dao.NewAPIAccessLogDao(db)

// 	result, _, _ := dao.FindReqLogs(&model.TAPIAccessReqLog{
// 		OrderNo: orderNo,
// 	}, time.Time{}, time.Time{}, nil)
// 	return result
// }
