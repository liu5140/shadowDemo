package vo

import (
	"shadowDemo/model/do"
	"time"
)

// swagger:parameters upmsMenu createUpmsMenu
type CreateUpmsMenuRequest struct {
	//in:body
	Body struct {
		
	}
}

// swagger:parameters upmsMenu searchUpmsMenu
type SearchUpmsMenuRequest struct {
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
//swagger:response searchUpmsMenuResponse
type SearchUpmsMenuResponse struct {
	//in: body
	Result SearchUpmsMenuResponseBody
}

// swagger:model
type SearchUpmsMenuResponseBody struct {
	Result []do.UpmsMenu
	Count  int
}

//详情入参
// swagger:parameters upmsMenu getUpmsMenu
type GetUpmsMenuRequest struct {
	//in: query
	ID int64
}

//详情出参
// swagger:response getUpmsMenuResponse
type GetUpmsMenuResponse struct {
	//in: body
	Result GetUpmsMenuResponseBody
}

//swagger:model
type GetUpmsMenuResponseBody struct {
	Result do.UpmsMenu
}

//修改
// swagger:parameters upmsMenu updateUpmsMenu
type UpdateUpmsMenuRequest struct {
	// in:query
	//id主键
	ID int64
	// in:body
	Body struct {
		
	}
}

//删除
// swagger:parameters upmsMenu deleteUpmsMenu
type DeleteUpmsMenuRequest struct {
	//in: query
	ID int64
}
