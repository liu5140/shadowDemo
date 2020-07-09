package do

import "time"

//商户信息
// sagger:model
type Merchant struct {
	ID         int64 `gorm:"primary_key"`
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	DeletedAt  *time.Time
	MerchantNo string
	Name       string
	// 帐号
	Account string `gorm:"not null;index" json:"account"`
	// 登录密码
	LoginPassword string `json:"loginPassword"`
	// 安全密码
	SecurePassword string `json:"securePassword"`
	// 登录密码的MD5加密字符串
	LoginPasswordMd5 string `json:"loginPasswordMd5"`

	AesKey string //aes加密

	Md5Secret string //用于api的md5加密
}
