package vo

import (
	"shadowDemo/model/do"
	"time"
)

// swagger:parameters upmsRole createUpmsRole
type CreateUpmsRoleRequest struct {
	//in:body
	Body struct {
		
	}
}

// swagger:parameters upmsRole searchUpmsRole
type SearchUpmsRoleRequest struct {
	//in:body
	Body struct {
		//id
		IDS []int64
		//创建开始时间
		CreateStartTime time.Time
		//创建截止时间
		CreateEndTime time.Time
		//页码
		PagingRequest
	}
}

//查询列表
//swagger:response searchUpmsRoleResponse
type SearchUpmsRoleResponse struct {
	//in: body
	Result SearchUpmsRoleResponseBody
}

// swagger:model
type SearchUpmsRoleResponseBody struct {
	Result []do.UpmsRole
	Count  int
}

//详情入参
// swagger:parameters upmsRole getUpmsRole
type GetUpmsRoleRequest struct {
	//in: query
	ID int64
}

//详情出参
// swagger:response getUpmsRoleResponse
type GetUpmsRoleResponse struct {
	//in: body
	Result GetUpmsRoleResponseBody
}

//swagger:model
type GetUpmsRoleResponseBody struct {
	Result do.UpmsRole
}

//修改
// swagger:parameters upmsRole updateUpmsRole
type UpdateUpmsRoleRequest struct {
	// in:query
	//id主键
	ID int64
	// in:body
	Body struct {
		
	}
}

//删除
// swagger:parameters upmsRole deleteUpmsRole
type DeleteUpmsRoleRequest struct {
	//in: query
	ID int64
}
