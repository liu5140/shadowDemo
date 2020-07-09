package service

import (
	"shadowDemo/model"
	"shadowDemo/model/dao"
	"shadowDemo/model/do"
	modelc "shadowDemo/zframework/model"
)

type IPWhiteListService struct{}

var iPWhiteListService *IPWhiteListService

func NewIPWhiteListService() *IPWhiteListService {
	if iPWhiteListService == nil {
		l.Lock()
		if iPWhiteListService == nil {
			iPWhiteListService = &IPWhiteListService{}
		}
		l.Unlock()
	}
	return iPWhiteListService
}

//创建
func (service *IPWhiteListService) CreateIPWhiteList(iPWhiteList *do.IPWhiteList) (err error) {
	return model.GetModel().IPWhiteListDao.Create(iPWhiteList)
}

//通过id获取详情
func (service *IPWhiteListService) GetIPWhiteListByID(id int64) (iPWhiteList *do.IPWhiteList, err error) {
	iPWhiteList.ID = id
	err = model.GetModel().IPWhiteListDao.Get(iPWhiteList)
	if err != nil {
		Log.Error(err)
		return iPWhiteList, err
	}
	return iPWhiteList, err
}

//通过id删除
func (service *IPWhiteListService) DeleteIPWhiteListByID(id int64) (err error) {
	if model.GetModel().IPWhiteListDao.Delete(&do.IPWhiteList{ID: id}); err != nil {
		Log.Error(err)
		return err
	}
	return err
}

func (service *IPWhiteListService) GetIPWhiteListByAccountNo(account string) (request []*do.IPWhiteList, err error) {

	return model.GetModel().IPWhiteListDao.Find(&do.IPWhiteList{AccountNo: account})
}

//通过id更新
func (service *IPWhiteListService) UpdateIPWhiteList(id int64, attrs map[string]interface{}) (err error) {
	if err = model.GetModel().IPWhiteListDao.Updates(id, attrs); err != nil {
		Log.Error(err)
		return err
	}
	return err
}

//查询
func (service *IPWhiteListService) SearchIPWhiteListPaging(condition *dao.IPWhiteListSearchCondition, pageNum int, pageSize int) (request []do.IPWhiteList, count int, err error) {
	rowbound := modelc.NewRowBound(pageNum, pageSize)
	return service.searchIPWhiteList(condition, &rowbound)
}

func (service *IPWhiteListService) SearchIPWhiteListWithOutPaging(condition *dao.IPWhiteListSearchCondition) (request []do.IPWhiteList, count int, err error) {
	return service.searchIPWhiteList(condition, nil)
}

func (service *IPWhiteListService) searchIPWhiteList(condition *dao.IPWhiteListSearchCondition, rowbound *modelc.RowBound) (request []do.IPWhiteList, count int, err error) {
	result, count, err := model.GetModel().IPWhiteListDao.SearchIPWhiteLists(condition, rowbound)
	if err != nil {
		Log.Error(err)
		return nil, 0, err
	}
	return result, count, err
}
