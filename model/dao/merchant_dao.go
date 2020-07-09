package dao

import (
	"errors"
	"shadowDemo/model/do"
	"shadowDemo/shadow-framework/utils"
	"sync"
	"github.com/jinzhu/gorm"
)

type MerchantDao struct {
	db    *gorm.DB
	mutex *sync.Mutex
}

var merchantDao *MerchantDao = nil

func NewMerchantDao(db *gorm.DB) *MerchantDao {
	merchantDao = &MerchantDao{
		db:    db,
		mutex: &sync.Mutex{},
	}
	return merchantDao
}

func GetMerchantDao() *MerchantDao {
	utils.ASSERT(merchantDao != nil)
	return merchantDao
}

func (dao *MerchantDao) Lock() {
	dao.mutex.Lock()
}

func (dao *MerchantDao) Unlock() {
	dao.mutex.Unlock()
}

func (dao *MerchantDao) Create(m *do.Merchant) error {
	return dao.db.Create(m).Error
}

func (dao *MerchantDao) Find(m *do.Merchant) (result []*do.Merchant, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *MerchantDao) FindOne(m *do.Merchant) error {
	return dao.db.First(m, m).Error
}

func (dao *MerchantDao) FindLast(m *do.Merchant) error {
	return dao.db.Last(m, m).Error
}

func (dao *MerchantDao) Get(m *do.Merchant) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *MerchantDao) BatchGet(idbatch []int64) (result []*do.Merchant, err error) {
	if len(idbatch) == 0 {
		return nil, errors.New("id is nil")
	}
	err = dao.db.Model(&do.Merchant{}).Where("id in (?)", idbatch).Find(&result).Error
	return
}

func (dao *MerchantDao) GetForUpdate(m *do.Merchant) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *MerchantDao) Save(m *do.Merchant) error {
	return dao.db.Save(m).Error
}

func (dao *MerchantDao) Delete(m *do.Merchant) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *MerchantDao) BatchDelete(idbatch []int64) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Where("id in (?)", idbatch).Delete(&do.Merchant{}).Error
}

func (dao *MerchantDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&do.Merchant{}).Where("id = ?", id).Updates(attrs).Error
}

func (dao *MerchantDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&do.Merchant{}).Where("id = ?", id).Update(attr, value).Error
}

func (dao *MerchantDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Model(&do.Merchant{}).Where("id in (?)", idbatch).Updates(attrs).Error
}

func (dao *MerchantDao) Found(m *do.Merchant) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
