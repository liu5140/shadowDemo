package do

import "time"

type APIAccessResLog struct {
	ID           int64 `gorm:"primary_key"`
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	DeletedAt    *time.Time
	CreatedBy    string
	RequestNo    string
	MerchantNo   string
	MerchantName string
	IPAddress    string
	APIURL       string
	Request      string
	Method       string
	RequestTime  time.Time
	Response     string
	ResponseTime time.Time
	OrderNo      string
}
