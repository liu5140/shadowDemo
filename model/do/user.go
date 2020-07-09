package do

import (
	"time"

	"shadowDemo/zframework/model"
)

// sagger:model
type User struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	// 帐号
	Account string `gorm:"not null;index" json:"account"`
	// 登录密码
	LoginPassword string `json:"loginPassword"`
	// 安全密码
	SecurePassword string `json:"securePassword"`
	// 登录密码的MD5加密字符串
	LoginPasswordMd5 string `json:"loginPasswordMd5"`
	// 是否设置安全密码
	IsSecurePwdSet bool `json:"lsSecurePwdSet"`
	// 真实姓名
	NickName string `json:"nickname"`
	// 层级ID
	UserLevelID int64 `gorm:"index" json:"userLevelID"`
	// 代理层级ID
	UserAgentLevelID int64 `gorm:"index" json:"userAgentLevelID"`
	// 手机
	Mobile string `json:"mobile"`
	// Qq
	Qq string `json:"qq"`
	// skype
	Skype string `json:"skype"`
	// 微信
	Wechat string `json:"wechat"`
	// 住址
	Address string `json:"address"`
	// 邮件
	Email string `json:"email"`
	// 时区
	TimeZone string `json:"timeZone"`
	// 生日
	Birthday *time.Time `json:"birthday"`
	// 性别
	// 国家地区
	Country string `json:"country"`
	// 来源终端(PC/后台注册/导入玩家/手机端/手机端H5/手机端ANDROID/手机端IOS)
	// 创建人ID
	Creator int64 `gorm:"creator"`
	// 注册URL
	RegistURL string `json:"registURL"`
	// 注册IP
	RegistIP string `json:"registIP"`
	// 推荐人的推广码, (如果玩家是被推荐的）
	RecommanderPCode string `json:"recommanderPCode"`
	// 帐号状态,  正常/锁定/冻结/停用
	// 状态修改时间
	StateTime *time.Time
	//状态
	State model.AccountState
	// 钱包ID
	WalletID int64 `gorm:"index"`
	// 最后登录时间
	LastLoginAt *time.Time
	// 最后登录IP
	LastLoginIP string `gorm:"index"`
	// 最后登录ip地址
	LastLoginIPAddr string
	// 最后登录的访问链接
	LastLogUrl string
	// 设备信息
	// 钱包
	//Wallet do.Wallet `gorm:"ForeignKey:WalletID" gorm:"association_foreignkey:UserID" json:"wallet"`
	Recommend string `gorm:"-"`
	//默认地址
	PhoneCode string `json:"phoneCode"`

	Sca int `gorm:"default:0;"`
}
