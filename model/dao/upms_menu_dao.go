package dao

import (
	"errors"
	"shadowDemo/model/do"
	"shadowDemo/zframework/utils"
	"sync"
	"github.com/jinzhu/gorm"
)

type UpmsMenuDao struct {
	db    *gorm.DB
	mutex *sync.Mutex
}

var upmsMenuDao *UpmsMenuDao = nil

func NewUpmsMenuDao(db *gorm.DB) *UpmsMenuDao {
	upmsMenuDao = &UpmsMenuDao{
		db:    db,
		mutex: &sync.Mutex{},
	}
	return upmsMenuDao
}

func GetUpmsMenuDao() *UpmsMenuDao {
	utils.ASSERT(upmsMenuDao != nil)
	return upmsMenuDao
}

func (dao *UpmsMenuDao) Lock() {
	dao.mutex.Lock()
}

func (dao *UpmsMenuDao) Unlock() {
	dao.mutex.Unlock()
}

func (dao *UpmsMenuDao) Create(m *do.UpmsMenu) error {
	return dao.db.Create(m).Error
}

func (dao *UpmsMenuDao) Find(m *do.UpmsMenu) (result []*do.UpmsMenu, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *UpmsMenuDao) FindOne(m *do.UpmsMenu) error {
	return dao.db.First(m, m).Error
}

func (dao *UpmsMenuDao) FindLast(m *do.UpmsMenu) error {
	return dao.db.Last(m, m).Error
}

func (dao *UpmsMenuDao) Get(m *do.UpmsMenu) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *UpmsMenuDao) BatchGet(idbatch []int64) (result []*do.UpmsMenu, err error) {
	if len(idbatch) == 0 {
		return nil, errors.New("id is nil")
	}
	err = dao.db.Model(&do.UpmsMenu{}).Where("id in (?)", idbatch).Find(&result).Error
	return
}

func (dao *UpmsMenuDao) GetForUpdate(m *do.UpmsMenu) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *UpmsMenuDao) Save(m *do.UpmsMenu) error {
	return dao.db.Save(m).Error
}

func (dao *UpmsMenuDao) Delete(m *do.UpmsMenu) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *UpmsMenuDao) BatchDelete(idbatch []int64) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Where("id in (?)", idbatch).Delete(&do.UpmsMenu{}).Error
}

func (dao *UpmsMenuDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&do.UpmsMenu{}).Where("id = ?", id).Updates(attrs).Error
}

func (dao *UpmsMenuDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&do.UpmsMenu{}).Where("id = ?", id).Update(attr, value).Error
}

func (dao *UpmsMenuDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Model(&do.UpmsMenu{}).Where("id in (?)", idbatch).Updates(attrs).Error
}

func (dao *UpmsMenuDao) Found(m *do.UpmsMenu) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
