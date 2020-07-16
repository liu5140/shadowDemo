package dao

import (
	"errors"
	"shadowDemo/model/do"
	"shadowDemo/zframework/utils"
	"sync"
	"github.com/jinzhu/gorm"
)

type IPWhiteListDao struct {
	db    *gorm.DB
	mutex *sync.Mutex
}

var iPWhiteListDao *IPWhiteListDao = nil

func NewIPWhiteListDao(db *gorm.DB) *IPWhiteListDao {
	iPWhiteListDao = &IPWhiteListDao{
		db:    db,
		mutex: &sync.Mutex{},
	}
	return iPWhiteListDao
}

func GetIPWhiteListDao() *IPWhiteListDao {
	utils.ASSERT(iPWhiteListDao != nil)
	return iPWhiteListDao
}

func (dao *IPWhiteListDao) Lock() {
	dao.mutex.Lock()
}

func (dao *IPWhiteListDao) Unlock() {
	dao.mutex.Unlock()
}

func (dao *IPWhiteListDao) Create(m *do.IPWhiteList) error {
	return dao.db.Create(m).Error
}

func (dao *IPWhiteListDao) Find(m *do.IPWhiteList) (result []*do.IPWhiteList, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *IPWhiteListDao) FindOne(m *do.IPWhiteList) error {
	return dao.db.First(m, m).Error
}

func (dao *IPWhiteListDao) FindLast(m *do.IPWhiteList) error {
	return dao.db.Last(m, m).Error
}

func (dao *IPWhiteListDao) Get(m *do.IPWhiteList) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *IPWhiteListDao) BatchGet(idbatch []int64) (result []*do.IPWhiteList, err error) {
	if len(idbatch) == 0 {
		return nil, errors.New("id is nil")
	}
	err = dao.db.Model(&do.IPWhiteList{}).Where("id in (?)", idbatch).Find(&result).Error
	return
}

func (dao *IPWhiteListDao) GetForUpdate(m *do.IPWhiteList) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *IPWhiteListDao) Save(m *do.IPWhiteList) error {
	return dao.db.Save(m).Error
}

func (dao *IPWhiteListDao) Delete(m *do.IPWhiteList) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *IPWhiteListDao) BatchDelete(idbatch []int64) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Where("id in (?)", idbatch).Delete(&do.IPWhiteList{}).Error
}

func (dao *IPWhiteListDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&do.IPWhiteList{}).Where("id = ?", id).Updates(attrs).Error
}

func (dao *IPWhiteListDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&do.IPWhiteList{}).Where("id = ?", id).Update(attr, value).Error
}

func (dao *IPWhiteListDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Model(&do.IPWhiteList{}).Where("id in (?)", idbatch).Updates(attrs).Error
}

func (dao *IPWhiteListDao) Found(m *do.IPWhiteList) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
