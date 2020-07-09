package dao

import (
	"errors"
	"shadowDemo/model/do"
	"shadowDemo/zframework/utils"
	"sync"
	"github.com/jinzhu/gorm"
)

type UpmsRoleDao struct {
	db    *gorm.DB
	mutex *sync.Mutex
}

var upmsRoleDao *UpmsRoleDao = nil

func NewUpmsRoleDao(db *gorm.DB) *UpmsRoleDao {
	upmsRoleDao = &UpmsRoleDao{
		db:    db,
		mutex: &sync.Mutex{},
	}
	return upmsRoleDao
}

func GetUpmsRoleDao() *UpmsRoleDao {
	utils.ASSERT(upmsRoleDao != nil)
	return upmsRoleDao
}

func (dao *UpmsRoleDao) Lock() {
	dao.mutex.Lock()
}

func (dao *UpmsRoleDao) Unlock() {
	dao.mutex.Unlock()
}

func (dao *UpmsRoleDao) Create(m *do.UpmsRole) error {
	return dao.db.Create(m).Error
}

func (dao *UpmsRoleDao) Find(m *do.UpmsRole) (result []*do.UpmsRole, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *UpmsRoleDao) FindOne(m *do.UpmsRole) error {
	return dao.db.First(m, m).Error
}

func (dao *UpmsRoleDao) FindLast(m *do.UpmsRole) error {
	return dao.db.Last(m, m).Error
}

func (dao *UpmsRoleDao) Get(m *do.UpmsRole) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *UpmsRoleDao) BatchGet(idbatch []int64) (result []*do.UpmsRole, err error) {
	if len(idbatch) == 0 {
		return nil, errors.New("id is nil")
	}
	err = dao.db.Model(&do.UpmsRole{}).Where("id in (?)", idbatch).Find(&result).Error
	return
}

func (dao *UpmsRoleDao) GetForUpdate(m *do.UpmsRole) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *UpmsRoleDao) Save(m *do.UpmsRole) error {
	return dao.db.Save(m).Error
}

func (dao *UpmsRoleDao) Delete(m *do.UpmsRole) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *UpmsRoleDao) BatchDelete(idbatch []int64) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Where("id in (?)", idbatch).Delete(&do.UpmsRole{}).Error
}

func (dao *UpmsRoleDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&do.UpmsRole{}).Where("id = ?", id).Updates(attrs).Error
}

func (dao *UpmsRoleDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&do.UpmsRole{}).Where("id = ?", id).Update(attr, value).Error
}

func (dao *UpmsRoleDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Model(&do.UpmsRole{}).Where("id in (?)", idbatch).Updates(attrs).Error
}

func (dao *UpmsRoleDao) Found(m *do.UpmsRole) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
