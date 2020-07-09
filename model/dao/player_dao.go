package dao

import (
	"errors"
	"shadowDemo/model/do"
	"shadowDemo/zframework/utils"
	"sync"

	"github.com/jinzhu/gorm"
)

type PlayerDao struct {
	db    *gorm.DB
	mutex *sync.Mutex
}

var playerDao *PlayerDao = nil

func NewPlayerDao(db *gorm.DB) *PlayerDao {
	playerDao = &PlayerDao{
		db:    db,
		mutex: &sync.Mutex{},
	}
	return playerDao
}

func GetPlayerDao() *PlayerDao {
	utils.ASSERT(playerDao != nil)
	return playerDao
}

func (dao *PlayerDao) Lock() {
	dao.mutex.Lock()
}

func (dao *PlayerDao) Unlock() {
	dao.mutex.Unlock()
}

func (dao *PlayerDao) Create(m *do.Player) error {
	return dao.db.Create(m).Error
}

func (dao *PlayerDao) Find(m *do.Player) (result []*do.Player, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *PlayerDao) FindOne(m *do.Player) error {
	return dao.db.First(m, m).Error
}

func (dao *PlayerDao) FindLast(m *do.Player) error {
	return dao.db.Last(m, m).Error
}

func (dao *PlayerDao) Get(m *do.Player) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *PlayerDao) BatchGet(idbatch []int64) (result []*do.Player, err error) {
	if len(idbatch) == 0 {
		return nil, errors.New("id is nil")
	}
	err = dao.db.Model(&do.Player{}).Where("id in (?)", idbatch).Find(&result).Error
	return
}

func (dao *PlayerDao) GetForUpdate(m *do.Player) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *PlayerDao) Save(m *do.Player) error {
	return dao.db.Save(m).Error
}

func (dao *PlayerDao) Delete(m *do.Player) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *PlayerDao) BatchDelete(idbatch []int64) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Where("id in (?)", idbatch).Delete(&do.Player{}).Error
}

func (dao *PlayerDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&do.Player{}).Where("id = ?", id).Updates(attrs).Error
}

func (dao *PlayerDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&do.Player{}).Where("id = ?", id).Update(attr, value).Error
}

func (dao *PlayerDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Model(&do.Player{}).Where("id in (?)", idbatch).Updates(attrs).Error
}

func (dao *PlayerDao) Found(m *do.Player) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
