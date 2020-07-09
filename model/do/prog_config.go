package do

import "time"

type ProgConfig struct {
	ID         int64 `gorm:"primary_key"`
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	ParamName  string `gorm:"size:64"`
	ParamValue string `gorm:"type:longtext"`
	Type       string `gorm:"size:32"`
	Disabled   bool   `gorm:""`
	Encrypted  bool   `gorm:""`
	Comment    string `gorm:"size:1024"`
}
