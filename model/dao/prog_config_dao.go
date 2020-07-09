package dao

import (
	"errors"
	"shadowDemo/model/do"
	"shadowDemo/shadow-framework/utils"
	"sync"
	"github.com/jinzhu/gorm"
)

type ProgConfigDao struct {
	db    *gorm.DB
	mutex *sync.Mutex
}

var progConfigDao *ProgConfigDao = nil

func NewProgConfigDao(db *gorm.DB) *ProgConfigDao {
	progConfigDao = &ProgConfigDao{
		db:    db,
		mutex: &sync.Mutex{},
	}
	return progConfigDao
}

func GetProgConfigDao() *ProgConfigDao {
	utils.ASSERT(progConfigDao != nil)
	return progConfigDao
}

func (dao *ProgConfigDao) Lock() {
	dao.mutex.Lock()
}

func (dao *ProgConfigDao) Unlock() {
	dao.mutex.Unlock()
}

func (dao *ProgConfigDao) Create(m *do.ProgConfig) error {
	return dao.db.Create(m).Error
}

func (dao *ProgConfigDao) Find(m *do.ProgConfig) (result []*do.ProgConfig, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *ProgConfigDao) FindOne(m *do.ProgConfig) error {
	return dao.db.First(m, m).Error
}

func (dao *ProgConfigDao) FindLast(m *do.ProgConfig) error {
	return dao.db.Last(m, m).Error
}

func (dao *ProgConfigDao) Get(m *do.ProgConfig) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *ProgConfigDao) BatchGet(idbatch []int64) (result []*do.ProgConfig, err error) {
	if len(idbatch) == 0 {
		return nil, errors.New("id is nil")
	}
	err = dao.db.Model(&do.ProgConfig{}).Where("id in (?)", idbatch).Find(&result).Error
	return
}

func (dao *ProgConfigDao) GetForUpdate(m *do.ProgConfig) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *ProgConfigDao) Save(m *do.ProgConfig) error {
	return dao.db.Save(m).Error
}

func (dao *ProgConfigDao) Delete(m *do.ProgConfig) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *ProgConfigDao) BatchDelete(idbatch []int64) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Where("id in (?)", idbatch).Delete(&do.ProgConfig{}).Error
}

func (dao *ProgConfigDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&do.ProgConfig{}).Where("id = ?", id).Updates(attrs).Error
}

func (dao *ProgConfigDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&do.ProgConfig{}).Where("id = ?", id).Update(attr, value).Error
}

func (dao *ProgConfigDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Model(&do.ProgConfig{}).Where("id in (?)", idbatch).Updates(attrs).Error
}

func (dao *ProgConfigDao) Found(m *do.ProgConfig) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
