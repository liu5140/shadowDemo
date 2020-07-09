package vo

import (
	"shadowDemo/model/do"
	"time"
)

// swagger:parameters user createUser
type CreateUserRequest struct {
	//in:body
	Body struct {
		
	}
}

// swagger:parameters user searchUser
type SearchUserRequest struct {
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
//swagger:response searchUserResponse
type SearchUserResponse struct {
	//in: body
	Result SearchUserResponseBody
}

// swagger:model
type SearchUserResponseBody struct {
	Result []do.User
	Count  int
}

//详情入参
// swagger:parameters user getUser
type GetUserRequest struct {
	//in: query
	ID int64
}

//详情出参
// swagger:response getUserResponse
type GetUserResponse struct {
	//in: body
	Result GetUserResponseBody
}

//swagger:model
type GetUserResponseBody struct {
	Result do.User
}

//修改
// swagger:parameters user updateUser
type UpdateUserRequest struct {
	// in:query
	//id主键
	ID int64
	// in:body
	Body struct {
		
	}
}

//删除
// swagger:parameters user deleteUser
type DeleteUserRequest struct {
	//in: query
	ID int64
}
