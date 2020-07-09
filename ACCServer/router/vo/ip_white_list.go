package vo

import (
	"shadowDemo/model/do"
	"time"
)

// swagger:parameters iPWhiteList createIPWhiteList
type CreateIPWhiteListRequest struct {
	//in:body
	Body struct {
		
	}
}

// swagger:parameters iPWhiteList searchIPWhiteList
type SearchIPWhiteListRequest struct {
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
//swagger:response searchIPWhiteListResponse
type SearchIPWhiteListResponse struct {
	//in: body
	Result SearchIPWhiteListResponseBody
}

// swagger:model
type SearchIPWhiteListResponseBody struct {
	Result []do.IPWhiteList
	Count  int
}

//详情入参
// swagger:parameters iPWhiteList getIPWhiteList
type GetIPWhiteListRequest struct {
	//in: query
	ID int64
}

//详情出参
// swagger:response getIPWhiteListResponse
type GetIPWhiteListResponse struct {
	//in: body
	Result GetIPWhiteListResponseBody
}

//swagger:model
type GetIPWhiteListResponseBody struct {
	Result do.IPWhiteList
}

//修改
// swagger:parameters iPWhiteList updateIPWhiteList
type UpdateIPWhiteListRequest struct {
	// in:query
	//id主键
	ID int64
	// in:body
	Body struct {
		
	}
}

//删除
// swagger:parameters iPWhiteList deleteIPWhiteList
type DeleteIPWhiteListRequest struct {
	//in: query
	ID int64
}
