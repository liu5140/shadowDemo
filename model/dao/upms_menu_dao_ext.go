package dao

import (
	"shadowDemo/model/do"
	"shadowDemo/zframework/model"
	"time"
)

type UpmsMenuSearchCondition struct {
	//id
	IDS []int64
	//创建开始时间
	CreateStartTime time.Time
	//创建截止时间
	CreateEndTime time.Time
	//路径
	Path string
}

func (dao *UpmsMenuDao) SearchUpmsMenus(condition *UpmsMenuSearchCondition, rowbound *model.RowBound) (result []do.UpmsMenu, count int, err error) {
	db := dao.db

	if len(condition.IDS) > 0 {
		db = db.Where("id in (?) ", condition.IDS)
	}

	//类型
	if condition.Path != "" {
		db = db.Where("menu.path  like  ? ", condition.Path+"%")
	}

	//创建时间
	if !condition.CreateStartTime.IsZero() && !condition.CreateEndTime.IsZero() {
		db = db.Where("created_at BETWEEN ? AND ?", condition.CreateStartTime, condition.CreateEndTime)
	} else if !condition.CreateStartTime.IsZero() {
		db = db.Where("created_at >= ?", condition.CreateStartTime)
	} else if !condition.CreateEndTime.IsZero() {
		db = db.Where("created_at <= ?", condition.CreateEndTime)
	}

	if rowbound == nil {
		err = db.Model(&do.UpmsMenu{}).Order("ID desc").Count(&count).Find(&result).Error
	} else {
		err = db.Model(&do.UpmsMenu{}).Order("ID desc").Count(&count).Offset(rowbound.Offset).Limit(rowbound.Limit).Find(&result).Error
	}

	if err != nil {
		return result, 0, err
	}

	return result, count, nil
}

func (dao *UpmsMenuDao) UpdatePath(frompath string, topath string) error {
	return dao.db.Exec("update menu set path =REPLACE(path,?,?),updated_at =Now() where path like ?  ", frompath, topath, frompath+"%").Error
}
