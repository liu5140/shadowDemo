package vo

import (
	"shadowDemo/model/do"
	"time"
)

// swagger:parameters aPIAccessReqLog createAPIAccessReqLog
type CreateAPIAccessReqLogRequest struct {
	//in:body
	Body struct {
		
	}
}

// swagger:parameters aPIAccessReqLog searchAPIAccessReqLog
type SearchAPIAccessReqLogRequest struct {
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
//swagger:response searchAPIAccessReqLogResponse
type SearchAPIAccessReqLogResponse struct {
	//in: body
	Result SearchAPIAccessReqLogResponseBody
}

// swagger:model
type SearchAPIAccessReqLogResponseBody struct {
	Result []do.APIAccessReqLog
	Count  int
}

//详情入参
// swagger:parameters aPIAccessReqLog getAPIAccessReqLog
type GetAPIAccessReqLogRequest struct {
	//in: query
	ID int64
}

//详情出参
// swagger:response getAPIAccessReqLogResponse
type GetAPIAccessReqLogResponse struct {
	//in: body
	Result GetAPIAccessReqLogResponseBody
}

//swagger:model
type GetAPIAccessReqLogResponseBody struct {
	Result do.APIAccessReqLog
}

//修改
// swagger:parameters aPIAccessReqLog updateAPIAccessReqLog
type UpdateAPIAccessReqLogRequest struct {
	// in:query
	//id主键
	ID int64
	// in:body
	Body struct {
		
	}
}

//删除
// swagger:parameters aPIAccessReqLog deleteAPIAccessReqLog
type DeleteAPIAccessReqLogRequest struct {
	//in: query
	ID int64
}
