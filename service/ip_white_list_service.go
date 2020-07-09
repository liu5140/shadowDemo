package service

import (
	"shadowDemo/model"
	"shadowDemo/model/do"
)

type IPWhiteListService struct{}

var ipWhiteListService *IPWhiteListService

func NewIPWhiteListService() *IPWhiteListService {
	if ipWhiteListService == nil {
		l.Lock()
		if ipWhiteListService == nil {
			ipWhiteListService = &IPWhiteListService{}
		}
		l.Unlock()
	}
	return ipWhiteListService
}

func (service *IPWhiteListService) CreateIPWhiteList(ipwhiteList *do.IPWhiteList) error {
	return model.GetModel().IPWhiteListDao.Create(ipwhiteList)
}

func (service *IPWhiteListService) DeleteIPWhiteList(ipwhiteList *do.IPWhiteList) error {
	return model.GetModel().IPWhiteListDao.Delete(ipwhiteList)
}

func (service *IPWhiteListService) GetIPWhiteListByAccountNo(accountNo string) (result []*do.IPWhiteList, err error) {
	return model.GetModel().IPWhiteListDao.Find(&do.IPWhiteList{
		AccountNo: accountNo,
		Enable:    true,
	})
}

func (service *IPWhiteListService) GetAllIPWhiteListByAccountNo(accountNo string) (result []*do.IPWhiteList, err error) {
	return model.GetModel().IPWhiteListDao.Find(&do.IPWhiteList{
		AccountNo: accountNo,
	})
}

// func (service *IPWhiteListService) UpdateIPWhiteListByAccountNo(accountNo string, enable bool) error {
// 	db := datasource.DatasourceServiceInstance(datasource.DATASOURCE_MANAGER).Datasource()
// 	bccIPWhiteListDao := dao.NewBccIPWhiteListDao(db)
// 	return bccIPWhiteListDao.UpdateIPWhiteListByAccountNo(accountNo, enable)
// }
