package dto

import "shadowDemo/model/do"

type UpdatePwd struct {
	//旧密码
	OldPwd string
	//新密码
	NewPwd string `binding:"required"`
	//确认密码
	EnsurePwd string `binding:"required"`
	//密码修改类型
	UpdatePwdType do.UpdatePwdType `binding:"required"`
}