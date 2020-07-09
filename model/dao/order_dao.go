package dao

import (
	"errors"
	"shadowDemo/model/do"
	"shadowDemo/zframework/utils"
	"sync"

	"github.com/jinzhu/gorm"
)

type OrderDao struct {
	db    *gorm.DB
	mutex *sync.Mutex
}

var orderDao *OrderDao = nil

func NewOrderDao(db *gorm.DB) *OrderDao {
	orderDao = &OrderDao{
		db:    db,
		mutex: &sync.Mutex{},
	}
	return orderDao
}

func GetOrderDao() *OrderDao {
	utils.ASSERT(orderDao != nil)
	return orderDao
}

func (dao *OrderDao) Lock() {
	dao.mutex.Lock()
}

func (dao *OrderDao) Unlock() {
	dao.mutex.Unlock()
}

func (dao *OrderDao) Create(m *do.Order) error {
	return dao.db.Create(m).Error
}

func (dao *OrderDao) Find(m *do.Order) (result []*do.Order, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *OrderDao) FindOne(m *do.Order) error {
	return dao.db.First(m, m).Error
}

func (dao *OrderDao) FindLast(m *do.Order) error {
	return dao.db.Last(m, m).Error
}

func (dao *OrderDao) Get(m *do.Order) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *OrderDao) BatchGet(idbatch []int64) (result []*do.Order, err error) {
	if len(idbatch) == 0 {
		return nil, errors.New("id is nil")
	}
	err = dao.db.Model(&do.Order{}).Where("id in (?)", idbatch).Find(&result).Error
	return
}

func (dao *OrderDao) GetForUpdate(m *do.Order) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *OrderDao) Save(m *do.Order) error {
	return dao.db.Save(m).Error
}

func (dao *OrderDao) Delete(m *do.Order) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *OrderDao) BatchDelete(idbatch []int64) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Where("id in (?)", idbatch).Delete(&do.Order{}).Error
}

func (dao *OrderDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&do.Order{}).Where("id = ?", id).Updates(attrs).Error
}

func (dao *OrderDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&do.Order{}).Where("id = ?", id).Update(attr, value).Error
}

func (dao *OrderDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Model(&do.Order{}).Where("id in (?)", idbatch).Updates(attrs).Error
}

func (dao *OrderDao) Found(m *do.Order) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
