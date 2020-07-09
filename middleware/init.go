package middleware

import (
	"shadowDemo/model"
	modelc "shadowDemo/zframework/model"

	"shadowDemo/zframework/logger"
)

const (
	MENUS   = "menus"
	PROFILE = "profile"
)

var Log *logger.Logger

func init() {
	Log = logger.InitLog()
}

type Profile struct {
	//ID
	ID int64
	//名称
	Username string
	//帐号
	Account string
	//玩家类型
	UserType model.UserType
	//帐号状态,  正常/锁定/冻结/停用
	State modelc.AccountState
	//是否锁定
	Locked bool
	//是否登录
	Logined bool
	//角色
	Role string
	//ip地址
	IP string
	//具体地址
	IPAddr string
	//设备信息
	DevInfo modelc.Device
	//是否已经设置谷歌验证码
	SetGoogleToken bool
	//当期语言
	Lang string
	//当前域名
	Host string
	//用户层级
	UserLevel int64
	//用户层级名称
	UserLevelName string
	//appid
	CurrentSite string
}
