package dao

import (
	"errors"
	"shadowDemo/model/do"
	"shadowDemo/zframework/utils"
	"sync"
	"github.com/jinzhu/gorm"
)

type UpmsAdminDao struct {
	db    *gorm.DB
	mutex *sync.Mutex
}

var upmsAdminDao *UpmsAdminDao = nil

func NewUpmsAdminDao(db *gorm.DB) *UpmsAdminDao {
	upmsAdminDao = &UpmsAdminDao{
		db:    db,
		mutex: &sync.Mutex{},
	}
	return upmsAdminDao
}

func GetUpmsAdminDao() *UpmsAdminDao {
	utils.ASSERT(upmsAdminDao != nil)
	return upmsAdminDao
}

func (dao *UpmsAdminDao) Lock() {
	dao.mutex.Lock()
}

func (dao *UpmsAdminDao) Unlock() {
	dao.mutex.Unlock()
}

func (dao *UpmsAdminDao) Create(m *do.UpmsAdmin) error {
	return dao.db.Create(m).Error
}

func (dao *UpmsAdminDao) Find(m *do.UpmsAdmin) (result []*do.UpmsAdmin, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *UpmsAdminDao) FindOne(m *do.UpmsAdmin) error {
	return dao.db.First(m, m).Error
}

func (dao *UpmsAdminDao) FindLast(m *do.UpmsAdmin) error {
	return dao.db.Last(m, m).Error
}

func (dao *UpmsAdminDao) Get(m *do.UpmsAdmin) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *UpmsAdminDao) BatchGet(idbatch []int64) (result []*do.UpmsAdmin, err error) {
	if len(idbatch) == 0 {
		return nil, errors.New("id is nil")
	}
	err = dao.db.Model(&do.UpmsAdmin{}).Where("id in (?)", idbatch).Find(&result).Error
	return
}

func (dao *UpmsAdminDao) GetForUpdate(m *do.UpmsAdmin) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *UpmsAdminDao) Save(m *do.UpmsAdmin) error {
	return dao.db.Save(m).Error
}

func (dao *UpmsAdminDao) Delete(m *do.UpmsAdmin) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *UpmsAdminDao) BatchDelete(idbatch []int64) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Where("id in (?)", idbatch).Delete(&do.UpmsAdmin{}).Error
}

func (dao *UpmsAdminDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&do.UpmsAdmin{}).Where("id = ?", id).Updates(attrs).Error
}

func (dao *UpmsAdminDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&do.UpmsAdmin{}).Where("id = ?", id).Update(attr, value).Error
}

func (dao *UpmsAdminDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Model(&do.UpmsAdmin{}).Where("id in (?)", idbatch).Updates(attrs).Error
}

func (dao *UpmsAdminDao) Found(m *do.UpmsAdmin) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
