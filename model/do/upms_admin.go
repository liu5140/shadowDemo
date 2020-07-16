package do

import (
	"shadowDemo/zframework/model"
	"time"
)

//后台登陆用户
// sagger:model
type UpmsAdmin struct {
	ID        int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	//创建人
	CreatedBy string
	// 性别(子账号)
	Gender Gender
	//用户类型
	UserType UserType
	//帐号
	Account string
	//登录密码
	LoginPassword string
	//安全密码
	SecurePassword string
	//真实姓名
	RealName string
	//角色ID可以多个，用逗号分隔
	RoleID string
	//角色
	RoleCode string
	//角色名称
	RoleName string
	//可以访问的siteID，总站长不受此限制
	SiteID string
	//备注
	Remark string
	//帐号状态,  正常/锁定/冻结/停用
	State model.AccountState
}

func (admin *UpmsAdmin) GetUsername() string {
	return admin.Account
}
func (admin *UpmsAdmin) GetPassword() string {
	return admin.LoginPassword
}
func (admin *UpmsAdmin) IsAccountExpired() bool {
	return admin.State == model.Deleted
}
func (admin *UpmsAdmin) IsAccountLocked() bool {
	return admin.State == model.Frozen
}
func (admin *UpmsAdmin) IsCredentialsExpired() bool {
	return false
}
