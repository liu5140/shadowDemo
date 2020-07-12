package vo

import (
	"github.com/shopspring/decimal"
	"shadowDemo/model/do"
)

// swagger:parameters role getPermissionByID
type RoleID struct {
	//in:query
	ID string `json:"ID"`
}

// swagger:parameters role createRole
type CreateRoleRequest struct {
	//in:body
	Body struct {
		//名称
		Name string `binding:"required"`
		//角色编码
		Code string `binding:"required"`
		//单次可操作最大金额
		MaxAmount decimal.Decimal `sql:"type:decimal(20,4);"`
	}
}

// swagger:parameters role createPermission
type CreateRolePermissionRequest struct {
	//in:query
	ID int64 `binding:"required"`
	//in:body
	//菜单
	Body []do.UpmsMenu
}

// swagger:parameters role UpdateRole
type UpdateRoleRequest struct {
	// in:query
	ID int64
	// in:body
	Body struct {
		//名字
		Name string
		//单次可操作最大金额
		MaxAmount decimal.Decimal `sql:"type:decimal(20,4);"`
	}
}

// swagger:parameters role deleteRole
type DeleteRoleRequest struct {
	// in:query
	ID int64
}

//角色信息
// swagger:response searchRoleResponse
type SearchRoleResponse struct {
	// in: body
	Result SearchRoleBody
}

//swagger:model
type SearchRoleBody struct {
	Data []*do.UpmsRole
}

// swagger:parameters role getRoleByID
type GetRoleByIDRequst struct {
	//in: query
	ID int64 `binding:"required"`
}

//菜单信息
// swagger:response searchRoleMenuResponse
type SearchRoleMenuResponse struct {
	// in: body
	Result []*do.UpmsMenu
}
