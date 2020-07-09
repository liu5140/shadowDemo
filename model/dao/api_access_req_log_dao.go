package dao

import (
	"errors"
	"shadowDemo/model/do"
	"shadowDemo/zframework/utils"
	"sync"

	"github.com/jinzhu/gorm"
)

type APIAccessReqLogDao struct {
	db    *gorm.DB
	mutex *sync.Mutex
}

var aPIAccessReqLogDao *APIAccessReqLogDao = nil

func NewAPIAccessReqLogDao(db *gorm.DB) *APIAccessReqLogDao {
	aPIAccessReqLogDao = &APIAccessReqLogDao{
		db:    db,
		mutex: &sync.Mutex{},
	}
	return aPIAccessReqLogDao
}

func GetAPIAccessReqLogDao() *APIAccessReqLogDao {
	utils.ASSERT(aPIAccessReqLogDao != nil)
	return aPIAccessReqLogDao
}

func (dao *APIAccessReqLogDao) Lock() {
	dao.mutex.Lock()
}

func (dao *APIAccessReqLogDao) Unlock() {
	dao.mutex.Unlock()
}

func (dao *APIAccessReqLogDao) Create(m *do.APIAccessReqLog) error {
	return dao.db.Create(m).Error
}

func (dao *APIAccessReqLogDao) Find(m *do.APIAccessReqLog) (result []*do.APIAccessReqLog, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *APIAccessReqLogDao) FindOne(m *do.APIAccessReqLog) error {
	return dao.db.First(m, m).Error
}

func (dao *APIAccessReqLogDao) FindLast(m *do.APIAccessReqLog) error {
	return dao.db.Last(m, m).Error
}

func (dao *APIAccessReqLogDao) Get(m *do.APIAccessReqLog) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *APIAccessReqLogDao) BatchGet(idbatch []int64) (result []*do.APIAccessReqLog, err error) {
	if len(idbatch) == 0 {
		return nil, errors.New("id is nil")
	}
	err = dao.db.Model(&do.APIAccessReqLog{}).Where("id in (?)", idbatch).Find(&result).Error
	return
}

func (dao *APIAccessReqLogDao) GetForUpdate(m *do.APIAccessReqLog) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *APIAccessReqLogDao) Save(m *do.APIAccessReqLog) error {
	return dao.db.Save(m).Error
}

func (dao *APIAccessReqLogDao) Delete(m *do.APIAccessReqLog) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *APIAccessReqLogDao) BatchDelete(idbatch []int64) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Where("id in (?)", idbatch).Delete(&do.APIAccessReqLog{}).Error
}

func (dao *APIAccessReqLogDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&do.APIAccessReqLog{}).Where("id = ?", id).Updates(attrs).Error
}

func (dao *APIAccessReqLogDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&do.APIAccessReqLog{}).Where("id = ?", id).Update(attr, value).Error
}

func (dao *APIAccessReqLogDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Model(&do.APIAccessReqLog{}).Where("id in (?)", idbatch).Updates(attrs).Error
}

func (dao *APIAccessReqLogDao) Found(m *do.APIAccessReqLog) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
