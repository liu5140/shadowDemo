package service

import (
	"shadowDemo/model"
	"shadowDemo/model/dao"
	"shadowDemo/model/do"
	modelc "shadowDemo/zframework/model"
)

type UpmsRoleService struct{}

var upmsRoleService *UpmsRoleService

func NewUpmsRoleService() *UpmsRoleService {
	if upmsRoleService == nil {
		l.Lock()
		if upmsRoleService == nil {
			upmsRoleService = &UpmsRoleService{}
		}
		l.Unlock()
	}
	return upmsRoleService
}

//创建
func (service *UpmsRoleService) CreateUpmsRole(upmsRole *do.UpmsRole) (err error) {
	return model.GetModel().UpmsRoleDao.Create(upmsRole)
}

//通过id获取详情
func (service *UpmsRoleService) GetUpmsRoleByID(id int64) (upmsRole *do.UpmsRole, err error) {
	upmsRole.ID = id
	err = model.GetModel().UpmsRoleDao.Get(upmsRole)
	if err != nil {
		Log.Error(err)
		return upmsRole, err
	}
	return upmsRole, err
}

//通过id删除
func (service *UpmsRoleService) DeleteUpmsRoleByID(id int64) (err error) {
	if model.GetModel().UpmsRoleDao.Delete(&do.UpmsRole{ID: id}); err != nil {
		Log.Error(err)
		return err
	}
	return err
}

//通过id更新
func (service *UpmsRoleService) UpdateUpmsRole(id int64, attrs map[string]interface{}) (err error) {
	if err = model.GetModel().UpmsRoleDao.Updates(id, attrs); err != nil {
		Log.Error(err)
		return err
	}
	return err
}

//查询
func (service *UpmsRoleService) SearchUpmsRolePaging(condition *dao.UpmsRoleSearchCondition, pageNum int, pageSize int) (request []do.UpmsRole, count int, err error) {
	rowbound := modelc.NewRowBound(pageNum, pageSize)
	return service.searchUpmsRole(condition, &rowbound)
}

func (service *UpmsRoleService) SearchUpmsRoleWithOutPaging(condition *dao.UpmsRoleSearchCondition) (request []do.UpmsRole, count int, err error) {
	return service.searchUpmsRole(condition, nil)
}

func (service *UpmsRoleService) searchUpmsRole(condition *dao.UpmsRoleSearchCondition, rowbound *modelc.RowBound) (request []do.UpmsRole, count int, err error) {
	result, count, err := model.GetModel().UpmsRoleDao.SearchUpmsRoles(condition, rowbound)
	if err != nil {
		Log.Error(err)
		return nil, 0, err
	}
	return result, count, err
}
