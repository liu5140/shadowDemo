package dao

import (
	"errors"
	"shadowDemo/model/do"
	"shadowDemo/shadow-framework/utils"
	"sync"
	"github.com/jinzhu/gorm"
)

type APIAccessResLogDao struct {
	db    *gorm.DB
	mutex *sync.Mutex
}

var aPIAccessResLogDao *APIAccessResLogDao = nil

func NewAPIAccessResLogDao(db *gorm.DB) *APIAccessResLogDao {
	aPIAccessResLogDao = &APIAccessResLogDao{
		db:    db,
		mutex: &sync.Mutex{},
	}
	return aPIAccessResLogDao
}

func GetAPIAccessResLogDao() *APIAccessResLogDao {
	utils.ASSERT(aPIAccessResLogDao != nil)
	return aPIAccessResLogDao
}

func (dao *APIAccessResLogDao) Lock() {
	dao.mutex.Lock()
}

func (dao *APIAccessResLogDao) Unlock() {
	dao.mutex.Unlock()
}

func (dao *APIAccessResLogDao) Create(m *do.APIAccessResLog) error {
	return dao.db.Create(m).Error
}

func (dao *APIAccessResLogDao) Find(m *do.APIAccessResLog) (result []*do.APIAccessResLog, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *APIAccessResLogDao) FindOne(m *do.APIAccessResLog) error {
	return dao.db.First(m, m).Error
}

func (dao *APIAccessResLogDao) FindLast(m *do.APIAccessResLog) error {
	return dao.db.Last(m, m).Error
}

func (dao *APIAccessResLogDao) Get(m *do.APIAccessResLog) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *APIAccessResLogDao) BatchGet(idbatch []int64) (result []*do.APIAccessResLog, err error) {
	if len(idbatch) == 0 {
		return nil, errors.New("id is nil")
	}
	err = dao.db.Model(&do.APIAccessResLog{}).Where("id in (?)", idbatch).Find(&result).Error
	return
}

func (dao *APIAccessResLogDao) GetForUpdate(m *do.APIAccessResLog) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *APIAccessResLogDao) Save(m *do.APIAccessResLog) error {
	return dao.db.Save(m).Error
}

func (dao *APIAccessResLogDao) Delete(m *do.APIAccessResLog) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *APIAccessResLogDao) BatchDelete(idbatch []int64) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Where("id in (?)", idbatch).Delete(&do.APIAccessResLog{}).Error
}

func (dao *APIAccessResLogDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&do.APIAccessResLog{}).Where("id = ?", id).Updates(attrs).Error
}

func (dao *APIAccessResLogDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&do.APIAccessResLog{}).Where("id = ?", id).Update(attr, value).Error
}

func (dao *APIAccessResLogDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Model(&do.APIAccessResLog{}).Where("id in (?)", idbatch).Updates(attrs).Error
}

func (dao *APIAccessResLogDao) Found(m *do.APIAccessResLog) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
