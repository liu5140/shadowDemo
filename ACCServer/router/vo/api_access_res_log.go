package vo

import (
	"shadowDemo/model/do"
	"time"
)

// swagger:parameters aPIAccessResLog createAPIAccessResLog
type CreateAPIAccessResLogRequest struct {
	//in:body
	Body struct {
		
	}
}

// swagger:parameters aPIAccessResLog searchAPIAccessResLog
type SearchAPIAccessResLogRequest struct {
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
//swagger:response searchAPIAccessResLogResponse
type SearchAPIAccessResLogResponse struct {
	//in: body
	Result SearchAPIAccessResLogResponseBody
}

// swagger:model
type SearchAPIAccessResLogResponseBody struct {
	Result []do.APIAccessResLog
	Count  int
}

//详情入参
// swagger:parameters aPIAccessResLog getAPIAccessResLog
type GetAPIAccessResLogRequest struct {
	//in: query
	ID int64
}

//详情出参
// swagger:response getAPIAccessResLogResponse
type GetAPIAccessResLogResponse struct {
	//in: body
	Result GetAPIAccessResLogResponseBody
}

//swagger:model
type GetAPIAccessResLogResponseBody struct {
	Result do.APIAccessResLog
}

//修改
// swagger:parameters aPIAccessResLog updateAPIAccessResLog
type UpdateAPIAccessResLogRequest struct {
	// in:query
	//id主键
	ID int64
	// in:body
	Body struct {
		
	}
}

//删除
// swagger:parameters aPIAccessResLog deleteAPIAccessResLog
type DeleteAPIAccessResLogRequest struct {
	//in: query
	ID int64
}
