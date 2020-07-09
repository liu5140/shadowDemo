package service

import (
	"shadowDemo/model"
	"shadowDemo/model/dao"
	"shadowDemo/model/do"
	modelc "shadowDemo/zframework/model"
)

type UserService struct{}

var userService *UserService

func NewUserService() *UserService {
	if userService == nil {
		l.Lock()
		if userService == nil {
			userService = &UserService{}
		}
		l.Unlock()
	}
	return userService
}

//创建
func (service *UserService) CreateUser(user *do.User) (err error) {
	return model.GetModel().UserDao.Create(user)
}

//通过id获取详情
func (service *UserService) GetUserByID(id int64) (user *do.User, err error) {
	user.ID = id
	err = model.GetModel().UserDao.Get(user)
	if err != nil {
		Log.Error(err)
		return user, err
	}
	return user, err
}

//通过id删除
func (service *UserService) DeleteUserByID(id int64) (err error) {
	if model.GetModel().UserDao.Delete(&do.User{ID: id}); err != nil {
		Log.Error(err)
		return err
	}
	return err
}

//通过id更新
func (service *UserService) UpdateUser(id int64, attrs map[string]interface{}) (err error) {
	if err = model.GetModel().UserDao.Updates(id, attrs); err != nil {
		Log.Error(err)
		return err
	}
	return err
}

//查询
func (service *UserService) SearchUserPaging(condition *dao.UserSearchCondition, pageNum int, pageSize int) (request []do.User, count int, err error) {
	rowbound := modelc.NewRowBound(pageNum, pageSize)
	return service.searchUser(condition, &rowbound)
}

func (service *UserService) SearchUserWithOutPaging(condition *dao.UserSearchCondition) (request []do.User, count int, err error) {
	return service.searchUser(condition, nil)
}

func (service *UserService) searchUser(condition *dao.UserSearchCondition, rowbound *modelc.RowBound) (request []do.User, count int, err error) {
	result, count, err := model.GetModel().UserDao.SearchUsers(condition, rowbound)
	if err != nil {
		Log.Error(err)
		return nil, 0, err
	}
	return result, count, err
}
