package dao

import (
	"errors"
	"shadowDemo/model/do"
	"shadowDemo/shadow-framework/utils"
	"sync"
	"github.com/jinzhu/gorm"
)

type UserDao struct {
	db    *gorm.DB
	mutex *sync.Mutex
}

var userDao *UserDao = nil

func NewUserDao(db *gorm.DB) *UserDao {
	userDao = &UserDao{
		db:    db,
		mutex: &sync.Mutex{},
	}
	return userDao
}

func GetUserDao() *UserDao {
	utils.ASSERT(userDao != nil)
	return userDao
}

func (dao *UserDao) Lock() {
	dao.mutex.Lock()
}

func (dao *UserDao) Unlock() {
	dao.mutex.Unlock()
}

func (dao *UserDao) Create(m *do.User) error {
	return dao.db.Create(m).Error
}

func (dao *UserDao) Find(m *do.User) (result []*do.User, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *UserDao) FindOne(m *do.User) error {
	return dao.db.First(m, m).Error
}

func (dao *UserDao) FindLast(m *do.User) error {
	return dao.db.Last(m, m).Error
}

func (dao *UserDao) Get(m *do.User) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *UserDao) BatchGet(idbatch []int64) (result []*do.User, err error) {
	if len(idbatch) == 0 {
		return nil, errors.New("id is nil")
	}
	err = dao.db.Model(&do.User{}).Where("id in (?)", idbatch).Find(&result).Error
	return
}

func (dao *UserDao) GetForUpdate(m *do.User) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *UserDao) Save(m *do.User) error {
	return dao.db.Save(m).Error
}

func (dao *UserDao) Delete(m *do.User) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *UserDao) BatchDelete(idbatch []int64) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Where("id in (?)", idbatch).Delete(&do.User{}).Error
}

func (dao *UserDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&do.User{}).Where("id = ?", id).Updates(attrs).Error
}

func (dao *UserDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&do.User{}).Where("id = ?", id).Update(attr, value).Error
}

func (dao *UserDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Model(&do.User{}).Where("id in (?)", idbatch).Updates(attrs).Error
}

func (dao *UserDao) Found(m *do.User) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
