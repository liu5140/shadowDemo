package do

import "time"

type APIAccessReqLog struct {
	ID           int64 `gorm:"primary_key"`
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	DeletedAt    *time.Time
	RequestNo    string
	MerchantNo   string
	MerchantName string
	IPAddress    string
	APIURL       string
	Request      string
	Method       string
	RequestTime  time.Time
	OrderNo      string
}

func (reqLog APIAccessReqLog) TableName() string {
	return "api_access_req_log"
}
