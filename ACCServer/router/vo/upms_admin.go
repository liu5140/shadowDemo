package vo

import (
	"shadowDemo/model/do"
	"time"
)

// swagger:parameters upmsAdmin createUpmsAdmin
type CreateUpmsAdminRequest struct {
	//in:body
	Body struct {
		
	}
}

// swagger:parameters upmsAdmin searchUpmsAdmin
type SearchUpmsAdminRequest struct {
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
//swagger:response searchUpmsAdminResponse
type SearchUpmsAdminResponse struct {
	//in: body
	Result SearchUpmsAdminResponseBody
}

// swagger:model
type SearchUpmsAdminResponseBody struct {
	Result []do.UpmsAdmin
	Count  int
}

//详情入参
// swagger:parameters upmsAdmin getUpmsAdmin
type GetUpmsAdminRequest struct {
	//in: query
	ID int64
}

//详情出参
// swagger:response getUpmsAdminResponse
type GetUpmsAdminResponse struct {
	//in: body
	Result GetUpmsAdminResponseBody
}

//swagger:model
type GetUpmsAdminResponseBody struct {
	Result do.UpmsAdmin
}

//修改
// swagger:parameters upmsAdmin updateUpmsAdmin
type UpdateUpmsAdminRequest struct {
	// in:query
	//id主键
	ID int64
	// in:body
	Body struct {
		
	}
}

//删除
// swagger:parameters upmsAdmin deleteUpmsAdmin
type DeleteUpmsAdminRequest struct {
	//in: query
	ID int64
}
