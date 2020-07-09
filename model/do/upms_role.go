package do

import (
	"time"
)

type UpmsRole struct {
	//ID
	ID int64 `gorm:"primary_key"`
	//创建时间
	CreatedAt *time.Time
	//修改时间
	UpdatedAt *time.Time
	//创建人
	CreatedBy string
	//名称
	Name string
	//角色编码
	Code string
}
