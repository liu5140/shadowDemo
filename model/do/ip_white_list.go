package do

import "time"

type IPWhiteList struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	CreatedBy string
	AccountNo string
	//	AccountType AccountType
	AccountName string
	IP          string
	Enable      bool
}

func (model IPWhiteList) TableName() string {
	return "ip_white_list"
}
