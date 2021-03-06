package dao

import (
	"shadowDemo/model/do"
	"shadowDemo/zframework/model"
	"time"
)

type APIAccessResLogSearchCondition struct {
	//id
	IDS []int64
	//创建开始时间
	CreateStartTime time.Time
	//创建截止时间
	CreateEndTime time.Time
}

func (dao *APIAccessResLogDao) SearchAPIAccessResLogs(condition *APIAccessResLogSearchCondition, rowbound *model.RowBound) (result []do.APIAccessResLog, count int, err error) {
	db := dao.db

	if len(condition.IDS) > 0 {
		db = db.Where("id in (?) ", condition.IDS)
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
		err = db.Model(&do.APIAccessResLog{}).Order("ID desc").Count(&count).Find(&result).Error
	} else {
		err = db.Model(&do.APIAccessResLog{}).Order("ID desc").Count(&count).Offset(rowbound.Offset).Limit(rowbound.Limit).Find(&result).Error
	}

	if err != nil {
		return result, 0, err
	}

	return result, count, nil
}
