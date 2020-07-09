package do

import "shadowDemo/zframework/logger"

var Log *logger.Logger = nil

func init() {
	Log = logger.InitLog()
}

// swagger:model
type (
	Gender   int //性别
	UserType int //用户类型
	UpdatePwdType               int //密码类型

)

const (
	_       Gender = iota
	Male           //性别: 男 1
	FeMale         //性别: 女 2
	Unkonwn        //性别：无 3
)

const (
	_                   UserType = iota
	UserTypeAdmin                //站长	1
	UserTypePlayer               //玩家	2
	UserTypeSubAccount           //子账号	3
)

const (
	_                   UpdatePwdType = iota
	UpdatePwdTypeLogin                //修改登录密码	1
	UpdatePwdTypeSecure               //修改安全密码	2
)